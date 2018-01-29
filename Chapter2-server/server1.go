package main

import (
	"fmt"
	"log"
	"net/http"
	"html"
)

func main() {
	http.HandleFunc("/hello", HandlerHello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func HandlerHello(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
