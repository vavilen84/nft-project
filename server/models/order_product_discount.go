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
//type OrderProductDiscount struct {
//	Id             uint32 `json:"id" column:"id"`
//	OrderProductId uint32 `json:"order_product_id" column:"order_product_id"`
//	DiscountId     uint32 `json:"discount_id" column:"discount_id"`
//}
//
//func (m OrderProductDiscount) GetId() uint32 {
//	return m.Id
//}
//
//func (OrderProductDiscount) GetTableName() string {
//	return constants.OrderProductDiscountDBTable
//}
//
//func (OrderProductDiscount) getValidationRules() validation.ScenarioRules {
//	return validation.ScenarioRules{
//		constants.ScenarioCreate: validation.FieldRules{
//			constants.OrderProductDiscountOrderProductIdField: "required",
//			constants.OrderProductDiscountDiscountIdField:     "required",
//		},
//	}
//}
//
//func (OrderProductDiscount) getValidator() (v *validator.Validate) {
//	v = validator.New()
//	return
//}
//
//func (m OrderProductDiscount) Create(ctx context.Context, conn *sql.Conn) (err error) {
//	err = validation.ValidateByScenario(constants.ScenarioCreate, m, m.getValidator(), m.getValidationRules())
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	err = orm.Insert(ctx, conn, m)
//	return
//}
