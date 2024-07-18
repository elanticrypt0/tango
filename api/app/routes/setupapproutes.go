package routes

import (
	"tango_pkg/tangoapp"
)

func SetupAppRoutes(tapp *tangoapp.TangoApp) {
	urlApiRootPath := tapp.Server.Group("api")

	// Auth WIP
	// tango_auth.AuthRoutes(tapp, rootPath)
	// tango_auth.UsersRoutes(tapp, rootPath)

	// books example
	BooksRoutes(tapp)

	// categories example
	categoriesRoutes(tapp, urlApiRootPath)

}
