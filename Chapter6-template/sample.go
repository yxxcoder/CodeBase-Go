package main

import (
	"net/http"
	"html/template"
)

func SayHello(w http.ResponseWriter, req *http.Request) {
	name := "克莱普斯"
	tmpl, _ := template.New("TXT").Parse("大家好，我是{{.}}")
	tmpl.Execute(w, name)
}

type Info struct {
	Title string
	Name  string
	Site  string
}

func SayHello2(w http.ResponseWriter, req *http.Request) {
	info := Info{"个人网站", "克莱普斯", "http://www.sample.com/"}
	tmpl, _ := template.ParseFiles("Chapter6-template/home.html")
	tmpl.Execute(w, info)
}

func main() {
	http.HandleFunc("/hello", SayHello)
	http.HandleFunc("/hello2", SayHello2)
	http.ListenAndServe(":8080", nil)
}