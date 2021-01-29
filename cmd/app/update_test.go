package app

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/ganboonhong/gotp/pkg/cmdutil"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/stretchr/testify/suite"
)

type updateAppSuite struct {
	suite.Suite
}

func (s *updateAppSuite) SetupSuite() {
	log.SetFlags(log.Llongfile)
	f = &cmdutil.Factory{
		GetConfig: cmdutil.GetConfigTest,
		Repo:      nil,
	}
	configDir = ConfigDir(f)
}

func (s *updateAppSuite) TearDownSuite() {
	// Remove directory created by previous test.
	os.RemoveAll(configDir)
}

func TestUpdateAppSuite(t *testing.T) {
	suite.Run(t, new(updateAppSuite))
}

func (s *updateAppSuite) TestUpdateApp() {
	var m *migrate.Migrate
	assert := s.Assert()
	configDir := ConfigDir(f)
	count := 0

	err := os.Mkdir(configDir, 0777)
	if err != nil {
		assert.FailNow(err.Error())
	}

	databasePath := DatabasePath(f)
	if _, err = os.Create(databasePath); err != nil {
		assert.FailNow(err.Error())
	}

	databaseURL := DatabaseURL(f)
	if m, err = migrate.New(SourceURL, databaseURL); err != nil {
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

	err = updateApp(f)
	if err != nil {
		assert.FailNow(err.Error())
	}
	row = db.QueryRow(sqlStmt)
	row.Scan(&count)
	s.Equal(1, count)
}
