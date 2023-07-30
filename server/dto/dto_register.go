package dto

import (
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/validation"
)

type Register struct {
	Nickname    string `json:"nickname"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	BillingPlan int    `json:"billing_plan"`
}

func (Register) GetValidationRules() interface{} {
	return validation.ScenarioRules{
		constants.ScenarioRegister: validation.FieldRules{
			"Nickname":    "min=3,max=255,required",
			"Email":       "min=3,max=255,email,required",
			"Password":    "min=8,max=255,required,customPasswordValidator",
			"BillingPlan": "lt=4,required",
		},
	}
}
