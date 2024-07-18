/*
Copyright © 2024 NAME HERE elanticrypt0@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"tango_cli/pkg/cmdrunner"
	"tango_cli/pkg/filemaker"
	"tango_cli/pkg/parser"

	"github.com/spf13/cobra"
)

// the API folder
const APIPATH = "./api/"
const FRONTENDROOTPATH = "/frontend"

var cmdRunner = cmdrunner.New()

// func init() {
// cmdRunner.AppendToRootPath(APIPATH)
// }

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tango_cli",
	Short: "CLI to create CRUD or make a build for Tango",
	Long:  `CLI to create CRUD or make a build for Tango`,
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creación de archivos individuales de features, models y routes",
	Long:  `Crear features, models, views, Api`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 0 {
		}
	},
}

var createPackApiCmd = &cobra.Command{
	Use:   "createapicrud",
	Short: "Creación de archivos paquetes de features, models y routes",
	Long:  `Creación de archivos paquetes de features, models y routes`,
	Run: func(cmd *cobra.Command, args []string) {

		p := parser.New()
		var fm filemaker.FileMaker
		var namespace string
		var templateSelected string = "api"

		if len(args) > 0 {
			namespace = args[0]
		}

		p.Read(namespace)
		fm = filemaker.New(cmdRunner.AppendToRootPath("/api"), "app", *p)
		fm.SelectTemplate(templateSelected)
		// forcemode=true will delete the files if exists
		fm.SetForceMode(true)

		// Creación
		fmt.Println("Making: ", os.Args[1])
		fmt.Println("Mode: ", os.Args[2])
		fmt.Println("Execuit it!")
		fm.MakeIt()

	},
}

var createHttpClient = &cobra.Command{
	Use:   "httpclient",
	Short: "Crea una clase para realizar las peticiones ajax al servidor.",
	Long:  `Crea una clase para realizar las peticiones ajax al servidor.`,
	Run: func(cmd *cobra.Command, args []string) {

		p := parser.New()
		var fm filemaker.FileMaker

		namespace := "_tangoclient"
		templateSelected := "httpclient"

		p.Read(namespace)
		fm = filemaker.New(cmdRunner.GetRootPath(), "frontend", *p)
		fm.SelectTemplate(templateSelected)
		// forcemode=true will delete the files if exists
		fm.SetForceMode(true)

		// Creación
		fmt.Println("Making: ", namespace)
		fmt.Println("Mode: ", "forcedMode = true")
		fmt.Println("Execuit it!")
		fm.MakeIt()

	},
}

var createModelCmd = &cobra.Command{
	Use:   "createmodel",
	Short: "Creación de archivo de modelo",
	Long:  `Creación de archivo de modelo`,
	Run: func(cmd *cobra.Command, args []string) {

		p := parser.New()
		var fm filemaker.FileMaker
		var namespace string
		var templateSelected string = "model"

		if len(args) > 0 {
			namespace = args[0]
		}

		p.Read(namespace)
		fm = filemaker.New(cmdRunner.AppendToRootPath("/api"), "app", *p)
		fm.SelectTemplate(templateSelected)
		// forcemode=true will delete the files if exists
		fm.SetForceMode(true)

		// Creación
		fmt.Println("Making: ", os.Args[1])
		fmt.Println("Mode: ", os.Args[2])
		fmt.Println("Execuit it!")
		fm.MakeIt()

	},
}

var showAppConfig = &cobra.Command{
	Use:   "appconfig",
	Short: "Muestra la configuracion de la app",
	Long:  `Muestra la configuracion de la app`,
	Run: func(cmd *cobra.Command, args []string) {

		appconfig := cmdRunner.LoadAppConfig()
		fmt.Printf("%+v\n", appconfig)

	},
}

var makeBuild = &cobra.Command{
	Use:   "build",
	Short: "Crea el build de la app",
	Long:  `Crea el build de la app`,
	Run: func(cmd *cobra.Command, args []string) {

		// TODO: agregar la opcion de compilar para diferentes plataformas.
		appconfig := cmdRunner.LoadAppConfig()
		appnamePlusVersion := parseAppNameAndVersion(appconfig.Name, appconfig.Version)

		MakeBuild(appnamePlusVersion)
	},
}

var remakeBuild = &cobra.Command{
	Use:   "rebuild",
	Short: "Recrea el build de la app",
	Long:  `Recrea el build de la app`,
	Run: func(cmd *cobra.Command, args []string) {

		// TODO: agregar la opcion de compilar para diferentes plataformas.
		appconfig := cmdRunner.LoadAppConfig()
		appnamePlusVersion := parseAppNameAndVersion(appconfig.Name, appconfig.Version)

		RemakeBuild(cmdRunner.GetRootPath(), appnamePlusVersion)
		MakeBuild(appnamePlusVersion)
		BuildLinux64Exe(cmdRunner.GetRootPath()+"/api", appnamePlusVersion)
	},
}

var unBuild = &cobra.Command{
	Use:   "unbuild",
	Short: "Elimina el build de la version",
	Long:  `Elimina el build de la version`,
	Run: func(cmd *cobra.Command, args []string) {

		// TODO: agregar la opcion de compilar para diferentes plataformas.
		appconfig := cmdRunner.LoadAppConfig()
		appnamePlusVersion := parseAppNameAndVersion(appconfig.Name, appconfig.Version)

		RemakeBuild(cmdRunner.GetRootPath(), appnamePlusVersion)
	},
}

var devApp = &cobra.Command{
	Use:   "dev",
	Short: "Start the application on DEV mode",
	Long:  `Start the application on DEV mode`,
	Run: func(cmd *cobra.Command, args []string) {

		var input string
		// TODO: agregar la opcion de compilar para diferentes plataformas.
		// appconfig := cmdRunner.LoadAppConfig()
		rootpath := cmdRunner.GetRootPath()
		// appconfig := cmdRunner.LoadAppConfig()
		// cmdRunner.Run("go", "run", rootpath+"/api")
		// go cmdRunner.Run("npm", "run", "astro", "dev")

		func() {
			// renderiza los templates TEMPL
			// RenderTemplSh(cmdRunner, rootpath)
			cmdRunner.Run("go", "run", rootpath+"/api")
			// si tienen un frontend en astro entonces lo levanta
			// if ...
			// go cmd
		}()
		fmt.Scanln(&input)
		// close command
	},
}

var AppRender = &cobra.Command{
	Use:   "render",
	Short: "Render the templ views",
	Long:  `Render the templ views`,
	Run: func(cmd *cobra.Command, args []string) {

		rootpath := cmdRunner.GetRootPath()
		appconfig := cmdRunner.LoadAppConfig()

		if appconfig.UseTempl {
			// TODO
			// TemplRender(rootpath)
			RenderTemplSh(cmdRunner, rootpath)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	appBanner("1.6.0")
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tango_cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(createPackApiCmd)
	rootCmd.AddCommand(createHttpClient)
	rootCmd.AddCommand(createModelCmd)
	rootCmd.AddCommand(showAppConfig)
	rootCmd.AddCommand(makeBuild)
	rootCmd.AddCommand(remakeBuild)
	rootCmd.AddCommand(unBuild)
	rootCmd.AddCommand(devApp)
	rootCmd.AddCommand(AppRender)
}
