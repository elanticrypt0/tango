# Setup

add this lines to the next files


/app/routes/setupapproutes.go
```
// Auth
tango_auth.AuthRoutes(tapp, rootPath)
tango_auth.UsersRoutes(tapp, rootPath)
```
/main.go
```
// OPTIONAL: connect to a diferent DB
// app_config.Connect2Db("your_DB")
// add the default conection to the auth package
app_config.DB.SetAuthDB(0)
```

app/setup.go
```
// migrate tables
if tapp.App.Config.App_debug_mode {
    tapp.App.DB.Primary.AutoMigrate(&models.Category{})
    // migrate auth
    tapp.App.DB.Auth.AutoMigrate(&tango_auth.User{}, &tango_auth.Auth{})
}
```

# JWT Authentication

change the key of the JWT in jwt_config.go