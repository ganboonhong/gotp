package app

import (
	"log"
	"os"
	"testing"

	"github.com/ganboonhong/gotp/pkg/cmdutil"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/stretchr/testify/suite"
)

type deleteAppSuite struct {
	suite.Suite
}

func (s *deleteAppSuite) SetupSuite() {
	log.SetFlags(log.Llongfile)
	f = &cmdutil.Factory{
		GetConfig: cmdutil.GetConfigTest,
		Repo:      nil,
	}
	configDir = ConfigDir(f)
}

func TestDeleteAppSuite(t *testing.T) {
	suite.Run(t, new(deleteAppSuite))
}

func (s *deleteAppSuite) TestDeleteApp() {
	assert := s.Assert()
	if err := initApp(f); err != nil {
		assert.FailNow(err.Error())
	}
	if err := deleteApp(f); err != nil {
		assert.FailNow(err.Error())
	}
	if _, err := os.Stat(configDir); !os.IsNotExist(err) {
		assert.FailNow("Failed to delete application")
	}
}
