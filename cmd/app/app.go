package app

import (
	"github.com/ganboonhong/gotp/pkg/config"
	"github.com/spf13/cobra"
)

func NewAppCmd(config *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "app",
		Short: "Initialize, update, remove gotp app",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	cmd.AddCommand(NewInitCommand(config))

	return cmd
}
