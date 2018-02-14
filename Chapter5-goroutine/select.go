package main

import (
	"time"
	"fmt"
)

// select 语句使得一个 goroutine 在多个通讯操作上等待
// select 会阻塞，直到条件分支中的某个可以继续执行，这时就会执行那个条件分支
// 当多个都准备好的时候，会随机选择一个
func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
