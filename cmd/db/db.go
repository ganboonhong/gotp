package db

import (
	"github.com/spf13/cobra"
)

// New returns a command that handle database related operations
func New() *cobra.Command {
	dbCmd := &cobra.Command{
		Use:   "db",
		Short: "database related manipulation",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	dbCmd.AddCommand(NewGetCommand())
	dbCmd.AddCommand(NewSetCommand())

	return dbCmd
}
