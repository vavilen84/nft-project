package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/helpers"
	"github.com/vavilen84/nft-project/validation"
	"gorm.io/gorm"
	"log"
	"time"
)

type Drop struct {
	Id                   int        `json:"id" column:"id" gorm:"primaryKey;autoIncrement:true"`
	CollectionName       string     `json:"collection_name"  `
	Blockchain           Blockchain `json:"blockchain" column:"blockchain"`
	WebsiteURL           string     `json:"website_url" column:"website_url"`
	TwitterURL           string     `json:"twitter_url" column:"twitter_url"`
	DiscordURL           string     `json:"discord_url" column:"discord_url"`
	PublicSaleDateTime   time.Time  `json:"public_sale_date" column:"public_sale_date"`
	TimeZone             string     `json:"time_zone" column:"time_zone"`
	PublicSalePrice      float64    `json:"public_sale_price" column:"public_sale_price"`
	TotalSupply          int        `json:"total_supply" column:"total_supply"`
	UserID               int        `json:"user_id" column:"user_id"`
	BillingPlan          int        `json:"billing_plan" column:"billing_plan"`
	BillingTransactionID string     `json:"billing_transaction_id" column:"billing_transaction_id"`
	Status               DropStatus `json:"status" column:"status"`
	PreviewImg           string     `json:"preview_img" column:"preview_img"`
	BannerImg            string     `json:"banner_img" column:"banner_img"`
}

func (m *Drop) TableName() string {
	return "drop"
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
			"PublicSalePrice":    "required",
			"TotalSupply":        "required",
			"UserID":             "required",
			"BillingPlan":        "required,gt=0,lt=3",
			"Status":             "required,gt=0,lt=3",
		},
	}
}

func InsertDrop(db *gorm.DB, m *Drop) (err error) {
	err = validation.ValidateByScenario(constants.ScenarioCreate, *m)
	if err != nil {
		log.Println(err)
		return
	}
	err = db.Create(m).Error
	if err != nil {
		helpers.LogError(err)
	}
	return
}
