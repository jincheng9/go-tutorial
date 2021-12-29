package main

import "fmt"

type Add func(int, int) int

func main() {
	var f Add
	defer f(1, 2)
	fmt.Println("end")
}
