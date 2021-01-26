package user

import (
	"github.com/ganboonhong/gotp/pkg/parameter"
)

const Table = "users"

type User struct {
	ID         uint   `json:"id,omitempty" gorm:"primaryKey"`
	Account    string `json:"account"`
	Password   string `json:"password"`
	CreatedAt  int    `json:"created_at,omitempty"`
	Parameters []parameter.Parameter
}
