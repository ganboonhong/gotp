package database

import "gorm.io/gorm"

type Reader interface {
	Find(int, *interface{}) error
}

type Writer interface {
	Create(*Repository) error
	Update(*Repository) error
	Delete(*interface{}, int) *gorm.DB
}

type Repository interface {
	Reader
	Writer
}
