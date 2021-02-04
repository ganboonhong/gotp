package app

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ganboonhong/gotp/pkg/config"
	"github.com/stretchr/testify/suite"
)

var suitename string

type initAppSuite struct {
	suite.Suite
}

func (s *initAppSuite) BeforeTest(suiteName, testName string) {
	suitename = suiteName
}

func (s *initAppSuite) AfterTest(suiteName, testName string) {
	configDir := config.NewTestConfig(suitename).Dir()
	os.RemoveAll(configDir)
}

func TestInitApp(t *testing.T) {
	suite.Run(t, new(initAppSuite))
}

func (s *initAppSuite) TestInitApp() {
	config := config.NewTestConfig(suitename)
	configDir := config.Dir()
	err := InitApp(config)
	s.Require().NoError(err)

	_, err = os.Stat(configDir)
	s.True(!os.IsNotExist(err))

	statusFilePath := filepath.Join(configDir, StatusFilename)
	_, err = os.Stat(statusFilePath)
	s.True(!os.IsNotExist(err))

	databasePath := config.DatabasePath()
	_, err = os.Stat(databasePath)
	s.True(!os.IsNotExist(err))
}
