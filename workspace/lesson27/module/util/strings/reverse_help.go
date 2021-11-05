package strings

import "fmt"

func init() {
	fmt.Println("reverse_help init")
}

func reverseRune(r []rune) []rune {
	size := len(r)
	for i:=0; i<size/2; i++ {
		r[i], r[size-1-i] = r[size-1-i], r[i]
	}
	return r
}