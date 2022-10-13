package main

import (
	accessLogDB "app.inherited.magic/internal/entity/db/access_logs"
	membersPhoneDB "app.inherited.magic/internal/entity/db/members_phone"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func main() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&membersPhoneDB.Table{}, &accessLogDB.Table{})
}

func New() (*gorm.DB, error) {
	const config string = "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s"
	sources := fmt.Sprintf(config,
		"localhost",
		"4432",
		"inherited",
		"7GsBgA%#WR5?",
		"magic_test",
		"disable",
	)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  sources,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		NowFunc: func() time.Time { return time.Now().UTC() },
		Logger:  logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
