package models

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/validation"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"testing"
	"time"
)

func Test_DropInsert_notOk_1(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}), &gorm.Config{})

	// should be error
	// Calculate the duration of 24 hours
	duration, err := time.ParseDuration("-24h")
	if err != nil {
		fmt.Println("Error parsing duration:", err)
		return
	}
	currentTime := time.Now()
	// Subtract 24 hours from the current time
	pastTime := currentTime.Add(duration)

	m := Drop{
		CollectionName:     "",
		Blockchain:         0,
		PublicSaleDateTime: pastTime,
		TimeZone:           "",
		PublicSalePrice:    0,
		TotalSupply:        0,
		BillingPlan:        0,
		Status:             UnPublishedDropStatus,
		UserID:             0,
	}
	err = InsertDrop(gormDB, &m)
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "Drop", "CollectionName"), v["CollectionName"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "Drop", "Blockchain"), v["Blockchain"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.FutureErrorMsg, "Drop", "PublicSaleDateTime"), v["PublicSaleDateTime"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "Drop", "TimeZone"), v["TimeZone"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "Drop", "PublicSalePrice"), v["PublicSalePrice"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "Drop", "TotalSupply"), v["TotalSupply"][0].Message)
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "Drop", "BillingPlan"), v["BillingPlan"][0].Message)

	//// should be error
	//m = User{
	//	BillingPlan: 4,
	//}
	//err = InsertUser(gormDB, &m)
	//v, ok = err.(validation.Errors)
	//if !ok {
	//	log.Fatalln("can not assert validation.Errors")
	//}
	//assert.Equal(t, fmt.Sprintf(constants.LowerThanTagErrorMsg, "User", "4"), v["BillingPlan"][0].Message)
	//
	//m.BillingPlan = constants.ProBillingPlan
	//err = InsertUser(gormDB, &m)
	//v, ok = err.(validation.Errors)
	//if !ok {
	//	log.Fatalln("can not assert validation.Errors")
	//}
	//
	//// no error
	//_, ok = v["BillingPlan"]
	//assert.False(t, ok)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
