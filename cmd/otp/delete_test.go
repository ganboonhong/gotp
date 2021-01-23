package otp

import (
	"testing"

	"github.com/ganboonhong/gotp/pkg/cmdutil"
	"github.com/ganboonhong/gotp/pkg/database"
	"github.com/ganboonhong/gotp/pkg/parameter"
	"github.com/ganboonhong/gotp/pkg/testutil"
	"github.com/ganboonhong/gotp/pkg/user"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/stretchr/testify/suite"
)

type deleteOTPSuite struct {
	suite.Suite
}

func (s *deleteOTPSuite) SetupSuite() {
	testutil.SetupDB()
}

func (s *deleteOTPSuite) TearDownSuite() {
	testutil.TearDownDB()
}

func TestDeleteOTPTestSuite(t *testing.T) {
	suite.Run(t, new(deleteOTPSuite))
}

func (s *deleteOTPSuite) TestDelete() {
	// arrange
	var parameters []parameter.Parameter
	gormDB, _ := gorm.Open(sqlite.Open(testutil.DSN), &gorm.Config{})
	repo := database.NewRepo(gormDB)
	u := user.User{
		Account:  "FakeAccount",
		Password: "FakePassword",
	}
	repo.Create(&u)
	parameters = []parameter.Parameter{
		{
			UserID:  u.ID,
			Secret:  "secret1",
			Issuer:  "issuer1",
			Account: "account1",
		},
		{
			UserID:  u.ID,
			Secret:  "secret2",
			Issuer:  "issuer2",
			Account: "account2",
		},
	}
	gormDB.Create(&parameters)
	f := &cmdutil.Factory{
		GetConfig: cmdutil.GetConfigTest,
		Repo:      repo,
	}

	userParameters := repo.DB.Model(&u).Association("Parameters")
	s.Equal(2, int(userParameters.Count()))

	// act
	delete(f, []parameter.Parameter{
		{
			UserID:  u.ID,
			Secret:  "secret1",
			Issuer:  "issuer1",
			Account: "account1",
		},
	})

	// assert
	userParameters = repo.DB.Model(&u).Association("Parameters")
	s.Equal(1, int(userParameters.Count()))
	repo.DB.Model(&u).Association("Parameters").Find(&parameters)
	p := parameters[0]
	s.Equal("secret2", p.Secret)
	s.Equal("issuer2", p.Issuer)
	s.Equal("account2", p.Account)
}
