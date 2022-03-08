package main

import "fmt"

func test() (a int) {
	a = 10
	return a
}

func test2() (a int) {
	a = 10
	return
}

func test3() (a int) {
	fmt.Println(a)
	a, b := 10, 1
	fmt.Println(a, b)
	return func() (a int) { return a }()
}

func test4() (a int) {
	fmt.Println(a)
	a, b := 10, 1
	fmt.Println(a, b)
	return func() int { return a }()
}

func test5() int {
	panic(0)
	//return 1
}

func test6() {
	ns := []int{010: 200, 005: 100}
	fmt.Println(ns)
	fmt.Println(len(ns))
}

func test7() {
	m := map[string][]string{}
	m["a"] = []string{"a"}
	result := m["b"]
	fmt.Println(result, result == nil)
}

func test8() {
	if a, b := 1, 2; a != 10 {
		fmt.Println(a, b)
	}
	// compile error: undefined a
	a = 10
	fmt.Println(a)
}

func test9() {
	var a int = 10
	var b int32 = 20
	// compile error: mismatched types int and int32
	fmt.Println(a == b)
}

func main() {
	// fmt.Println("test:", test())
	// fmt.Println("test2:", test2())
	// fmt.Println("test3:", test3())
	// fmt.Println("test4:", test4())
	// test6()
	test7()
}
