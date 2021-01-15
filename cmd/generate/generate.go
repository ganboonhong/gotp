package generate

import (
	"fmt"
	"os"

	survey "github.com/AlecAivazis/survey/v2"
	gotp "github.com/ganboonhong/gotp/pkg"
	"github.com/ganboonhong/gotp/pkg/cmdutil"
	errMsg "github.com/ganboonhong/gotp/pkg/error"
	"github.com/ganboonhong/gotp/pkg/user"
	"github.com/spf13/cobra"
)

// New returns command to generate OTP
func New(f *cmdutil.Factory) *cobra.Command {
	var chooseType bool
	var genCmd = &cobra.Command{
		Use:   "gen",
		Short: "Generate an OPT",
		RunE: func(cmd *cobra.Command, args []string) error {
			msg, err := generate(f, chooseType)
			if err != nil {
				fmt.Errorf(err.Error())
			}
			fmt.Println(msg)

			return nil
		},
	}

	genCmd.Flags().BoolVarP(&chooseType, "type", "t", false, "Select OTP type (defaults to Time-based OTP)")

	return genCmd
}

func generate(f *cmdutil.Factory, chooseType bool) (string, error) {
	var OTPType int
	var err error
	var msg string

	cfg := f.GetConfig()
	if cfg.DbUser == "" {
		msg := errMsg.NoAccount()
		fmt.Fprintf(os.Stderr, msg)
		os.Exit(1)
	}

	if chooseType == true {
		prompt := &survey.Select{
			Message: "Which kind of OTP would you like to generate?",
			Options: []string{
				"TOTP (Time-based)",
				"HOTP (HMAC-based)",
			},
		}

		err = survey.AskOne(prompt, &OTPType)
		if err != nil {
			return msg, err
		}
	}

	DB := f.DB

	u := &user.User{}
	DB.Find(1, u)

	if OTPType == 0 {
		otp := gotp.NewDefaultTOTP("MCWFKC6VWWVIDGYC4ZULRKSLQWC7GROF")
		msg = fmt.Sprintf("Your OTP: %s", otp.Now())
	} else {
		msg = "HOTP not implemented yet"
	}

	return msg, nil
}
