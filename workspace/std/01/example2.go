// example2.go
package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	// 日志文件名
	fileName := fmt.Sprintf("app_%s.log", time.Now().Format("20060102"))
	// 创建文件
	f, err := os.OpenFile(fileName, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}
	// main退出之前，关闭文件
	defer f.Close()
	// 调用SetOutput设置日志输出的地方
	log.SetOutput(f)
	//log.SetOutput(io.MultiWriter(os.Stdout, f))
	// 通过SetFlags设置Logger结构体里的flag属性
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile | log.Lmsgprefix)
	// 通过SetPrefix设置Logger结构体里的prefix属性
	log.SetPrefix("INFO:")
	// 调用辅助函数Println打印日志到指定文件
	log.Println("your message")
}