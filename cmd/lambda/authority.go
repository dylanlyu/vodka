package main

// authority lambda
import (
	"github.com/apex/gateway"
	"vodka.app/internal/pkg/dao/gorm"
	"vodka.app/internal/pkg/log"
	"vodka.app/internal/v1/router"
	"vodka.app/internal/v1/router/account"
	"vodka.app/internal/v1/router/company"
)

func main() {
	db, err := gorm.New()
	if err != nil {
		log.Error(err)
		return
	}

	route := router.Default()
	route = account.GetRoute(route, db)
	route = company.GetRoute(route, db)
	log.Fatal(gateway.ListenAndServe(":8080", route))
}
