package main

import "fmt"

func test() {
	r := recover()
	fmt.Println(r)
}

func main() {
	var s []int
	fmt.Println(len(s), cap(s))
	defer func() {
		//defer recover()
		test()
	}()
	//defer recover()
	panic(1)
}
