package cmdrunner

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"tango_pkg/tangoapp"
)

// This run OS commands

type CmdRunner struct {
	rootPath string
}

func New() CmdRunner {
	instance := CmdRunner{}
	instance.SetRootPath(instance.PWD())
	return instance
}

func (cr *CmdRunner) SetRootPath(rpath string) {
	cr.rootPath = rpath
}

func (cr *CmdRunner) SetRootPathWithPWD(rpath string) {
	cr.rootPath = cr.PWD()
}

func (cr *CmdRunner) AppendToRootPath(path string) string {
	// cr.rootPath = cr.rootPath + path
	return cr.rootPath + path
}

func (cr *CmdRunner) GetRootPath() string {
	return cr.rootPath
}

func (cr *CmdRunner) RunLines(lines []string) {
	// chequea primero si el comando está dentro de los comandos espaciales
	// sino está lo ejecuta
	for _, command := range lines {
		command_splitted := strings.Split(command, " ")
		command_name := command_splitted[0]
		command_args := command_splitted[1:]
		if !cr.isSpecialCmdAndExecute(command_name, command_args) {
			cr.RunSliceArgs(command_name, command_args)
		}
	}

}

func (cr *CmdRunner) isSpecialCmdAndExecute(cmd string, args []string) bool {
	isSpecial := true

	switch cmd {
	case "mkdir":
		cr.Mkdir(args[0])
	case "cd":
		cr.Cd(args[0])
	case "echo":
		cr.Echo(strings.Join(args, " "))
	case "pwd":
		cr.Run("pwd")
	default:
		isSpecial = false
	}
	return isSpecial

}

func (cr *CmdRunner) RunSliceArgs(name string, arg []string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Printf("> %s\n", cmd.String())
	return cmd.Run()
}

func (cr *CmdRunner) Run(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (cr *CmdRunner) RunWithOutput(name string, arg ...string) CmdOutput {
	cmd := exec.Command(name, arg...)
	// capture the output or error
	var outb, errb bytes.Buffer
	cmd.Stdin = os.Stdin
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
	return NewCmdOutput(outb.String(), errb.String())
}

func (cr *CmdRunner) RunSh(shpath string, arg ...string) CmdOutput {
	cmd := exec.Command(shpath, arg...)
	// capture the output or error
	var outb, errb bytes.Buffer
	cmd.Stdin = os.Stdin
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
	return NewCmdOutput(outb.String(), errb.String())
}

func (cr *CmdRunner) PWD() string {
	// output := cr.RunWithOutput("pwd")
	// pwd := strings.ReplaceAll(output.Output, "\n", "")
	pwd, _ := os.Getwd()
	return pwd
}

func (cr *CmdRunner) Mkdir(newDir string) error {
	// checks is a directory exists
	// if is not then create
	_, err := os.Stat(newDir)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		err = cr.Run("mkdir", newDir)
		if err != nil {
			return err
		}
	}

	return nil
}

func (cr *CmdRunner) CopyAll(sourcepath, destinypath string) error {
	_, err := os.Stat(sourcepath)
	if err == nil {
		return nil
	}

	if os.IsNotExist(err) {
		err = cr.Run("cp", "-rf", sourcepath, destinypath)
		if err != nil {
			return err
		}
	}

	return nil
}

func (cr *CmdRunner) Cd(dirpath string) error {
	// checks is a directory exists
	// if is not then create
	dirpath = cr.rootPath + "/" + dirpath
	_, err := os.Stat(dirpath)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		err = cr.Run("cd", dirpath)
		if err != nil {
			return err
		}
	}

	return nil
}

func (cr *CmdRunner) Echo(msg string) {
	fmt.Println(msg)
}

func (cr *CmdRunner) LoadAppConfig() *tangoapp.AppConfig {
	configPath := cr.AppendToRootPath("/api/config/app.toml")
	appConfig := &tangoapp.AppConfig{}
	LoadTomlFile(configPath, appConfig)
	return appConfig
}
