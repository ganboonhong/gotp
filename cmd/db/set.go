package db

import (
	survey "github.com/AlecAivazis/survey/v2"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// NewSetCommand sets database username and password in .env file
func NewSetCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "set",
		Short: "Setup database username and password",
		RunE: func(cmd *cobra.Command, args []string) error {
			// todo: update single key
			qs := []*survey.Question{
				{
					Name: "username",
					Prompt: &survey.Input{
						Message: "Please key in your username",
					},
				},
				{
					Name: "password",
					Prompt: &survey.Password{
						Message: "Please key in your password",
					},
				},
				{
					Name: "repeat password",
					Prompt: &survey.Password{
						Message: "Please repeat your password",
					},
				},
			}
			ans := struct {
				Username string
				Password string
				Confirm  string
			}{}

			survey.Ask(qs, &ans)

			// TODO: compare passwords

			envVars := map[string]string{
				"DB_USER":     ans.Username,
				"DB_PASSWORD": ans.Password,
			}
			err := godotenv.Write(envVars, ".env")
			if err != nil {
				return err
			}
			return nil
		},
	}
}
