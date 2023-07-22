package interfaces

type Model interface {
	GetTableName() string
	GetValidator() interface{}
	GetValidationRules() interface{}
}
