package dto

import (
	"github.com/go-playground/validator/v10"
	"github.com/vavilen84/nft-project/helpers"
	"github.com/vavilen84/nft-project/validation"
)

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Error  string      `json:"error"`
	Errors []string    `json:"errors"`

	FormErrors map[string][]string `json:"formErrors"`
}

type ResponseData map[string]interface{}

func GetValidator() *validator.Validate {
	v := validator.New()
	err := v.RegisterValidation("customPasswordValidator", validation.CustomPasswordValidator)
	if err != nil {
		helpers.LogError(err)
		return nil
	}
	return v
}
