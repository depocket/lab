package main

import (
	"github.com/depocket/lab/cli/templates"
	"text/template"
)

func main() {
	tpl, err := template.New(templates.DAppsIntegrationTemplate)
	tpl.Option()
}
