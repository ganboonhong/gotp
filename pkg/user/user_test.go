package user

import (
	"testing"

	"github.com/ganboonhong/gotp/pkg/testutil"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
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
func (suite *UserSuite) TestHashPassword() {
	cleartext := "Secret123!!"
	password := []byte(cleartext)
	hashedPassword := []byte(HashPassword(password))
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	rq.NoError(err)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(UserSuite))
}
