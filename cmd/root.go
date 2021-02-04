package cmd

import (
	"github.com/ganboonhong/gotp/cmd/app"
	"github.com/ganboonhong/gotp/cmd/generate"
	"github.com/ganboonhong/gotp/cmd/otp"
	"github.com/ganboonhong/gotp/cmd/user"
	"github.com/ganboonhong/gotp/pkg/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gotp",
	Short: "CLI OTP generator",
}

// Execute executes the root command.
func Execute() error {
	config := config.New()

	rootCmd.AddCommand(generate.NewGenereateCmd(config))
	rootCmd.AddCommand(user.NewUserCmd(config))
	rootCmd.AddCommand(otp.NewOTPCmd(config))
	rootCmd.AddCommand(app.NewAppCmd(config))

	return rootCmd.Execute()
}
