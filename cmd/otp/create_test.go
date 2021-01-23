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

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

var (
	t  *testing.T
	rq *require.Assertions
)

type s struct {
	suite.Suite
}

func (suite *s) SetupSuite() {
	testutil.SetupDB()

	t = suite.T()
	rq = suite.Require()
}

func (suite *s) TearDownSuite() {
	testutil.TearDownDB()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(s))
}

func (suite *s) TestCreate() {
	var parameters []parameter.Parameter
	gormDB, _ := gorm.Open(sqlite.Open(testutil.DSN), &gorm.Config{})
	repo := database.NewRepo(gormDB)
	u := user.User{
		Account:  "FakeAccount",
		Password: "FakePassword",
	}
	repo.Create(&u)
	secret := "HXDMVJECJJWSRB3HWIZR4IFUGFTMXBOZ"
	issuer := "Google"
	account := "user@google.com"
	a := &answer{
		Secret:  secret,
		Issuer:  issuer,
		Account: account,
	}
	f := &cmdutil.Factory{
		GetConfig: cmdutil.GetConfigTest,
		Repo:      repo,
	}
	create(f, a)

	userParameters := repo.DB.Model(&u).Association("Parameters")
	suite.Equal(1, int(userParameters.Count()))

	repo.DB.Model(&u).Association("Parameters").Find(&parameters)
	p := parameters[0]
	suite.Equal(secret, p.Secret)
	suite.Equal(issuer, p.Issuer)
	suite.Equal(account, p.Account)
}
