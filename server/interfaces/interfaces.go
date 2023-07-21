package interfaces

type Model interface {
	GetTableName() string
	GetId() uint32
	SetId(id uint32) Model
	SetCreatedAt() Model
	SetUpdatedAt() Model
	SetDeletedAt() Model
	GetValidator() interface{}
	GetValidationRules() interface{}
}
