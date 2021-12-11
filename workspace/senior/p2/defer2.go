// defer2.go
package main

import (
	"fmt"
)

func test1() {
	fmt.Println("test")
}

func test2() {
	panic(1)
}
func main() {
	fmt.Println("main start")
	defer test1()
	test2()
	fmt.Println("main end")
}
