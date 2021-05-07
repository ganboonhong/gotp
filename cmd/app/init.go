package app

import (
	"errors"
	"os"
	"path/filepath"

	pconfig "github.com/ganboonhong/gotp/pkg/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/spf13/cobra"
)

type UserStatus struct {
	UserId int
	Status string
}

const (
	StatusFilename = "status.yaml"
	dbFilename     = "db.sqlite"
)

func NewInitCommand(config *pconfig.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize application",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := InitApp(config); err != nil {
				return err
			}
			return nil
		},
	}
}

// InitApp initializes this application by creating required directories and files.
func InitApp(config *pconfig.Config) error {
	configDir := config.Dir()

	// Prevent from accidentally overwriting existing configuration.
	if _, err := os.Stat(configDir); !os.IsNotExist(err) {
		return errors.New("Configuration already exists.: " + configDir)
	}

	err := os.Mkdir(configDir, 0777)
	if err != nil {
		return err
	}

	statusFilePath := filepath.Join(configDir, StatusFilename)
	if _, err = os.Create(statusFilePath); err != nil {
		return err
	}

	databasePath := config.DatabasePath()
	if _, err = os.Create(databasePath); err != nil {
		return err
	}

	databaseURL := config.DSN()
	m, err := migrate.New(config.MigrationSource, databaseURL)
	if err != nil {
		return err
	}
	m.Up()

	return nil
}
