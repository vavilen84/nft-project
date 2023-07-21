package models

//
//import (
//	"github.com/stretchr/testify/assert"
//	"github.com/vavilen84/nft-project/constants"
//	"github.com/vavilen84/nft-project/validation"
//	"testing"
//)
//
//func TestOrder_ValidateOnCreate(t *testing.T) {
//	m := Order{}
//	err := validation.ValidateByScenario(constants.ScenarioCreate, &m, m.getValidator(), m.getValidationRules())
//	assert.NotNil(t, err)
//	assert.NotEmpty(t, err[constants.OrderCustomerIdField])
//
//	m = Order{
//		CustomerId: 1,
//	}
//	err = validation.ValidateByScenario(constants.ScenarioCreate, &m, m.getValidator(), m.getValidationRules())
//	assert.NotNil(t, err)
//}
