package usecases

import "github.com/Erikaa81/Banco-api/store"

type Accounts struct {
	store store.AccountStore
}

func NewAccounts(store store.AccountStore) Accounts {
	return Accounts{
		store: store,
	}
}
