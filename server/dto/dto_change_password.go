package dto

type ChangePassword struct {
	OldPassword string `column:"password" validate:"min=3,max=255,required"`
	NewPassword string `column:"password" validate:"min=3,max=255,required"`
}
