package cmd

import (
	"fmt"
	"log"
	"strings"
	"tango_cli/pkg/cmdrunner"
)

const buildDir = "/builds/"

func RemakeBuild(rootPath, appnamePlusVersion string) {
	cr := cmdrunner.New()
	buildpath := rootPath + buildDir + appnamePlusVersion
	cr.Run("rm", "-rf", buildpath)
}

func MakeBuild(appnamePlusVersion string) {
	// fmt.Printf("%+v\n", appconfig)

	rootPath := cmdRunner.GetRootPath()

	fmt.Printf("%+v\n", appnamePlusVersion)
	// tiene que crear la carpeta build
	buildpath := rootPath + buildDir
	cmdRunner.Mkdir(buildpath)
	// luego crear una carpeta con el nombre de app + version (app_version)
	buildpath += appnamePlusVersion
	cmdRunner.Mkdir(buildpath)
	// logs, configuracion, public,cookies,uploads, _db
	cmdRunner.Mkdir(buildpath + "/logs")
	fmt.Printf("	> Creando la carpeta %s\n", "/logs")
	cmdRunner.Mkdir(buildpath + "/_db")
	fmt.Printf("	> Creando la carpeta %s\n", "/_db")
	cmdRunner.Mkdir(buildpath + "/cookies")
	fmt.Printf("	> Creando la carpeta %s\n", "/cookies")
	cmdRunner.Mkdir(buildpath + "/uploads")
	fmt.Printf("	> Creando la carpeta %s\n", "/uploads")
	cmdRunner.Mkdir(buildpath + "/config")
	fmt.Printf("	> Creando la carpeta %s\n", "/config")
	cmdRunner.Mkdir(buildpath + "/public")
	fmt.Printf("	> Creando la carpeta %s\n", "/public")
	// err := cmdRunner.CopyAll()
	err := cmdrunner.CopyDirectory(rootPath+"/api/config", buildpath+"/config")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("	> Copiando la carpeta %s\n", "/config")
	// err = cmdRunner.CopyAll(rootPath+"/api/public", buildpath+"/public")
	err = cmdrunner.CopyDirectory(rootPath+"/api/public", buildpath+"/public")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("	> Copiando la carpeta %s\n", "/public")
	// crear el ejecutable
	// buildExe(cmdRunner.AppendToRootPath("/api"),appnamePlusVersion,"linux")
	BuildLinux64Exe(cmdRunner.AppendToRootPath("/api/"), appnamePlusVersion)
	// dentro el ejecutable con el mismo nombre
}

func BuildLinux64Exe(path, name string) {
	buildExe(path, name, "linux", "amd64")
}

func BuildWindows64Exe(path, name string) {
	buildExe(path, name, "windows", "amd64")
}

func buildExe(path, name, platform, arch string) {
	commandZero := fmt.Sprintf("GOOS=%s GOARCH=%s go", platform, arch)
	commandArgs := fmt.Sprintf("build %s -o %s -gccgoflags \"-w -s\"", path, name)
	commandSplited := strings.Split(commandArgs, " ")
	cr := cmdrunner.New()
	cr.RunSliceArgs(commandZero, commandSplited)

}
