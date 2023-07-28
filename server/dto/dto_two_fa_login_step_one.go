package dto

type TwoFaLoginFirstStep struct {
	Email string `json:"email" validation:"max=255,email,required"`
}
