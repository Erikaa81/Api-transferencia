package routes

import (
	"net/http"

	"github.com/Erikaa81/Banco-api/controllers"

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
	Banco controllers.Accounts
	http.Handler
}

func NewServer(controllers controllers.Accounts) Server {
	server := Server{
		Banco:   controllers,
		Handler: nil,
	}
	router := mux.NewRouter()

	router.HandleFunc("/", controllers.Accounts)
	router.HandleFunc("/accounts", controllers.Accounts).Methods("GET", "POST")
	router.HandleFunc("accounts/{id}/balance", controllers.Accounts).Methods("GET")
	router.HandleFunc("/login", controllers.login).Methods("POST")
	router.HandleFunc("/transfers", controllers.Transfer).Queries("GET", "POST")

	server.Handler = router
	return server
}
