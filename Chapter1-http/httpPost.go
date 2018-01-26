package main

import (
	"net/http"
	"time"
	"io"
	"net/url"
)

func SamplePost(url string, contentType string, body io.Reader) int {
	// 得到结果
	response, _ := http.Post(url, contentType, body)

	// 程序在使用完回复后必须关闭回复的主体
	defer response.Body.Close()

	//返回的状态码
	return response.StatusCode
}

func ClientPost(url string) int {
	//生成client 可设置参数
	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 100,
		},
	}

	//提交请求
	request, err := http.NewRequest("POST", url, nil)

	if err != nil {
		panic(err)
	}

	//返回结果
	response, _ := client.Do(request)

	// 程序在使用完回复后必须关闭回复的主体
	defer response.Body.Close()

	//返回的状态码
	return response.StatusCode
}

func ClientPost2(url string, contentType string, body io.Reader) int {
	//生成client 可设置参数
	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 100,
		},
	}

	//返回结果
	response, _ := client.Post(url, contentType, body)

	// 程序在使用完回复后必须关闭回复的主体
	defer response.Body.Close()

	//返回的状态码
	return response.StatusCode
}

func PostForm(url string, data url.Values) int {
	//得到结果
	response, _ := http.PostForm(url, data)

	// 程序在使用完回复后必须关闭回复的主体
	defer response.Body.Close()

	//返回的状态码
	return response.StatusCode
}

func PostForm2(url string, data url.Values) int {
	//生成client 可设置参数
	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 100,
		},
	}

	//返回结果
	response, _ := client.PostForm(url, data)

	// 程序在使用完回复后必须关闭回复的主体
	defer response.Body.Close()

	//返回的状态码
	return response.StatusCode
}

