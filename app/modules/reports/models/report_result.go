package models

import "gorm.io/gorm"

type ReportResult struct {
	gorm.Model
	Level                uint8
	RequiredExperience   uint16
	CumulativeExperience uint16
	CombatExperience     uint16
	CombatCounter        uint16
	LevelingTime         uint32
}
