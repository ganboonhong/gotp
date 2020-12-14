package user

import (
	"os"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	dsn       = "test.sqlite"
	gormDB, _ = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	suiteRepo = NewRepo(gormDB)
	t         *testing.T
	rq        *require.Assertions
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

// TestCRUUser tests (C)reate, (R)ead, (U)pdate user entity
func (suite *UserSuite) TestCRUUser() {
	// create
	u := &User{Name: "Test"}
	u, err := suiteRepo.Create(u)
	rq.NoError(err)
	suite.Equal(1, int(u.ID))

	// find
	u, err = suiteRepo.Find(1)
	require.NoError(t, err)
	suite.Equal(uint(1), u.ID)

	// update
	expected := "Test2"
	u.Name = expected
	actualUser, err := suiteRepo.Update(u)
	suite.Equal(expected, actualUser.Name)
}

func (suite *UserSuite) TestDeleteUser() {
	gormDB, _ = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	suiteRepo = NewRepo(gormDB)
	tx := suiteRepo.Delete(1)
	suite.Equal(1, int(tx.RowsAffected))
	require.NoError(t, tx.Error)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(UserSuite))
}
