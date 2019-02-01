package main

import (
	"html/template"
	"io"
	"strings"
)

const (
	markdownTemplate = `
### {{toUpper .ProjectName}}
#
{{.Subtitle}}
#

#### Usage:
- one
- two
- three

#
### Author
#### {{.Author}}
	`
)

// TemplateMD template
type TemplateMD struct {
	ProjectName string
	Subtitle    string
	Author      string
}

// ParseOutput template
func ParseOutput(in TemplateMD, out io.Writer) error {
	funcMap := template.FuncMap{
		"toUpper": func(v string) string {
			return strings.ToUpper(v)
		},
	}

	t := template.New("tmpl").Funcs(funcMap)

	t, err := t.Parse(markdownTemplate)
	if err != nil {
		return err
	}

	err = t.Execute(out, in)
	if err != nil {
		return err
	}

	return nil

}
