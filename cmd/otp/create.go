package otp

import (
	"fmt"
	"os"
	"strings"

	survey "github.com/AlecAivazis/survey/v2"

	"github.com/ganboonhong/gotp/pkg/config"
	"github.com/ganboonhong/gotp/pkg/crypto"
	errMsg "github.com/ganboonhong/gotp/pkg/error"
	"github.com/ganboonhong/gotp/pkg/orm"
	"github.com/ganboonhong/gotp/pkg/parameter"

	"github.com/spf13/cobra"
)

type answer struct {
	Secret  string
	Issuer  string
	Account string
}

// NewCreateCommand creates new user
func NewCreateCommand(config *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create OTP",
		RunE: func(cmd *cobra.Command, args []string) error {
			q := questions()
			a := answer{}
			survey.Ask(q, &a)
			return create(config, &a)
		},
	}
}

func questions() []*survey.Question {
	return []*survey.Question{
		{
			Name: "secret",
			Prompt: &survey.Input{
				Message: "Please type in \"secret\" of your OTP: ",
			},
		},
		{
			Name: "issuer",
			Prompt: &survey.Input{
				Message: "Please type in \"issuer\" of your OTP, e.g.: Google, GitHub, AWS: ",
			},
		},
		{
			Name: "account",
			Prompt: &survey.Input{
				Message: "Please type in \"account\" as an identifier of same issuer, e.g.: email, account name: ",
			},
		},
	}
}

func create(c *config.Config, a *answer) error {
	if c.UserID == 0 {
		fmt.Fprintf(os.Stderr, errMsg.NoAccount())
		os.Exit(1)
	}

	orm := orm.New(c)
	secret := strings.ReplaceAll(a.Secret, " ", "")
	secret = crypto.Encrypt(secret, config.Key)
	p := &parameter.Parameter{
		UserID:  uint(c.UserID),
		Secret:  secret,
		Issuer:  a.Issuer,
		Account: a.Account,
	}

	err := orm.Create(p)
	if err != nil {
		return err
	}

	return nil
}
