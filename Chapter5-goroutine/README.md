# Chapter5-goroutine 协程

GO语言协程的使用方法

### 1. goroutine
goroutine 是由 Go 运行时环境管理的轻量级线程<br>
<code>go f(x, y)</code>开启一个新的 goroutine 执行<br>
详情请见main.go
````go
func main() {
	go f()
}

func f() {
	// do something
}
````

### 2. select
select 语句使得一个 goroutine 在多个通讯操作上等待
select 会阻塞，直到条件分支中的某个可以继续执行，这时就会执行那个条件分支。当多个都准备好的时候，会随机选择一个
详情请见select.go
````go
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
````

