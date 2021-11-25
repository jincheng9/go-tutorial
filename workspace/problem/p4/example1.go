package main

import (
	"bytes"
	"fmt"
)

type A struct {
	s []int
	b map[string]int
	c int
}

func main() {
	var a bytes.Buffer
	fmt.Println(a)

	c := []int{}
	fmt.Println(c == nil)

	d := new([]int)
	fmt.Println(d)
}