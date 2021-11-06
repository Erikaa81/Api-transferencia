package store

import (
	"context"
	"time"

	"github.com/Erikaa81/Banco-api/domain/usecases"
)

type (
	// FindAllAccountUseCase input port
	FindAllAccountUseCase interface {
		Execute(context.Context) ([]FindAllAccountOutput, error)
	}

	// FindAllAccountPresenter output port
	FindAllAccountPresenter interface {
		Output([]usecases.Account) []FindAllAccountOutput
	}

	// FindAllAccountOutput outputData
	FindAllAccountOutput struct {
		ID        string  `json:"id"`
		Name      string  `json:"name"`
		CPF       string  `json:"cpf"`
		Balance   float64 `json:"balance"`
		CreatedAt string  `json:"created_at"`
	}

	findAllAccountInteractor struct {
		repo       usecases.AccountRepository
		presenter  FindAllAccountPresenter
		ctxTimeout time.Duration
	}
)

// NewFindAllAccountInteractor creates new findAllAccountInteractor with its dependencies
func NewFindAllAccountInteractor(
	repo usecases.AccountRepository,
	presenter FindAllAccountPresenter,
	t time.Duration,
) FindAllAccountUseCase {
	return findAllAccountInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

// Execute orchestrates the use case
func (a findAllAccountInteractor) Execute(ctx context.Context) ([]FindAllAccountOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, a.ctxTimeout)
	defer cancel()

	accounts, err := a.repo.FindAll(ctx)
	if err != nil {
		return a.presenter.Output([]usecases.Account{}), err
	}

	return a.presenter.Output(accounts), nil
}
