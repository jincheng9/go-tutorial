package main

import "fmt"

func main() {
    slice := []int{1,2,3}
    // 方式1
    for index := range slice {
        fmt.Printf("index=%d, value=%d\n", index, slice[index])
    }
    // 方式2
    for index, value := range slice {
        fmt.Printf("index=%d, value=%d\n", index, value)
    }
}