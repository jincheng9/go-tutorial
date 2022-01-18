package main

import "fmt"

const N = 6

var n = 6

func f() string {
	return string(1 << N)
}

// func g() string {
// 	return string(1 << n)
// }

func main() {
	fmt.Println(f())
	fmt.Println(string(2))
	// fmt.Println(g())
}
