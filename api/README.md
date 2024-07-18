# Tango

Tango es un SDK que incluye Go + Templ + Htmlx
Ideal para crear APIs

## Diseño RFM (Routes Features Models)

### Make automatically Route, Features, Models and views

To make automatically this execute the tango_cli. Is easy

    tango_cli [PACKAGE_NAME] [MODE]

e.g.

    tango_cli photo basic

Watch the full documentation in: https://github.com/k23dev/tango_cli

### Routes

Las rutas se definen en la carpeta routes dentro del archivo setuproutes allí se agregan las funciones. Para poder acceder a las funcionalidades de Fiber o Gorm o la configuración de la app se debe pasar una variable

    tapp *tangoapp.TangoApp

### Features

Las "features" son los controladores de las rutas y reciben 2 parámetros: 1) el contexto de Fiber y otra variable de *tangoapp.TangoApp para poder acceder a la db y configuración.

### Models

Los modelos deben recibir también la variable que apunte a *tangoapp.TangoApp

## Public

los archivos publicos están dentro de la carpeta public y por defecto tienen varios accesos. Esto se puede ver en webcore_features/routes.go

# Makefile

## Instalar todas las depencias de tango

Este comando instala las dependencias de go y tailwind

    make tango-install

## Instalar dependencias de go

Para instalar las dependecias de go y air para poder tener hotreload ejecutar

    make go-install-deps

## Instalar dependencias de tailwind

Para instalar las dependecias de go y air para poder tener hotreload ejecutar

    make tailwind-install

## Run in dev mode

Ejecutar

    make dev

## Run tailwind in dev mode

    make tailwind-dev

## Generar vistas

El sistema de vistas se realiza con la lib templ. Es una de las dependencias que se instalan al con el comando _make deps_ de todas formas y se ejecuta antes de hacer un build, o cuando se ejecuta _make dev_.
Pero también puede hacerse manualmente con

    make templates

## Generar un build

Para genenerar un build ejecutar el comando

    make build

## DEBUG

Hay un paquete de debug junto a los errores manejados por tango.

Una función útil para debuguear es:

    tango_errors.Debug("nombre",fmt.Sprintf("%v",variable))

## HTMLX

Para instalar la librería htmlx sólo hay que ejecutar

    make htmlx-install

y para actualizar la librería

    make htmlx-update

# Auth

TODO

Librería de autenticación
