package interfaces

import (
	"github.com/go-playground/validator/v10"
	"github.com/vavilen84/nft-project/validation"
)

type Model interface {
	GetValidator() *validator.Validate
	GetValidationRules() validation.ScenarioRules
}
