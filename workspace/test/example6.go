package main

import (
	"fmt"
)

func testAddr() [2]int {
	return [2]int{1,2}
}

type S1 struct{
	a int
}

func(S1) print() {
	fmt.Println("doing sth")
}

func sliceTest1(s *[]int) {
	s1 := s
	fmt.Println(s, *s)
	*s = make([]int, 10)
	s2 := s
	fmt.Println(s, *s)
	fmt.Println(*s1, *s2)
}

func main() {
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(len(x), cap(x))
	y := x[1:3]
	fmt.Println(len(y), cap(y))
	fmt.Println(y[0:4])
}
