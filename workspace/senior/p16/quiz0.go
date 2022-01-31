// quiz0.go
package main

import "fmt"

func main() {
	s := []string{"a", "b", "c"}

	copy(s[1:], s)

	fmt.Println(s)
}
