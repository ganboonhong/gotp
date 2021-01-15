package database

import (
	"path/filepath"

	"gorm.io/gorm"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
)

// DB is a wrapper of gorm.DB
type DB struct {
	*gorm.DB
}

var dsn = "data/DB.sqlite"

func NewDB(DBArg *gorm.DB) *DB {
	newDB := DBArg
	var err error
	if newDB == nil {
		dsn, _ := filepath.Abs(dsn)
		newDB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err.Error())
		}
	}

	return &DB{
		DB: newDB,
	}
}
