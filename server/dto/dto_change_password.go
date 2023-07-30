package dto

import (
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/validation"
)

type ChangePassword struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func (ChangePassword) GetValidationRules() interface{} {
	return validation.ScenarioRules{
		constants.ScenarioChangePassword: validation.FieldRules{
			"OldPassword": "min=8,max=255,required,customPasswordValidator",
			"NewPassword": "min=8,max=255,required,customPasswordValidator",
		},
	}
}
