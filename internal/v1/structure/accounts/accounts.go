package accounts

import (
	"time"
	model "vodka.app/internal/v1/structure"
)

// Table struct is companies database table struct
type Table struct {
	//編號
	AccountID string `gorm:"primaryKey;column:account_id;uuid_generate_v4()type:UUID;" json:"account_id,omitempty"`
	//公司ID
	CompanyID string `gorm:"column:company_id;type:UUID;" json:"company_id,omitempty"`
	//帳號
	Account string `gorm:"column:account;type:VARCHAR;" json:"account,omitempty"`
	//中文名稱
	Name string `gorm:"column:name;type:VARCHAR;" json:"name,omitempty"`
	//密碼
	Password string `gorm:"column:pwd;type:VARCHAR;" json:"password,omitempty"`
	//角色編號
	RoleID string `gorm:"column:role_id;type:VARCHAR;" json:"role_id,omitempty"`
	//是否刪除
	IsDeleted bool `gorm:"column:is_deleted;type:bool;false" json:"is_deleted,omitempty"`
	//創建時間
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;" json:"created_at"`
	//創建者
	CreatedBy string `gorm:"column:created_by;type:UUID;" json:"created_by,omitempty"`
	//更新時間
	UpdatedAt *time.Time `gorm:"column:updated_at;type:TIMESTAMP;" json:"updated_at,omitempty"`
	//更新者
	UpdatedBy *string `gorm:"column:updated_by;type:UUID;" json:"updated_by,omitempty"`
}

type Base struct {
	//編號
	AccountID string `json:"account_id,omitempty"`
	//公司ID
	CompanyID string `json:"company_id,omitempty"`
	//帳號
	Account string `json:"account,omitempty"`
	//中文名稱
	Name string `json:"name,omitempty"`
	//密碼
	Password string `json:"password,omitempty"`
	//角色編號
	RoleID string `json:"role_id,omitempty"`
	//是否刪除
	IsDeleted bool `json:"is_deleted,omitempty"`
	//創建時間
	CreatedAt time.Time `json:"created_at"`
	//創建者
	CreatedBy string `json:"created_by,omitempty"`
	//更新時間
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	//更新者
	UpdatedBy *string `json:"updated_by,omitempty"`
}

type Single struct {
	//編號
	AccountID string `json:"account_id,omitempty"`
	//公司ID
	CompanyID string `json:"company_id,omitempty"`
	//帳號
	Account string `json:"account,omitempty"`
	//中文名稱
	Name string `json:"name,omitempty"`
	//密碼
	Password string `json:"password,omitempty"`
	//角色編號
	RoleID string `json:"role_id,omitempty"`
	//創建時間
	CreatedAt time.Time `json:"created_at"`
	//創建者
	CreatedBy string `json:"created_by,omitempty"`
	//更新時間
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	//更新者
	UpdatedBy *string `json:"updated_by,omitempty"`
}

type Created struct {
	//公司ID
	CompanyID string `json:"company_id" binding:"required,omitempty" swaggerignore:"true"`
	//帳號
	Account string `json:"account,omitempty" binding:"required,email" validate:"required"`
	//中文名稱
	Name string `json:"name" binding:"required" validate:"required"`
	//密碼
	Password string `json:"password" binding:"required" validate:"required"`
	//角色編號
	RoleID string `json:"role_id" binding:"required" validate:"required"`
	//創建者
	CreatedBy string `json:"created_by" swaggerignore:"true"`
}

type Field struct {
	//編號
	AccountID string `json:"account_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//公司ID
	CompanyID *string `json:"company_id,omitempty" form:"company_id" binding:"omitempty,uuid4"`
	//帳號
	Account *string `json:"account,omitempty" form:"account"`
	//密碼
	Password *string `json:"password" swaggerignore:"true"`
	//中文名稱
	Name *string `json:"name,omitempty" form:"name"`
	//角色編號
	RoleID *string `json:"role_id,omitempty" form:"role_id"`
	//是否刪除
	IsDeleted *bool `json:"is_deleted,omitempty" swaggerignore:"true"`
}

type Fields struct {
	Field
	model.InPage
}

type List struct {
	Accounts []*struct {
		//編號
		AccountID string `json:"account_id,omitempty"`
		//公司ID
		CompanyID string `json:"company_id,omitempty"`
		//帳號
		Account string `json:"account,omitempty"`
		//中文名稱
		Name string `json:"name,omitempty"`
		//角色編號
		RoleID string `json:"role_id,omitempty"`
	} `json:"accounts"`
	model.OutPage
}

type Updated struct {
	//編號
	AccountID string `json:"account_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//組織ID
	CompanyID *string `json:"company_id,omitempty" binding:"omitempty,uuid4"`
	//中文名稱
	Name *string `json:"name,omitempty"`
	//密碼
	Password *string `json:"password,omitempty"`
	//角色編號
	RoleID *string `json:"role_id,omitempty"`
	//更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//是否刪除
	IsDeleted *bool `json:"is_deleted,omitempty" swaggerignore:"true"`
}

func (a *Table) TableName() string {
	return "accounts"
}
