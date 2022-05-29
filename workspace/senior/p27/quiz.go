// quiz.go
package main

import "fmt"

func test1() {
	var y = 5.2
	const z = 2
	fmt.Printf("%T, %T, %v, %v, %v\n", y, z, y, z, y/z)
}

// func test2() {
// 	const t = 4.8
// 	var u = 2
// 	fmt.Println(t / u)
// }

// func test3() {
// 	var (
// 		a = 1.0
// 		b = 2
// 	)
// 	fmt.Println(a / b)
// }

// func test3_1() {
// 	a := 1.0
// 	b := 2
// 	fmt.Println(a / b)
// }

func test4() {
	const (
		x = 5.0
		y = 4
	)
	fmt.Printf("%T, %T, %v, %v, %v\n", x, y, x, y, x/y)
}

func test5() {
	const a = 7.0
	var b = 2
	fmt.Println(a / b)
}

func main() {
	test1()
	test4()
	test5()
}
