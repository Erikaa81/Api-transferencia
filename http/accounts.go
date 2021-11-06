package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

type AccountResponse struct {
	ID        uuid.UUID                 `gorm:"type:uuid" json:"id"`
	Name      string                    `json:"name"`
	CPF       string                    `json:"cpf"`
	Secret    string                    `json:"secret"`
	Balance   float64                   `json:"balance" validate:"required"`
	CreatedAt time.Time                 `json:"created_at"`
	Transfers interface{ fmt.Stringer } `json:"-" gorm:"foreignKey:AccountOriginID"`
}

func (s Server) ListAccounts(w http.ResponseWriter, r *http.Request) {
	list, err := s.accounts.ListAccounts()
	if err != nil {
		log.Printf("Erro na listagem das contas %s\n", err.Error())
		response := Error{Reason: "internal server error"}
		w.Header().Set(ContentType, JSONContentType)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Build response based on the usecase output
	response := make([]AccountResponse, len(list))
	for i, client := range list {
		response[i].ID = uuid.UUID(client.ID)
		response[i].Name = string(client.Name)
		response[i].CPF = string(client.Cpf)
		response[i].Secret = string(client.secret)
		response[i].Balance = float64(client.balance)
		response[i].CreatedAt = client.CreatedAt.Format(time.Time)
	}

	w.Header().Set(ContentType, JSONContentType)
	json.NewEncoder(w).Encode(response)
	log.Printf("sent list all accounts successful response with %d", len(response))
}

func ListAccounts() {
	panic("unimplemented")
}
