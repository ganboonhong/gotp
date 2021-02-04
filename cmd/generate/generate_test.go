package generate

import (
	"testing"

	"github.com/ganboonhong/gotp/pkg/config"
	"github.com/ganboonhong/gotp/pkg/orm"
	"github.com/ganboonhong/gotp/pkg/parameter"
	"github.com/ganboonhong/gotp/pkg/testutil"
	"github.com/ganboonhong/gotp/pkg/user"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stretchr/testify/suite"
)

var suitename string

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

func GenerateTestSuite(t *testing.T) {
	suite.Run(t, new(generateSuite))
}

func (s *generateSuite) TestGenerateTOTP() {
	config := config.NewTestConfig(suitename)
	orm := orm.New(config)
	u := &user.User{
		Account:  "Test",
		Password: "hashedpassword",
	}
	orm.Create(u)
	orm.Create(&parameter.Parameter{
		UserID:  u.ID,
		Secret:  "HXDMVJECJJWSRB3HWIZR4IFUGFTMXBOZ",
		Issuer:  "Google",
		Account: "user@gmail.com",
	})

	chooseType := false

	msg, err := generate(config, chooseType)
	if err != nil {
		s.Fail(err.Error())
	}

	s.Contains(msg, "Your OTP: ")
}
