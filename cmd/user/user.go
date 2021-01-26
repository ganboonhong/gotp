package user

import (
	"github.com/ganboonhong/gotp/pkg/cmdutil"
	"github.com/spf13/cobra"
)

// NewUserCmd returns a command that handle database related operations
func NewUserCmd(f *cmdutil.Factory) *cobra.Command {
	dbCmd := &cobra.Command{
		Use:   "user",
		Short: "database related manipulation",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	dbCmd.AddCommand(NewGetCommand())
	dbCmd.AddCommand(NewCreateCommand(f))

	return dbCmd
}
