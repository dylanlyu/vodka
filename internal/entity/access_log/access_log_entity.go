package access_log

import (
	model "app.inherited.magic/internal/entity/db/access_logs"
	"app.inherited.magic/internal/interactor/util/log"
	"encoding/json"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Entity interface {
	WithTrx(trx *gorm.DB) Entity
	Create(input *model.Base) (err error)
	GetByList(input *model.Base) (quantity int64, output []*model.Table, err error)
	GetBySingle(input *model.Base) (output *model.Table, err error)
	GetByQuantity(input *model.Base) (quantity int64, err error)
	Delete(input *model.Base) (err error)
	Update(input *model.Base) (err error)
}

type storage struct {
	db *gorm.DB
}

func Init(db *gorm.DB) Entity {
	return &storage{
		db: db,
	}
}

func (s *storage) WithTrx(trx *gorm.DB) Entity {
	return &storage{
		db: trx,
	}
}

func (s *storage) Create(input *model.Base) (err error) {
	marshal, err := json.Marshal(input)
	if err != nil {
		log.Error(err)
		return err
	}

	data := &model.Table{}
	err = json.Unmarshal(marshal, data)
	if err != nil {
		log.Error(err)
		return err
	}

	err = s.db.Model(&model.Table{}).Omit(clause.Associations).Create(&data).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (s *storage) GetByList(input *model.Base) (quantity int64, output []*model.Table, err error) {
	query := s.db.Model(&model.Table{}).Preload(clause.Associations)
	if input.ID != nil {
		query.Where("id = ?", input.ID)
	}

	if input.IPAddress != nil {
		query.Where("ip_address like %?%", input.IPAddress)
	}

	if input.MemberID != nil {
		query.Where("member_id = ?", input.MemberID)
	}

	if input.Method != nil {
		query.Where("method = ?", input.Method)
	}

	if input.API != nil {
		query.Where("api like %?%", input.API)
	}

	if input.Msg != nil {
		query.Where("msg like %?%", input.Msg)
	}

	if input.Status != nil {
		query.Where("status = ?", input.Status)
	}

	if input.CreatedAt != nil {
		query.Where("created_at = ?", input.CreatedAt)
	}

	if input.UpdatedAt != nil {
		query.Where("updated_at = ?", input.UpdatedAt)
	}

	if input.DeletedAt != nil {
		query.Unscoped().Where("deleted_at = ?", input.DeletedAt)
	}

	if input.StartAt != nil {
		query.Where("updated_at >= ?", input.StartAt)
	}

	if input.EndAt != nil {
		query.Where("updated_at <= ?", input.EndAt)
	}

	if input.DelStartAt != nil {
		query.Unscoped().Where("deleted_at >= ?", input.DelStartAt)
	}

	if input.DelEndAt != nil {
		query.Unscoped().Where("deleted_at <= ?", input.DelEndAt)
	}

	if input.OrderBy != nil {
		query.Order(*input.OrderBy)
	}

	err = query.Count(&quantity).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Find(&output).Error
	if err != nil {
		log.Error(err)
		return 0, nil, err
	}

	return quantity, output, nil
}

func (s *storage) GetBySingle(input *model.Base) (output *model.Table, err error) {
	query := s.db.Model(&model.Table{}).Preload(clause.Associations)
	if input.ID != nil {
		query.Where("id = ?", input.ID)
	}

	if input.IPAddress != nil {
		query.Where("ip_address like %?%", input.IPAddress)
	}

	if input.MemberID != nil {
		query.Where("member_id = ?", input.MemberID)
	}

	if input.Method != nil {
		query.Where("method like %?%", input.Method)
	}

	if input.API != nil {
		query.Where("api like %?%", input.API)
	}

	if input.Msg != nil {
		query.Where("msg like %?%", input.Msg)
	}

	if input.Status != nil {
		query.Where("status = ?", input.Status)
	}

	if input.CreatedAt != nil {
		query.Where("created_at = ?", input.CreatedAt)
	}

	if input.UpdatedAt != nil {
		query.Where("updated_at = ?", input.UpdatedAt)
	}

	if input.DeletedAt != nil {
		query.Unscoped().Where("deleted_at = ?", input.DeletedAt)
	}

	if input.StartAt != nil {
		query.Where("updated_at >= ?", input.StartAt)
	}

	if input.EndAt != nil {
		query.Where("updated_at <= ?", input.EndAt)
	}

	if input.DelStartAt != nil {
		query.Unscoped().Where("deleted_at >= ?", input.DelStartAt)
	}

	if input.DelEndAt != nil {
		query.Unscoped().Where("deleted_at <= ?", input.DelEndAt)
	}

	err = query.First(&output).Error
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return output, nil
}

func (s *storage) GetByQuantity(input *model.Base) (quantity int64, err error) {
	query := s.db.Model(&model.Table{})
	if input.ID != nil {
		query.Where("id = ?", input.ID)
	}

	if input.IPAddress != nil {
		query.Where("ip_address like %?%", input.IPAddress)
	}

	if input.MemberID != nil {
		query.Where("member_id = ?", input.MemberID)
	}

	if input.Method != nil {
		query.Where("method like %?%", input.Method)
	}

	if input.API != nil {
		query.Where("api like %?%", input.API)
	}

	if input.Msg != nil {
		query.Where("msg like %?%", input.Msg)
	}

	if input.Status != nil {
		query.Where("status = ?", input.Status)
	}

	if input.CreatedAt != nil {
		query.Where("created_at = ?", input.CreatedAt)
	}

	if input.UpdatedAt != nil {
		query.Where("updated_at = ?", input.UpdatedAt)
	}

	if input.DeletedAt != nil {
		query.Unscoped().Where("deleted_at = ?", input.DeletedAt)
	}

	if input.StartAt != nil {
		query.Where("updated_at >= ?", input.StartAt)
	}

	if input.EndAt != nil {
		query.Where("updated_at <= ?", input.EndAt)
	}

	if input.DelStartAt != nil {
		query.Unscoped().Where("deleted_at >= ?", input.DelStartAt)
	}

	if input.DelEndAt != nil {
		query.Unscoped().Where("deleted_at <= ?", input.DelEndAt)
	}

	err = query.Count(&quantity).Select("*").Error
	if err != nil {
		log.Error(err)
		return 0, err
	}

	return quantity, nil
}

func (s *storage) Update(input *model.Base) (err error) {
	query := s.db.Model(&model.Table{}).Omit(clause.Associations)
	data := map[string]any{}
	if input.Method != nil {
		data["method"] = input.Method
	}

	if input.API != nil {
		data["api"] = input.API
	}

	if input.Msg != nil {
		data["msg"] = input.Msg
	}

	if input.Status != nil {
		data["status"] = input.Status
	}

	if input.IPAddress != nil {
		query.Where("ip_address = ?", input.IPAddress)
	}

	if input.MemberID != nil {
		query.Where("member_id = ?", input.MemberID)
	}

	if input.ID != nil {
		query.Where("id = ?", input.ID)
	}

	err = query.Select("*").Updates(data).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (s *storage) Delete(input *model.Base) (err error) {
	query := s.db.Model(&model.Table{}).Omit(clause.Associations)
	if input.ID != nil {
		query.Where("id = ?", input.ID)
	}

	if input.IPAddress != nil {
		query.Where("ip_address = ?", input.IPAddress)
	}

	if input.MemberID != nil {
		query.Where("member_id = ?", input.MemberID)
	}

	if input.Method != nil {
		query.Where("method = ?", input.Method)
	}

	if input.API != nil {
		query.Where("api = ?", input.API)
	}

	if input.Msg != nil {
		query.Where("msg like %?%", input.Msg)
	}

	if input.Status != nil {
		query.Where("status = ?", input.Status)
	}

	if input.CreatedAt != nil {
		query.Where("created_at = ?", input.CreatedAt)
	}

	if input.UpdatedAt != nil {
		query.Where("updated_at = ?", input.UpdatedAt)
	}

	if input.DeletedAt != nil {
		query.Unscoped().Where("deleted_at = ?", input.DeletedAt)
	}

	if input.StartAt != nil {
		query.Where("updated_at >= ?", input.StartAt)
	}

	if input.EndAt != nil {
		query.Where("updated_at <= ?", input.EndAt)
	}

	err = query.Delete(&model.Table{}).Error
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
