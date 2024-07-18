package templates

import (
	"tango_cli/pkg/template_feature"
	"tango_cli/pkg/template_frontend"
	"tango_cli/pkg/template_model"
	"tango_cli/pkg/template_route"
)

func (t *Templates) GetAPI() {

	files := NewTemplateType()

	files.add("feature", TemplateFileData{
		Data:      template_feature.FeatureAPI(),
		Dir:       "features",
		Extension: "go",
		IsPlural:  false,
	})
	files.add("model", TemplateFileData{
		Data:      template_model.ModelAPI(),
		Dir:       "models",
		Extension: "go",
		IsPlural:  false,
	})
	files.add("route", TemplateFileData{
		Data:      template_route.RouteAPI(),
		Dir:       "routes",
		Extension: "go",
		IsPlural:  false,
	})

	t.TemplateParsed = files
}

func (t *Templates) GetHttpClient() {

	files := NewTemplateType()

	files.add("httpclient", TemplateFileData{
		Data:      template_frontend.HttpClient(),
		Dir:       "src",
		Extension: "ts",
		IsPlural:  false,
	})

	t.TemplateParsed = files
}

func (t *Templates) GetModel() {

	files := NewTemplateType()

	files.add("model", TemplateFileData{
		Data:      template_model.ModelAPI(),
		Dir:       "models",
		Extension: "go",
		IsPlural:  false,
	})

	t.TemplateParsed = files
}
