package testutil

import (
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var DSN = "test.sqlite"

func SetupDB() {
	os.Create(DSN)
	m, _ := migrate.New(
		"file://../../migration",
		"sqlite3://"+DSN,
	)
	m.Up()
}

func TearDownDB() {
	os.Remove(DSN)
}
