package dto

import (
	"github.com/go-playground/validator/v10"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/interfaces"
	"github.com/vavilen84/nft-project/validation"
)

type Login struct {
	interfaces.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (*Login) GetValidationRules() interface{} {
	return validation.ScenarioRules{
		constants.ScenarioSignIn: validation.FieldRules{
			"Email":    "max=255,email,required",
			"Password": "max=255,required",
		},
	}
}

func (*Login) GetValidator() interface{} {
	v := validator.New()
	return v
}
