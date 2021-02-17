package otp

import (
	"log"
	"testing"

	"github.com/ganboonhong/gotp/pkg/config"
	"github.com/ganboonhong/gotp/pkg/orm"
	"github.com/ganboonhong/gotp/pkg/parameter"
	"github.com/ganboonhong/gotp/pkg/testutil"
	"github.com/ganboonhong/gotp/pkg/user"
	_ "github.com/mattn/go-sqlite3"

	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
)

type otpCreateSuite struct {
	suite.Suite
}

func (s *otpCreateSuite) BeforeTest(suiteName, testName string) {
	suitename = suiteName
	testutil.SetupDB(suitename)
}

func (s *otpCreateSuite) AfterTest(suiteName, testName string) {
	testutil.TearDownDB(suitename)
}

func TestOTPCreate(t *testing.T) {
	log.SetFlags(log.Llongfile)
	suite.Run(t, new(otpCreateSuite))
}

func (s *otpCreateSuite) TestCreate() {
	var parameters []parameter.Parameter
	config := config.NewTestConfig(suitename)
	orm := orm.New(config)
	u := user.User{
		Account:  "FakeAccount",
		Password: "FakePassword",
	}
	orm.Create(&u)
	secret := "HXDMVJECJJWSRB3HWIZR4IFUGFTMXBOZ"
	issuer := "Google"
	account := "user@google.com"
	a := &answer{
		Secret:  secret,
		Issuer:  issuer,
		Account: account,
	}

	create(config, a)

	userParameters := orm.DB.Model(&u).Association("Parameters")
	s.Equal(1, int(userParameters.Count()))

	orm.DB.Model(&u).Association("Parameters").Find(&parameters)
	p := parameters[0]
	err := bcrypt.CompareHashAndPassword([]byte(p.Secret), []byte(secret))
	s.Require().NoError(err)
	s.Equal(issuer, p.Issuer)
	s.Equal(account, p.Account)
}
