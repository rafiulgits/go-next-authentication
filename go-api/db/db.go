package db

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"auth-api/models/domains"

	gormlog "gorm.io/gorm/logger"
)

type DB struct {
	*gorm.DB
}

var dbInstance *DB

func ConnectDB() error {
	gormCfg := &gorm.Config{
		DisableNestedTransaction: true,
		SkipDefaultTransaction:   true,
		Logger: gormlog.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			gormlog.Config{
				SlowThreshold:             time.Second,  // Slow SQL threshold
				LogLevel:                  gormlog.Info, // Log level
				IgnoreRecordNotFoundError: true,         // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,         // Disable color
			}),
	}

	conn, err := gorm.Open(sqlite.Open("auth.db"), gormCfg)
	if err != nil {
		return err
	}
	dbInstance = &DB{conn}
	return nil
}

//Migration : auto migrate data models
func (db *DB) Migration() {
	db.AutoMigrate(
		&domains.User{},
	)
}

func GetInstance() *DB {
	return dbInstance
}
