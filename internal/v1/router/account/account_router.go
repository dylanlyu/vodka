package account

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"vodka.app/internal/v1/middleware"
	presenter "vodka.app/internal/v1/presenter/account"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := presenter.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("account")
	{
		v10.POST("", middleware.Verify(), middleware.Transaction(db), controller.Created)
		v10.GET("", middleware.Verify(), controller.List)
		v10.GET(":accountID", middleware.Verify(), controller.GetByID)
		v10.DELETE(":accountID", middleware.Verify(), controller.Delete)
		v10.PATCH(":accountID", middleware.Verify(), controller.Updated)
	}

	return route
}
