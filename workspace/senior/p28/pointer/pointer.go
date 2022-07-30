package main

import "fmt"

func test1() *int {
	n := 10
	return &n
}

func test2() *int {
	n := 10
	return &n
}

var a int

func main() {
	a := test1()
	fmt.Println(*a)

	test2()
}
