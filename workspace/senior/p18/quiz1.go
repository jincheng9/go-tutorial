// quiz1.go
package main

import "fmt"

func main() {
	defer func() { fmt.Println(recover()) }()
	defer panic(1)
	panic(2)
}
