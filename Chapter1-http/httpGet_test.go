package main

import "testing"

func TestSampleGet(t *testing.T)  {
	//生成要访问的url
	url := "http://www.baidu.com"

	if code := SampleGet(url); code == 200 {
		t.Logf("测试通过，返回 %v", code)
	} else {
		t.Errorf("测试失败，返回 %v", code)
	}
}

func TestClientGet(t *testing.T)  {
	//生成要访问的url
	url := "http://www.baidu.com"

	if code := ClientGet(url); code == 200 {
		t.Logf("测试通过，返回 %v", code)
	} else {
		t.Errorf("测试失败，返回 %v", code)
	}
}

func TestClientGet2(t *testing.T)  {
	//生成要访问的url
	url := "http://www.baidu.com"

	if code := ClientGet2(url); code == 200 {
		t.Logf("测试通过，返回 %v", code)
	} else {
		t.Errorf("测试失败，返回 %v", code)
	}
}
