package models

import (
	"context"
	"database/sql"
	"github.com/joho/godotenv"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/helpers"
	"github.com/vavilen84/nft-project/store"
	"log"
	"os"
)

func beforeTestRun() {
	setTestAppEnv()
	err := godotenv.Load("./../.env.development")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	store.InitTestDB()
}

func setTestAppEnv() {
	err := os.Setenv(constants.AppEnvEnvVar, constants.TestingAppEnv)
	if err != nil {
		helpers.LogError(err)
	}
}

/**
 * ! IMPORTANT - dont use for production DB !
 */
func prepareTestDB(ctx context.Context, conn *sql.Conn) {
	dropAllTablesFromTestDB(ctx, conn)
	err := CreateMigrationsTableIfNotExists(ctx, conn)
	if err != nil {
		log.Println(err)
	}

	err = MigrateUp(ctx, conn)
	if err != nil {
		log.Println(err)
	}

	LoadFixtures(ctx, conn)
	return
}

/**
 * ! IMPORTANT - dont use for production DB !
 */
func dropAllTablesFromTestDB(ctx context.Context, conn *sql.Conn) {
	tables := []string{
		constants.MigrationDBTable,
		constants.ProductDBTable,
	}
	for i := 0; i < len(tables); i++ {
		_, err := conn.ExecContext(ctx, "DROP TABLE IF EXISTS "+tables[i])
		if err != nil {
			log.Println(err)
		}
	}
}
