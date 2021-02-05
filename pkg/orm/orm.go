package orm

import (
	"fmt"

	"github.com/ganboonhong/gotp/pkg/config"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ORM is a wrapper of gorm.io/gorm with additional Create, Find, Update, Delete methods.
type ORM struct {
	*gorm.DB
}

func New(c *config.Config) *ORM {
	db, err := gorm.Open(sqlite.Open(c.DatabasePath()), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("%s: %s", err.Error(), c.DSN()))
	}

	return &ORM{
		DB: db,
	}
}
