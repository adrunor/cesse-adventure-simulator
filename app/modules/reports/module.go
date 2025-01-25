package reports

import (
	"simulator/app/modules/reports/controllers"
	"simulator/app/modules/reports/repository"
	"simulator/systems"
)

type ReportModule struct {
	Repository repository.ReportConfigurationRepository
}

func SetupModule(context *systems.Context) {
	module := ReportModule{
		&repository.ReportConfigurationRepositoryImpl{Database: context.Database()},
	}

	controllers.Setup(context.Router(), module.Repository)
}
