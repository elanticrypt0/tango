// task.go
//go:build ignore

package main

import (
	"log"
	"os"
	"os/exec"
)

const BINARY_NAME = "app_api"
const BINARY_NAME_WIN = "app_api.exe"

const BUILD_DIR = "../build"

const BUILD_DIR_LINUX = "../build/linux"
const BUILD_DIR_LINUXARM64 = "../build/arm64"
const BUILD_DIR_WIN = "../build/windows"

func cmdRun(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Printf("$ %s\n", cmd.String())
	return cmd.Run()
}

func cmdMkdir(newDir string) error {
	// checks is a directory exists
	// if is not then create
	_, err := os.Stat(newDir)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		err = cmdRun("mkdir", newDir)
		if err != nil {
			return err
		}
	}

	return nil
}

func cmdRmDir(oldDir string) error {
	// checks is a directory exists
	// if is not then create
	_, err := os.Stat(oldDir)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		err = cmdRun("rm -rf", oldDir)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	log.SetFlags(0)
	var taskName string
	if len(os.Args) >= 2 {
		taskName = os.Args[1]
	} else {
		log.Fatalln("no task")
	}
	task, ok := map[string]func() error{
		"install":     Install,
		"test":        Test,
		"cli_install": CliInstall,
		"build":       Build,
		"build_arm":   BuildArm64,
		"build_x64":   BuildX64,

		// Add more tasks here!
	}[taskName]
	if !ok {
		log.Fatalln("no such task")
	}
	err := task()
	if err != nil {
		log.Fatalln(err)
	}
}

// Tasks comes here

func Install() error {

	err := cmdMkdir("_db")
	if err != nil {
		return err
	}

	CliInstall()

	err = cmdRun("go", "mod", "tidy")
	if err != nil {
		return err
	}

	return nil
}

func Test() error {
	err := cmdRun("go", "test", "./tests")
	if err != nil {
		return err
	}

	return nil
}

func Build() error {
	err := cmdMkdir(BUILD_DIR_LINUX)
	if err != nil {
		return err
	}
	return nil
}

func BuildX64() error {

	err := cmdMkdir(BUILD_DIR_LINUX)
	if err != nil {
		return err
	}
	return nil
}

func BuildArm64() error {
	err := cmdMkdir(BUILD_DIR_LINUXARM64)
	if err != nil {
		return err
	}
	return nil
}

func CliInstall() error {
	err := cmdRun("git", "clone", "https://github.com/k23dev/tango_cli", "cli")
	if err != nil {
		return err
	}
	return nil
}

func clean() error {
	err := cmdRun("go", "clean")
	if err != nil {
		return err
	}

	err = cmdRmDir(BUILD_DIR)
	if err != nil {
		return err
	}

	return nil
}
