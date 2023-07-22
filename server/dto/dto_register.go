package dto

import (
	"github.com/go-playground/validator/v10"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/interfaces"
	"github.com/vavilen84/nft-project/validation"
)

type SignUp struct {
	interfaces.Model
	Nickname    string `json:"lastName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	BillingPlan int    `json:"billing_plan"`
	Role        int    `json:"role"`
}

func (*SignUp) GetValidationRules() interface{} {
	return validation.ScenarioRules{
		constants.ScenarioSignUp: validation.FieldRules{
			"Nickname":    "min=3,max=255,required",
			"Email":       "min=3,max=255,required,email",
			"Password":    "min=6,max=255,required",
			"BillingPlan": "lt=4,required",
			"Role":        "lt=4,required",
		},
	}
}

func (*SignUp) GetValidator() interface{} {
	v := validator.New()
	return v
}
