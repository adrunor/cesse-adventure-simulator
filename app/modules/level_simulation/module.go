package level_simulation

import (
	"simulator/app/modules/level_simulation/controllers"
	"simulator/app/modules/level_simulation/repository"
	"simulator/systems"
)

type ReportModule struct {
	Repository repository.LevelConfigRepository
}

func SetupModule(context *systems.Context) {
	module := ReportModule{
		&repository.LevelConfigRepositoryImpl{Database: context.Database()},
	}

	controllers.Setup(context.Router(), module.Repository)
}
