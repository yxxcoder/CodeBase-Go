package main

import (
	"testing"
	"net/url"
)

const (
	URL = "http://www.baidu.com"
	ContentType = "img/bd_logo1.png"
)

func TestSamplePost(t *testing.T)  {

	if code := SamplePost(URL, ContentType, nil); code == 200 {
		t.Logf("测试通过，返回 %v", code)
	} else {
		t.Errorf("测试失败，返回 %v", code)
	}
}

func TestClientPost(t *testing.T)  {

	if code := ClientPost(URL); code == 200 {
		t.Logf("测试通过，返回 %v", code)
	} else {
		t.Errorf("测试失败，返回 %v", code)
	}
}

func TestClientPost2(t *testing.T)  {

	if code := ClientPost2(URL, ContentType, nil); code == 200 {
		t.Logf("测试通过，返回 %v", code)
	} else {
		t.Errorf("测试失败，返回 %v", code)
	}
}

func TestPostForm(t *testing.T)  {

	if code := PostForm(URL, url.Values{"key": {"Value"}, "id": {"123"}}); code == 200 {
		t.Logf("测试通过，返回 %v", code)
	} else {
		t.Errorf("测试失败，返回 %v", code)
	}
}

func TestPostForm2(t *testing.T)  {

	if code := PostForm2(URL, url.Values{"key": {"Value"}, "id": {"123"}}); code == 200 {
		t.Logf("测试通过，返回 %v", code)
	} else {
		t.Errorf("测试失败，返回 %v", code)
	}
}
