package models

type Blockchain int

type DropStatus int

type BillingPlan int

const (
	_ Blockchain = iota
	Ethereum
	Polygon
	Solana
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
