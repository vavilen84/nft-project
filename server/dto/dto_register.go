package dto

type SignUp struct {
	Nickname    string `json:"nickname" validation:"min=3,max=255,required"`
	Email       string `json:"email" validation:"min=3,max=255,required,email"`
	Password    string `json:"password" validation:"min=6,max=255,required"`
	BillingPlan int    `json:"billing_plan" validation:"lt=4,required"`
}
