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
