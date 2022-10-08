package main

import "fmt"

var i = data()

const j = 0

var k = 0

func data() int {
	return 0
}

func varTest() (r int) {
	fmt.Printf("%v, %p\n", &r, &r)
	r, a := 10, 11
	fmt.Printf("%v, %p\n", &r, &r)
	fmt.Printf("%v, %p\n", &a, &a)
	return r
}

func main() {
	varTest()
	fmt.Println(i)
	a := 10
	fmt.Printf("%v, %p\n", &a, &a)
	{
		// a is a new local variable, differrent with variable `a` in line 17
		a, b := 1, 2
		fmt.Println(&a, &b)
	}
	fmt.Printf("%v, %p\n", &a, &a)
}
