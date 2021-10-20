package main

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/Erikaa81/Banco-api/controllers/exit"
	"github.com/Erikaa81/Banco-api/controllers/logger"
	"github.com/Erikaa81/Banco-api/controllers/server"
	"github.com/Erikaa81/Banco-api/routes"

	"github.com/Erikaa81/Banco-api/app"
)

var api *app.App

func HandlerFunc(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

		s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		if len(s) != 2 {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}

		b, err := base64.StdEncoding.DecodeString(s[1])
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		pair := strings.SplitN(string(b), ":", 2)
		if len(pair) != 2 {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}

		if pair[0] != "username" || pair[1] != "password" {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}

		h.ServeHTTP(w, r)
	}
}
	func main() {

		srv := server.
			GetServer().
			WithRouter(routes.GetRouter(api)).
			WithLogger(logger.Error)

		go func() {
			api.Log.Info("Iniciando servidor na porta ", api.Cfg.GetAPIPort())
			if err := srv.StartServer(); err != nil {
				api.Log.Fatal(err.Error())
			}

			exit.Init(func() {
				if err := srv.CloseServer(); err != nil {
					api.Log.Error(err.Error())
				}

				if err := api.DB.CloseDB(); err != nil {
					api.Log.Error(err.Error())
				}
			})
		}()
	}