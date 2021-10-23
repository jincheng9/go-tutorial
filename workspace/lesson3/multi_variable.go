package main

import "fmt"

/*
一次性定义多个变量
*/
var a, b, c int = 1, 2, 3
var e, f, g = 1, 2, "12"
var h, i bool


func main() {
	fmt.Println(a, b, c)
	fmt.Println(e, f, g)
	fmt.Println(h, i)

	var a1, b1 int		
	fmt.Println(a1, b1)

	a1, b1 = 11, 22
	fmt.Println(a1, b1)

	var a2, b2 int
	a2 = 111
	b2 = 222
	fmt.Println(a2, b2)

	a3, b3 := 10, 20
	fmt.Println(a3, b3)

	a4, b4 := 1, "str"
	fmt.Println(a4, b4)
}