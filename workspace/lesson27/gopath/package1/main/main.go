package main

import (
	"fmt"
	"package1/package2"
	"package1/package3"
)

func init() {
	fmt.Println("test1")
}

func init() {
	fmt.Println("test2")
}

func main() {
	fmt.Println(package2.Add(1, 2))
	fmt.Println(package3.GetStr("abc"))
	fmt.Println(sub(1, 2))
}

