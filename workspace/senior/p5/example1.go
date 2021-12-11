// example1.go
package main

import (
	"fmt"
)

func main() {
	a := struct{name string; age int}{"bob", 10}
	b := struct{
		school string
		city string
	}{"THU", "Beijing"}
	fmt.Println(a, b)
}
