package access_logs

import (
	"app.inherited.magic/internal/interactor/models/special"
)

// Table struct is access_logs database table struct
type Table struct {
	//IP位子
	IPAddress string `gorm:"column:ip_address;type:text;not null;" json:"ipAddress"`
	//使用者編號
	MemberID string `gorm:"column:member_id;type:uuid;not null;" json:"memberID"`
	//呼叫方法
	Method string `gorm:"column:method;type:text;not null;" json:"method"`
	//使用者呼叫API
	API string `gorm:"column:api;type:text;not null;" json:"api"`
	//訊息
	Msg string `gorm:"column:msg;type:text;not null;" json:"msg"`
	//狀態
	Status string `gorm:"column:status;type:text;not null;" json:"status"`
	//引入後端專用
	special.Table
}

// Base struct is corresponding to access_logs table structure file
type Base struct {
	//IP位子
	IPAddress *string `json:"ipAddress,omitempty"`
	//使用者編號
	MemberID *string `json:"memberID,omitempty"`
	//呼叫方法
	Method *string `json:"method,omitempty"`
	//使用者呼叫API
	API *string `json:"api,omitempty"`
	//訊息
	Msg *string `json:"msg,omitempty"`
	//狀態
	Status *string `json:"status,omitempty"`
	//引入後端專用
	special.Base
}

func (t *Table) TableName() string {
	return "access_logs"
}
