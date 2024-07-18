package template_route

func RouteAPI() string {

	template := `
package routes

import (
	"tango_api/app/features"
	"tango_pkg/tangoapp"
	"github.com/labstack/echo/v4"
)

func $PL$Routes(tapp *tangoapp.TangoApp, rootPath *echo.Group) {
	$PL$ := rootPath.Group("/$PL$/")
	feat:=features.New$PC$Feature(tapp)

	$PL$.GET(":id", func(ctx echo.Context) error {
		feat.SetCtx(ctx)
		return feat.FindOne()
	})

	$PL$.GET("", func(ctx echo.Context) error {
		feat.SetCtx(ctx)
		return feat.FindAll()
	})

	$PL$.POST("create", func(ctx echo.Context) error {
		feat.SetCtx(ctx)
		return feat.Create()
	})

	$PL$.PUT("update/:id", func(ctx echo.Context) error {
		feat.SetCtx(ctx)
		return feat.Update()
	})

	$PL$.DELETE("delete/:id", func(ctx echo.Context) error {
		feat.SetCtx(ctx)
		return feat.Delete()
	})
}
	`
	return template
}
