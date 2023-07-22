package dto

type ResetPassword struct {
	Token       string `json:"token" validate:"min=3,max=255,required"`
	NewPassword string `json:"new_password" validate:"min=3,max=255,required"`
}
