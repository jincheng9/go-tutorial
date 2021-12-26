package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := a[1:]
	copy(a, b)
	fmt.Println(a, b)
}
