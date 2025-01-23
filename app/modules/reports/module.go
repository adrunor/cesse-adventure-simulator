package reports

import (
	"simulator/app/modules/reports/controllers"
	"simulator/systems"
)

func SetupModule(context *systems.Context) {
	api := context.Router().GetRouter().Group("/api")

	controllers.AddRoutes(api)

}
