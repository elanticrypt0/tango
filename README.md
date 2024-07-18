# Tango
## Framework to build web applications with go.

Tango is a framwork to build web applications with go fast and easy. Tango is ispired in Laravel, but with a go mentality.

Tango uses almost and MVC but in Tango the view is optional. You can choose use Tango as a api backend or use templ to generate the views.

Tango have  CLI to make build fast CRUDs super easy.

# Let's go!

First copy the repo:

```bash
git clone https://github.com/elanticrypt0/tango
```

## Get Packages

Before start you need get the packages to work

### Linux and Mac

Just run

```bash
./sh/get_pkgs.sh
```

### Windows

Just run

```bash
cd ./cli 
go mod tidy
cd ../api
go mod tidy
cd ../tango_pkg
go mod tidy
```

# Config

Your tango app's config by default is in: **api/config** there is two files:

-app.toml
-db.toml

## Default app URL

The apps default port is 9000 and the features works like API.

http://localhost:9000/api/categories/

## app.toml

There is your tango app configs like your apps title, version and if use astro, templ or both

## db.toml

There is your database(s) configurations files. By default your apps connect to sqlite but you can use:

- MySQL
- Postgres
- Sqlite

If you need other one download the package [dbman](https://github.com/elanticrypt0/dbman) and change it as you need.

# Use the CLI

Use the CLI to run your application

## Run dev

```bash
go run ./cli dev
```

## Create CRUD

This command create a Feature, Model, And Routes.

```bash
go run ./cli createapicrud [NAME]
```

The only thing that you need to write is the setup of your routes.

### Routes Setup

Edit **./api/app/routes/setupapproutes.go**

Add this lines

```go
[NAME]Routes(tapp, urlApiRootPath)
```

Now you can access this at: http://localhost:9000/api/[NAME]/

If you use templ like a regular app just edit the file like this:

```go
[NAME]Routes(tapp)
```

The access this at: http://localhost:9000/[NAME]/

If you do this you need edit the router file.


## Create httpclient for your frontend

This command create a httpclient to make request to your tango API. The file is created in: **./frontend/src/_tangoclient.ts**

```bash
go run ./cli httpclient
```

# Migrations (WIP)

# Authentications (WIP)

# Build your app (WIP)

Work in progress. So now you must make this in the old way:

```bash
cd api && go build -o your_api
```

And then create the copy the folders:

- config
- _db (if you're using sqlite)
- public