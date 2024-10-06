package routes

import (
	"tango_api/app/features"
	"tango_pkg/tangoapp"

	"github.com/labstack/echo/v4"
)

func BooksRoutes(tapp *tangoapp.TangoApp) {

	// this goes to the / dir
	// because is not parte of the API
	tapp.Server.GET("books", func(ctx echo.Context) error {
		return features.FindAllBooks(ctx, tapp)
	})

	tapp.Server.GET("books/first", func(ctx echo.Context) error {
		return features.FindFirstBook(ctx, tapp)
	})

	tapp.Server.GET("books/:id", func(ctx echo.Context) error {
		return features.FindOneBook(ctx, tapp)
	})

}
