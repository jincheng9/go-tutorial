// named_return.go
package main

import (
	"fmt"
)

var print = fmt.Println

func aaa() (done func()) {
	return func() { print("aaa: done") }
}

func bbb() (done func()) {
	fmt.Println(done)
	done = aaa()
	return
	//return func() { print("bbb: surprise!"); time.Sleep(time.Second * 5); done() }
}

func main() {
	done := bbb()
	done()
}
