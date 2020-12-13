package user

import "gorm.io/gorm"

type Reader interface {
	Find(id int) (*User, error)
}

type Writer interface {
	Create(user *User) (int, error)
	Delete(id int) *gorm.DB
}

type Repository interface {
	Reader
	Writer
}
