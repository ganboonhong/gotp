package app

import (
	"errors"
	"os"

	"github.com/ganboonhong/gotp/pkg/cmdutil"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/spf13/cobra"
)

func NewUpdateCommand(f *cmdutil.Factory) *cobra.Command {
	return &cobra.Command{
		Use:   "update",
		Short: "Update application",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := updateApp(f); err != nil {
				return err
			}
			return nil
		},
	}
}

// updateApp runs migrations to update database schema
func updateApp(f *cmdutil.Factory) error {
	databasePath := DatabasePath(f)
	if _, err := os.Stat(databasePath); os.IsNotExist(err) {
		return errors.New("Application not initialized. Run `app init`")
	}

	m, err := migrate.New(SourceURL, DatabaseURL(f))
	if err != nil {
		return err
	}
	m.Up()

	return nil
}
