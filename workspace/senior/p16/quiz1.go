// quiz1.go
package main

import "fmt"

func main() {
	a := make([]int, 20)

	b := a[18:]
	b = append(b, 2022)

	fmt.Println(len(b), cap(b))
}
