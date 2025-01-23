package app

import (
	"simulator/app/modules/reports"
	"simulator/systems"
	"simulator/systems/database"
	"simulator/systems/env"
	"simulator/systems/router"
)

func Run() {
	context := systems.NewContext(
		env.NewEnv(),
		database.NewDatabase(),
		router.NewRouter(),
	)

	reports.SetupModule(context)

	context.Router().Run()
}
