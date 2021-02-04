package otp

import (
	"github.com/ganboonhong/gotp/pkg/config"
	"github.com/spf13/cobra"
)

// NewOTPCmd returns a command that handle database related operations
func NewOTPCmd(config *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "otp",
		Short: "Create, show, update, delete an OTP info",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	cmd.AddCommand(NewCreateCommand(config))
	cmd.AddCommand(NewDeleteCommand(config))

	return cmd
}
