package app

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/ganboonhong/gotp/pkg/cmdutil"
	// "github.com/golang-migrate/migrate/v4"
	// _ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	// _ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

type UserStatus struct {
	UserId int
	Status string
}

const (
	statusFilename = "status.yaml"
	dbFilename     = "db.sqlite"
)

func StatusFilename() string {
	return statusFilename
}

func DBFilename() string {
	return dbFilename
}

func NewInitCommand(f *cmdutil.Factory) *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize application",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := initApp(f); err != nil {
				return err
			}
			return nil
		},
	}
}

// initApp initializes this application by creating required directories and files.
func initApp(f *cmdutil.Factory) error {
	configDir, err := ConfigDir(f)
	if err != nil {
		return err
	}

	// Prevent from accidentally overwriting existing configuration.
	if _, err = os.Stat(configDir); !os.IsNotExist(err) {
		return errors.New("Configuration already exists.")
	}

	err = os.Mkdir(configDir, 0744)
	if err != nil {
		return err
	}

	statusFilePath := filepath.Join(configDir, statusFilename)
	if _, err = os.Create(statusFilePath); err != nil {
		return err
	}

	DBFilePath := filepath.Join(configDir, dbFilename)
	if _, err = os.Create(DBFilePath); err != nil {
		return err
	}

	// m, _ := migrate.New(
	// 	"file://../../migration",
	// 	"sqlite3://"+DBFilePath,
	// )
	// m.Up()

	return nil
}

// configDir returns application configuration absolute path.
// Note that configuration directory is OS-specific.
// See https://pkg.go.dev/os#UserConfigDir for more information.
func ConfigDir(f *cmdutil.Factory) (string, error) {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(userConfigDir, f.GetConfig().AppName), nil
}
