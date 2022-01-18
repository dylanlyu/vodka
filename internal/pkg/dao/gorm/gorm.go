package gorm

import (
	"fmt"
	"os"
	"strconv"
	"vodka.app/internal/pkg/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

func New() (*gorm.DB, error) {
	const config string = "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s"
	sources := fmt.Sprintf(config,
		os.Getenv("SOURCES_HOST"),
		os.Getenv("SOURCES_PORT"),
		os.Getenv("SOURCES_USER"),
		os.Getenv("SOURCES_PASSWORD"),
		os.Getenv("SOURCES_DATABASE"),
		os.Getenv("SOURCES_SSLMODE"),
	)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  sources,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	dbReplicas, err := strconv.Atoi(os.Getenv("DB_REPLICAS"))
	if err != nil {
		log.Debug(err)
	}
	if dbReplicas >= 1 {
		var dialectics []gorm.Dialector
		for i := 1; i <= dbReplicas; i++ {
			replicas := fmt.Sprintf(config,
				os.Getenv("REPLICAS_HOST_"+strconv.Itoa(i)),
				os.Getenv("REPLICAS_PORT_"+strconv.Itoa(i)),
				os.Getenv("REPLICAS_USER_"+strconv.Itoa(i)),
				os.Getenv("REPLICAS_PASSWORD_"+strconv.Itoa(i)),
				os.Getenv("REPLICAS_DATABASE_"+strconv.Itoa(i)),
				os.Getenv("REPLICAS_SSLMODE_"+strconv.Itoa(i)),
			)
			director := postgres.New(postgres.Config{
				DSN:                  replicas,
				PreferSimpleProtocol: true,
			})
			dialectics = append(dialectics, director)
		}
		err = db.Use(dbresolver.Register(dbresolver.Config{Replicas: dialectics}))
		if err != nil {
			log.Error(err)
		}
	}

	return db, nil
}
