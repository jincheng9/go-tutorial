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
	var b[][5]int
	b[0] = [5]int{1}
	fmt.Println(b == nil)

	c := []int{}
	fmt.Println(c == nil)

}