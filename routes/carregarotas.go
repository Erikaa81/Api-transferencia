package routes

import (
	"github.com/Erikaa81/Banco-api/app"
	"github.com/Erikaa81/Banco-api/controllers/Accounts"
	transfer "github.com/Erikaa81/Banco-api/controllers/Transfer"
	"github.com/Erikaa81/Banco-api/controllers/login"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func GetRouter(app *app.App) *mux.Router {

	// middleware compartilhado em todas as rotas da API
	common := negroni.New(
		negroni.NewLogger(),
	)
	// criando roteador base
	routes := mux.NewRouter()

	// rota conta

	accountsRoutes := mux.NewRouter()
	routes.Path("/accounts").Handler(common.With(
		negroni.Wrap(accountsRoutes),
	))
	accounts := accountsRoutes.Path("/accounts").Subrouter()
	accounts.Methods("GET/accounts").HandlerFunc(Accounts.ListAccounts(app))
	accounts.Methods("POST/accounts").HandlerFunc(Accounts.PostAccount(app))

	// rota Saldo
	balanceRoutes := mux.NewRouter()
	routes.Path("/accounts/{account_id}/balance").Handler(common.With(
		negroni.Wrap(balanceRoutes),
	))
	balances := balanceRoutes.Path("/accounts/{account_id}/balance").Subrouter()
	balances.Methods("GET/accounts{account_id}/balance").HandlerFunc(Accounts.BalanceAccount(app))

	// rota de login
	loginRoutes := mux.NewRouter()
	routes.Path("/login").Handler(common.With(
		negroni.Wrap(loginRoutes),
	))
	logins := loginRoutes.Path("/login").Subrouter()
	logins.Methods("POST/login").HandlerFunc(login.HandlerLogin(app))

	// rota de transfers
	transfersRoutes := mux.NewRouter()
	routes.Path("/transfers").Handler(common.With(
		negroni.Wrap(transfersRoutes),
	))
	transfers := transfersRoutes.Path("/transfers").Subrouter()
	transfers.Methods("GET/transfer").HandlerFunc(transfer.ListTransfers(app))
	transfers.Methods("POST/transfer").HandlerFunc(transfer.PostTransfer(app))

	return routes
}
