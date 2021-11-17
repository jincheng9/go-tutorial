// example3.go
package main

import "fmt"

func main() {
	a := 10
	var p1 *int = &a
	var p2 *int = &a
	fmt.Println("p1 value:", p1, " address:", &p1)
	fmt.Println("p2 value:", p2, " address:", &p2)
}
