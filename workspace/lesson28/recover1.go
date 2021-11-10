package main

import (
	"fmt"
)

func a() {
	defer func() {
		/*捕获函数a内部的panic*/
		r := recover()
		fmt.Println("panic recover", r)
	}()
	panic(1)
}

func main() {
	defer func() {
		/*因为函数a的panic已经被函数a内部的recover捕获了
		所以main里的recover捕获不到异常，r的值是nil*/
		r := recover()
		fmt.Println("main recover", r)
	}()
	a()
	fmt.Println("main")
}
