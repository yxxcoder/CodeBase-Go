# Chapter7-log 日志

Golang's log模块主要提供了3类接口。分别是 “Print 、Panic 、Fatal”
对每一类接口其提供了3中调用方式，分别是 "Xxxx 、 Xxxxln 、Xxxxf"，基本和fmt中的相关函数类似
### 1. 基本使用
详情请见log.go
````go
func main(){

	arr := []int {2,3}
	log.Print("Print array ",arr,"\n")
	log.Println("Println array",arr)
	log.Printf("Printf array with item [%d,%d]\n",arr[0],arr[1])

	//deferFatal()

	deferPanic()
}
/**
 * 对于 log.Fatal 接口，会先将日志内容打印到标准输出，
 * 接着调用系统的 os.exit(1) 接口，退出程序并返回状态 1
 * 但是有一点需要注意，由于是直接调用系统接口退出，defer函数不会被调用
 */
func deferFatal(){
	defer func() {
		fmt.Println("--first--")
	}()
	log.Fatalln("test for defer Fatal")
}

/**
 * 对于log.Panic接口，该函数把日志内容刷到标准错误后调用 panic 函数
 * 注意: 在Panic之后声明的defer是不会执行的
 */
func deferPanic(){
	defer func() {
		log.Println("--first--")
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	log.Panicln("test for defer Panic")
	defer func() {
		log.Println("--second--")
	}()
}


// 输出结果:
// 2018/02/19 18:58:05 Print array [2 3]
// 2018/02/19 18:58:05 Println array [2 3]
// 2018/02/19 18:58:05 Printf array with item [2,3]
// 2018/02/19 18:58:05 test for defer Panic
// 2018/02/19 18:58:05 --first--
// 2018/02/19 18:58:05 test for defer Panic
````

### 2. 自定义Logger
log.Logger提供了一个New方法用来创建对象, 详情请见customize.go

````go
func main(){

	fileName := "Chapter7-log/Info_First.log"
	logFile,err  := os.Create(fileName)

	defer logFile.Close()
	if err != nil {
		log.Fatalln("open file error")
	}

	debugLog := log.New(logFile,"[Info]",log.Llongfile)
	debugLog.Println("A Info message here")

	debugLog.SetPrefix("[Debug]")
	debugLog.Println("A Debug Message here ")
}


// Info_First.log 内容:
// [Info]/Users/yuxuan/Project/GoProject/CodeBase-Go/Chapter7-log/customize.go:32: A Info message here
// [Debug]/Users/yuxuan/Project/GoProject/CodeBase-Go/Chapter7-log/customize.go:35: A Debug Message here 
````