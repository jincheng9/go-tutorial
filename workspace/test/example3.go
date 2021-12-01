package main

import "fmt"

func main() {
	var a chan int
	value := <- a
	fmt.Println(value)
}