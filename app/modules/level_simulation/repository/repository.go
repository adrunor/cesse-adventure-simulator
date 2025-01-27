package repository

import (
	"simulator/app/modules/level_simulation/models"
	"simulator/systems"
)

type LevelConfigRepository interface {
	Create(configuration *models.LevelConfig) (interface{}, error)
	Update(configuration *models.LevelConfig) (interface{}, error)
	Delete(configuration *models.LevelConfig) (interface{}, error)
	Find(config *models.LevelConfig, id int) error
}

type LevelConfigRepositoryImpl struct {
	Database systems.Database
}

func (d *LevelConfigRepositoryImpl) Create(config *models.LevelConfig) (interface{}, error) {
	return d.Database.Create(config)
}

func (d *LevelConfigRepositoryImpl) Update(config *models.LevelConfig) (interface{}, error) {
	return d.Database.Update(config)
}

func (d *LevelConfigRepositoryImpl) Delete(config *models.LevelConfig) (interface{}, error) {
	return d.Database.Delete(config)
}

func (d *LevelConfigRepositoryImpl) Find(config *models.LevelConfig, id int) error {
	return d.Database.Find(config, id)
}
