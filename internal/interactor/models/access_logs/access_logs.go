package access_logs

import (
	"app.inherited.magic/internal/interactor/models/page"
	"app.inherited.magic/internal/interactor/models/section"
)

// Create struct is used to create access_log
type Create struct {
	//IP位子
	IPAddress string `json:"ipAddress,omitempty" binding:"required,ip" validate:"required,ip"`
	//使用者編號
	MemberID string `json:"memberID,omitempty" binding:"required,uuid4" validate:"required,uuid4"`
	//呼叫方法
	Method string `json:"method,omitempty" binding:"required,oneof=OPTIONS HEAD POST GET PATCH PUT DELETE" validate:"required,oneof=OPTIONS HEAD POST GET PATCH PUT DELETE"`
	//使用者呼叫API
	API string `json:"api,omitempty" binding:"required,uri" validate:"required,uri"`
	//訊息
	Msg string `json:"msg,omitempty"`
	//狀態
	Status string `json:"status,omitempty" binding:"required,oneof=None Success Fail" validate:"required,oneof=None Success Fail"`
}

// Field is structure file for search
type Field struct {
	//編號
	ID string `json:"ID,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	//IP位子
	IPAddress *string `json:"ipAddress,omitempty" form:"ipAddress"`
	//使用者編號
	MemberID *string `json:"memberID,omitempty" form:"memberID" binding:"omitempty,uuid4" validate:"omitempty,uuid4"`
	//呼叫方法
	Method *string `json:"method,omitempty" form:"method" binding:"omitempty,oneof=OPTIONS HEAD POST GET PATCH PUT DELETE" validate:"omitempty,oneof=OPTIONS HEAD POST GET PATCH PUT DELETE"`
	//使用者呼叫API
	API *string `json:"api,omitempty" form:"API"`
	//訊息
	Msg *string `json:"msg,omitempty" form:"msg"`
	//狀態
	Status *string `json:"status,omitempty" form:"status"`
}

// Fields is the searched structure file (including pagination)
type Fields struct {
	//搜尋結構檔
	Field
	//分頁搜尋結構檔
	page.Pagination
}

// List is multiple return structure files
type List struct {
	//多筆
	AccessLogs []*struct {
		//編號
		ID string `json:"ID,omitempty"`
		//IP位子
		IPAddress string `json:"ipAddress,omitempty"`
		//使用者編號
		MemberID string `json:"memberID,omitempty"`
		//呼叫方法
		Method string `json:"method,omitempty"`
		//使用者呼叫API
		API string `json:"api,omitempty"`
		//訊息
		Msg string `json:"msg,omitempty"`
		//狀態
		Status string `json:"status,omitempty"`
		//時間戳記
		section.TimeAt
	} `json:"accessLogs"`
	//分頁返回結構檔
	page.Total
}

// Single return structure file
type Single struct {
	//編號
	ID string `json:"ID,omitempty"`
	//IP位子
	IPAddress string `json:"ipAddress,omitempty"`
	//使用者編號
	MemberID string `json:"memberID,omitempty"`
	//呼叫方法
	Method string `json:"method,omitempty"`
	//使用者呼叫API
	API string `json:"api,omitempty"`
	//訊息
	Msg string `json:"msg,omitempty"`
	//狀態
	Status string `json:"status,omitempty"`
	//時間戳記
	section.TimeAt
}

// Update struct is used to update achieves
type Update struct {
	//編號
	ID string `json:"ID,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	//IP位子
	IPAddress string `json:"ipAddress,omitempty" binding:"omitempty,ip" validate:"omitempty,ip" swaggerignore:"true"`
	//使用者編號
	MemberID string `json:"memberID,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	//呼叫方法
	Method *string `json:"method,omitempty" binding:"omitempty,oneof=OPTIONS HEAD POST GET PATCH PUT DELETE" validate:"omitempty,oneof=OPTIONS HEAD POST GET PATCH PUT DELETE"`
	//使用者呼叫API
	API *string `json:"api,omitempty" binding:"omitempty,uri" validate:"omitempty,uri"`
	//訊息
	Msg *string `json:"msg,omitempty"`
	//狀態
	Status *string `json:"status,omitempty" binding:"omitempty,oneof=None Success Fail" validate:"omitempty,oneof=None Success Fail"`
}
