package login

import (
	"gorm.io/gorm"
	"vodka.app/internal/v1/service/account"
	"vodka.app/internal/v1/service/jwe"
	jweModel "vodka.app/internal/v1/structure/jwe"
	loginsModel "vodka.app/internal/v1/structure/logins"
)

type Resolver interface {
	Web(input *loginsModel.Login) interface{}
	Refresh(input *jweModel.Refresh) interface{}
}

type resolver struct {
	Account account.Service
	JWE     jwe.Service
}

func New(db *gorm.DB) Resolver {
	return &resolver{
		Account: account.New(db),
		JWE:     jwe.New(),
	}
}
