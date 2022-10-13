package section

import (
	"time"
)

type StartEnd struct {
	//開始時間
	StartAt *time.Time `json:"startAt,omitempty" form:"startAt"`
	//結束時間
	EndAt *time.Time `json:"endAt,omitempty" form:"endAt"`
}

type TimeAt struct {
	//創建時間
	CreatedAt *time.Time `json:"createdAt"`
	//更新時間
	UpdatedAt *time.Time `json:"updatedAt"`
	//刪除時間
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}

type ManagementExclusive struct {
	//刪除的開始時間
	DelStartAt *time.Time `json:"delStartAt,omitempty" form:"delStartAt"`
	//刪除的結束時間
	DelEndAt *time.Time `json:"delEndAt,omitempty" form:"delEndAt"`
}
