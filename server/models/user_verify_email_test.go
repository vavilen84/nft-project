package models

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestUser_SetUserEmailVerified(t *testing.T) {
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

	mock.ExpectExec("UPDATE user").
		WillReturnResult(sqlmock.NewResult(1, 1))

	m := User{}
	err = SetUserEmailVerified(gormDB, &m)
	assert.Nil(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
