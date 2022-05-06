package main

import "fmt"

func main() {
	var a [3]int
	var b []int
	c := new([]int)
	fmt.Println(a[0], b == nil)
	fmt.Println(*c == nil, c)
}
