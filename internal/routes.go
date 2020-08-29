package internal

import (
	"./pages"
	"net/http"
)

func DefineURIHandlers() {
	http.HandleFunc("/", getView(pages.HomePage))
}

func getView(handler func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		isValid := validateURIs(request.URL.Path)
		if !isValid {
			http.NotFound(responseWriter, request)
			return
		}
		handler(responseWriter, request)
	}
}

func validateURIs(uri string) bool {
	return uri == "/"
}
