package company

import (
	"encoding/json"
	"vodka.app/internal/pkg/log"
	"vodka.app/internal/pkg/util"
	model "vodka.app/internal/v1/structure/companies"
)

func (s *service) Created(input *model.Created) (output *model.Base, err error) {
	marshal, err := json.Marshal(input)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	output.CompanyID = util.GenerateUUID()
	output.CreatedAt = util.NowToUTC()
	output.IsDeleted = false
	marshal, err = json.Marshal(output)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	table := &model.Table{}
	err = json.Unmarshal(marshal, &table)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	err = s.Entity.Created(table)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return output, nil
}

func (s *service) List(input *model.Fields) (quantity int64, output []*model.Base, err error) {
	amount, fields, err := s.Entity.List(input)
	if err != nil {
		log.Error(err)
		return 0, output, err
	}

	marshal, err := json.Marshal(fields)
	if err != nil {
		log.Error(err)
		return 0, output, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)
		return 0, output, err
	}

	return amount, output, err
}

func (s *service) GetByID(input *model.Field) (output *model.Base, err error) {
	field, err := s.Entity.GetByID(input)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	marshal, err := json.Marshal(field)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	err = json.Unmarshal(marshal, &output)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return output, nil
}

func (s *service) Deleted(input *model.Updated) (err error) {
	field, err := s.Entity.GetByID(&model.Field{CompanyID: input.CompanyID,
		IsDeleted: util.PointerBool(false)})
	if err != nil {
		log.Error(err)
		return err
	}

	field.UpdatedBy = input.UpdatedBy
	field.UpdatedAt = util.PointerTime(util.NowToUTC())
	field.IsDeleted = true
	err = s.Entity.Updated(field)

	return err
}

func (s *service) Updated(input *model.Updated) (err error) {
	field, err := s.Entity.GetByID(&model.Field{CompanyID: input.CompanyID,
		IsDeleted: util.PointerBool(false)})
	if err != nil {
		log.Error(err)
		return err
	}

	marshal, err := json.Marshal(input)
	if err != nil {
		log.Error(err)
		return err
	}

	err = json.Unmarshal(marshal, &field)
	if err != nil {
		log.Error(err)
		return err
	}

	err = s.Entity.Updated(field)
	return err
}
