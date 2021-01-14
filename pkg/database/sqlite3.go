package database

import (
	"gorm.io/gorm"

	_ "github.com/mattn/go-sqlite3"
)

func (DB *DB) Create(value interface{}) error {
	gormDB := DB.DB
	if err := gormDB.Create(value).Error; err != nil {
		return err
	}
	return nil
}

func (DB *DB) Find(ID int, i interface{}) error {
	gormDB := DB.DB
	if err := gormDB.First(i, uint(ID)).Error; err != nil {
		return err
	}
	return nil
}

func (DB *DB) Update(i interface{}) error {
	gormDB := DB.DB
	if err := gormDB.Model(i).Updates(i).Error; err != nil {
		return err
	}
	return nil
}

func (DB *DB) Delete(i interface{}, ID int) *gorm.DB {
	gormDB := DB.DB
	return gormDB.Delete(i, uint(ID))
}
