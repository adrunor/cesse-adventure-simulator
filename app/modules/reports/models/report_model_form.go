package models

import "time"

type ReportConfigurationForm struct {
	MaxLevel               uint8         `json:"maxLevel"`
	DurationToMaxLevel     time.Duration `json:"maxTimeToMaxLevel"`
	BattleDurationAvg      time.Duration `json:"battleDurationAvg"`
	ExplorationDurationAvg time.Duration `json:"explorationDurationAvg"`
}
