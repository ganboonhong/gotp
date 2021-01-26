package user

import "gorm.io/gorm"

type Reader interface {
	Find(int) (*User, error)
}

type Writer interface {
	Create(*User) (*User, error)
	Update(*User) (*User, error)
	Delete(int) *gorm.DB
}

type Repository interface {
	Reader
	Writer
	Db() *gorm.DB
	SetTransaction(*gorm.DB)
}
