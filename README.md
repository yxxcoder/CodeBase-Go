# CodeBase-Go
关于Go语言的CodeBase，可能包括Web框架、模板

## Chapter1-http请求
发送post、get请求的各种方式

### 1. get请求
````go
http.Get(url)
...
client := &http.Client{}
client.Get(url)
...
req, err := http.NewRequest("GET", url, body)
client.Do(req)
````
### 2. post请求
````go
http.Post(url, contentType, body)
...
client := &http.Client{}
client.Post(url, contentType, body)
...
req, err := http.NewRequest("POST", url, body)
client.Do(req)
...
http.PostForm(url,
	url.Values{"key": {"Value"}, "id": {"123"}})
...
client.PostForm(url, url.Values{"key": {"Value"}, "id": {"123"}})
````


## Chapter2-http服务

创建http服务常见的几种方式

### 1. 简单的HTTP服务器
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

## Chapter3-测试
单元测试与性能测试的方法，详情请见calculate.go
### 1. 执行单元测试
```shell
go test -v
```
### 2. 执行性能测试
go test不会默认执行压力测试的函数，如果要执行压力测试需要带上参数-test.bench
```shell
# 测试全部的压力测试函数
go test -test.bench=".*"
```