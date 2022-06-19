package main

import (
	"fmt"
)

type T[T any] struct {
	m1 T
}

func(t T[T]) print() {
	fmt.Println(t.m1)
}

func main() {
	a := T[int]{}
	fmt.Println(a.m1)
}