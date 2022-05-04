package main

import "fmt"

func test1() {
	const C1 = iota
	const C2 = iota
	fmt.Println("C1=", C1, " C2=", C2)
}

func main() {
	test1()
}
