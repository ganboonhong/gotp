package otp

import (
	"testing"

	"github.com/ganboonhong/gotp/pkg/config"
	"github.com/ganboonhong/gotp/pkg/orm"
	"github.com/ganboonhong/gotp/pkg/parameter"
	"github.com/ganboonhong/gotp/pkg/testutil"
	"github.com/ganboonhong/gotp/pkg/user"
	_ "github.com/mattn/go-sqlite3"

	"github.com/stretchr/testify/suite"
)

var suitename string

type deleteOTPSuite struct {
	suite.Suite
}

func (s *deleteOTPSuite) BeforeTest(suiteName, testName string) {
	suitename = suiteName
	testutil.SetupDB(suitename)
}

func (s *deleteOTPSuite) AfterTest(suiteName, testName string) {
	testutil.TearDownDB(suitename)
}

func TestDeleteOTP(t *testing.T) {
	suite.Run(t, new(deleteOTPSuite))
}

func (s *deleteOTPSuite) TestDeleteOTP() {
	// arrange
	config := config.NewTestConfig(suitename)
	orm := orm.New(config)
	var parameters []parameter.Parameter
	u := user.User{
		Account:  "FakeAccount",
		Password: "FakePassword",
	}
	orm.Create(&u)
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
	orm.Create(&parameters)

	userParameters := orm.DB.Model(&u).Association("Parameters")
	s.Equal(2, int(userParameters.Count()))

	// act
	delete(orm, []parameter.Parameter{
		{
			UserID:  u.ID,
			Secret:  "secret1",
			Issuer:  "issuer1",
			Account: "account1",
		},
	})

	// assert
	userParameters = orm.DB.Model(&u).Association("Parameters")
	s.Equal(1, int(userParameters.Count()))
	orm.DB.Model(&u).Association("Parameters").Find(&parameters)
	p := parameters[0]
	s.Equal("secret2", p.Secret)
	s.Equal("issuer2", p.Issuer)
	s.Equal("account2", p.Account)
}
