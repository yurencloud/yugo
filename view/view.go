package view

import (
	"github.com/CloudyKit/jet"
	"bytes"
	"net/http"
	"yugo/config"
)

func Render(writer http.ResponseWriter, templateName string, varMap jet.VarMap) {
	var View = jet.NewHTMLSet(config.Get("template.path"))

	template, err := View.GetTemplate(templateName + config.Get("template.suffix"))
	if err != nil {
		// template could not be loaded
	}
	var html bytes.Buffer // needs to conform to io.Writer interface (like gin's context.Writer for example)
	if err = template.Execute(&html, varMap, nil); err != nil {
		// error when executing template
	}
	writer.Write(html.Bytes())
}
