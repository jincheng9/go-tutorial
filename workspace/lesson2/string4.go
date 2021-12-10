// string4.go
package main

import "fmt"

func strTest() {
	s := "abc"
	fmt.Println(len(s)) // 3
	s1 := s[:]
	s2 := s[:1]
	s3 := s[0:]
	s4 := s[0:2]
	fmt.Println(s, s1, s2, s3, s4) // abc abc a abc ab
}

func main() {
	strTest()
}
