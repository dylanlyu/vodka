package connect

import (
	"app.inherited.magic/internal/interactor/util/log"
	"database/sql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
	"time"
)

type PostgresConfig struct {
	//Customize Driver
	DriverName *string
	//Data Source Name
	DSN *string
	//DB connect pool interface
	Conn *sql.DB
	//Disables implicit prepared statement usage
	PreferSimpleProtocol *bool
	//Creates a prepared statement when executing any SQL and caches them to speed up future calls
	PrepareStmt *bool
	//Allow to change GORM’s default logger by overriding this option
	Logger logger.Interface
	//Change the function to be used when creating a new timestamp
	NowFunc func() time.Time
	//DBResolver adds multiple databases support
	Replicas []*string
}

func (pc *PostgresConfig) Connect() (db *gorm.DB, err error) {
	postgresConfig := postgres.Config{}
	gormConfig := gorm.Config{}

	if pc.DSN != nil {
		postgresConfig.DSN = *pc.DSN
	}

	if pc.DriverName != nil {
		postgresConfig.DriverName = *pc.DriverName
	}

	if pc.Conn != nil {
		postgresConfig.Conn = pc.Conn
	}

	if pc.PreferSimpleProtocol != nil {
		postgresConfig.PreferSimpleProtocol = *pc.PreferSimpleProtocol
	}

	if pc.PrepareStmt != nil {
		gormConfig.PrepareStmt = *pc.PrepareStmt
	}

	if pc.Logger != nil {
		gormConfig.Logger = pc.Logger
	}

	if pc.NowFunc != nil {
		gormConfig.NowFunc = pc.NowFunc
	}

	db, err = gorm.Open(postgres.New(postgresConfig), &gormConfig)
	if err != nil {
		return nil, err
	}

	var dialectics []gorm.Dialector
	for _, replica := range pc.Replicas {
		director := postgres.New(postgres.Config{
			DSN:                  *replica,
			PreferSimpleProtocol: true,
		})
		dialectics = append(dialectics, director)
	}

	if pc.Replicas != nil {
		err = db.Use(dbresolver.Register(dbresolver.Config{Replicas: dialectics}))
		if err != nil {
			log.Error(err)
		}
	}

	return db, nil
}
