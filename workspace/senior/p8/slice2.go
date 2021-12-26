package main

import "fmt"

func main() {
	a := []int{1, 2}
	b := append(a, 3)

	c := append(b, 4)
	d := append(b, 5)

	fmt.Println(a, b, c[3], d[3])
}
