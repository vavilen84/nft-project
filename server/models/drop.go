package models

import "time"

type Drop struct {
	CollectionName     string     `json:"collection_name"`
	Blockchain         Blockchain `json:"blockchain"`
	WebsiteURL         string     `json:"website_url"`
	TwitterURL         string     `json:"twitter_url"`
	DiscordURL         string     `json:"discord_url"`
	PublicSaleDateTime time.Time  `json:"public_sale_date"`
	PublicSalePrice    float64    `json:"public_sale_price"`
	TotalSupply        int        `json:"total_supply"`
	UserID             int        `json:"user_id"`
}
