package cmd

import (
	"github.com/ganboonhong/gotp/cmd/generate"
	"github.com/ganboonhong/gotp/cmd/otp"
	"github.com/ganboonhong/gotp/cmd/user"
	"github.com/ganboonhong/gotp/pkg/cmdutil"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gotp",
		Short: "one-time passcode generator",
		Long:  "Generate OTP to authenticate your application from CLI",
	}
)

// Execute executes the root command.
func Execute() error {
	f := cmdutil.NewFactory()

	rootCmd.AddCommand(generate.NewGenereateCmd(f))
	rootCmd.AddCommand(user.NewUserCmd(f))
	rootCmd.AddCommand(otp.NewOTPCmd(f))

	return rootCmd.Execute()
}
