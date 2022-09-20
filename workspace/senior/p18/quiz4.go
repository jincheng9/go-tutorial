// quiz4.go
package main

import "fmt"

func main() {
	defer func() {
		func() { fmt.Println(recover()) }()
	}()
	panic(1)
}
