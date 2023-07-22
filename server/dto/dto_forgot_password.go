package dto

type ForgotPassword struct {
	Email string `json:"email" validate:"min=3,max=255,email,required"`
}
