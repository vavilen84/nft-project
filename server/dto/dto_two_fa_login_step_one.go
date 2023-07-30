package dto

import (
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/validation"
)

type TwoFaLoginStepOne struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (TwoFaLoginStepOne) GetValidationRules() interface{} {
	return validation.ScenarioRules{
		constants.ScenarioTwoFaLoginStepOne: validation.FieldRules{
			"Email":    "max=255,email,required",
			"Password": "min=8,max=255,required,customPasswordValidator",
		},
	}
}
