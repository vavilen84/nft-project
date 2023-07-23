package validation

import (
	"gopkg.in/go-playground/validator.v9"
)

func Validate(model interface{}) error {
	v := validator.New()
	err := v.Struct(model)
	if err != nil {
		result := make(Errors)
		result.setStructValidationErrors(err, model)
		return result
	}
	return nil
}
