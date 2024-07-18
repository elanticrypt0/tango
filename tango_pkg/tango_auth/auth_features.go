package tango_auth

import (
	"net/http"

	"tango_pkg/tangoapp"

	"github.com/labstack/echo/v4"
)

func AuthLogin(ctx echo.Context, tapp *tangoapp.TangoApp) error {

	authDTO := AuthDTOLogin{}
	if err := ctx.Bind(&authDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, "")
	}

	auth := NewAuth()

	user, err := auth.Login(tapp.DBAuth, &authDTO)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	// generate a JWT
	token, err := GenerateTokensAndSetCookies(user, ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	user.Token = token

	return ctx.JSON(http.StatusOK, user)
}

func AuthCheck(ctx echo.Context, tapp *tangoapp.TangoApp) error {
	return nil
}

func AuthLogout(ctx echo.Context, tapp *tangoapp.TangoApp) error {
	return nil
}
