package main

import (
	"os"
	"log"
)


// 自定义Logger类型, 使用log.Logger提供了一个New方法用来创建对象
// 该函数一共有三个参数：
//（1）输出位置out，是一个io.Writer对象，该对象可以是一个文件也可以是实现了该接口的对象。通常我们可以用这个来指定日志输出到哪个文件。
//（2）prefix 我们在前面已经看到，就是在日志内容前面的东西。我们可以将其置为 "[Info]" 、 "[Warning]"等来帮助区分日志级别。
//（3） flags 是一个选项，显示日志开头的东西，可选的值有：
//  Ldate         = 1 << iota     // 形如 2009/01/23 的日期
//  Ltime                         // 形如 01:23:23   的时间
//  Lmicroseconds                 // 形如 01:23:23.123123   的时间
//  Llongfile                     // 全路径文件名和行号: /a/b/c/d.go:23
//  Lshortfile                    // 文件名和行号: d.go:23
//  LstdFlags     = Ldate | Ltime // 日期和时间

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