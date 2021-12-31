package main

import "fmt"

func test() {
	r := recover()
	fmt.Println(r)
}

func main() {
	defer func() {
		//defer recover()
		test()
	}()
	//defer recover()
	panic(1)
}
