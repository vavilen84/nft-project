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
//type OrderDiscount struct {
//	Id         uint32 `json:"id" column:"id"`
//	OrderId    uint32 `json:"order_id" column:"order_id"`
//	DiscountId uint32 `json:"discount_id" column:"discount_id"`
//}
//
//func (m OrderDiscount) GetId() uint32 {
//	return m.Id
//}
//
//func (OrderDiscount) GetTableName() string {
//	return constants.OrderDiscountDBTable
//}
//
//func (OrderDiscount) getValidationRules() validation.ScenarioRules {
//	return validation.ScenarioRules{
//		constants.ScenarioCreate: validation.FieldRules{
//			constants.OrderDiscountOrderIdField:    "required",
//			constants.OrderDiscountDiscountIdField: "required",
//		},
//	}
//}
//
//func (OrderDiscount) getValidator() (v *validator.Validate) {
//	v = validator.New()
//	return
//}
//
//func (m OrderDiscount) Create(ctx context.Context, conn *sql.Conn) (err error) {
//	err = validation.ValidateByScenario(constants.ScenarioCreate, m, m.getValidator(), m.getValidationRules())
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	err = orm.Insert(ctx, conn, m)
//	return
//}
