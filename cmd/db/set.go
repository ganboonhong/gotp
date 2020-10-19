package db

import (
	"database/sql"
	"fmt"

	survey "github.com/AlecAivazis/survey/v2"
	"github.com/ganboonhong/gotp/pkg/cmdutil"
	"github.com/ganboonhong/gotp/pkg/user"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

// NewSetCommand sets database username and password in .env file
func NewSetCommand(f *cmdutil.Factory) *cobra.Command {
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
			}
			ans := struct {
				Username string
			}{}

			survey.Ask(qs, &ans)

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
			fmt.Println(id)
			return nil

		},
	}
}
