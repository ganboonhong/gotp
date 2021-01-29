package testutil

import (
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/github"
)

var DSN = "test.sqlite"

func SetupDB() {
	owner := "ganboonhong"
	repo := "gotp"
	path := "migration"
	sourceURL := fmt.Sprintf("github://%s/%s/%s", owner, repo, path)
	databaseURL := "sqlite3://" + DSN

	os.Create(DSN)
	m, err := migrate.New(sourceURL, databaseURL)
	if err != nil {
		panic(err.Error())
	}
	m.Up()
}

func TearDownDB() {
	os.Remove(DSN)
}
