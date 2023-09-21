package web

import (
	"fmt"
	"jnaves/api/backend"
	"net/http"
	s "strings"
)

var handlers = make(map[string]http.HandlerFunc)

func fooHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello! handler map is working")
}

func jsonToCss(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir("./web/public/"))
	if r.Method == "GET" {
		r.URL.Path += ".html"
		fs.ServeHTTP(w, r)
	}
	if r.Method == "POST" {
		json_struct, err := backend.ProcessBodyJson(r.Body)
		if err != nil {
			fmt.Println(err)
		}
		backend_err := backend.JsonToCss(json_struct)
		if backend_err != nil {
			fmt.Println("ERROR IN BACKEND")
			fmt.Println(backend_err)
		}
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir("./web/public/"))
	if !s.HasSuffix(r.URL.Path, ".html") && !s.HasSuffix(r.URL.Path, ".js") && !s.HasSuffix(r.URL.Path, ".css") {
		r.URL.Path += ".html"
	}
	// fmt.Println(r.URL.Path)
	fs.ServeHTTP(w, r)
}

func GetHandlers() map[string]http.HandlerFunc {
	handlers["/"] = mainHandler
	handlers["/foo"] = fooHandler
	handlers["/json"] = jsonToCss
	return handlers
}
