package companies

import (
	"time"
	model "vodka.app/internal/v1/structure"
)

// Table struct is companies database table struct
type Table struct {
	//公司編號
	CompanyID string `gorm:"primaryKey;uuid_generate_v4();column:company_id;type:UUID;" json:"company_id,omitempty"`
	//公司名稱
	Name string `gorm:"column:name;type:VARCHAR;" json:"name,omitempty"`
	//公司統一編號
	UniformNumber int64 `gorm:"column:uniform_number;type:INT4;" json:"uniform_number,omitempty"`
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

// Base struct is corresponding to companies table structure file
type Base struct {
	//公司編號
	CompanyID string `json:"company_id,omitempty"`
	//公司名稱
	Name string `json:"name,omitempty"`
	//公司統一編號
	UniformNumber int64 `json:"uniform_number,omitempty"`
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

// Single return structure file
type Single struct {
	//公司編號
	CompanyID string `json:"company_id,omitempty"`
	//公司名稱
	Name string `json:"name,omitempty"`
	//公司統一編號
	UniformNumber int64 `json:"uniform_number,omitempty"`
	//創建時間
	CreatedAt time.Time `json:"created_at"`
	//創建者
	CreatedBy string `json:"created_by,omitempty"`
	//更新時間
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	//更新者
	UpdatedBy *string `json:"updated_by,omitempty"`
}

// Created struct is used to create companies
type Created struct {
	//公司名稱
	Name string `json:"name,omitempty" binding:"required" validate:"required"`
	//公司統一編號
	UniformNumber int64 `json:"uniform_number,omitempty" binding:"required" validate:"required"`
	//創建者
	CreatedBy string `json:"created_by,omitempty" swaggerignore:"true"`
}

// Updated struct is used to update companies
type Updated struct {
	//公司編號
	CompanyID string `json:"company_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	//公司名稱
	Name string `json:"name,omitempty"`
	//公司統一編號
	UniformNumber int64 `json:"uniform_number,omitempty"`
	//更新者
	UpdatedBy *string `json:"updated_by,omitempty" swaggerignore:"true"`
	//是否刪除
	IsDeleted bool `json:"is_deleted,omitempty" swaggerignore:"true"`
}

// Field is structure file for search
type Field struct {
	//公司編號
	CompanyID string `json:"company_id,omitempty"  binding:"omitempty,uuid4" swaggerignore:"true"`
	//公司名稱
	Name *string `json:"name,omitempty" form:"name"`
	//公司統一編號
	UniformNumber *int64 `json:"uniform_number,omitempty" form:"uniformNumber"`
	//是否刪除
	IsDeleted *bool `json:"is_deleted,omitempty" swaggerignore:"true"`
}

// Fields is the searched structure file (including pagination)
type Fields struct {
	//搜尋結構檔
	Field
	//分頁搜尋結構檔
	model.InPage
}

// List is multiple return structure files
type List struct {
	//多筆公司
	Companies []*struct {
		//公司編號
		CompanyID string `json:"company_id,omitempty"`
		//公司名稱
		Name string `json:"name,omitempty"`
		//公司統一編號
		UniformNumber int64 `json:"uniform_number,omitempty"`
		//創建時間
		CreatedAt time.Time `json:"created_at"`
		//創建者
		CreatedBy string `json:"created_by,omitempty"`
	} `json:"companies"`
	//分頁返回結構檔
	model.OutPage
}

// TableName sets the insert table name for this struct type
func (c *Table) TableName() string {
	return "companies"
}
