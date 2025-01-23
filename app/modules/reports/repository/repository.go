package repository

import "simulator/systems/db"

type ReportRepository interface {
}

type ReportRepositoryImpl struct {
	db.Database
}
