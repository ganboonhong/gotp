package otp

import (
	"testing"

	"github.com/atotto/clipboard"
	"github.com/ganboonhong/gotp/pkg/config"
	"github.com/ganboonhong/gotp/pkg/crypto"
	"github.com/ganboonhong/gotp/pkg/orm"
	"github.com/ganboonhong/gotp/pkg/parameter"
	"github.com/ganboonhong/gotp/pkg/testutil"
	"github.com/ganboonhong/gotp/pkg/user"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stretchr/testify/suite"
)

type generateSuite struct {
	suite.Suite
}

func (s *generateSuite) BeforeTest(suiteName, testName string) {
	suitename = suiteName
	testutil.SetupDB(suitename)
}

func (s *generateSuite) AfterTest(suiteName, testName string) {
	testutil.TearDownDB(suitename)
}

func TestGenerateSuite(t *testing.T) {
	suite.Run(t, new(generateSuite))
}

func (s *generateSuite) TestGenerateTOTP() {
	c := config.NewTestConfig(suitename)
	orm := orm.New(c)
	u := &user.User{
		Account:  "Test",
		Password: "hashedpassword",
	}
	orm.Create(u)

	secret := crypto.Encrypt("HXDMVJECJJWSRB3HWIZR4IFUGFTMXBOZ", config.Key)
	orm.Create(&parameter.Parameter{
		UserID:  u.ID,
		Secret:  secret,
		Issuer:  "Google",
		Account: "user@gmail.com",
	})

	chooseType := false

	msg, err := generate(c, chooseType)
	if err != nil {
		s.Fail(err.Error())
	}

	// test workflow run failing
	// test again 1
	s.Contains(msg, "(copied)xx")

	code, err := clipboard.ReadAll()
	s.Equal(6, len(code))
}
