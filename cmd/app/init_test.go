package app

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/ganboonhong/gotp/pkg/cmdutil"
	"github.com/stretchr/testify/suite"
)

type initAppSuite struct {
	suite.Suite
}

var (
	f         *cmdutil.Factory
	configDir string
)

func (s *initAppSuite) SetupSuite() {
	log.SetFlags(log.Llongfile)
	f = &cmdutil.Factory{
		GetConfig: cmdutil.GetConfigTest,
		Repo:      nil,
	}
	configDir = ConfigDir(f)
}

func (s *initAppSuite) TearDownSuite() {
	// Remove directory created by previous test.
	os.RemoveAll(configDir)
}

func TestInitAppSuite(t *testing.T) {
	suite.Run(t, new(initAppSuite))
}

func (s *initAppSuite) TestInitApp() {
	err := initApp(f)
	s.Require().NoError(err)

	_, err = os.Stat(configDir)
	s.Equal(true, !os.IsNotExist(err))

	statusFilePath := filepath.Join(configDir, StatusFilename)
	_, err = os.Stat(statusFilePath)
	s.Equal(true, !os.IsNotExist(err))

	databasePath := DatabasePath(f)
	_, err = os.Stat(databasePath)
	s.Equal(true, !os.IsNotExist(err))
}
