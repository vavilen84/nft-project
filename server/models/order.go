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
//type Order struct {
//	Id         uint32 `json:"id" column:"id"`
//	CustomerId uint32 `json:"customer_id" column:"customer_id"`
//
//	CreatedAt int64 `json:"created_at" column:"created_at"`
//}
//
//func (m Order) GetId() uint32 {
//	return m.Id
//}
//
//func (Order) GetTableName() string {
//	return constants.OrderDBTable
//}
//
//func (Order) getValidationRules() validation.ScenarioRules {
//	return validation.ScenarioRules{
//		constants.ScenarioCreate: validation.FieldRules{
//			constants.OrderCustomerIdField: "required",
//		},
//	}
//}
//
//func (Order) getValidator() (v *validator.Validate) {
//	v = validator.New()
//	return
//}
//
//func (m Order) Create(ctx context.Context, conn *sql.Conn) (err error) {
//	err = validation.ValidateByScenario(constants.ScenarioCreate, m, m.getValidator(), m.getValidationRules())
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	err = orm.Insert(ctx, conn, m)
//	return
//}
