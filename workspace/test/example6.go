package main

import "fmt"

func testAddr() [2]int {
	return [2]int{1,2}
}

func main() {
	var a interface{}
	a = 10
	switch a.(type) {
	case int:
		fmt.Println("int", a)
	default:
		fmt.Println("default", a)
	}
}
