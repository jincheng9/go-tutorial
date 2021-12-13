package main

import "fmt"

type ST struct {
	ST2
}

type ST2 struct{

}

func(st *ST2) close(a int) int {
	return a
}

func add(a, b  int) int{
	return a+b
}


func main() {
	s := ST{ST2{}}
	fmt.Println(s.close(10))

	result := add(1, 2)
	fmt.Println(&result)

	sl := make([]int, 3, 10)
	fmt.Println(sl)
}

