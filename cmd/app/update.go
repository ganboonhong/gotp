package app

import (
	"errors"
	"os"

	pkgConfig "github.com/ganboonhong/gotp/pkg/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/spf13/cobra"
)

func NewUpdateCommand(config *pkgConfig.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "update",
		Short: "Update application",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := updateApp(config); err != nil {
				return err
			}
			return nil
		},
	}
}

// updateApp runs migrations to update database schema
func updateApp(config *pkgConfig.Config) error {
	databasePath := config.DatabasePath()
	if _, err := os.Stat(databasePath); os.IsNotExist(err) {
		return errors.New("Application not initialized. Run `app init`")
	}

	databaseURL := config.DSN()
	m, err := migrate.New(pkgConfig.SourceURL, databaseURL)
	if err != nil {
		return err
	}
	m.Up()

	return nil
}
