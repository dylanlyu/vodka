package account

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"vodka.app/internal/pkg/code"
	"vodka.app/internal/pkg/log"
	"vodka.app/internal/pkg/util"
	preset "vodka.app/internal/v1/presenter"
	"vodka.app/internal/v1/resolver/account"
	"vodka.app/internal/v1/structure/accounts"
)

type Presenter interface {
	Created(ctx *gin.Context)
	List(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Updated(ctx *gin.Context)
}

type presenter struct {
	Account account.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		Account: account.New(db),
	}
}

// Created
// @Summary 新增使用者
// @description 新增使用者
// @Tags Account
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body accounts.Created true "新增使用者"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/accounts [post]
func (p *presenter) Created(ctx *gin.Context) {
	//Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	input := &accounts.Created{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	input.CompanyID = ctx.MustGet("company_id").(string)
	input.CreatedBy = ctx.MustGet("account_id").(string)
	codeMessage := p.Account.Created(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// List
// @Summary 條件搜尋使用者
// @description 條件使用者
// @Tags Account
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param organizationID query string false "組織ID"
// @param account query string false "帳號"
// @param chineseName query string false "中文名稱"
// @param roleName query string false "角色名稱"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=accounts.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/accounts [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &accounts.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.Account.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary 取得單一使用者
// @description 取得單一使用者
// @Tags Account
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param accountID path string true "使用者ID"
// @success 200 object code.SuccessfulMessage{body=accounts.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/accounts/{accountID} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	accountID := ctx.Param("accountID")
	input := &accounts.Field{}
	input.AccountID = accountID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.Account.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary 刪除單一使用者
// @description 刪除單一使用者
// @Tags Account
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param accountID path string true "使用者ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/accounts/{accountID} [delete]
func (p *presenter) Delete(ctx *gin.Context) {
	//Todo 將UUID改成登入的使用者
	updatedBy := util.GenerateUUID()
	accountID := ctx.Param("accountID")
	input := &accounts.Updated{}
	input.AccountID = accountID
	input.UpdatedBy = util.PointerString(updatedBy)
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.Account.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Updated
// @Summary 更新單一使用者
// @description 更新單一使用者
// @Tags Account
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param accountID path string true "使用者ID"
// @param * body accounts.Updated true "更新使用者"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/accounts/{accountID} [patch]
func (p *presenter) Updated(ctx *gin.Context) {
	//Todo 將UUID改成登入的使用者
	updatedBy := util.GenerateUUID()
	accountID := ctx.Param("accountID")
	input := &accounts.Updated{}
	input.AccountID = accountID
	input.UpdatedBy = util.PointerString(updatedBy)
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.Account.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
