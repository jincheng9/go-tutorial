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

type Struct1 struct {
	v1 int
	v2 [3]int
	v3 []int
	v4 map[int]int
}

func copyTest(v Struct1) {
	v.v1 = 10
	v.v2[0] = 10
	v.v3[0] = 10
	v.v4[0] = 10
	fmt.Println(v)
}

func main() {
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(len(x), cap(x))
	y := x[1:3]
	fmt.Println(len(y), cap(y))
	fmt.Println(y[0:4])

	s1 := Struct1{v3:[]int{1}, v4:map[int]int{0:1}}
	fmt.Println(s1)
	copyTest(s1)
	fmt.Println(s1)
}
