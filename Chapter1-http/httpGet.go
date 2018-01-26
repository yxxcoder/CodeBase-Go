package main

import (
	"net/http"
	"time"
)

func SampleGet(url string) int {
	// 得到结果
	response, _ := http.Get(url)

	// 程序在使用完回复后必须关闭回复的主体
	defer response.Body.Close()

	// 返回的状态码
	return response.StatusCode
}

func ClientGet(url string) int {
	// 生成client 可设置参数
	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 100,
		},
	}

	// 返回结果
	response, _ := client.Get(url)

	// 程序在使用完回复后必须关闭回复的主体
	defer response.Body.Close()

	//返回的状态码
	return response.StatusCode
}

func ClientGet2(url string) int {
	//生成client 可设置参数
	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 100,
		},
	}

	//提交请求
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	// 添加http Header  如果需要的话
	// request.Header.Add("If-None-Match", `W/"wyzzy"`)

	//返回结果
	response, _ := client.Do(request)

	// 程序在使用完回复后必须关闭回复的主体
	defer response.Body.Close()

	//返回的状态码
	return response.StatusCode
}




