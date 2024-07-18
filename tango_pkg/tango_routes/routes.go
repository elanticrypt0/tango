package tango_routes

import (
	"tango_pkg/tangoapp"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(tapp *tangoapp.TangoApp) {

	// setup
	setup := tapp.Server.Group("/setup")

	setup.GET("/", func(c echo.Context) error {
		return Setup(c, tapp)
	})

	//status
	setup.GET("/status", func(c echo.Context) error {
		return Status(c)
	})

}
