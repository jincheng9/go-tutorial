// quiz2.go
package main

import "fmt"

func main() {
	fmt.Print("1 ")
	defer fmt.Println(recover())
	fmt.Print("2 ")
	var a []int
	_ = a[0]
	fmt.Print("3 ")
}
