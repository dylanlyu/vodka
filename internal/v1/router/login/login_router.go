package login

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"vodka.app/internal/v1/presenter/login"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := login.New(db)
	route.Group("authority").Group("v1.0").
		Group("login").POST("web", controller.Web)
	route.Group("authority").Group("v1.0").
		Group("refresh").POST("", controller.Refresh)

	return route
}
