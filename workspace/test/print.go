package main

import "fmt"

func test() (a int) {
	a = 10
	return func() int { return a }()
}
func main() {
	fmt.Println(test())
}
