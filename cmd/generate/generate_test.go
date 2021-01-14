package generate

import (
	"testing"

	"github.com/ganboonhong/gotp/pkg/cmdutil"
	"github.com/ganboonhong/gotp/pkg/database"
	"github.com/ganboonhong/gotp/pkg/testutil"
	"github.com/ganboonhong/gotp/pkg/user"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	t  *testing.T
	rq *require.Assertions
)

type Suite struct {
	suite.Suite
}

func (suite *Suite) SetupSuite() {
	testutil.SetupDB()

	t = suite.T()
	rq = suite.Require()
}

func (suite *Suite) TearDownSuite() {
	testutil.TearDownDB()
}

// func TestGenerateTOTP(t *testing.T) {
func (suite *Suite) TestGenerateTOTP() {
	gormDB, _ := gorm.Open(sqlite.Open(testutil.DSN), &gorm.Config{})
	suiteRepo := database.NewDb(gormDB)
	f := &cmdutil.Factory{
		GetConfig: cmdutil.GetConfigTest,
		DB:        suiteRepo,
	}

	u := &user.User{
		Account:  "Test",
		Password: "hashedpassword",
	}
	err := suiteRepo.Create(&u)

	chooseType := false

	msg, err := generate(f, chooseType)
	if err != nil {
		t.Error(err.Error())
	}

	suite.Contains(msg, "Your OTP: ")
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
