package main

import (
	"errors"
	"fmt"
)
/**
 * 基本用法
 */
func requireDual(n int) (int, error) {
	if n&1 == 1 {
		return -1, errors.New("您输入不是偶数") //生成一个简单的 error 类型
	}

	return n, nil
}

/**
 * 自定义参数的错误输出
 */
type dualError struct {
	Num     int
	problem string
}

func (e dualError) Error() string {
	return fmt.Sprintf("参数不正确，因为\"%d\"不是偶数", e.Num)
}

func requireDual2(n int) (int, error) {
	if n&1 == 1 {
		return -1, dualError{Num: n}
	}

	return n, nil
}

func main() {
	if result, err := requireDual(101); err != nil {
		fmt.Println("错误：", err)
	} else {
		fmt.Println("结果：", result)
	}

	if result, err := requireDual2(101); err != nil {
		fmt.Println("错误：", err)
	} else {
		fmt.Println("结果：", result)
	}
}
