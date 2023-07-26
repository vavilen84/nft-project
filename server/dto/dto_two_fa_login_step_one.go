package dto

type TwoFaLoginStepOne struct {
	Email string `json:"email" validate:"min=3,max=255,email,required"`
}
