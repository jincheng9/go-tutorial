// example1.go
package main

import (
	"log"
)

func main() {
	// 通过SetFlags设置Logger结构体里的flag属性
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile | log.Lmsgprefix)
	// 通过SetPrefix设置Logger结构体里的prefix属性
	log.SetPrefix("INFO:")
	// 调用辅助函数Println打印日志到标准输出
	log.Print("your message")
}