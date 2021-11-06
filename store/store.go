package store

import (
	"errors"

	"github.com/Erikaa81/Banco-api/domain/usecases"
)

var (
	ErrEmptyID = errors.New("ID cannot be empty")
)

type AccountStore struct {
	accountStore map[string]usecases.Account
}

func NewAccountStore() AccountStore {
	ns := make(map[string]usecases.Account)
	return AccountStore{
		accountStore: ns,
	}
}
