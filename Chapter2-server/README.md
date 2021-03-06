# Chapter2-http服务

创建http服务常见的几种方式

### 1. 简单的HTTP服务器
详情请见server1.go
````go
func main() {

	http.HandleFunc("/hello", HandlerHello)
	http.Handle("/hello2", http.HandlerFunc(HandlerHello))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
func HandlerHello(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
````

### 2. 自定义hanlder
详情请见server2.go
````go
//自定义hanlder
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
````

### 3. 自定义server对象
详情请见server3.go
```go
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
```

### 4. 控制路由访问
详情请见server4.go
```go
func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandle)
	mux.HandleFunc("/index", indexHandle)
	
    log.Fatal(http.ListenAndServe(":8080", mux))
}
func indexHandle(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
```