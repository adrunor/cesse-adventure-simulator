package repository

import (
	"simulator/systems"
)

type ImplRepository struct {
	database *systems.Database
}

func (d *ImplRepository) Create(model interface{}) (interface{}, error) {
	return d.database.Create(model)
}
