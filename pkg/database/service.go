package database

import (
	"path/filepath"

	"gorm.io/gorm"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
)

// Repo is a wrapper of gorm.Repo
type Repo struct {
	*gorm.DB
}

var dsn = "data/db.sqlite"

func NewRepo(DB *gorm.DB) *Repo {
	newDB := DB
	var err error
	if newDB == nil {
		dsn, _ := filepath.Abs(dsn)
		newDB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err.Error())
		}
	}

	return &Repo{
		DB: newDB,
	}
}
