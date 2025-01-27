package models

import "gorm.io/gorm"

type LevelResult struct {
	gorm.Model
	Level                int `json:"level"`
	RequiredExperience   int `json:"required_experience"`
	CumulativeExperience int `json:"cumulative_experience"`
	CombatExperience     int `json:"combat_experience"`
	CombatCounter        int `json:"combat_counter"`
	LevelingTime         int `json:"leveling_time"`
}
