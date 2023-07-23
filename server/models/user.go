package models

import (
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/go-playground/validator/v10"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/helpers"
	"github.com/vavilen84/nft-project/validation"
	"gorm.io/gorm"
	"log"
	"regexp"
	"time"
	"unicode/utf8"
)

const alphaNumericRegexString = "^[a-zA-Z0-9]+$"

type User struct {
	Id                         int        `json:"id" column:"id" gorm:"primaryKey;autoIncrement:true"`
	Email                      string     `json:"email" column:"email"`
	Nickname                   string     `json:"nickname" column:"nickname"`
	Password                   string     `json:"password" column:"password"`
	PasswordSalt               string     `column:"password_salt"`
	PasswordResetToken         string     `column:"password_reset_token"`
	PasswordResetTokenExpireAt *time.Time `column:"password_reset_token_expire_at"`
	BillingPlan                int        `json:"billing_plan" column:"billing_plan" validate:"min=1,max=5,required"`
	Role                       int        `json:"role" column:"role" json:"role" validate:"lt=4,required"`
}

func CustomPasswordValidator(fl validator.FieldLevel) bool {
	p := fl.Field().String()
	length := utf8.RuneCountInString(p)
	if length < 8 {
		return false
	}
	r, err := regexp.Match(alphaNumericRegexString, []byte(p))
	if err != nil {
		fmt.Println(err.Error())
	}
	return r
}

func (m *User) TableName() string {
	return "user"
}

func (User) GetValidationRules() interface{} {
	return validation.ScenarioRules{
		constants.ScenarioCreate: validation.FieldRules{
			"Email":        "min=3,max=255,email,required",
			"Nickname":     "min=3,max=255,required",
			"Password":     "min=8,max=255,required,customPasswordValidator",
			"PasswordSalt": "min=3,max=5000,required",
			"BillingPlan":  "required,gt=0,lt=4",
			"Role":         "required,gt=0,lt=2", // we can create only users, admin should be created separately
		},
	}
}

func (User) GetValidator() interface{} {
	v := validator.New()
	err := v.RegisterValidation("customPasswordValidator", CustomPasswordValidator)
	if err != nil {
		helpers.LogError(err)
		return nil
	}
	return v
}

func InsertUser(db *gorm.DB, m *User) (err error) {
	m.encodePassword()
	err = validation.ValidateByScenario(constants.ScenarioCreate, *m)
	if err != nil {
		log.Println(err)
		return
	}
	err = db.Create(m).Error
	if err != nil {
		helpers.LogError(err)
	}
	return
}

func UpdateUser(db *gorm.DB, m *User) (err error) {
	m.encodePassword()
	validate := validator.New()
	err = validate.Struct(m)
	if err != nil {
		helpers.LogError(err)
		return
	}
	err = db.Save(m).Error
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

func FindUserByResetPasswordToken(db *gorm.DB, token string) (*User, error) {
	m := User{}
	err := db.
		Where("reset_password_token = ?", token).
		Where("reset_password_token_expire_at > ?", time.Now().Unix()).
		First(&m).Error
	if err != nil {
		helpers.LogError(err)
	}
	return &m, err
}
