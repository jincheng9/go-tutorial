package main

import "fmt"
import "reflect"


func printSlice(param []int) {
    fmt.Printf("param len:%d, cap:%d, value:%v\n", len(param), cap(param), param)
}

func main() {
    slice := []int{}
    var slice2 []int
    
    fmt.Println("slice==nil", slice==nil) // false
    printSlice(slice)
    
    fmt.Println("slice2==nil", slice2==nil) // true
    printSlice(slice2)
    
    // 对数组做切片
    array := [3]int{1,2,3} // array是数组
    slice3 := array[1:3] // slice3是切片
    fmt.Println("slice3 type:", reflect.TypeOf(slice3))
    fmt.Println("slice3=", slice3) // slice3= [2 3]
    
    slice4 := slice3[1:2]
    fmt.Println("slice4=", slice4) // slice4= [3]
    
    /* slice5->slice4->slice3->array
    对slice5的修改，会影响到slice4, slice3和array
    */
    slice5 := slice4[:]
    fmt.Println("slice5=", slice5) // slice5= [3]
    
    slice5[0] = 10
    fmt.Println("array=", array) // array= [1 2 10]
    fmt.Println("slice3=", slice3) // slice3= [2 10]
    fmt.Println("slice4=", slice4) // slice4= [10]
    fmt.Println("slice5=", slice5) // slice5= [10]
}