package http

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// Credentials struct para armazenar o cpf e secret no corpo do request
type Credentials struct {
	CPF    string `json:"cpf" validate:"required"`
	Secret string `json:"secret" validate:"required"`
}

// Claims struct que será criptografado em um token JWT
type Claims struct {
	CPF string `json:"cpf" validate:"required"`
	jwt.StandardClaims
}

func (s Server) HandlerLogin(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

	if len(auth) != 2 || auth[0] != "Basic" {
		http.Error(w, "autorização falhou", http.StatusUnauthorized)
		return
	}

	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)

	if len(pair) != 2 || !validate(pair[0], pair[1]) {
		http.Error(w, "autorização falhou", http.StatusUnauthorized)
		return
	}

	pass(w, r)
}

func pass(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func validate(username, password string) bool {
	if username == "test" && password == "test" {
		return true
	}
	return false
}
