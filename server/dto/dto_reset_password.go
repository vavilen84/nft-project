package dto

import (
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/validation"
)

type ResetPassword struct {
	Token       string `json:"token"`
	NewPassword string `json:"new_password"`
}

func (ResetPassword) GetValidationRules() interface{} {
	return validation.ScenarioRules{
		constants.ScenarioResetPassword: validation.FieldRules{
			"Token":       "min=6,max=255,required",
			"NewPassword": "min=8,max=255,required,customPasswordValidator",
		},
	}
}
