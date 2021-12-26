package main

import "fmt"

func test() {
	a := [...]int{1, 2, 3, 4, 5}
	x := a[2:4]
	fmt.Println(len(x), cap(x), x)
	y := x[0:1]
	fmt.Println(len(y), cap(y), y)
}

func test2() {
	var a []int
	var oldCap = 0
	for i := 0; i < 2048; i++ {
		a = append(a, i)
		if cap(a) != oldCap {
			fmt.Println(len(a), cap(a))
			oldCap = cap(a)
		}
	}
}

func test3() {
	a := make([]int, 1, 4)
	fmt.Println(&a[0])
	b := append(a, 1)
	fmt.Println(&b[0])
	fmt.Println(a, b)
	fmt.Println(append(a, 2), a, b)
	c := a[1:4]
	fmt.Println(len(c), cap(c), c)
}

func test4() {
	a := make([]int, 0, 4)
	b := a[:]  // 等价于 b := a[0:0], b的长度是0，容量是4
	c := a[:1] // 等价于 c := a[0:1], b的长度是1，容量是4
	// d := a[1:]  // 编译报错 panic: runtime error: slice bounds out of range
	e := a[1:4] // e的长度3，容量3
	fmt.Println(cap(b), cap(c), cap(e))
}

func main() {
	test4()
	// var s []int
	// fmt.Println(len(s), cap(s))
	// a := [...]int{0, 1, 2, 3}
	// fmt.Printf("%T %d %d\n", a, len(a), cap(a))
	// x := a[:1]
	// fmt.Printf("%T %d %d\n", x, len(x), cap(x))
	// y := a[2:]
	// fmt.Printf("%T %d %d\n", y, len(y), cap(y))
	// x = append(x, y...)
	// fmt.Printf("%v %d %d\n", a, len(a), cap(a))
	// fmt.Printf("%v %d %d\n", x, len(x), cap(x))
	// fmt.Printf("%v %d %d\n", y, len(y), cap(y))
	// x = append(x, y...)
	// fmt.Printf("%v %d %d\n", a, len(a), cap(a))
	// fmt.Printf("%v %d %d\n", x, len(x), cap(x))
	// fmt.Printf("%v %d %d\n", y, len(y), cap(y))
	// fmt.Println(&y[0], &a[2])
}
