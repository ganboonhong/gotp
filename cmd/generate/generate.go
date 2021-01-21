package generate

import (
	"fmt"
	"os"

	survey "github.com/AlecAivazis/survey/v2"
	gotp "github.com/ganboonhong/gotp/pkg"
	"github.com/ganboonhong/gotp/pkg/cmdutil"
	errMsg "github.com/ganboonhong/gotp/pkg/error"
	"github.com/ganboonhong/gotp/pkg/parameter"
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
				panic(err.Error())
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
	var secret string
	var parameters []parameter.Parameter

	cfg := f.GetConfig()
	if cfg.UserID == 0 {
		fmt.Fprintf(os.Stderr, errMsg.NoAccount())
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
	DB.Find(cfg.UserID, u)

	userParameters := DB.Model(&u).Association("Parameters")
	parameterCount := userParameters.Count()
	DB.Model(&u).Association("Parameters").Find(&parameters)

	if parameterCount == 0 {
		fmt.Fprintf(os.Stderr, errMsg.NoParameter())
		os.Exit(1)
	}

	if parameterCount > 1 {
		options := make([]string, parameterCount)
		for i, v := range parameters {
			option := fmt.Sprintf("%s (%s)", v.Account, v.Issuer)
			options[i] = option
		}

		var selectedOption int
		prompt := &survey.Select{
			Message: "Select an account",
			Options: options,
		}
		survey.AskOne(prompt, &selectedOption)
		secret = parameters[selectedOption].Secret
	} else {
		secret = parameters[0].Secret
	}

	if OTPType == 0 {
		otp := gotp.NewDefaultTOTP(secret)
		msg = fmt.Sprintf("Your OTP: %s", otp.Now())
	} else {
		msg = "HOTP not implemented yet"
	}

	return msg, nil
}
