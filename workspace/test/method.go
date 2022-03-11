package main

import "fmt"

type A struct {
	x int
}

func (a A) caller1() {
	a.x = 10
}

func (a *A) caller2() {
	a.x = 10
	// fmt.Println(a)
}

func main() {
	var a = new(A)
	a.caller1()
	fmt.Println(a.x)
}
