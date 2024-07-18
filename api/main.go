package main

import (
	"log"
	"strings"
	"tango_cli/pkg/cmdrunner"

	"tango_pkg/tango_helpers"
	"tango_pkg/tango_log"
	"tango_pkg/tango_middlewares"
	"tango_pkg/tango_routes"
	"tango_pkg/tangoapp"

	app "tango_api/app"
)

var rootPath string

func init() {
	cmdRuner := cmdrunner.New()
	rootPath = cmdRuner.GetRootPath()
	if !strings.Contains(rootPath+"/api", "/api/api") {
		rootPath += "/api"
	}
	tango_log.LogPrefix = "TANGO"
	tango_log.Print("Starting up")
}

func main() {
	tapp := tangoapp.NewTangoApp(rootPath)
	err := tapp.DB.Connect("local")
	if err != nil {
		log.Fatal(err)
	}

	tapp.PrintAppInfo()

	// Middleware
	tango_middlewares.Setup(tapp)

	//  Tango Routes
	if tapp.Config.SetupEnabled && tapp.Config.NotInProduction {
		tango_routes.SetupRoutes(tapp)
	}

	tango_routes.SetupStaticRoutes(tapp)

	// App routes
	app.AppSetup(tapp)

	// open app in default browser
	if tapp.Config.OpenInBrowser {
		tango_helpers.OpenInBrowser("http://" + tapp.GetAppUrl())
	}

	// Start server
	tapp.Server.Logger.Fatal(tapp.Server.Start(":" + tapp.GetPortAsStr()))

}
