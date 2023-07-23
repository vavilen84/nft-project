package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/go-playground/validator/v10"
	"github.com/vavilen84/nft-project/auth"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/dto"
	"github.com/vavilen84/nft-project/helpers"
	"github.com/vavilen84/nft-project/models"
	"github.com/vavilen84/nft-project/store"
	"gorm.io/gorm"
	"net/http"
)

func (c *SecurityController) ChangePassword(w http.ResponseWriter, r *http.Request) {
	db := store.GetDB()
	token := r.Header.Get("Authorization")
	isValid, err := auth.VerifyJWT(db, []byte(token))
	if err != nil || token == "" || !isValid {
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.UnauthorizedError, http.StatusUnauthorized)
		return
	}
	dec := json.NewDecoder(r.Body)
	dtoModel := dto.ChangePassword{}
	err = dec.Decode(&dtoModel)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.BadRequestError, http.StatusBadRequest)
		return
	}
	validate := validator.New()
	err = validate.Struct(dtoModel)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.BadRequestError, http.StatusBadRequest)
		return
	}
	p, err := auth.ParseJWTPayload([]byte(token))
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.UnauthorizedError, http.StatusUnauthorized)
		return
	}

	jwtInfo, err := models.FindJWTInfoById(db, p.JWTInfoId)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.UnauthorizedError, http.StatusUnauthorized)
		return
	}

	u, err := models.FindUserByEmail(db, jwtInfo.User.Email)
	if err != nil {
		helpers.LogError(err)
		if err == gorm.ErrRecordNotFound {
			err = errors.New(fmt.Sprintf("user with email %s not found", jwtInfo.User.Email))
			c.WriteErrorResponse(w, err, http.StatusNotFound)
		} else {
			c.WriteErrorResponse(w, constants.ServerError, http.StatusInternalServerError)
		}
		return
	}
	passwordIsValid := password.Verify(dtoModel.OldPassword, u.PasswordSalt, u.Password, nil)
	if !passwordIsValid {
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.UnauthorizedError, http.StatusUnauthorized)
		return
	}
	u.Password = dtoModel.NewPassword
	err = models.UserChangePassword(db, u)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.ServerError, http.StatusInternalServerError)
		return
	}
	resp := make(dto.ResponseData)
	c.WriteSuccessResponse(w, resp, http.StatusOK)
}
