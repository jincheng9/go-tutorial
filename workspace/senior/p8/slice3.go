package main

import "fmt"

func main() {
	s := []int{1, 2}
	s = append(s, 4, 5, 6)
	fmt.Println(len(s), cap(s))
}
