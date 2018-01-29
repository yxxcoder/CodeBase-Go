# CodeBase-Go
关于Go语言的CodeBase，可能包括Web框架、模板

## Chapter1-http
发送post、get请求的各种方式

### 1. get请求
````
http.Get(url)
...
client := &http.Client{}
client.Get(url)
...
req, err := http.NewRequest("GET", url, body)
client.Do(req)
````
### 2. post请求
````
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
