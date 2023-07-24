package models

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestUser_FindUserById(t *testing.T) {
	customMatcher := CustomMatcher{}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(customMatcher))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}), &gorm.Config{})

	columns := []string{"id"}
	mock.ExpectQuery("SELECT * FROM `user`").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows(columns).FromCSVString("1"))

	_, err = FindUserById(gormDB, 1)
	assert.Nil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

//func TestUser_FindUserBy2FAToken(t *testing.T) {
//	customMatcher := CustomMatcher{}
//	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(customMatcher))
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//	gormDB, err := gorm.Open(mysql.New(mysql.Config{
//		SkipInitializeWithVersion: true,
//		Conn:                      db,
//	}), &gorm.Config{})
//
//	columns := []string{"email_2fa_code"}
//	mock.ExpectQuery("SELECT * FROM `user`").
//		WithArgs("token").
//		WillReturnRows(sqlmock.NewRows(columns).FromCSVString("token"))
//
//	_, err = FindUserBy2FAToken(gormDB, "token")
//	assert.Nil(t, err)
//
//	if err := mock.ExpectationsWereMet(); err != nil {
//		t.Errorf("there were unfulfilled expectations: %s", err)
//	}
//}

func Test_FindUserByEmail(t *testing.T) {
	customMatcher := CustomMatcher{}
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(customMatcher))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}), &gorm.Config{})

	columns := []string{"email"}
	mock.ExpectQuery("SELECT * FROM `user`").
		WithArgs("email@example.com").
		WillReturnRows(sqlmock.NewRows(columns).FromCSVString("email@example.com"))

	_, err = FindUserByEmail(gormDB, "email@example.com")
	assert.Nil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

//
//func TestUser_FindUserByResetPasswordToken(t *testing.T) {
//	customMatcher := CustomMatcher{}
//	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(customMatcher))
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//	gormDB, err := gorm.Open(mysql.New(mysql.Config{
//		SkipInitializeWithVersion: true,
//		Conn:                      db,
//	}), &gorm.Config{})
//
//	duration, err := time.ParseDuration("24h")
//	if err != nil {
//		fmt.Println("Error parsing duration:", err)
//		return
//	}
//	currentTime := time.Now()
//	// Add 24 hours to the current time
//	futureTime := currentTime.Add(duration)
//
//	columns := []string{"password_reset_token"}
//	mock.ExpectQuery("SELECT * FROM `user`").
//		WithArgs("token", futureTime).
//		WillReturnRows(sqlmock.NewRows(columns).FromCSVString("token"))
//
//	_, err = FindUserByResetPasswordToken(gormDB, "token")
//	assert.Nil(t, err)
//
//	if err := mock.ExpectationsWereMet(); err != nil {
//		t.Errorf("there were unfulfilled expectations: %s", err)
//	}
//}
