package web

import (
	"html/template"
	"net/http"
)

type PageContent struct {
}

func RenderTemplate(response http.ResponseWriter, templateName string, content *PageContent) {
	templates := template.Must(template.ParseFiles("./templates/" + templateName))
	err := templates.ExecuteTemplate(response, templateName, content)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}
