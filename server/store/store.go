package store

import (
	"context"
	"database/sql"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/helpers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"time"
)

func createDbIfNotExists(ctx context.Context, conn *sql.Conn, dbName string) (err error) {
	_, err = conn.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbName)
	if err != nil {
		helpers.LogError(err)
		return err
	}
	return nil
}

func GetDefaultDBContext() context.Context {
	parentCtx := context.Background()
	ctx, _ := context.WithTimeout(parentCtx, constants.DefaultStoreTimeout)
	return ctx
}

func processInitDb(sqlServerDsn, mysqlDbName, DbDsn string) (db *gorm.DB) {
	sqlDriver := os.Getenv("SQL_DRIVER")
	// use credentials without db in order to create db
	sqlDb, err := sql.Open(sqlDriver, sqlServerDsn)
	if err != nil {
		panic("failed to connect sql server: " + err.Error())
	}
	ctx := GetDefaultDBContext()
	conn, err := sqlDb.Conn(ctx)
	if err != nil {
		helpers.LogError(err)
	}
	defer conn.Close()
	err = createDbIfNotExists(ctx, conn, mysqlDbName)
	if err != nil {
		panic("failed to create test db: " + err.Error())
	}

	db, err = gorm.Open(mysql.Open(DbDsn), &gorm.Config{})
	if err != nil {
		panic("failed to database: " + err.Error())
	}
	db.Debug()
	sqlDB, err := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
	return
}
