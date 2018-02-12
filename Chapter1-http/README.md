# Chapter1-http请求
常见的发送post、get请求的方式

### 1. get请求
详情请见httpGet.go
````go
// 基本的使用
http.Get(url)

// 自定义Client
client := &http.Client{}
client.Get(url)

// 构造request参数
req, err := http.NewRequest("GET", url, body)
client.Do(req)
````
### 2. post请求
````go
// 基本的使用
http.Post(url, contentType, body)

// 自定义Client
client := &http.Client{}
client.Post(url, contentType, body)

// 构造request参数
req, err := http.NewRequest("POST", url, body)
client.Do(req)

// 发送表单
http.PostForm(url,
	url.Values{"key": {"Value"}, "id": {"123"}})

client.PostForm(url, url.Values{"key": {"Value"}, "id": {"123"}})
````