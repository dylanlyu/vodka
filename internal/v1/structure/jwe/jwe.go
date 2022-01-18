package jwe

type JWE struct {
	//公司ID
	CompanyID string `json:"company_id,omitempty"`
	//中文名稱
	Name string `json:"name,omitempty"`
	//編號
	AccountID string `json:"account_id,omitempty"`
}

type Token struct {
	//授權令牌
	AccessToken string `json:"access_token,omitempty"`
	//刷新令牌
	RefreshToken string `json:"refresh_token,omitempty"`
}

type Refresh struct {
	//刷新令牌
	RefreshToken string `json:"refresh_token,omitempty" binding:"required" validate:"required"`
}
