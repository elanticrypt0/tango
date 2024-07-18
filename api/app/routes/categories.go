package routes

import (
	"tango_pkg/tangoapp"

	"tango_api/app/features"

	"github.com/labstack/echo/v4"
)

func categoriesRoutes(tapp *tangoapp.TangoApp, rootPath *echo.Group) {
	categories := rootPath.Group("/categories/")

	categories.GET(":id", func(ctx echo.Context) error {
		return features.FindOneCategory(ctx, tapp)
	})

	categories.GET("", func(ctx echo.Context) error {
		return features.FindAllCategories(ctx, tapp)
	})

	categories.POST("create", func(ctx echo.Context) error {
		return features.CreateCategory(ctx, tapp)
	})

	categories.PUT("update/:id", func(ctx echo.Context) error {
		return features.UpdateCategory(ctx, tapp)
	})

	categories.DELETE("delete/:id", func(ctx echo.Context) error {
		return features.DeleteCategory(ctx, tapp)
	})
}
