package main

import (
	"fmt"
)

func test1() {
	ch := make(<-chan int)
	close(ch)
	fmt.Printf("%T\n", ch)
}

func test2() {
	ch := make(chan<- int)
	close(ch)
	fmt.Printf("%T\n", ch)
}

func main() {
	test1()
	test2()
}
