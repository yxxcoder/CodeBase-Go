package main

import (
	"net/http"
	"fmt"
	"html"
	"log"
)

type Handler struct {}

func main() {
	server := &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: &Handler{},
	}
	log.Fatal(server.ListenAndServe())
}

func (handler Handler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/hello" {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	} else {
		fmt.Fprintln(w, "Hello World")
	}

}


