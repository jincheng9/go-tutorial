// example1.go
package main

import "fmt"

func main() {
	a := *new([]int)
	fmt.Printf("%T, %v\n", a, a==nil)

	b := *new(map[string]int)
	fmt.Printf("%T, %v\n", b, b==nil)

	c := *new(chan int)
	fmt.Printf("%T, %v\n", c, c==nil)
}