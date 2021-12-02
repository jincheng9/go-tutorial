// example4.go
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	// 打开文件
	fileName := fmt.Sprintf("app_%s.log", time.Now().Format("20060102"))
	f, err := os.OpenFile(fileName, os.O_RDWR | os.O_APPEND | os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}
	// 通过New方法自定义Logger，New的参数对应的是Logger结构体的output, prefix和flag字段
	logger := log.New(io.MultiWriter(os.Stdout, os.Stderr, f), "[INFO] ", log.LstdFlags | log.Lshortfile | log.Lmsgprefix)
	// 调用Logger的方法Println打印日志到指定文件
	logger.Println("your message")
}