package dto

type ResetPassword struct {
	NewPassword string `column:"password" validate:"min=3,max=255,required"`
}
