package dto

import (
	"github.com/go-playground/validator/v10"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/validation"
	"time"
)

type Drop struct {
	CollectionName     string    `json:"collection_name"  `
	Blockchain         int       `json:"blockchain"`
	WebsiteURL         string    `json:"website_url"`
	TwitterURL         string    `json:"twitter_url"`
	DiscordURL         string    `json:"discord_url"`
	PublicSaleDateTime time.Time `json:"public_sale_date"`
	TimeZone           string    `json:"time_zone"`
	PublicSalePrice    float64   `json:"public_sale_price"`
	TotalSupply        int       `json:"total_supply"`
	BillingPlan        int       `json:"billing_plan"`
}

func (Drop) GetValidator() interface{} {
	v := validator.New()
	return v
}

func (Drop) GetValidationRules() interface{} {
	return validation.ScenarioRules{
		constants.ScenarioCreate: validation.FieldRules{
			"CollectionName":     "min=3,max=255,required",
			"Blockchain":         "required",
			"PublicSaleDateTime": "required",
			"TimeZone":           "required",
			"PublicSalePrice":    "required",
			"TotalSupply":        "required",
			"BillingPlan":        "required,gt=0,lt=3",
		},
	}
}
