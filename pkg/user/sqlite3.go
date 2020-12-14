package user

import (
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
	db := r.db.First(result, uint(id))
	if db.Error != nil {
		return nil, db.Error
	}
	return result, nil
}

func (r *repo) Create(u *User) (*User, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	result := r.db.Create(u)
	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	db := tx.Commit()
	if db.Error != nil {
		return nil, db.Error
	}

	return u, nil
}

func (r *repo) Update(u *User) (*User, error) {
	db := r.db.Model(u).Updates(User{
		Name: u.Name,
	})

	if db.Error != nil {
		return nil, db.Error
	}

	return u, nil
}

func (r *repo) Delete(id int) *gorm.DB {
	return r.db.Delete(&User{}, uint(id))
}
