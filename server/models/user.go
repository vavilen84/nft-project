package models

import (
	"github.com/anaskhan96/go-password-encoder"
	"github.com/biter777/countries"
	"github.com/go-playground/validator/v10"
	"github.com/vavilen84/nft-project/helpers"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        int    `json:"id" column:"id" gorm:"primaryKey;autoIncrement:true"`
	FirstName string `json:"first_name" column:"first_name" validate:"min=3,max=255,required"`
	LastName  string `json:"last_name" column:"last_name" validate:"min=3,max=255,required"`
	Email     string `json:"email" column:"email" validate:"min=3,max=255,email,required"`
	Twitter   string `json:"twitter" column:"twitter" validate:"max=255"`
	Facebook  string `json:"facebook" column:"facebook" validate:"max=255"`
	Skype     string `json:"skype" column:"skype" validate:"max=255"`
	Telegram  string `json:"telegram" column:"telegram" validate:"max=255"`
	Instagram string `json:"instagram" column:"instagram" validate:"max=255"`
	Password  string `column:"password" validate:"min=3,max=5000,required"`
	Avatar    string `column:"avatar" validate:"max=255"`
	Timezone  string `column:"timezone" validate:"max=255,required"`

	PasswordSalt      string `column:"password_salt" validate:"max=5000,required"`
	PasswordResetHash string `column:"password_reset_hash"`
	TextPresentation  string `column:"text_presentation" validate:"max=5000"`
	VideoPresentation string `column:"video_presentation" validate:"max=255"`

	Role   int `json:"role" column:"role" validate:"lt=4,required"`
	Gender int `validate:"lt=4,required"`

	Birthday time.Time `validate:"required"`

	CountryOfBirth countries.CountryCode `column:"country_of_birth"`
	CurrentCountry countries.CountryCode `column:"current_country"`

	CreatedAt time.Time `json:"created_at" column:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" column:"updated_at" gorm:"autoUpdateTime"`
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
