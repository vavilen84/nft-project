package handlers_test

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/vavilen84/nft-project/handlers"
	"github.com/vavilen84/nft-project/store"
	"log"
	"net/http/httptest"
)

func initApp() *httptest.Server {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Can not load .env file")
	}
	store.InitDB()
	db := store.GetDB()
	if err := db.Exec("delete from jwt_info").Error; err != nil {
		fmt.Println("Error deleting entities:", err)
	}
	if err := db.Exec("delete from user").Error; err != nil {
		fmt.Println("Error deleting entities:", err)
	}
	handler := handlers.MakeHandler()
	ts := httptest.NewServer(handler)
	return ts
}
