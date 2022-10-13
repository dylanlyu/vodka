package members_phone

import (
	"app.inherited.magic/internal/interactor/models/page"
	"app.inherited.magic/internal/interactor/models/section"
)

// Create struct is used to create access_log
type Create struct {
	//會員手機國碼
	PhoneCode string `json:"phoneCode,omitempty" binding:"required" validate:"required"`
	//會員手機號碼
	PhoneNumber string `json:"phoneNumber,omitempty" binding:"required" validate:"required"`
}

// Field is structure file for search
type Field struct {
	//編號
	ID *string `json:"ID,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	//會員手機國碼
	PhoneCode *string `json:"phoneCode,omitempty" form:"phoneCode"`
	//會員手機號碼
	PhoneNumber *string `json:"phoneNumber,omitempty" form:"phoneNumber"`
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
	MembersPhone []*struct {
		//編號
		ID string `json:"ID,omitempty"`
		//會員手機國碼
		PhoneCode string `json:"phoneCode,omitempty"`
		//會員手機號碼
		PhoneNumber string `json:"phoneNumber,omitempty"`
		//時間戳記
		section.TimeAt
	} `json:"membersPhone"`
	//分頁返回結構檔
	page.Total
}

// Single return structure file
type Single struct {
	//編號
	ID string `json:"ID,omitempty"`
	//會員手機國碼
	PhoneCode string `json:"phoneCode,omitempty"`
	//會員手機號碼
	PhoneNumber string `json:"phoneNumber,omitempty"`
	//時間戳記
	section.TimeAt
}

// Update struct is used to update achieves
type Update struct {
	//編號
	ID *string `json:"ID,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	//會員手機國碼
	PhoneCode *string `json:"phoneCode,omitempty"`
	//會員手機號碼
	PhoneNumber *string `json:"phoneNumber,omitempty"`
}
