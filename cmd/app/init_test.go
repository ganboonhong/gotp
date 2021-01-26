package app

import (
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
	f = &cmdutil.Factory{
		GetConfig: cmdutil.GetConfigTest,
		Repo:      nil,
	}
	configDir, _ = ConfigDir(f)
}

func (s *initAppSuite) TearDownSuite() {
	// Remove directory created by previous test.
	// test new repo
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

	statusFilePath := filepath.Join(configDir, statusFilename)
	_, err = os.Stat(statusFilePath)
	s.Equal(true, !os.IsNotExist(err))

	DBFilePath := filepath.Join(configDir, dbFilename)
	_, err = os.Stat(DBFilePath)
	s.Equal(true, !os.IsNotExist(err))
}
