package user

import (
	"log"
	"strconv"

	survey "github.com/AlecAivazis/survey/v2"
	"github.com/ganboonhong/gotp/pkg/config"
	"github.com/ganboonhong/gotp/pkg/crypto"
	"github.com/ganboonhong/gotp/pkg/orm"
	"github.com/ganboonhong/gotp/pkg/user"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

// NewCreateCommand creates new user
func NewCreateCommand(config *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create new user",
		RunE: func(cmd *cobra.Command, args []string) error {
			qs := []*survey.Question{
				{
					Name: "Account",
					Prompt: &survey.Input{
						Message: "Please key in your Account Name",
					},
				},
				{
					Name: "Password",
					Prompt: &survey.Password{
						Message: "Please key in your Password",
					},
				},
				{
					Name: "EnableBackup",
					Prompt: &survey.Confirm{
						Message: "Enable Backup on Private Cloud Drive?",
						Default: true,
					},
				},
			}
			answer := struct {
				Account      string
				Password     string
				EnableBackup bool
			}{}

			survey.Ask(qs, &answer)

			if answer.EnableBackup {
				var cloudDrive string
				googleDrive := "Google Drive"
				oneDrive := "One Drive"
				prompt := &survey.Select{
					Message: "Select cloud storage",
					Options: []string{
						googleDrive,
						oneDrive,
						"Cancel",
					},
				}
				survey.AskOne(prompt, &cloudDrive)

				if cloudDrive == googleDrive {

					log.Println("Google Drive token granted")
				} else if cloudDrive == oneDrive {
					// get one drive token
					log.Println("One Drive token granted")
				}
			}

			orm := orm.New(config)
			password := []byte(answer.Password)
			u := &user.User{
				Account:  answer.Account,
				Password: string(crypto.HashPassword(password)),
			}

			err := orm.Create(u)
			if err != nil {
				return err
			}
			envVars := map[string]string{
				"UserID": strconv.FormatUint(uint64(u.ID), 10),
			}
			err = godotenv.Write(envVars, ".env")
			if err != nil {
				return err
			}

			log.Printf("User %s (id: %d) created", u.Account, u.ID)
			return nil
		},
	}
}
