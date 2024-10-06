package features

import (
	"log"
	"net/http"
	"tango_api/app/views"

	"tango_pkg/tango_view"
	"tango_pkg/tangoapp"

	"github.com/labstack/echo/v4"
)

var books = []interface{}{"the lord of the rings", "Nibola", "Let us C", "Golang for beginners"}

func FindAllBooks(ctx echo.Context, tapp *tangoapp.TangoApp) error {
	view := tango_view.New()
	view.Data.SetSlice("books", books)
	log.Println("DEBUG --- BEGIN SLICE ---")
	log.Printf("%v\n\n\n", books)
	log.Printf("%v\n\n\n", view.Data.GetSliceAsStrings("books"))
	log.Printf("%v\n\n\n", view.Data.GetAllSlices())
	log.Println("DEBUG --- END SLICE ---")
	view.SetComponent(views.BooksShowAll(view.Data))
	return view.Render(ctx, http.StatusOK)
}

func FindFirstBook(ctx echo.Context, tapp *tangoapp.TangoApp) error {
	view := tango_view.New()
	view.Data.Set("book", "The lord of the rings")
	view.SetComponent(views.BooksShowFirst(view.Data))
	return view.Render(ctx, http.StatusOK)
}

func FindOneBook(ctx echo.Context, tapp *tangoapp.TangoApp) error {
	idStr := ctx.Param("id")
	view := tango_view.New()
	view.Data.Set("id", idStr)
	view.Data.SetSlice("books", books)
	view.SetComponent(views.BooksShowOne(view.Data))
	return view.Render(ctx, http.StatusOK)
}
