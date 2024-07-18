package templates

import (
	"fmt"
	"log"
	"strings"
	"tango_cli/pkg/parser"
)

type TemplatesFunc struct {
	Name string
	Fn   func()
}

// README
// Para reemplazar los archivos tienen una connotación especial
// $[TIPO: Singular | Plural][Capitalized | Lowercase]$
// EJ: $PC$ (Plural Capitalized)

type Templates struct {
	Parser           *parser.Parser
	Replacements     *strings.Replacer
	TemplateSelected string
	TemplateParsed   TemplateType
}

func New(p *parser.Parser, templateSelected string) *Templates {
	t := &Templates{
		Parser:           p,
		TemplateSelected: templateSelected,
	}
	t.setReplacements()
	if templateSelected != "" {
		t.SelectTemplate()
		return t
	} else {
		return nil
	}
}

func (t *Templates) SelectTemplate() error {
	switch t.TemplateSelected {
	case "API":
		t.GetAPI()
	case "HTTPCLIENT":
		t.GetHttpClient()
	case "MODEL":
		t.GetModel()
	default:
		log.Printf("Must select one template \n")

	}

	if t.TemplateSelected != "" {
		t.parseTemplateType()
		return nil
	} else {
		return fmt.Errorf("no template selected")
	}

}

func (t *Templates) setReplacements() {

	pc := t.Parser.ConvertToTitle(t.Parser.NamePlural)
	pl := t.Parser.NamePlural
	sc := t.Parser.ConvertToTitle(t.Parser.NameSingular)
	sl := t.Parser.NameSingular
	fl := t.Parser.FirstChar

	t.Replacements = strings.NewReplacer("$PC$", pc, "$PL$", pl, "$SC$", sc, "$SL$", sl, "$FL$", fl)

}

func (t *Templates) parseTemplateType() {
	for key, val := range t.TemplateParsed.File {
		val.Data = t.Replacements.Replace(val.Data)
		// t.TemplateParsed.add(key, t.Replacements.Replace(val))
		t.TemplateParsed.File[key] = val
	}
}
