package util

import "fmt"

func init() {
	fmt.Println("init add...")
}

func Add(a, b int) int {
	return a + b
}
