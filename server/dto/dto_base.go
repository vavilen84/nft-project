package dto

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/vavilen84/nft-project/helpers"
	"regexp"
	"unicode/utf8"
)

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Error  string      `json:"error"`
	Errors []string    `json:"errors"`

	FormErrors map[string][]string `json:"formErrors"`
}

type ResponseData map[string]interface{}

const alphaNumericRegexString = "^[a-zA-Z0-9]+$"

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

func GetValidator() *validator.Validate {
	v := validator.New()
	err := v.RegisterValidation("customPasswordValidator", CustomPasswordValidator)
	if err != nil {
		helpers.LogError(err)
		return nil
	}
	return v
}
