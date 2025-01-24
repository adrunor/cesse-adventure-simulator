package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type Impl struct {
	Db *gorm.DB
}

func NewDatabase() *Impl {
	return setUpDatabase()
}

func setUpDatabase() *Impl {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return &Impl{db}
}

func (db *Impl) Create(model interface{}) (interface{}, error) {
	result := db.Db.Create(model)
	return result.RowsAffected, result.Error
}

func (db *Impl) Update(model interface{}) (interface{}, error) {
	return nil, nil
}

func (db *Impl) Delete(model interface{}) (interface{}, error) {
	return nil, nil
}
