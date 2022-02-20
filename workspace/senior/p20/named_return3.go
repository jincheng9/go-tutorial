// named_return3.go
package main

import "fmt"

var print = fmt.Println

func aaa() (done func(), err error) {
	return func() { print("aaa: done") }, nil
}

func bbb() (done func(), _ error) {
	f, err := aaa()
	return func() { print("bbb: surprise!"); f() }, err
}

func main() {
	done, _ := bbb()
	// 下面的函数调用会正常结束，不会进入死循环
	done()
}
