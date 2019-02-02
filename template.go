package main

import (
	"html/template"
	"io"
	"strings"
)

const (
	markdownTemplate = `
<div align="center">
<h1>{{toUpper .ProjectName}}</h1>
{{.Website}}
</div>

### What is {{.ProjectName}}?
{{.Description}}

> #### Important Notice
> {{.Notice}}

### Usage:
- one
- two
- three

#
### Author
**{{.Author}} {{.Year}}**
	`
)

// TemplateMD template
type TemplateMD struct {
	ProjectName string
	Website     string
	Description string
	Notice      string
	Author      string
	Year        string
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
