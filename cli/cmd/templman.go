package cmd

import (
	"fmt"
	"log"
	"os"
	templmanager "tango_cli/pkg/templ_manager"
	"tango_pkg/tango_log"
)

func TemplRender(rootpath string) {
	// TODO
	// name, _ := os.Executable()
	name, _ := os.Getwd()
	log.Printf("%q", name)
	cmd := "templ"
	tango_log.LogPrefix = "CLI"
	tango_log.Print("Templ rendering...")
	templCommandArgs := templmanager.GetRenderArgs(rootpath)
	cmd2run := fmt.Sprintf("%s %s", cmd, templCommandArgs)
	tango_log.Print(cmd2run)
	output := cmdRunner.RunWithOutput(cmd, templCommandArgs)
	log.Printf("%s", output)
}
