// named_return4.go
package main

func aaa() (done func()) {
	return func() { print("aaa: done") }
}

func bbb() (done func()) {
	done = aaa()
	return done
}

func main() {
	done := bbb()
	// 下面的函数调用会正常结束，打印"aaa: done"
	done()
}
