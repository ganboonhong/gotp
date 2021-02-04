package app

import (
	"database/sql"
	"os"
	"testing"

	pkgConfig "github.com/ganboonhong/gotp/pkg/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/stretchr/testify/suite"
)

type updateAppSuite struct {
	suite.Suite
}

func (s *updateAppSuite) BeforeTest(suiteName, testName string) {
	suitename = suiteName
}

func (s *updateAppSuite) AfterTest(suiteName, testName string) {
	configDir := pkgConfig.NewTestConfig(suitename).Dir()
	os.RemoveAll(configDir)
}

func TestUpdateAppSuite(t *testing.T) {
	suite.Run(t, new(updateAppSuite))
}

func (s *updateAppSuite) TestUpdateApp() {
	var m *migrate.Migrate
	config := pkgConfig.NewTestConfig(suitename)
	assert := s.Assert()
	configDir := config.Dir()
	count := 0

	err := os.Mkdir(configDir, 0777)
	if err != nil {
		assert.FailNow(err.Error())
	}

	databasePath := config.DatabasePath()
	if _, err = os.Create(databasePath); err != nil {
		assert.FailNow(err.Error())
	}

	databaseURL := config.DSN()
	if m, err = migrate.New(pkgConfig.SourceURL, databaseURL); err != nil {
		assert.FailNow(err.Error())
	}
	m.Steps(1)

	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		assert.FailNow(err.Error())
	}

	sqlStmt := `
	select count(*) as count
	from sqlite_master
	where
	type = 'table' and
	name = 'parameters';
	`
	row := db.QueryRow(sqlStmt)
	row.Scan(&count)
	s.Equal(0, count)

	err = updateApp(config)
	if err != nil {
		assert.FailNow(err.Error())
	}
	row = db.QueryRow(sqlStmt)
	row.Scan(&count)
	s.Equal(1, count)
}
