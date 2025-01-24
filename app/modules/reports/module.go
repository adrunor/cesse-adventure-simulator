package reports

import (
	"simulator/app/modules/reports/controllers"
	"simulator/systems"
)

func SetupModule(context *systems.Context) {

	controllers.Setup(context)
}
