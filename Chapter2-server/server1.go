package main

import (
	"fmt"
	"log"
	"net/http"
	"html"
)

func main() {

	http.HandleFunc("/hello", helloHandle)
	http.Handle("/hello2", http.HandlerFunc(helloHandle))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
func helloHandle(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
