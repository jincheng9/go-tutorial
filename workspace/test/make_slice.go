package main

import "fmt"

func main() {
	s := make([]int, 2, 3)
	fmt.Println(s)
	s = append(s, 1, 2, 3, 4, 5)
	fmt.Println(len(s), cap(s))
}
