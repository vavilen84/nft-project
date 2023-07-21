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
//type Customer struct {
//	Id        uint32 `json:"id" column:"id"`
//	FirstName string `json:"first_name" column:"first_name"`
//	LastName  string `json:"last_name" column:"last_name"`
//	Email     string `json:"email" column:"email"`
//
//	CreatedAt int64 `json:"created_at" column:"created_at"`
//	UpdatedAt int64 `json:"updated_at" column:"updated_at"`
//	DeletedAt int64 `json:"deleted_at" column:"deleted_at"`
//}
//
//func (m Customer) GetId() uint32 {
//	return m.Id
//}
//
//func (Customer) GetTableName() string {
//	return constants.CustomerDBTable
//}
//
//func (Customer) getValidationRules() validation.ScenarioRules {
//	return validation.ScenarioRules{
//		constants.ScenarioCreate: validation.FieldRules{
//			constants.CustomerEmailField:     "required,min=1,max=255,email",
//			constants.CustomerFirstNameField: "required,min=1,max=255",
//			constants.CustomerLastNameField:  "required,min=1,max=255",
//		},
//	}
//}
//
//func (Customer) getValidator() (v *validator.Validate) {
//	v = validator.New()
//	return
//}
//
//func (m Customer) Create(ctx context.Context, conn *sql.Conn) (err error) {
//	err = validation.ValidateByScenario(constants.ScenarioCreate, m, m.getValidator(), m.getValidationRules())
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	err = orm.Insert(ctx, conn, m)
//	return
//}
