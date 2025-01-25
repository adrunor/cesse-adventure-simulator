package migration

import (
	"gorm.io/gorm"
	"log"
	"simulator/app/modules/reports/models"
)

func Migrate(db *gorm.DB) {
	if err := db.AutoMigrate(
		models.ReportConfiguration{},
		models.ReportResult{},
	); err != nil {
		log.Fatal(err)
	}
	log.Println("Migration done.")
}
