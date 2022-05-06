// quiz4.go
package main

import "fmt"

type Circle struct {
	a int
}

func main() {
	var c Circle
	fmt.Printf("%p %v\n", &c, c)
}
