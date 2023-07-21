package auth

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gbrlsnchs/jwt/v3"
	_ "github.com/go-sql-driver/mysql"
	"github.com/vavilen84/nft-project/helpers"
	"github.com/vavilen84/nft-project/models"
	"gorm.io/gorm"
	"regexp"
	"time"
)

type JWTPayload struct {
	jwt.Payload
	JWTInfoId int `json:"jwt_info_id"`
}

func insertJWTInfo(db *gorm.DB, u *models.User) (jwtInfo models.JWTInfo, err error) {
	jwtInfo = models.JWTInfo{
		User:      *u,
		ExpiresAt: helpers.GetDefaultJWTExpiresAt(),
	}
	err = models.InsertJWTInfo(db, &jwtInfo)
	if err != nil {
		helpers.LogError(err)
	}
	return
}

func generateJWT(jwtInfo models.JWTInfo) (token []byte, err error) {
	algorithm := jwt.NewHS256([]byte(jwtInfo.Secret))
	payload := JWTPayload{
		Payload: jwt.Payload{
			ExpirationTime: jwt.NumericDate(jwtInfo.ExpiresAt),
			IssuedAt:       jwt.NumericDate(jwtInfo.CreatedAt),
		},
		JWTInfoId: jwtInfo.Id,
	}
	token, err = jwt.Sign(payload, algorithm)
	if err != nil {
		helpers.LogError(err)
	}
	return
}

func CreateJWT(db *gorm.DB, u *models.User) (token []byte, err error) {
	jwtInfo, err := insertJWTInfo(db, u)
	if err != nil {
		helpers.LogError(err)
		return
	}
	token, err = generateJWT(jwtInfo)
	if err != nil {
		helpers.LogError(err)
	}
	return
}

func parseJWTPayload(token []byte) (jwtPayload JWTPayload, err error) {
	re, err := regexp.Compile(`(.*)\.(?P<payload>.*)\.(.*)`)
	if err != nil {
		helpers.LogError(err)
		return
	}
	matches := re.FindStringSubmatch(string(token))
	i := re.SubexpIndex("payload")
	payloadData, err := base64.StdEncoding.DecodeString(matches[i])
	if err != nil {
		helpers.LogError(err)
	}
	err = json.Unmarshal(payloadData, &jwtPayload)
	if err != nil {
		helpers.LogError(err)
	}
	return
}

func VerifyJWT(db *gorm.DB, token []byte) (isValid bool, err error) {
	payload, err := parseJWTPayload(token)
	if err != nil {
		helpers.LogError(err)
		return
	}
	jwtInfo, err := models.FindJWTInfoById(db, payload.JWTInfoId)
	if err != nil {
		helpers.LogError(err)
		return
	}
	algorithm := jwt.NewHS256([]byte(jwtInfo.Secret))
	_, err = jwt.Verify(token, algorithm, &payload)
	if err != nil {
		helpers.LogError(err)
		return
	}
	if payload.ExpirationTime.Before(time.Now()) {
		return
	}
	isValid = true
	return
}
