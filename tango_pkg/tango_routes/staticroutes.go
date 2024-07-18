package tango_routes

import (
	"strings"

	"tango_pkg/tangoapp"
)

func SetupStaticRoutes(tapp *tangoapp.TangoApp) {

	publicPath := tapp.GetRootPath() + "/" + strings.Replace(tapp.Config.PublicPath, ".", "", 1)
	publicAssetsPath := tapp.GetRootPath() + "/" + strings.Replace(tapp.Config.PublicAssetsPath, ".", "", 1)

	tapp.Server.Static("/", publicPath)
	tapp.Server.Static("/public", publicPath)
	tapp.Server.Static("/assets", publicAssetsPath)
	tapp.Server.Static("/assets/js", publicAssetsPath+"/js")
	tapp.Server.Static("/assets/css", publicAssetsPath+"/css")
	tapp.Server.Static("/images", publicAssetsPath+"/images")

}
