package company

import (
	"gorm.io/gorm"
	entity "vodka.app/internal/v1/entity/company"
	model "vodka.app/internal/v1/structure/companies"
)

type Service interface {
	WithTrx(tx *gorm.DB) Service
	Created(input *model.Created) (output *model.Base, err error)
	List(input *model.Fields) (quantity int64, output []*model.Base, err error)
	GetByID(input *model.Field) (output *model.Base, err error)
	Deleted(input *model.Updated) (err error)
	Updated(input *model.Updated) (err error)
}

type service struct {
	Entity entity.Entity
}

func New(db *gorm.DB) Service {

	return &service{
		Entity: entity.New(db),
	}
}

func (s *service) WithTrx(tx *gorm.DB) Service {

	return &service{
		Entity: s.Entity.WithTrx(tx),
	}
}
