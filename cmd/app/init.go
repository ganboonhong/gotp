package app

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ganboonhong/gotp/pkg/cmdutil"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/spf13/cobra"
)

type UserStatus struct {
	UserId int
	Status string
}

const (
	StatusFilename = "status.yaml"
	SourceURL      = "github://ganboonhong/gotp/migration"
	dbFilename     = "db.sqlite"
)

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
	configDir := ConfigDir(f)

	// Prevent from accidentally overwriting existing configuration.
	if _, err := os.Stat(configDir); !os.IsNotExist(err) {
		return errors.New("Configuration already exists.")
	}

	err := os.Mkdir(configDir, 0777)
	if err != nil {
		return err
	}

	statusFilePath := filepath.Join(configDir, StatusFilename)
	if _, err = os.Create(statusFilePath); err != nil {
		return err
	}

	databasePath := DatabasePath(f)
	if _, err = os.Create(databasePath); err != nil {
		return err
	}

	databaseURL := DatabaseURL(f)
	m, err := migrate.New(SourceURL, databaseURL)
	if err != nil {
		return err
	}
	m.Up()

	return nil
}

func DatabasePath(f *cmdutil.Factory) string {
	configDir := ConfigDir(f)
	return filepath.Join(configDir, dbFilename)
}

func DatabaseURL(f *cmdutil.Factory) string {
	databasePath := DatabasePath(f)
	schema := "sqlite3://file:"
	return fmt.Sprintf("%s%s", schema, databasePath)
}

// ConfigDir returns application configuration absolute path.
// Note that configuration directory is OS-specific.
// See https://pkg.go.dev/os#UserConfigDir for more information.
func ConfigDir(f *cmdutil.Factory) string {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		panic(err.Error())
	}
	return filepath.Join(userConfigDir, f.GetConfig().AppName)
}
