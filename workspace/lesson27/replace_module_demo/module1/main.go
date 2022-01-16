package main

import (
	"fmt"
	// 模块module1要使用本地模块module2里的Add函数
	// 这里被import的本地模块的名称要和module2/go.mod里保持一致
	"module2"
)

func main() {
	a := 1
	b := 2
	sum := module2.Add(a, b)
	fmt.Printf("sum of %d and %d is %d\n", a, b, sum)
}
