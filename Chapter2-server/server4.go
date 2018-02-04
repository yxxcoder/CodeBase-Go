package main

import (
	"net/http"
	"net/http/pprof"
	"fmt"
	"html"
	"log"
)

// 自定义Mux控制路由访问
func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandle)
	mux.HandleFunc("/debug/pprof", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)

	server := &http.Server{
		Addr:        "0.0.0.0:8080",
		Handler:     mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Print(err)
		return
	}


}

func indexHandle(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
