package app

import (
	"github.com/ganboonhong/gotp/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewAppCmd(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "app",
		Short: "Initialize, update, remove gotp app",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	cmd.AddCommand(NewInitCommand(f))

	return cmd
}
