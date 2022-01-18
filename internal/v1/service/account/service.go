package account

import (
	"gorm.io/gorm"
	"vodka.app/internal/v1/entity/account"
	model "vodka.app/internal/v1/structure/accounts"
)

type Service interface {
	WithTrx(tx *gorm.DB) Service
	Created(input *model.Created) (output *model.Base, err error)
	List(input *model.Fields) (quantity int64, output []*model.Base, err error)
	GetByID(input *model.Field) (output *model.Base, err error)
	Deleted(input *model.Updated) (err error)
	Updated(input *model.Updated) (err error)
	AcknowledgeAccount(input *model.Field) (acknowledge bool, output []*model.Base, err error)
}

type service struct {
	Entity account.Entity
}

func New(db *gorm.DB) Service {

	return &service{
		Entity: account.New(db),
	}
}

func (s *service) WithTrx(tx *gorm.DB) Service {

	return &service{
		Entity: s.Entity.WithTrx(tx),
	}
}
