package otp

import (
	survey "github.com/AlecAivazis/survey/v2"
	"github.com/ganboonhong/gotp/pkg/cmdutil"

	"github.com/ganboonhong/gotp/pkg/database"
	"github.com/ganboonhong/gotp/pkg/parameter"

	"github.com/spf13/cobra"
)

// NewCreateCommand creates new user
func NewCreateCommand(f *cmdutil.Factory) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create OTP",
		RunE: func(cmd *cobra.Command, args []string) error {

			// TODO: get current login user ID (from cache)?
			userId := uint(1)

			qs := []*survey.Question{
				{
					Name: "Secret",
					Prompt: &survey.Input{
						Message: "Please type in \"secret\" of your OTP: ",
					},
				},
				{
					Name: "Issuer",
					Prompt: &survey.Input{
						Message: "Please type in \"issuer\" of your OTP, e.g.: Google, GitHub, AWS: ",
					},
				},
				{
					Name: "Account",
					Prompt: &survey.Input{
						Message: "Please type in \"account\" as an identifier of same issuer, e.g.: email, account name: ",
					},
				},
			}
			answer := struct {
				Secret  string
				Issuer  string
				Account string
			}{}

			survey.Ask(qs, &answer)

			db := database.NewDb(nil)
			p := &parameter.Parameter{
				UserID:  userId,
				Secret:  answer.Secret,
				Issuer:  answer.Issuer,
				Account: answer.Account,
			}

			err := db.Create(p)
			if err != nil {
				return err
			}

			return nil
		},
	}
}
