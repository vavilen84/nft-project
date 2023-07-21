package models

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/vavilen84/nft-project/helpers"
	"gorm.io/gorm"
	"time"
)

type JWTInfo struct {
	Id        int       `json:"id" column:"id" gorm:"primaryKey;autoIncrement:true"`
	UserId    int       `column:"user_id"`
	User      User      `gorm:"foreignKey:UserId"`
	Secret    string    `column:"secret" validate:"min=3,max=5000,required"`
	CreatedAt time.Time `json:"created_at" column:"created_at" gorm:"autoCreateTime"`
	ExpiresAt time.Time `json:"expires_at" column:"expires_at" validate:"required"`
}

func (m *JWTInfo) TableName() string {
	return "jwt_info"
}

func InsertJWTInfo(db *gorm.DB, m *JWTInfo) (err error) {
	m.generateSecret()
	validate := validator.New()
	err = validate.Struct(m)
	if err != nil {
		helpers.LogError(err)
		return
	}
	err = db.Create(m).Error
	if err != nil {
		helpers.LogError(err)
	}
	return
}

func (m *JWTInfo) generateSecret() {
	m.Secret = helpers.GenerateRandomString(64)
}

func FindJWTInfoById(db *gorm.DB, id int) (*JWTInfo, error) {
	m := JWTInfo{}
	err := db.First(&m, id).Error
	if err != nil {
		helpers.LogError(err)
	}
	return &m, err
}
