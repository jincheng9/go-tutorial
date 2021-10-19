package main

import "fmt"

func main() {
    fmt.Println("Hello World!")
    test()
}

func init() {
    fmt.Println("Init!")
}

func test() {
     fmt.Println("test!")
}