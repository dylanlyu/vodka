package page

type Pagination struct {
	//頁數(請從1開始帶入)
	Page int64 `json:"page" binding:"required,gt=0" validate:"required,gt=0" form:"page"`
	//筆數(請從1開始帶入,最高上限20)
	Limit int64 `json:"limit" binding:"required,gt=0" validate:"required,gt=0" form:"limit"`
}

type Total struct {
	//頁數結構
	Pagination
	//總筆數
	Total int64 `json:"total"`
	//總頁數
	Pages int64 `json:"pages"`
}
