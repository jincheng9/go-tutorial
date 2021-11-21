// string3.go
package main

import "fmt"

func main() {
	str := "abc"
	/*
	the following code has compile error:
	cannot take the address of str[0]
	*/
	fmt.Println(&str[0])
}

