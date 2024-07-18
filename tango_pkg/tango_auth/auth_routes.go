package tango_auth

import (
	"tango_pkg/tangoapp"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(tapp *tangoapp.TangoApp, rootPath *echo.Group) {
	users := rootPath.Group("/auth/")

	users.POST("login", func(ctx echo.Context) error {
		return AuthLogin(ctx, tapp)
	})

	users.GET("logout", func(ctx echo.Context) error {
		return AuthLogout(ctx, tapp)
	})
}
