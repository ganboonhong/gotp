package user

import (
	"gorm.io/gorm"

	_ "github.com/mattn/go-sqlite3"
)

type repo struct {
	db *gorm.DB
	tx *gorm.DB
}

func (r *repo) Create(u *User) (*User, error) {
	if err := r.tx.Create(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (r *repo) Db() *gorm.DB {
	return r.db
}

func (r *repo) Find(id int) (u *User, err error) {
	u = &User{}
	db := r.tx.First(u, uint(id))
	if db.Error != nil {
		return nil, db.Error
	}
	return u, nil
}

func NewRepo(db *gorm.DB) Repository {
	return &repo{
		db: db,
		tx: nil,
	}
}

func (r *repo) SetTransaction(tx *gorm.DB) {
	r.tx = tx
}

func (r *repo) Update(u *User) (*User, error) {
	db := r.tx.Model(u).Updates(User{
		Name: u.Name,
	})

	if db.Error != nil {
		return nil, db.Error
	}

	return u, nil
}

func (r *repo) Delete(id int) *gorm.DB {
	return r.tx.Delete(&User{}, uint(id))
}
