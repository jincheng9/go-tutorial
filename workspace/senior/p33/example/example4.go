package main

import "fmt"

func IsEqual[T comparable](a T, b T) bool {
	return a == b
}

func main() {
	var a interface{} = 1
	var b interface{} = []int{1}
	fmt.Println(a == b)
	fmt.Println(IsEqual(a, b))
}
