package models

type InPage struct {
	//頁數(請從1開始帶入)
	Page int64 `json:"page" binding:"required,gt=0" validate:"required" form:"page"`
	//筆數(請從1開始帶入,最高上限20)
	Limit int64 `json:"limit" binding:"required,gt=0" validate:"required" form:"limit"`
}

type OutPage struct {
	//頁數結構
	InPage
	//總頁數
	Pages int64 `json:"pages"`
}
