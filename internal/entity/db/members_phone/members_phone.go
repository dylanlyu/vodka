package members_phone

import (
	"app.inherited.magic/internal/interactor/models/special"
)

// Table struct is members_phone database table struct
type Table struct {
	//會員手機國碼
	PhoneCode string `gorm:"column:phone_code;type:text;not null;" json:"phoneCode"`
	//會員手機號碼
	PhoneNumber string `gorm:"column:phone_number;type:text;not null;" json:"phoneNumber"`
	//引入後端專用
	special.Table
}

// Base struct is corresponding to members_phone table structure file
type Base struct {
	//會員手機國碼
	PhoneCode *string `json:"phoneCode,omitempty"`
	//會員手機號碼
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	//引入後端專用
	special.Base
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "members_phone"
}
