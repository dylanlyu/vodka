package company

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"vodka.app/internal/v1/middleware"
	"vodka.app/internal/v1/presenter/company"
)

func GetRoute(route *gin.Engine, db *gorm.DB) *gin.Engine {
	controller := company.New(db)
	v10 := route.Group("authority").Group("v1.0").Group("companies")
	{
		v10.POST("", middleware.Verify(), middleware.Transaction(db), controller.Created)
		v10.GET("", middleware.Verify(), controller.List)
		v10.GET(":companyID", middleware.Verify(), controller.GetByID)
		v10.DELETE(":companyID", middleware.Verify(), controller.Delete)
		v10.PATCH(":companyID", middleware.Verify(), controller.Updated)
	}

	return route
}
