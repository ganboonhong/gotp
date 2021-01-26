package otp

import (
	"github.com/ganboonhong/gotp/pkg/cmdutil"
	"github.com/spf13/cobra"
)

// NewOTPCmd returns a command that handle database related operations
func NewOTPCmd(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "otp",
		Short: "Create, show, update, delete an OTP info",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	cmd.AddCommand(NewCreateCommand(f))
	cmd.AddCommand(NewDeleteCommand(f))

	return cmd
}
