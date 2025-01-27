package models

import (
	"gorm.io/gorm"
	"time"
)

type LevelConfig struct {
	gorm.Model
	MaxLevel               int           `json:"max_level" validator:"required,numeric"`
	BaseExperience         int           `json:"base_experience" validator:"required,numeric"`
	IncreaseRate           float32       `json:"increase_rate" validator:"required,numeric"`
	CombatExperienceRate   float32       `json:"combat_experience_rate" validator:"required,numeric"`
	CombatDurationAvg      time.Duration `json:"combat_duration_avg" validator:"required,numeric"`
	ExplorationDurationAvg time.Duration `json:"exploration_duration_avg" validator:"required,numeric"`
}
