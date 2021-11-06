package http

import (
	"net/http"

	"github.com/Erikaa81/Banco-api/domain/usecases"
	"github.com/gorilla/mux"
)

const (
	ContentType     = "Content-Type"
	JSONContentType = "application/json"
	DateLayout      = "2006-01-02T15:04:05Z"
)

type Error struct {
	Reason string `json:"reason"`
}

type Server struct {
	account usecases.Account
	http.Handler
}

func NewServer(usecases usecases.Account) Server {
	server := Server{account: usecases}

	router := mux.NewRouter()

	router.HandleFunc("/accounts", server.ListAccounts).Methods(http.MethodGet)
	router.HandleFunc("/accounts", server.ListAccounts).Methods(http.MethodPost)
	router.HandleFunc("accounts/{id}/balance", server.ListAccounts).Methods(http.MethodGet)
	router.HandleFunc("/login", server.HandlerLogin).Methods(http.MethodPost)
	router.HandleFunc("/transfers", server.HandlerLogin).Queries(http.MethodGet, http.MethodPost)

	server.Handler = router
	return server
}
