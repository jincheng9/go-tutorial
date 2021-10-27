package main

import "fmt"

// 这个可以交换外部传入的2个实参的值
func swap(a *int, b *int) {
    *a, *b = *b, *a
}

// 这个无法交换外部传入的2个实参的值
func swap2(a *int, b *int) {
    a, b = b, a
}


func main() {
    a, b := 1, 2
    swap(&a, &b)
    fmt.Println("a=", a, " b=", b) // a= 2  b= 1
    
    swap2(&a, &b)
    fmt.Println("a=", a, " b=", b) // a= 2  b= 1
}