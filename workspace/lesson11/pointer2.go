package main

import "fmt"

const SIZE = 5

func main() {
    var ptrArray [SIZE]*int
    a := [5]int{1,2,3,4,5}
    for i:=0; i<SIZE; i++ {
        ptrArray[i] = &a[i]
    }
    
    for i:=0; i<SIZE; i++ {
        fmt.Printf("%d ", *ptrArray[i])
    }
    fmt.Println()
}