package main

import (
	"net/http"
	"time"
	"log"
	"fmt"
	"jnaves/api/web"
)

type httpServer struct{
	address   string
	readt 	  time.Duration
	writet    time.Duration
	maxh      int
	handlers  map[string]http.HandlerFunc
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func StartHttpServer(def httpServer){
	s := &http.Server{
		Addr:           def.address,
		Handler:        logRequest(http.DefaultServeMux),
		ReadTimeout:    def.readt * time.Second,
		WriteTimeout:   def.writet * time.Second,
		MaxHeaderBytes: 1 << def.maxh,
	}
	
	for path, handle := range def.handlers{
		http.HandleFunc(path, handle)
	}

	log.Fatal(s.ListenAndServe())
}

func main() {

	def := httpServer{
		address:   ":8080",
		readt:     10,
		writet:    10,
		maxh:      20,
		handlers:  web.GetHandlers(),
	}
	
	fmt.Printf("Staring server with the following config: %v\n", def)
	
	StartHttpServer(def)

	

	
}