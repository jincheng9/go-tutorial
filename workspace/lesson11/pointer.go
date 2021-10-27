package main

import "fmt"
import "reflect"

func main() {
    i := 10
    // 方式1
    var intPtr *int = &i
    fmt.Println("pointer value:", intPtr, " point to: ", *intPtr)
    fmt.Println("type of pointer:", reflect.TypeOf(intPtr))
    
    // 方式2
    intPtr2 := &i
    fmt.Println(*intPtr2)
    fmt.Println("type of pointer:", reflect.TypeOf(intPtr2))
    
    // 方式3
    var intPtr3 = &i;
    fmt.Println(*intPtr3)
    fmt.Println("type of pointer:", reflect.TypeOf(intPtr3))
    
    // 方式4
    var intPtr4 *int
    intPtr4 = &i
    fmt.Println(*intPtr4)
    fmt.Println("type of pointer:", reflect.TypeOf(intPtr4))
    
    // 不赋值的时候，默认值为nil
    var intPtr5 *int
    fmt.Println("intPtr5==nil:", intPtr5==nil)
    
    array := [3]int{1,2,3}
    var arrayPtr *[3]int = &array
    for i:=0; i<len(array); i++ {
        fmt.Printf("arrayPtr[%d]=%d\n", i, arrayPtr[i])
    }
 }