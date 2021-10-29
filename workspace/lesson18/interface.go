package main

import "fmt"

func checkType(x int) {
    v := x
    switch v {
    case int:
        fmt.Printf("type: int, value: %v\n", v)
    case string:
        fmt.Printf("type: string，value: %v\n", v)
    case bool:
        fmt.Printf("type: bool, value: %v\n", v)
    default:
        fmt.Printf("type: %T, value: %v\n", x, x)
    }
}


func main() {
    x := 10
    checkType(x)
}