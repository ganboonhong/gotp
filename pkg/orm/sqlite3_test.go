package orm

import (
	"testing"

	"github.com/ganboonhong/gotp/pkg/config"
	"github.com/ganboonhong/gotp/pkg/testutil"
	"github.com/ganboonhong/gotp/pkg/user"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var suitename string

type sqlite3Suite struct {
	suite.Suite
}

func (s *sqlite3Suite) BeforeTest(suiteName, testName string) {
	suitename = suiteName
	testutil.SetupDB(suitename)
}

func (s *sqlite3Suite) AfterTest(suiteName, testName string) {
	testutil.TearDownDB(suitename)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(sqlite3Suite))
}

// TestCRUDUser tests (C)reate, (R)ead, (U)pdate, (D)elete user entity
func (s *sqlite3Suite) TestCRUDUser() {
	rq := s.Require()
	config := config.NewTestConfig(suitename)
	orm := New(config)

	orm.DB.Transaction(func(tx *gorm.DB) error {
		// create
		password, err := bcrypt.GenerateFromPassword([]byte("plainpassword"), bcrypt.MinCost)
		rq.NoError(err)
		u := &user.User{
			Account:  "Test",
			Password: string(password),
		}
		err = orm.Create(&u)
		rq.NoError(err)
		s.Equal(1, int(u.ID))

		// find
		u = &user.User{}
		err = orm.Find(1, u)
		rq.NoError(err)
		s.Equal(uint(1), u.ID)

		// update
		expected := "Test2"
		u.Account = expected
		err = orm.Update(u)
		s.Equal(expected, u.Account)

		// delete
		execution := orm.DB.Delete(user.User{}, 1)
		s.Equal(1, int(execution.RowsAffected))
		rq.NoError(execution.Error)
		return nil
	})
}
