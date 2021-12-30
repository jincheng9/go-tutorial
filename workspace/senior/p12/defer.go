package main

import "fmt"

// 函数类型FuncType
type FuncType func(int) int

// 定义变量f1, 类型是FuncType
var f1 FuncType = func(a int) int { return a }

// 定义变量f, 类型是一个函数类型，函数签名是func(int) int
var f func(int) int

func main() {
	fmt.Println(f1(1))
	// 定义变量f2，类型是FuncType
	var f2 FuncType
	f2 = func(a int) int {
		a++
		return a
	}
	fmt.Println(f2(1))

	// 给函数类型变量f赋值
	f = func(a int) int {
		return 10
	}
	fmt.Println(f(1))
}
