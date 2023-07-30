package dto

type TwoFaLoginSecondStep struct {
	EmailTwoFaCode string `json:"email_to_fa_code" validate:"min=6,max=6,required"`
}
