package dto

type Login struct {
	Email    string `json:"email" validation:"max=255,email,required"`
	Password string `json:"password" validation:"max=255,required"`
}
