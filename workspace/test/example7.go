package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3}
	fmt.Printf("%T %d %d\n", a, len(a), cap(a))
	x := a[:1]
	fmt.Printf("%T %d %d\n", x, len(x), cap(x))
	y := a[2:]
	fmt.Printf("%T %d %d\n", y, len(y), cap(y))
	x = append(x, y...)
	fmt.Printf("%v %d %d\n", a, len(a), cap(a))
	fmt.Printf("%v %d %d\n", x, len(x), cap(x))
	fmt.Printf("%v %d %d\n", y, len(y), cap(y))
	x = append(x, y...)
	fmt.Printf("%v %d %d\n", a, len(a), cap(a))
	fmt.Printf("%v %d %d\n", x, len(x), cap(x))
	fmt.Printf("%v %d %d\n", y, len(y), cap(y))
	fmt.Println(&y[0], &a[2])
}
