package migrations

import (
	"log"
	"net/http"

	"tango_pkg/tango_helpers"

	"tango_api/app/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const seedDir = "./migrations/seeds/"

func Seed(c echo.Context, db *gorm.DB) error {
	seedCategories(db)
	return c.JSON(http.StatusOK, "OK")
}

func seedCategories(db *gorm.DB) {
	cat_list := []models.Category{}
	tango_helpers.ReadAndParseJson(seedDir+"categories.json", &cat_list)

	db.Save(&cat_list)
	log.Println("Categories seeded")
}
