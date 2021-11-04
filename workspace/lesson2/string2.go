package main

import "fmt"

func main() {
    str := "abc"
    /*下标访问*/
    size := len(str)
    for i:=0; i<size; i++ {
        fmt.Printf("%d ", str[i])
    }
    fmt.Println()
    
    /*range迭代访问*/
    for _, value := range str {
        fmt.Printf("%d ", value)
    }
    fmt.Println()
}