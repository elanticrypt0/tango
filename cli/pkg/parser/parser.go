package parser

import (
	"fmt"
	"strings"
)

type Parser struct {
	NamespaceOriginal string
	NamePlural        string `json:"name_plural"`
	NameSingular      string `json:"name_singular"`
	FirstChar         string `json:"first_char"`
	LastCharPos       int
}

func New() *Parser {
	return &Parser{}
}

func (p *Parser) Read(namespace string) {

	p.NamespaceOriginal = namespace

	// Convert the input
	p.convertToPlural()
	p.convertToSingular()
	p.getFirstChar()

}

func (p *Parser) ConvertToPluralAndGet(str2plural string) string {

	// todo
	// si termina en vocal agrega S
	// si termina en Y agrega "ies"
	p.LastCharPos = len(str2plural) - 1
	lastChar := string(str2plural[p.LastCharPos])
	pluralEnd := "s"

	buf := str2plural
	if lastChar == "y" {
		pluralEnd = "ies"
		buf = str2plural[0:p.LastCharPos]
	} else if lastChar == "s" {
		pluralEnd = ""
	}

	namePlural := fmt.Sprintf("%s%s", buf, pluralEnd)
	return strings.ToLower(namePlural)

}

func (p *Parser) convertToPlural() {

	// todo
	// si termina en vocal agrega S
	// si termina en Y agrega "ies"
	p.LastCharPos = len(p.NamespaceOriginal) - 1
	lastChar := string(p.NamespaceOriginal[p.LastCharPos])
	pluralEnd := "s"

	buf := p.NamespaceOriginal
	if lastChar == "y" {
		pluralEnd = "ies"
		buf = p.NamespaceOriginal[0:p.LastCharPos]
	} else if lastChar == "s" {
		pluralEnd = ""
	}

	p.NamePlural = fmt.Sprintf("%s%s", buf, pluralEnd)
	p.NamePlural = strings.ToLower(p.NamePlural)
}

func (p *Parser) convertToSingular() {
	if string(p.NamespaceOriginal[p.LastCharPos]) == "s" {
		p.NameSingular = p.NamespaceOriginal[0:p.LastCharPos]
	} else {
		p.NameSingular = p.NamespaceOriginal
	}
	p.NameSingular = strings.ToLower(p.NameSingular)
}

func (p *Parser) ConvertToTitle(s string) string {
	return strings.Title(s)
}

func (p *Parser) getFirstChar() {
	p.FirstChar = string(p.NamespaceOriginal[0:1])
	p.FirstChar = strings.ToLower(p.FirstChar)
}
