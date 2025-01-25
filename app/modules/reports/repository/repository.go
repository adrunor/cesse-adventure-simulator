package repository

import (
	"simulator/app/modules/reports/models"
	"simulator/systems"
)

type ReportConfigurationRepository interface {
	Create(configuration *models.ReportConfiguration) (interface{}, error)
	Update(configuration *models.ReportConfiguration) (interface{}, error)
	Delete(configuration *models.ReportConfiguration) (interface{}, error)
}

type ReportConfigurationRepositoryImpl struct {
	Database systems.Database
}

func (d *ReportConfigurationRepositoryImpl) Create(configuration *models.ReportConfiguration) (interface{}, error) {
	return d.Database.Create(configuration)
}

func (d *ReportConfigurationRepositoryImpl) Update(configuration *models.ReportConfiguration) (interface{}, error) {
	return d.Database.Update(configuration)
}

func (d *ReportConfigurationRepositoryImpl) Delete(configuration *models.ReportConfiguration) (interface{}, error) {
	return d.Database.Delete(configuration)
}
