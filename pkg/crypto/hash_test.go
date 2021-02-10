package crypto

import (
	"testing"

	"github.com/ganboonhong/gotp/pkg/testutil"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
)

var suitename string

type hashSuite struct {
	suite.Suite
}

func (suite *hashSuite) BeforeTest(suiteName, testName string) {
	suitename = suiteName
	testutil.SetupDB(suitename)
}

func (s *hashSuite) AfterTest(suiteName, testName string) {
	testutil.TearDownDB(suitename)
}

func (suite *hashSuite) TestHashPassword() {
	rq := suite.Require()
	cleartext := "Secret123!!"
	password := []byte(cleartext)
	hashedPassword := []byte(HashPassword(password))
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	rq.NoError(err)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(hashSuite))
}
