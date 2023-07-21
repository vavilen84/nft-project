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
//type OrderProductTax struct {
//	Id             uint32 `json:"id" column:"id"`
//	OrderProductId uint32 `json:"order_product_id" column:"order_product_id"`
//	TaxId          uint32 `json:"tax_id" column:"tax_id"`
//}
//
//func (m OrderProductTax) GetId() uint32 {
//	return m.Id
//}
//
//func (OrderProductTax) GetTableName() string {
//	return constants.OrderProductTaxDBTable
//}
//
//func (OrderProductTax) getValidationRules() validation.ScenarioRules {
//	return validation.ScenarioRules{
//		constants.ScenarioCreate: validation.FieldRules{
//			constants.OrderProductTaxOrderProductIdField: "required",
//			constants.OrderProductTaxTaxIdField:          "required",
//		},
//	}
//}
//
//func (OrderProductTax) getValidator() (v *validator.Validate) {
//	v = validator.New()
//	return
//}
//
//func (m OrderProductTax) Create(ctx context.Context, conn *sql.Conn) (err error) {
//	err = validation.ValidateByScenario(constants.ScenarioCreate, m, m.getValidator(), m.getValidationRules())
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	err = orm.Insert(ctx, conn, m)
//	return
//}
