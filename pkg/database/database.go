package database

import (
	"log"

	"gorm.io/gorm"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
)

type database struct {
	tx *gorm.DB
}

func NewDb(config *Config) *database {

	if config.Database == nil {
		dsn := "data/db.sqlite"
		db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal(err.Error())
		}
		config.Database = db
	}

	return &database{
		tx: config.Database,
	}
}
