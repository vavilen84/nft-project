package dto

type TwoFaLoginFirstStep struct {
	Email    string `json:"email" validation:"max=255,email,required"`
	Password string `json:"password" validate:"min=3,max=255,required,customPasswordValidator"`
}
