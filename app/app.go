package app

import (
	"simulator/app/modules/level_simulation"
	"simulator/systems"
	"simulator/systems/database"
	"simulator/systems/env"
	"simulator/systems/migration"
	"simulator/systems/router"
)

func Run() {
	e := env.NewEnv()
	d := database.NewDatabase()
	r := router.NewRouter()
	ctx := systems.NewContext(e, d, r)

	migration.Migrate(d.Db)

	level_simulation.SetupModule(ctx)

	r.Run()
}
