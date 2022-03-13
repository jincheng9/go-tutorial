package main

import "fmt"

func main() {
	a := 10
	var num *int = &a
	fmt.Println(&num, num)
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", num)
}
