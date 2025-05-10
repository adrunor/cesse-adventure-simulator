package migration

import (
	"gorm.io/gorm"
	"log"
	"simulator/app/modules/level_simulation/models"
)

func Migrate(db *gorm.DB) {
	if err := db.AutoMigrate(
		models.LevelConfig{},
		models.LevelResult{},
	); err != nil {
		log.Fatal(err)
	}
	log.Println("Migration done.")
}
