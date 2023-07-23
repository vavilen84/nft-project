package models

import (
	"github.com/anaskhan96/go-password-encoder"
	"github.com/go-playground/validator/v10"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/helpers"
	"github.com/vavilen84/nft-project/validation"
	"gorm.io/gorm"
	"regexp"
	"time"
	"unicode/utf8"
)

const alphaNumericRegexString = "^[a-zA-Z0-9]+$"

type User struct {
	Id                         int        `json:"id" column:"id" gorm:"primaryKey;autoIncrement:true"`
	Email                      string     `json:"email" column:"email" validate:"min=3,max=255,email,required"`
	Nickname                   string     `json:"nickname" column:"nickname" validate:"min=3,max=255,required"`
	Password                   string     `json:"password" column:"password" validate:"min=8,max=5000,required,customPasswordValidator"`
	PasswordSalt               string     `column:"password_salt" validate:"min=3,max=5000,required"`
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

	// password must have should have at least 1 small letter, 1 uppercase letter,
	// 1 digit, and at least 1 symbol from the set +-_()*^%$#@!
	regex := regexp.MustCompile(alphaNumericRegexString)

	return regex.MatchString(p)
}

func (m *User) TableName() string {
	return "user"
}

func (User) getValidationRules() validation.ScenarioRules {
	return validation.ScenarioRules{
		constants.ScenarioCreate: validation.FieldRules{
			constants.CustomerEmailField:     "required,min=1,max=255,email",
			constants.CustomerFirstNameField: "required,min=1,max=255",
			constants.CustomerLastNameField:  "required,min=1,max=255",
		},
	}
}

func (User) getValidator() (v *validator.Validate) {
	v = validator.New()
	return
}

func InsertUser(db *gorm.DB, m *User) (err error) {
	m.encodePassword()
	validate := validator.New()
	err = validate.RegisterValidation("customPasswordValidator", CustomPasswordValidator)
	if err != nil {
		helpers.LogError(err)
		return
	}
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
