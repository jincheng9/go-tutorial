package main

import (
	"fmt"
)

func test() {
	i := -100
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println(i)
}

func main() {
	test()
}
