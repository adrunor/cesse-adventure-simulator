package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type DatabaseImpl struct {
	Db *gorm.DB
}

func NewDatabase() *DatabaseImpl {
	return setUpDatabase()
}

func setUpDatabase() *DatabaseImpl {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return &DatabaseImpl{db}
}

func (db *DatabaseImpl) Create(model interface{}) (interface{}, error) {
	result := db.Db.Create(model)
	return result.RowsAffected, result.Error
}

func (db *DatabaseImpl) Update(model interface{}) (interface{}, error) {
	return nil, nil
}

func (db *DatabaseImpl) Delete(model interface{}) (interface{}, error) {
	return nil, nil
}
