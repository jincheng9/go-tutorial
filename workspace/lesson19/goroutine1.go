package main

import "fmt"

func hello() {
    fmt.Println("hello")
}

func main() {
    go hello()
    fmt.Println("main end")
}