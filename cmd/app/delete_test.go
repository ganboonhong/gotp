package app

import (
	"os"
	"testing"

	"github.com/ganboonhong/gotp/pkg/config"
	pkgConfig "github.com/ganboonhong/gotp/pkg/config"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/stretchr/testify/suite"
)

type deleteAppSuite struct {
	suite.Suite
}

func (s *deleteAppSuite) BeforeTest(suiteName, testName string) {
	suitename = suiteName
	config := pkgConfig.NewTestConfig(suitename)
	if err := InitApp(config); err != nil {
		panic(err)
	}
}

func (s *deleteAppSuite) AfterTest(suiteName, testName string) {
	configDir := pkgConfig.NewTestConfig(suitename).Dir()
	os.RemoveAll(configDir)
}

func TestDeleteAppSuite(t *testing.T) {
	suite.Run(t, new(deleteAppSuite))
}

func (s *deleteAppSuite) TestDeleteApp() {
	config := config.NewTestConfig(suitename)
	configDir := config.Dir()
	assert := s.Assert()

	if err := deleteApp(config); err != nil {
		assert.FailNow(err.Error())
	}
	if _, err := os.Stat(configDir); !os.IsNotExist(err) {
		assert.FailNow("Failed to delete application")
	}
}
