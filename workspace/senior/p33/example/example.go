package main

import (
	"fmt"
)

func main() {
	s := make([]byte, 2, 4)
	s[0] = 100
	s0 := (*[0]byte)(s) // s0 != nil
	fmt.Printf("%T, %v, %p\n", s0, &s0, s)
	s1 := (*[1]byte)(s[1:]) // &s1[0] == &s[1]
	s2 := (*[2]byte)(s)     // &s2[0] == &s[0]
	// s4 := (*[4]byte)(s)     // panics: len([4]byte) > len(s)
	fmt.Printf("%T, %v, %p, %p\n", s1, s1[0], &s1[0], &s[1])
	fmt.Printf("%T, %v, %v, %p\n", s2, s2[0], &s2[0], s)
}
