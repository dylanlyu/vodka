package login

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"vodka.app/internal/pkg/code"
	"vodka.app/internal/pkg/log"
	"vodka.app/internal/v1/resolver/login"
	"vodka.app/internal/v1/structure/jwe"
	"vodka.app/internal/v1/structure/logins"
)

type Presenter interface {
	Web(ctx *gin.Context)
	Refresh(ctx *gin.Context)
}

type presenter struct {
	Login login.Resolver
}

func New(db *gorm.DB) Presenter {
	return &presenter{
		Login: login.New(db),
	}
}

// Web
// @Summary 使用者登入
// @description 使用者登入
// @Tags Login
// @version 1.0
// @Accept json
// @produce json
// @param * body logins.Login true "登入帶入"
// @success 200 object code.SuccessfulMessage{body=jwe.Token} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/login/web [post]
func (p *presenter) Web(ctx *gin.Context) {
	input := &logins.Login{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))
		return
	}

	codeMessage := p.Login.Web(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Refresh
// @Summary 換新的令牌
// @description 換新的令牌
// @Tags Login
// @version 1.0
// @Accept json
// @produce json
// @param * body jwe.Refresh true "登入帶入"
// @success 200 object code.SuccessfulMessage{body=jwe.Token} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /authority/v1.0/login/refresh [post]
func (p *presenter) Refresh(ctx *gin.Context) {
	input := &jwe.Refresh{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))
		return
	}

	codeMessage := p.Login.Refresh(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
