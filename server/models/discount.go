package models

//
//import (
//	"context"
//	"orm/sql"
//	"github.com/vavilen84/nft-project/constants"
//	"github.com/vavilen84/nft-project/orm"
//	"github.com/vavilen84/nft-project/validation"
//	"gopkg.in/go-playground/validator.v9"
//	"log"
//)
//
//type Discount struct {
//	Id         uint32 `json:"id" column:"id"`
//	Title      string `json:"title" column:"title"`
//	Amount     uint32 `json:"amount" column:"amount"`
//	Percentage uint8  `json:"percentage" column:"percentage"`
//	Type       uint8  `json:"type" column:"type"`
//
//	CreatedAt int64 `json:"created_at" column:"created_at"`
//	UpdatedAt int64 `json:"updated_at" column:"updated_at"`
//	DeletedAt int64 `json:"deleted_at" column:"deleted_at"`
//}
//
//func (m Discount) GetId() uint32 {
//	return m.Id
//}
//
//func (Discount) GetTableName() string {
//	return constants.TaxDBTable
//}
//
//func (Discount) getValidationRules() validation.ScenarioRules {
//	return validation.ScenarioRules{
//		constants.ScenarioCreate: validation.FieldRules{
//			constants.DiscountTitleField:      "required",
//			constants.DiscountAmountField:     "min=1",
//			constants.DiscountPercentageField: "min=1,max=100",
//			constants.DiscountTypeField:       "required",
//		},
//	}
//}
//
//func (Discount) getValidator() (v *validator.Validate) {
//	v = validator.New()
//	return
//}
//
//func (m Discount) Create(ctx context.Context, conn *sql.Conn) (err error) {
//	err = validation.ValidateByScenario(constants.ScenarioCreate, m, m.getValidator(), m.getValidationRules())
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	err = orm.Insert(ctx, conn, m)
//	return
//}
