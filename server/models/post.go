package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/vavilen84/nft-project/helpers"
	"gorm.io/gorm"
	"time"
)

type Post struct {
	Id          int       `json:"id" column:"id" gorm:"primaryKey;autoIncrement:true"`
	Description string    `json:"description" column:"description" validate:"required,min=2,max=5000"`
	CreatedAt   time.Time `json:"created_at" column:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" column:"updated_at" gorm:"autoUpdateTime"`
}

func (Post) TableName() string {
	return "posts"
}

func GetAllPosts(db *gorm.DB) ([]*Post, error) {
	res := []*Post{}
	err := db.Find(&res).Error
	if err != nil {
		helpers.LogError(err)
	}
	return res, err
}

func GetOnePostByID(db *gorm.DB, id int) (*Post, error) {
	m := Post{}
	err := db.First(&m, id).Error
	if err != nil {
		helpers.LogError(err)
	}
	return &m, err
}

func CreatePost(db *gorm.DB, m *Post) (err error) {
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

func UpdatePost(db *gorm.DB, m *Post) (err error) {
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

func DeletePost(db *gorm.DB, m *Post) (err error) {
	validate := validator.New()
	err = validate.Struct(m)
	if err != nil {
		helpers.LogError(err)
		return
	}
	err = db.Delete(m).Error
	if err != nil {
		helpers.LogError(err)
	}
	return
}
