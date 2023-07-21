package dto

import (
	"github.com/go-playground/validator/v10"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/interfaces"
	"github.com/vavilen84/nft-project/validation"
)

type SignUp struct {
	interfaces.Model
	FirstName string `json:"firstName" validate:""`
	LastName  string `json:"lastName" validate:"min=3,max=255,required"`
	Email     string `json:"email" validate:"min=3,max=255,email,required"`
	Password  string `json:"password" validate:"min=3,max=255,required"`
	Birthday  string `json:"birthday" validate:"min=3,max=255,required"`
	Gender    int    `json:"gender" validate:"lt=4,required"`
	Timezone  string `json:"timezone" validate:"min=3,max=255,required"`
	Role      int    `json:"role" validate:"lt=4,required"`
}

func (*SignUp) GetValidationRules() interface{} {
	return validation.ScenarioRules{
		constants.ScenarioSignUp: validation.FieldRules{
			"FirstName": "max=255,required",
			"LastName":  "max=255,required",
			"Email":     "max=255,required,email",
			"Password":  "max=255,required",
			"Birthday":  "max=255,required",
			"Gender":    "lt=4,required",
			"Timezone":  "max=255,required",
			"Role":      "lt=4,required",
		},
	}
}

func (*SignUp) GetValidator() interface{} {
	v := validator.New()
	return v
}
