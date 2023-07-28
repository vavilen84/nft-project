package dto

type ChangePassword struct {
	OldPassword string `json:"old_password" validate:"min=3,max=255,required,customPasswordValidator"`
	NewPassword string `json:"new_password" validate:"min=3,max=255,required,customPasswordValidator"`
}
