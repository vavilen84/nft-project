package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/helpers"
	"github.com/vavilen84/nft-project/models"
	"github.com/vavilen84/nft-project/store"
	"log"
	"os"
)

func main() {

	if os.Getenv(constants.AppEnvEnvVar) == constants.DevelopmentAppEnv {
		err := godotenv.Load("./.env.development")
		if err != nil {
			log.Fatal("Error loading .env.development file")
		}
	}

	store.InitDB()
	db := store.GetDB()
	ctx := store.GetDefaultDBContext()

	conn, err := db.Conn(ctx)
	if err != nil {
		helpers.LogFatal(err)
	}
	defer conn.Close()

	err = models.CreateMigrationsTableIfNotExists(ctx, conn)
	if err != nil {
		log.Println(err)
	}

	err = models.MigrateUp(ctx, conn)
	if err != nil {
		log.Println(err)
	}
}
