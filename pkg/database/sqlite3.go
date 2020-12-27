package database

import (
	"gorm.io/gorm"

	_ "github.com/mattn/go-sqlite3"
)

func (db *database) Create(value interface{}) error {
	if err := db.tx.Create(value).Error; err != nil {
		return err
	}
	return nil
}

func (db *database) getDb() *gorm.DB {
	return db.tx
}

func (db *database) Find(id int, i interface{}) error {
	tx := db.tx.First(i, uint(id))
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (db *database) SetTransaction(tx *gorm.DB) {
	db.tx = tx
}

func (db *database) Update(i interface{}) error {
	tx := db.tx.Model(i).Updates(i)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (db *database) Delete(i interface{}, id int) *gorm.DB {
	return db.tx.Delete(i, uint(id))
}
