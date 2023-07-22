package models

import (
	"github.com/anaskhan96/go-password-encoder"
	"github.com/go-playground/validator/v10"
	"github.com/vavilen84/nft-project/helpers"
	"gorm.io/gorm"
)

type User struct {
	Id                         int    `json:"id" column:"id" gorm:"primaryKey;autoIncrement:true"`
	Email                      string `json:"email" column:"email" validate:"min=3,max=255,email,required"`
	Nickname                   string `json:"nickname" column:"nickname" validate:"min=3,max=255,required"`
	Password                   string `column:"password" validate:"min=3,max=5000,required"`
	PasswordSalt               string `column:"password_salt" validate:"min=3,max=5000,required"`
	PasswordResetToken         string `column:"password_reset_token" validate:"min=3,max=5000"`
	PasswordResetTokenExpireAt int64  `column:"password_reset_token_expire_at"`
	BillingPlan                int    `column:"billing_plan" validate:"min=1,max=5,required"`
	Role                       int    `column:"role" json:"role" validate:"lt=4,required"`
}

func (m *User) TableName() string {
	return "user"
}

func InsertUser(db *gorm.DB, m *User) (err error) {
	m.encodePassword()
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

func (m *User) encodePassword() {
	salt, encodedPwd := password.Encode(m.Password, nil)
	m.Password = encodedPwd
	m.PasswordSalt = salt
}

func FindUserById(db *gorm.DB, id int) (*User, error) {
	m := User{}
	err := db.First(&m, id).Error
	if err != nil {
		helpers.LogError(err)
	}
	return &m, err
}

func FindUserByEmail(db *gorm.DB, email string) (*User, error) {
	m := User{}
	err := db.Where("email = ?", email).First(&m).Error
	if err != nil {
		helpers.LogError(err)
	}
	return &m, err
}
