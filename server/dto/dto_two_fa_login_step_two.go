package dto

import (
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/validation"
)

type TwoFaLoginStepTwo struct {
	EmailTwoFaCode string `json:"email_to_fa_code" validate:"min=6,max=6,required"`
}

func (TwoFaLoginStepTwo) GetValidationRules() interface{} {
	return validation.ScenarioRules{
		constants.ScenarioTwoFaLoginStepTwo: validation.FieldRules{
			"EmailTwoFaCode": "min=6,max=6,required",
		},
	}
}
