package account

import (
	model "vodka.app/internal/v1/structure/accounts"
)

func (e *entity) Created(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Create(&input).Error

	return err
}

func (e *entity) List(input *model.Fields) (amount int64, output []*model.Table, err error) {
	db := e.db.Model(&model.Table{})
	if input.CompanyID != nil {
		db.Where("company_id = ?", input.CompanyID)
	}

	if input.Account != nil {
		db.Where("account = ?", input.Account)
	}

	if input.Name != nil {
		db.Where("name like %?%", *input.Name)
	}

	if input.RoleID != nil {
		db.Where("role_id = ?", input.RoleID)
	}

	if input.IsDeleted != nil {
		db.Where("is_deleted = ?", input.IsDeleted)
	}

	err = db.Count(&amount).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Order("created_at desc").Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByID(input *model.Field) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{}).Where("account_id = ?", input.AccountID)
	if input.IsDeleted != nil {
		db.Where("is_deleted = ?", input.IsDeleted)
	}

	err = db.First(&output).Error

	return output, err
}

func (e *entity) Deleted(input *model.Field) (err error) {
	err = e.db.Model(&model.Table{}).Delete(&input).Error

	return err
}

func (e *entity) Updated(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("account_id = ?", input.AccountID).Save(&input).Error

	return err
}
