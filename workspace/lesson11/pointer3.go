package main

import "fmt"

func main() {
    var a int = 100
    var ptr1 *int = &a
    var ptr2 **int = &ptr1
    var ptr3 ***int = &ptr2
    
    fmt.Println("*ptr1=", *ptr1)
    fmt.Println("**ptr2=", **ptr2)
    fmt.Println("***ptr3=", ***ptr3)
}