package main

import (
	"log"
	"net/http"

	"github.com/Erikaa81/Banco-api/domain/usecases"
	api "github.com/Erikaa81/Banco-api/http"
	"github.com/Erikaa81/Banco-api/store"
)

const addr = ":1300"

func main() {
	accountStore := store.NewAccountStore()
	accountsUsecases := usecases.NewAccounts(accountStore)
	server := api.NewServer(accountsUsecases)

	log.Printf("starting server on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, server))
}
