package main

import "fmt"


// 计算n的阶乘
func factorial(n int) int {
    if n == 0 || n == 1 {
        return 1
    } else {
        return n * factorial(n-1)
    }
}

func main() {
    sum := factorial(5)
    fmt.Println("5!=", sum)
}