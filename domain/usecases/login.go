package usecases

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/Erikaa81/Banco-api/app"
)

func HandlerLogin(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
