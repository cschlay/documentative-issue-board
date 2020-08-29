package pages

import (
	"../web"
	"net/http"
)

func HomePage(response http.ResponseWriter, request *http.Request) {
	web.RenderTemplate(response, "home.html", &web.PageContent{})
}
