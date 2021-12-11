// example2.go
package main

import "fmt"

// DBConfig 声明全局匿名结构体变量
var DBConfig struct {
	user string
	pwd string
	host string
	port int
	db string
}

// SysConfig 全局匿名结构体变量也可以在声明的时候直接初始化赋值
var SysConfig = struct{
	sysName string
	mode string
}{"tutorial", "debug"}

func main() {
	// 给匿名结构体变量DBConfig赋值
	DBConfig.user = "root"
	DBConfig.pwd = "root"
	DBConfig.host = "127.0.0.1"
	DBConfig.port = 3306
	DBConfig.db = "test_db"
	fmt.Println(DBConfig)
}
