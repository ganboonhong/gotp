package user

import (
	"log"

	survey "github.com/AlecAivazis/survey/v2"
	"github.com/ganboonhong/gotp/pkg/cmdutil"
	"github.com/ganboonhong/gotp/pkg/user"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// NewCreateCommand creates new user
func NewCreateCommand(f *cmdutil.Factory) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create new user",
		RunE: func(cmd *cobra.Command, args []string) error {
			qs := []*survey.Question{
				{
					Name: "Username",
					Prompt: &survey.Input{
						Message: "Please key in your username",
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
				Username     string
				EnableBackup bool
			}{}

			survey.Ask(qs, &answer)
			log.Println(answer)

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

			envVars := map[string]string{
				"DB_USER": answer.Username,
			}
			err := godotenv.Write(envVars, ".env")
			if err != nil {
				return err
			}

			dsn := "data/db.sqlite"
			// os.Remove(dbPath)

			db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})

			if err != nil {
				return err
			}

			repo := user.NewRepo(db)
			u := &user.User{
				Name: answer.Username,
			}

			db.Transaction(func(tx *gorm.DB) error {
				repo.SetTransaction(tx)
				u, err = repo.Create(u)
				if err != nil {
					return err
				}
				return nil
			})

			log.Printf("User %s (id: %d) created", u.Name, u.ID)
			return nil
		},
	}
}
