package main

import (
"net/http"
"log"
	"fmt"
	"html"
)

type MyHandler struct {}

func main() {
	handler := &MyHandler{}
	http.Handle("/hello", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// 任何结构体，只要实现了ServeHTTP方法，这个结构就可以称之为handler对象
func (handler MyHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
