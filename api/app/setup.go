package app

import (
	"tango_pkg/tangoapp"

	"tango_api/app/routes"
)

func AppSetup(tapp *tangoapp.TangoApp) {

	routes.SetupAppRoutes(tapp)
}
