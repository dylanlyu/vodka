package special

import (
	"app.inherited.magic/internal/interactor/models/page"
	"app.inherited.magic/internal/interactor/models/section"
	"gorm.io/gorm"
	"time"
)

// Table is the common file of the backend table structure.
type Table struct {
	//編號
	ID string `gorm:"column:id;type:uuid;not null;primaryKey;" json:"ID"`
	//創建時間
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;not null;" json:"createdAt"`
	//更新時間
	UpdatedAt time.Time `gorm:"column:updated_at;type:TIMESTAMP;not null;" json:"updatedAt"`
	//刪除時間
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:TIMESTAMP;" json:"deletedAt,omitempty"`
}

// Base is the common file of the backend base structure.
type Base struct {
	//編號
	ID *string `json:"ID,omitempty"`
	//基本時間
	section.TimeAt
	//引入page
	page.Pagination
	//開始結束時間
	section.StartEnd
	//開始結束時間
	section.ManagementExclusive
	//SQL OrderBy 區段
	OrderBy *string
}
