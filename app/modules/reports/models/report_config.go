package models

import (
	"gorm.io/gorm"
	"time"
)

type ReportConfiguration struct {
	gorm.Model
	BaseExperience         uint16        `json:"primary_key"`
	IncreaseRate           uint16        `json:"increase_rate"`
	MaxLevel               uint8         `json:"max_level"`
	CombatExperienceRate   uint16        `json:"combat_experience_rate"`
	CombatDurationAvg      time.Duration `json:"combat_duration_avg"`
	ExplorationDurationAvg time.Duration `json:"exploration_duration_avg"`
}
