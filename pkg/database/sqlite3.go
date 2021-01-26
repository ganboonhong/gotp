package database

import (
	"gorm.io/gorm"

	_ "github.com/mattn/go-sqlite3"
)

func (repo *Repo) Create(value interface{}) error {
	gormDB := repo.DB
	if err := gormDB.Create(value).Error; err != nil {
		return err
	}
	return nil
}

func (repo *Repo) Find(ID int, i interface{}) error {
	gormDB := repo.DB
	if err := gormDB.First(i, uint(ID)).Error; err != nil {
		return err
	}
	return nil
}

func (repo *Repo) Update(i interface{}) error {
	gormDB := repo.DB
	if err := gormDB.Model(i).Updates(i).Error; err != nil {
		return err
	}
	return nil
}

func (repo *Repo) Delete(i interface{}, ID int) *gorm.DB {
	gormDB := repo.DB
	return gormDB.Delete(i, uint(ID))
}
