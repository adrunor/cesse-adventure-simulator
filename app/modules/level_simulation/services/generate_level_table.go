package services

import "simulator/app/modules/level_simulation/models"

func GenerateLevelTable(config *models.LevelConfig) *[]models.LevelResult {
	var levels []models.LevelResult

	var level, experience, cumulativeExperience, combatExperience, combatCounter, levelingTime = 1, config.BaseExperience, 0, 0, 0, 0
	levels = append(levels, newLevelResult(level, experience, cumulativeExperience, combatExperience, combatCounter, levelingTime))

	for level = 2; level <= config.MaxLevel; level++ {
		experience = computeNextRequiredExperience(experience, config.IncreaseRate)

		levels = append(levels, newLevelResult(level, experience, cumulativeExperience, combatExperience, combatCounter, levelingTime))
	}

	return &levels
}

func computeNextRequiredExperience(experience int, increaseRate float32) int {
	newExperience := float32(experience) * increaseRate
	return int(newExperience)
}

func newLevelResult(level, requiredExperience, cumulativeExperience, combatExperience, combatCounter, levelingTime int) models.LevelResult {
	return models.LevelResult{
		Level:                level,
		RequiredExperience:   requiredExperience,
		CumulativeExperience: cumulativeExperience,
		CombatExperience:     combatExperience,
		CombatCounter:        combatCounter,
		LevelingTime:         levelingTime,
	}
}
