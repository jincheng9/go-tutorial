package main

import "fmt"

func test() {
	num := 0
	f := func() func(int, int) {
		num = 1
		//fmt.Println("f...", num)
		return func(int, int) {}
	}
	g := func() (int, int) {
		fmt.Println("g...", num)
		return num, num
	}
	f()(g())
}

func main() {
	test2(f)
}
