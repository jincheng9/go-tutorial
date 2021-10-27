package main

import "fmt"


func printSlice(param []int) {
    fmt.Printf("slice len:%d, cap:%d, value:%v\n", len(param), cap(param), param)
}

func main() {
    slice1 := []int{1}
    slice2 := make([]int, 3, 100)
    printSlice(slice1)
    printSlice(slice2)
}