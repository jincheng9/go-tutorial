// named_return3.go
package main

import "fmt"

func test() (done func()) {
	done = func() { fmt.Println("done") }
	return func() { fmt.Println("test"); done() }
}

func main() {
	done := test()
	// 下面的函数调用会进入死循环，不断打印test
	done()
}
