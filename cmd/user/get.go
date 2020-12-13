package user

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var dbEnv = map[string]string{
	"username": "DB_USERNAME",
	"password": "DB_PASSWORD",
}

func NewGetCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "get",
		Short: "Get database setting",
		RunE: func(cmd *cobra.Command, args []string) error {
			k := args[0]
			godotenv.Load(".env")
			envKey := dbEnv[k]
			v := os.Getenv(envKey)
			fmt.Printf("%s: %s\n", k, v)

			// fmt.Printf("%v\n", args)
			return nil
		},
	}
}
