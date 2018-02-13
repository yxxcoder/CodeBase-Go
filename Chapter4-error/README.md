# Chapter4-error 异常处理

异常的处理方式，包括异常的使用方法、异常的捕获与恢复等

### 1. 自定义异常
详情请见customize.go
````go
// 生成一个简单的 error 类型
errors.New("您输入不是偶数")


// 自定义参数的错误输出
type dualError struct {
	Num     int
	problem string
}
// 实现Error接口
func (e dualError) Error() string {
	return fmt.Sprintf("参数不正确，因为\"%d\"不是偶数", e.Num)
}

func requireDual2(n int) (int, error) {
	if n&1 == 1 {
		return -1, dualError{Num: n}
	}

	return n, nil
}
````

### 2. 异常的捕获与恢复
详情请见catch.go
````go
func main() {
	defer func() {
		fmt.Println("c")
		if r := recover(); r != nil {
			fmt.Println(err)
		}
		fmt.Println("d")
	}()
	
	f()
}
func f() {
	fmt.Println("a")
	panic("yx")
	fmt.Println("b")
}

// 输出结果:
// a
// c
// yx
// d
// Process finished with exit code 0
````