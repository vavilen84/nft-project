package models

type Blockchain int

type DropStatus int

type BillingPlan int

const (
	_ Blockchain = iota
	Ethereum
	Polygon
	Solana
	Arbitrum
	Avalanche
	Binance
	Cardano
	Celo
	Cronos
	Fantom
	Flare
	IMX
	Tron
	Wax
	Ziliqa
	Ton
	Aptos
	OtherBlockchain
)

const (
	_ DropStatus = iota
	UnPublished
	InReview
	Published
)

const (
	_ BillingPlan = iota
	Standard
	Featured
)
