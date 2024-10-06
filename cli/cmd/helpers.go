package cmd

import (
	"fmt"
	"tango_cli/pkg/cmdrunner"
)

func RenderTemplSh(cmdrunner cmdrunner.CmdRunner, rootpath string) {
	fmt.Println("Renderizando...")
	response := cmdrunner.RunSh(rootpath + "/sh/templ_render.sh")
	fmt.Printf("%q\n", response)
}
