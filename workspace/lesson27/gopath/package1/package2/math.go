package package2

import "fmt"

func Add(a, b int) int {
	fmt.Println(multi(a, b))
	fmt.Printf("add of %d and %d is below:\n", a, b)
	return a+b
}
