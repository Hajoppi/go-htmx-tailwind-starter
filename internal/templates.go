package internal

import (
	"html/template"
	"io"
)

var templates *template.Template

func initTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

func renderTemplate(w io.Writer, name string, data interface{}) error {
	return templates.ExecuteTemplate(w, name, data)
}
