package orm

import (
	"gorm.io/gorm"

	_ "github.com/mattn/go-sqlite3"
)

func (orm *ORM) Create(value interface{}) error {
	gormDB := orm.DB
	if err := gormDB.Create(value).Error; err != nil {
		return err
	}
	return nil
}

func (orm *ORM) Find(ID int, i interface{}) error {
	gormDB := orm.DB
	if err := gormDB.First(i, uint(ID)).Error; err != nil {
		return err
	}
	return nil
}

func (orm *ORM) Update(i interface{}) error {
	gormDB := orm.DB
	if err := gormDB.Model(i).Updates(i).Error; err != nil {
		return err
	}
	return nil
}

func (orm *ORM) Delete(i interface{}, ID int) *gorm.DB {
	gormDB := orm.DB
	return gormDB.Delete(i, uint(ID))
}
