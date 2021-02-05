package app

import (
	"errors"
	"os"

	"github.com/ganboonhong/gotp/pkg/config"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/spf13/cobra"
)

func NewDeleteCommand(config *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "Delete application",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := deleteApp(config); err != nil {
				return err
			}
			return nil
		},
	}
}

// deleteApp runs migrations to update database schema
func deleteApp(config *config.Config) error {
	configDir := config.Dir()
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		return errors.New("Application not found")
	}

	if err := os.RemoveAll(configDir); err != nil {
		return err
	}

	return nil
}
