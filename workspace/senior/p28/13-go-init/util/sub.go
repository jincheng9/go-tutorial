package util

import "fmt"

func init() {
	fmt.Println("init sub...")
}

func Sub(a, b int) int {
	return a - b
}
