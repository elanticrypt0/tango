package tango_auth

import (
	"tango_pkg/tangoapp"

	"github.com/labstack/echo/v4"
)

func UsersRoutes(tapp *tangoapp.TangoApp, rootPath *echo.Group) {
	users := rootPath.Group("/users/")

	// use jwt
	JWTMiddleware(users)

	users.GET(":id", func(ctx echo.Context) error {

		return FindOneUser(ctx, tapp)
	})

	users.GET("", func(ctx echo.Context) error {
		// token, _ := ctx.Cookie(tango_jwt.AccessTokenCookieName)
		// tango_jwt.GetTokenData(token)
		return FindAllUsers(ctx, tapp)
	})

	users.GET("activate/:code", func(ctx echo.Context) error {
		return ActivateUser(ctx, tapp)
	})

	users.POST("create", func(ctx echo.Context) error {
		return CreateUser(ctx, tapp)
	})

	users.PUT("update/:id", func(ctx echo.Context) error {
		return UpdateUser(ctx, tapp)
	})

	users.POST("change", func(ctx echo.Context) error {
		return ChangePasswordUser(ctx, tapp)
	})

	users.DELETE("delete/:id", func(ctx echo.Context) error {
		return DeleteUser(ctx, tapp)
	})
}
