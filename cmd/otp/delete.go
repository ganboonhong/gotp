package otp

import (
	"fmt"
	"os"

	survey "github.com/AlecAivazis/survey/v2"
	"github.com/ganboonhong/gotp/pkg/cmdutil"

	errMsg "github.com/ganboonhong/gotp/pkg/error"
	"github.com/ganboonhong/gotp/pkg/parameter"
	"github.com/ganboonhong/gotp/pkg/user"

	"github.com/spf13/cobra"
)

// NewDeleteCommand creates new user
func NewDeleteCommand(f *cmdutil.Factory) *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "Delete OTP",
		RunE: func(cmd *cobra.Command, args []string) error {
			options, parameters := availableAccounts(f)
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
			delete(f, parametersToDelete)
			return nil
		},
	}
}

// availableAccounts provides OTP accounts of currently logged in user
func availableAccounts(f *cmdutil.Factory) ([]string, []parameter.Parameter) {
	cfg := f.GetConfig()
	if cfg.UserID == 0 {
		fmt.Fprintf(os.Stderr, errMsg.NoAccount())
		os.Exit(1)
	}
	u := user.User{}
	f.DB.Find(cfg.UserID, &u)
	var parameters []parameter.Parameter
	userParameters := f.DB.DB.Model(&u).Association("Parameters")
	userParameters.Find(&parameters)
	options := make([]string, userParameters.Count())

	for i, v := range parameters {
		options[i] = fmt.Sprintf("%s (%s)", v.Account, v.Issuer)
	}

	return options, parameters
}

func delete(f *cmdutil.Factory, parameters []parameter.Parameter) {
	for _, v := range parameters {
		f.DB.DB.
			Where("user_id = ?", v.UserID).
			Where("secret = ?", v.Secret).
			Where("issuer = ?", v.Issuer).
			Where("account = ?", v.Account).
			Delete(parameter.Parameter{})
	}
}
