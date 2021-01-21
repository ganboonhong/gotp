package otp

import (
	"fmt"
	"os"

	survey "github.com/AlecAivazis/survey/v2"
	"github.com/ganboonhong/gotp/pkg/cmdutil"

	"github.com/ganboonhong/gotp/pkg/database"
	errMsg "github.com/ganboonhong/gotp/pkg/error"
	"github.com/ganboonhong/gotp/pkg/parameter"

	"github.com/spf13/cobra"
)

type answer struct {
	Secret  string
	Issuer  string
	Account string
}

// NewCreateCommand creates new user
func NewCreateCommand(f *cmdutil.Factory) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create OTP",
		RunE: func(cmd *cobra.Command, args []string) error {
			q := questions()
			a := answer{}
			survey.Ask(q, &a)
			return create(f, &a)
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

func create(f *cmdutil.Factory, a *answer) error {
	cfg := f.GetConfig()
	if cfg.UserID == 0 {
		fmt.Fprintf(os.Stderr, errMsg.NoAccount())
		os.Exit(1)
	}

	db := database.NewDB(f.DB.DB)
	p := &parameter.Parameter{
		UserID:  uint(cfg.UserID),
		Secret:  a.Secret,
		Issuer:  a.Issuer,
		Account: a.Account,
	}

	err := db.Create(p)
	if err != nil {
		return err
	}

	return nil
}
