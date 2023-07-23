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
)

func TestUser_CreateScenario_emailValidation(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})

	// should be error
	m := User{
		Email: "not_valid_email",
	}
	err = InsertUser(gormDB, &m)
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.EmailErrorMsg, "User"), v["Email"][0].Message)

	// no error
	m.Email = "valid.email@example.com"
	err = InsertUser(gormDB, &m)
	v, ok = err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	_, ok = v["Email"]
	assert.False(t, ok)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUser_CreateScenario_nicknameValidation(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})

	// should be error
	m := User{
		Nickname: "",
	}
	err = InsertUser(gormDB, &m)
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.MinValueErrorMsg, "User", "Nickname", "3"), v["Nickname"][0].Message)

	m.Nickname = "valid.nickname"
	err = InsertUser(gormDB, &m)
	v, ok = err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}

	// no error
	_, ok = v["Nickname"]
	assert.False(t, ok)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUser_CreateScenario_passwordValidation(t *testing.T) {
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
	m := User{
		Password: "1234567",
	}
	err = InsertUser(gormDB, &m)
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.MinValueErrorMsg, "User", "Password", "8"), v["Password"][0].Message)

	// should be error
	m = User{
		Password: "1234567+",
	}
	err = InsertUser(gormDB, &m)
	v, ok = err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.CustomPasswordValidatorTagErrorMsg, "User"), v["Password"][0].Message)

	// no error
	rowValidPassword := "12345678"
	m.Password = rowValidPassword
	err = InsertUser(gormDB, &m)
	v, ok = err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	_, ok = v["Password"]
	assert.False(t, ok)
	//assert.NotEmpty(t, m.PasswordSalt)
	//assert.True(t, password.Verify(rowValidPassword, m.PasswordSalt, m.Password, nil))

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUser_CreateScenario_billingPlanValidation(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})

	// should be error
	m := User{
		BillingPlan: 0,
	}
	err = InsertUser(gormDB, &m)
	v, ok := err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.RequiredErrorMsg, "User", "BillingPlan"), v["BillingPlan"][0].Message)

	// should be error
	m = User{
		BillingPlan: 4,
	}
	err = InsertUser(gormDB, &m)
	v, ok = err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}
	assert.Equal(t, fmt.Sprintf(constants.LowerThanTagErrorMsg, "User", "4"), v["BillingPlan"][0].Message)

	m.BillingPlan = constants.ProBillingPlan
	err = InsertUser(gormDB, &m)
	v, ok = err.(validation.Errors)
	if !ok {
		log.Fatalln("can not assert validation.Errors")
	}

	// no error
	_, ok = v["BillingPlan"]
	assert.False(t, ok)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
