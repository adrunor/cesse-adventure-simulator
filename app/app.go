package app

import (
	"simulator/app/modules/reports"
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

	reports.SetupModule(ctx)

	r.Run()
}
