package user

import (
	"database/sql"
	"log"

	survey "github.com/AlecAivazis/survey/v2"
	"github.com/ganboonhong/gotp/pkg/cmdutil"
	"github.com/ganboonhong/gotp/pkg/user"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

// NewCreateCommand creates new user
func NewCreateCommand(f *cmdutil.Factory) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create new user",
		RunE: func(cmd *cobra.Command, args []string) error {
			qs := []*survey.Question{
				{
					Name: "username",
					Prompt: &survey.Input{
						Message: "Please key in your username",
					},
				},
				{
					Name: "Enable Cloud Backup",
					Prompt: &survey.Confirm{
						Message: "Enable Backup on Private Cloud Drive?",
					},
				},
			}
			ans := struct {
				Username     string
				EnableBackup bool
			}{}

			survey.Ask(qs, &ans)
			log.Println(ans.EnableBackup)

			if ans.EnableBackup {
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
				"DB_USER": ans.Username,
			}
			err := godotenv.Write(envVars, ".env")
			if err != nil {
				return err
			}

			dbPath := "./data/db.sqlite"
			// os.Remove(dbPath)

			db, err := sql.Open("sqlite3", dbPath)
			if err != nil {
				return err
			}

			defer db.Close()

			repo := user.NewRepo(db)
			u := &user.User{
				Name: ans.Username,
			}
			id, err := repo.Store(u)
			if err != nil {
				return err
			}
			log.Printf("User %s (id: %d) created", ans.Username, id)
			return nil
		},
	}
}
