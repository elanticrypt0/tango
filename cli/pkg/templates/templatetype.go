package templates

type TemplateType struct {
	File map[string]TemplateFileData
}

type TemplateFileData struct {
	Data      string
	Dir       string
	Extension string
	IsPlural  bool
}

func NewTemplateType() TemplateType {
	return TemplateType{
		File: make(map[string]TemplateFileData),
	}
}

func NewTemplateFileData() TemplateFileData {
	return TemplateFileData{
		Extension: "go",
		IsPlural:  false,
	}
}

func (tempty *TemplateType) add(key string, tfd TemplateFileData) {

	tempty.File[key] = tfd
	// tempty.File

}
