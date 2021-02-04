package otp

import (
	"fmt"
	"os"

	survey "github.com/AlecAivazis/survey/v2"

	"github.com/ganboonhong/gotp/pkg/config"
	errMsg "github.com/ganboonhong/gotp/pkg/error"
	"github.com/ganboonhong/gotp/pkg/orm"
	"github.com/ganboonhong/gotp/pkg/parameter"
	"github.com/ganboonhong/gotp/pkg/user"

	"github.com/spf13/cobra"
)

// NewDeleteCommand creates new user
func NewDeleteCommand(config *config.Config) *cobra.Command {
	orm := orm.New(config)
	return &cobra.Command{
		Use:   "delete",
		Short: "Delete OTP",
		RunE: func(cmd *cobra.Command, args []string) error {
			options, parameters := availableAccounts(config.UserID, orm)
			q := &survey.MultiSelect{
				Message: "Select account(s) to delete:",
				Options: options,
			}
			a := []int{}
			survey.AskOne(q, &a)
			var parametersToDelete []parameter.Parameter
			for _, v := range a {
				parametersToDelete = append(parametersToDelete, parameters[v])
			}
			delete(orm, parametersToDelete)
			return nil
		},
	}
}

// availableAccounts provides OTP accounts of currently logged in user
func availableAccounts(userID int, orm *orm.ORM) ([]string, []parameter.Parameter) {
	if userID == 0 {
		fmt.Fprintf(os.Stderr, errMsg.NoAccount())
		os.Exit(1)
	}
	u := user.User{}
	orm.Find(userID, &u)
	var parameters []parameter.Parameter
	userParameters := orm.DB.Model(&u).Association("Parameters")
	userParameters.Find(&parameters)
	options := make([]string, userParameters.Count())

	for i, v := range parameters {
		options[i] = fmt.Sprintf("%s (%s)", v.Account, v.Issuer)
	}

	return options, parameters
}

func delete(orm *orm.ORM, parameters []parameter.Parameter) {
	for _, v := range parameters {
		orm.DB.
			Where("user_id = ?", v.UserID).
			Where("secret = ?", v.Secret).
			Where("issuer = ?", v.Issuer).
			Where("account = ?", v.Account).
			Delete(parameter.Parameter{})
	}
}
