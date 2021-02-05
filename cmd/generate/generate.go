package generate

import (
	"fmt"
	"os"

	survey "github.com/AlecAivazis/survey/v2"
	"github.com/ganboonhong/gotp/pkg/config"
	errMsg "github.com/ganboonhong/gotp/pkg/error"
	"github.com/ganboonhong/gotp/pkg/orm"
	"github.com/ganboonhong/gotp/pkg/otp"
	"github.com/ganboonhong/gotp/pkg/parameter"
	"github.com/ganboonhong/gotp/pkg/user"
	"github.com/spf13/cobra"
)

// NewGenereateCmd returns command to generate OTP
func NewGenereateCmd(config *config.Config) *cobra.Command {
	var chooseType bool
	var genCmd = &cobra.Command{
		Use:   "gen",
		Short: "Generate an OPT",
		RunE: func(cmd *cobra.Command, args []string) error {
			msg, err := generate(config, chooseType)
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

func generate(config *config.Config, chooseType bool) (string, error) {
	var OTPType int
	var err error
	var msg string
	var secret string
	var parameters []parameter.Parameter

	if config.UserID == 0 {
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

	orm := orm.New(config)
	u := &user.User{}
	orm.Find(config.UserID, u)

	userParameters := orm.DB.Model(&u).Association("Parameters")
	parameterCount := userParameters.Count()
	orm.DB.Model(&u).Association("Parameters").Find(&parameters)

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
		otp := otp.NewDefaultTOTP(secret)
		msg = fmt.Sprintf("Your OTP: %s", otp.Now())
	} else {
		msg = "HOTP not implemented yet"
	}

	return msg, nil
}
