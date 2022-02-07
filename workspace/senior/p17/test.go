// test.go
package main

import (
	"constraints"
	"fmt"
)

// return the min value
func min[T constraints.Ordered](a, b T) T {
	fmt.Printf("%T ", a)
	if a < b {
		return a
	}
	return b
}

func main() {
	minInt := min(1, 2)
	fmt.Println(minInt)

	minFloat := min(1.0, 2.0)
	fmt.Println(minFloat)

	minStr := min("a", "b")
	fmt.Println(minStr)
}