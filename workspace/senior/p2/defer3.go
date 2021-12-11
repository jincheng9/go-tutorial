// defer3.go
package main

import (
	"fmt"
	"os"
)

func test1() {
	fmt.Println("test")
}

func main() {
	fmt.Println("main start")
	defer test1()
	fmt.Println("main end")
	os.Exit(0)
}
