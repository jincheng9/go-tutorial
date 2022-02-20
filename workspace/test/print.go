package main

import "fmt"

func test() (a int) {
	a = 10
	return a
}

func test2() (a int) {
	a = 10
	return
}

func test3() (a int) {
	fmt.Println(a)
	a, b := 10, 1
	fmt.Println(a, b)
	return func() (a int) { return a }()
}

func test4() (a int) {
	fmt.Println(a)
	a, b := 10, 1
	fmt.Println(a, b)
	return func() int { return a }()
}

func main() {
	fmt.Println("test:", test())
	fmt.Println("test2:", test2())
	fmt.Println("test3:", test3())
	fmt.Println("test4:", test4())
	var a func()
	a()
}
