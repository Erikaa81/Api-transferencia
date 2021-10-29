package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/Erikaa81/Banco-api/app"
	"github.com/Erikaa81/Banco-api/controllers/secret"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Account modelo para conta do usuário
type Accounts struct {
	gorm.Model `json:"-"`
	ID         uuid.UUID                 `gorm:"type:uuid" json:"id"`
	Name       string                    `json:"name"`
	Cpf        string                    `json:"cpf"`
	Secret     string                    `json:"secret"`
	Balance    float64                   `json:"balance" validate:"required"`
	CreatedAt  time.Time                 `json:"created_at"`
	Transfers  interface{ fmt.Stringer } `json:"-" gorm:"foreignKey:AccountOriginID"`
}

// CreateAccount cria uma conta de usuário
func (a *Accounts) CreateAccount(app *app.App) (*Accounts, error) {

	account := &Accounts{
		ID:        a.ID,
		Name:      a.Name,
		Cpf:       a.Cpf,
		Secret:    a.Secret,
		Balance:   a.Balance,
		CreatedAt: a.CreatedAt,
		Transfers: a.Transfers,
	}

	result := app.DB.Client.Create(account)

	if result.Error != nil {
		erroCriarConta := errors.New("Erro na criação da conta")

		return nil, erroCriarConta
	}

	return account, nil

}

func (a *Accounts) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New()
	a.Secret, err = secret.HashPassword(a.Secret)
	if err != nil {
		ErroConversãoSenha := errors.New("Erro ao criptografar senha")
		return ErroConversãoSenha
	}
	return
}
