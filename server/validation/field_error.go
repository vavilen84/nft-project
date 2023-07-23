package validation

import (
	"fmt"
	"github.com/vavilen84/nft-project/constants"
)

type FieldError struct {
	Tag     string
	Field   string
	Message string
	Value   string
	Param   string
	Name    string
}

func (s *FieldError) setErrorMessage() {
	switch s.Tag {
	case constants.RequiredTag:
		s.Message = fmt.Sprintf(constants.RequiredErrorMsg, s.Name, s.Field)
	case constants.MinTag:
		s.Message = fmt.Sprintf(constants.MinValueErrorMsg, s.Name, s.Field, s.Param)
	case constants.MaxTag:
		s.Message = fmt.Sprintf(constants.MaxValueErrorMsg, s.Name, s.Field, s.Param)
	case constants.EmailTag:
		s.Message = fmt.Sprintf(constants.EmailErrorMsg, s.Name)
	}
}
