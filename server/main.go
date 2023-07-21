package main

import (
	"github.com/vavilen84/nft-project/handlers"
	"github.com/vavilen84/nft-project/store"
	"log"
)

func main() {
	store.InitDB()
	handler := handlers.MakeHandler()
	httpServer := handlers.InitHttpServer(handler)
	log.Fatal(httpServer.ListenAndServe())
}
