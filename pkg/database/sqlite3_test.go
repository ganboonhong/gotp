package database

import (
	"os"
	"testing"

	"github.com/ganboonhong/gotp/pkg/user"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	dsn = "test.sqlite"
	t   *testing.T
	rq  *require.Assertions
)

type UserSuite struct {
	suite.Suite
}

func (suite *UserSuite) SetupSuite() {
	os.Remove(dsn)
	m, _ := migrate.New(
		"file://../../migration",
		"sqlite3://"+dsn,
	)
	m.Steps(1)

	t = suite.T()
	rq = suite.Require()
}

// TestCRUDUser tests (C)reate, (R)ead, (U)pdate, (D)elete user entity
func (suite *UserSuite) TestCRUDUser() {
	gormDB, _ := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	config := &Config{
		Database: gormDB,
	}
	suiteRepo := NewDb(config)
	db := suiteRepo.getDb()

	db.Transaction(func(tx *gorm.DB) error {
		suiteRepo.SetTransaction(tx)
		// create
		u := &user.User{Name: "Test"}
		err := suiteRepo.Create(&u)
		rq.NoError(err)
		suite.Equal(1, int(u.ID))

		// find
		u = &user.User{}
		err = suiteRepo.Find(1, u)
		require.NoError(t, err)
		suite.Equal(uint(1), u.ID)

		// update
		expected := "Test2"
		u.Name = expected
		err = suiteRepo.Update(u)
		suite.Equal(expected, u.Name)

		// delete
		execution := suiteRepo.Delete(user.User{}, 1)
		suite.Equal(1, int(execution.RowsAffected))
		require.NoError(t, execution.Error)
		return nil
	})
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(UserSuite))
}
