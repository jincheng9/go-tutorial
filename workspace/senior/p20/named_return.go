// named_return.go
package main

import "fmt"

var print = fmt.Println

func aaa() (done func(), err error) {
	return func() { print("aaa: done") }, nil
}

func bbb() (done func(), _ error) {
	donot, err := aaa()
	return func() { print("bbb: surprise!"); donot() }, err
}

func main() {
	done, _ := bbb()
	done()
}
