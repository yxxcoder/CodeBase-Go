# Chapter6-template 模版引擎

template包（html/template）实现了数据驱动的模板，用于生成可对抗代码注入的安全HTML输出

### 1. 基本使用
详情请见sample.go
````go
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
````

