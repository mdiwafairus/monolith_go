package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// RenderTemplate renders templates using HTML template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	var pathToTemplates = "./templates"
	parsedTemplate, _ := template.ParseFiles(pathToTemplates + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
	}
}