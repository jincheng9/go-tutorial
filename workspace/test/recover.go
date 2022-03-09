package main

import "fmt"

func A() {
	r := recover()
	fmt.Println(r)
}

func main() {
	defer A()
	panic(1)
}
