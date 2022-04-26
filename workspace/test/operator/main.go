package main

import (
	"fmt"
)

import "os"

var _s = test()

func test() string {
	return "test"
}

func main() {
	result := false && true || true
	fmt.Println(result)
	fmt.Println(_s)
	fmt.Println(os.Getwd())
}
