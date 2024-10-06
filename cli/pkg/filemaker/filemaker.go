package filemaker

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
	"tango_cli/pkg/parser"
	"tango_cli/pkg/templates"

	"tango_pkg/tangoapp"
)

type FileMaker struct {
	RootPath         string
	Parser           parser.Parser
	TemplateSelected string
	Workingdir       string
	forceMode        bool
	filePerms        fs.FileMode
	Appconfig        tangoapp.AppConfig
}

func New(rootpath, workingdir string, parser parser.Parser) FileMaker {

	return FileMaker{
		RootPath:   rootpath,
		Workingdir: workingdir,
		Parser:     parser,
		forceMode:  false,
		filePerms:  0666,
	}
}

func (fm *FileMaker) SetAppConfig(config tangoapp.AppConfig) {
	fm.Appconfig = config
}

func (fm *FileMaker) SelectTemplate(template string) {
	fm.TemplateSelected = strings.ToUpper(template)
}

func (fm *FileMaker) SetForceMode(mode bool) {
	fm.forceMode = mode
}

func (fm *FileMaker) MakeIt() {
	t := templates.New(&fm.Parser, fm.TemplateSelected)
	for _, current := range t.TemplateParsed.File {
		fm.builder(current.Dir, current.Extension, current.IsPlural, current.Data)
	}
}

func (fm *FileMaker) builder(directory, extension string, isPlural bool, fileData string) {
	filepath := fm.GetFilePath(directory, extension, isPlural)
	if fm.TemplateSelected != "" {
		// Creo el archivo
		fm.saveFile(filepath, fileData)
	} else {
		fmt.Println(" > No hay existe ese template de archivo")
	}
}

func (fm *FileMaker) GetFilePath(dir, extension string, isPlural bool) string {

	filename := fm.GetFilePathFilename("", "", isPlural)
	return fm.RootPath + "/" + fm.Workingdir + "/" + dir + "/" + filename + "." + extension
}

func (fm *FileMaker) saveFile(filepath string, content string) bool {
	if !fm.forceMode {
		fmt.Printf("the file %q alredy exists", filepath)
		return false
	} else {
		if err := os.WriteFile(filepath, []byte(content), fm.filePerms); err != nil {
			fmt.Println(err)
			return false
		} else {
			fmt.Println(" > Archivo creado: ", filepath)
			return true
		}
	}
}

func (fm *FileMaker) GetFilePathFilename(prefix, posfix string, isPlural bool) string {

	filename := fm.Parser.NameSingular
	if isPlural {
		filename = fm.Parser.NamePlural
	}
	if prefix != "" {
		filename = prefix + filename
	}
	if posfix != "" {
		filename = filename + posfix
	}
	return filename
}
