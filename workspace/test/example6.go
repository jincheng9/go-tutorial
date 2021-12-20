package main

import (
	"fmt"
	"unsafe"
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
	fmt.Printf("%p %p %v\n", s, &s, unsafe.Sizeof(s))
	b := *s
	fmt.Printf("%p %p\n", b, &b)
	*s = make([]int, 10)
	fmt.Printf("%p %p\n", s, &s)
}

func main() {
	s := make([]int, 5)
	sliceTest1(&s)
	s2 := s
	s2[0] = 10
	fmt.Println(s, s2)
}
