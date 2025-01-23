package models

import (
	"gorm.io/gorm"
	"time"
)

type ReportConfiguration struct {
	gorm.Model
	MaxLevel               uint8 `gorm:"<-create;uniqueIndex:idx_max_level"`
	DurationToMaxLevel     time.Duration
	BattleDurationAvg      time.Duration
	ExplorationDurationAvg time.Duration
}
