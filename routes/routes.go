package routes

import (
	"fmt"
	"go-static-website/page"
	"go-static-website/text/de"
	"go-static-website/text/en"
	. "maragu.dev/gomponents"
	"net/http"
	"strings"
)

func CreateApi() *http.ServeMux {
	routes := GetRoutes()

	server := http.NewServeMux()
	staticHandler := http.StripPrefix("/static", http.FileServer(http.Dir("./static")))
	server.HandleFunc("GET /", renderIndex(staticHandler))

	for route, content := range routes {
		server.HandleFunc(fmt.Sprint("GET ", route+".html"), callRenderOnNode(content))
	}

	return server
}

func GetRoutes() page.Routes {
	return en.Routes().AddRoutes(de.Routes())
}

func callRenderOnNode(content Node) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := content.Render(w)
		if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
		}
	}
}

func renderIndex(staticHandler http.Handler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		// all paths that have no directly matching HandleFunc path will land here
		// if the path starts with /static, the static files are served
		if strings.HasPrefix(path, "/static") {
			staticHandler.ServeHTTP(w, r)
			return
		}
		// the en index path is set to be the default path
		if path == "/" {
			err := en.IndexPage().Render(w)
			if err != nil {
				http.Error(w, "Internal error", http.StatusInternalServerError)
			}
			return
		}
		// if the path is not found, a 404 error is returned
		http.Error(w, "Not found", http.StatusNotFound)
	}
}
