package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"vodka.app/internal/pkg/code"
	"vodka.app/internal/pkg/jwe"
	"vodka.app/internal/pkg/log"
)

func Verify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		j := &jwe.JWT{
			PrivateKey: os.Getenv("PRIVATE"),
			Token:      ctx.GetHeader("Authorization"),
		}

		if len(j.Token) == 0 {
			ctx.AbortWithStatusJSON(http.StatusOK, code.GetCodeMessage(code.JWTRejected, "jwe is null"))
			return
		}

		j, err := j.Verify()
		if err != nil {
			log.Error(err)
			ctx.AbortWithStatusJSON(http.StatusOK, code.GetCodeMessage(code.JWTRejected, err.Error()))
			return
		}

		ctx.Set("account_id", j.Other["account_id"])
		ctx.Set("company_id", j.Other["company_id"])
		ctx.Next()
	}
}
