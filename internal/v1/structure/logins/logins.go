package logins

type Login struct {
	//公司ID
	CompanyID string `json:"company_id,omitempty" binding:"required" validate:"required"`
	//帳號
	Account string `json:"account,omitempty" binding:"required,email" validate:"required"`
	//密碼
	Password string `json:"password,omitempty" binding:"required" validate:"required"`
}
