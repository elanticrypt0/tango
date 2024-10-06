package migrations

import (
	"tango_pkg/tangoapp"

	"tango_api/app/models"
)

func Migrate() {

}

func Setup(tapp *tangoapp.TangoApp) {

	// is DebugMode == true
	// migrate tables
	if tapp.Config.NotInProduction {
		tapp.DB.Primary.AutoMigrate(&models.Category{})
		// migrateAuth(tapp.DBAuth)
	}

}

// func migrateAuth(dbAuth *gorm.DB) {
// 	// migrate auth
// 	dbAuth.AutoMigrate(&tango_auth.User{}, &tango_auth.Auth{})
// }
