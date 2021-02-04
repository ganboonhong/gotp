package config

import (
	"fmt"
	"os"
	"path/filepath"
	// "github.com/joho/godotenv"
)

const (
	SourceURL      = "github://ganboonhong/gotp/migration"
	StatusFilename = "status.yaml"

	appName     = "gotp"
	dbFilename  = "db.sqlite"
	testAppName = "gotp_test"
)

type Config struct {
	UserID  int
	AppName string
}

func New() *Config {
	// godotenv.Load()
	// userID, err := strconv.Atoi(os.Getenv("UserID"))
	// if err != nil {
	// 	panic(err)
	// }

	// TODO; Get UserID from status.yaml
	return &Config{
		UserID:  1,
		AppName: appName,
	}
}

// NewTestConfig uses customized application name (which will be sub directory name)
// to isolate unit test
func NewTestConfig(suiteName string) *Config {
	return &Config{
		UserID:  1,
		AppName: suiteName + "_" + testAppName,
	}
}

// DatabasePath returns absolute path of where sqlite file locates.
func (c *Config) DatabasePath() string {
	configDir := c.Dir()
	return filepath.Join(configDir, dbFilename)
}

// DSN returns data source name for sqlite.
func (c *Config) DSN() string {
	databasePath := c.DatabasePath()
	schema := "sqlite3://file:"
	return fmt.Sprintf("%s%s", schema, databasePath)
}

// Dir returns application configuration directory.
// Note that configuration directory is OS-specific.
// See https://pkg.go.dev/os#UserConfigDir for more information.
func (c *Config) Dir() string {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		panic(err.Error())
	}
	return filepath.Join(userConfigDir, c.AppName)
}
