package main

// main is run all api form localhost port 8080
import (
	"fmt"
	"net/http"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "vodka.app/api"
	"vodka.app/internal/pkg/dao/gorm"
	"vodka.app/internal/pkg/log"
	"vodka.app/internal/v1/router"
	"vodka.app/internal/v1/router/account"
	"vodka.app/internal/v1/router/company"
	"vodka.app/internal/v1/router/login"
)

// @title Vodka SYSTEM API
// @version 0.1
// @description 企業系統整合管理平台
// @termsOfService https://vodka.app/

// @contact.name API System Support
// @contact.url https://vodka.app/
// @contact.email mingzong.lyu@gmail.com

// @license.name AGPL 3.0
// @license.url https://www.gnu.org/licenses/agpl-3.0.en.html

// @host api.vodka.app
// @BasePath /
// @schemes https
func main() {
	db, err := gorm.New()
	if err != nil {
		log.Error(err)
		return
	}

	route := router.Default()
	route = account.GetRoute(route, db)
	route = company.GetRoute(route, db)
	route = login.GetRoute(route, db)
	url := ginSwagger.URL(fmt.Sprintf("http://localhost:8080/swagger/doc.json"))
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	log.Fatal(http.ListenAndServe(":8080", route))
}
