// https://github.com/golang/go/issues/9862
package main

import "fmt"

var a [1<<31]byte

func main() {
	fmt.Println(a[0])
}