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

func (suite *s) TestGenerateTOTP() {
	gormDB, _ := gorm.Open(sqlite.Open(testutil.DSN), &gorm.Config{})
	DB := database.NewDB(gormDB)
	f := &cmdutil.Factory{
		GetConfig: cmdutil.GetConfigTest,
		DB:        DB,
	}

	DB.Create(&user.User{
		Account:  "Test",
		Password: "hashedpassword",
	})

	chooseType := false

	msg, err := generate(f, chooseType)
	if err != nil {
		t.Error(err.Error())
	}

	suite.Contains(msg, "Your OTP: ")
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(s))
}
