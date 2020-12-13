package user

import (
	"fmt"

	"gorm.io/gorm"

	_ "github.com/mattn/go-sqlite3"
)

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) Repository {
	return &repo{
		db: db,
	}
}

func (r *repo) Find(id int) (u *User, err error) {
	result := &User{}
	r.db.First(result, uint(id))
	return result, nil
}

func (r *repo) Create(u *User) (int, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return 0, tx.Error
	}

	result := r.db.Create(u)
	if result.Error != nil {
		tx.Rollback()
		return 0, result.Error
	}

	fmt.Println(u.ID)
	db := tx.Commit()
	if db.Error != nil {
		return 0, db.Error
	}

	return int(u.ID), nil
}

func (r *repo) Delete(id int) *gorm.DB {
	return r.db.Delete(&User{}, uint(id))
}
