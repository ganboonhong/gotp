package database

import (
	"testing"

	"github.com/ganboonhong/gotp/pkg/testutil"
	"github.com/ganboonhong/gotp/pkg/user"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	t  *testing.T
	rq *require.Assertions
)

type UserSuite struct {
	suite.Suite
}

func (suite *UserSuite) SetupSuite() {
	testutil.SetupDB()

	t = suite.T()
	rq = suite.Require()
}

// TestCRUDUser tests (C)reate, (R)ead, (U)pdate, (D)elete user entity
func (suite *UserSuite) TestCRUDUser() {
	gormDB, _ := gorm.Open(sqlite.Open(testutil.DSN), &gorm.Config{})
	DB := NewDB(gormDB)

	DB.Transaction(func(tx *gorm.DB) error {
		// create
		password, _ := bcrypt.GenerateFromPassword([]byte("plainpassword"), bcrypt.MinCost)
		u := &user.User{
			Account:  "Test",
			Password: string(password),
		}
		err := DB.Create(&u)
		rq.NoError(err)
		suite.Equal(1, int(u.ID))

		// find
		u = &user.User{}
		err = DB.Find(1, u)
		require.NoError(t, err)
		suite.Equal(uint(1), u.ID)

		// update
		expected := "Test2"
		u.Account = expected
		err = DB.Update(u)
		suite.Equal(expected, u.Account)

		// delete
		execution := DB.Delete(user.User{}, 1)
		suite.Equal(1, int(execution.RowsAffected))
		require.NoError(t, execution.Error)
		return nil
	})
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(UserSuite))
}
