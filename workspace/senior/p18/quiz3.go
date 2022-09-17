// quiz3.go
package main

import "fmt"

func main() {
	defer func() {
		fmt.Println(recover())
	}()
	panic(1)
}
