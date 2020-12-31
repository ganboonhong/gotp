package cmd

import (
	genCmd "github.com/ganboonhong/gotp/cmd/generate"
	otpCmd "github.com/ganboonhong/gotp/cmd/otp"
	userCmd "github.com/ganboonhong/gotp/cmd/user"
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

	rootCmd.AddCommand(genCmd.New(f))
	rootCmd.AddCommand(userCmd.New(f))
	rootCmd.AddCommand(otpCmd.New(f))

	return rootCmd.Execute()
}
