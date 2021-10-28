package main

import "fmt"


func main() {
    num := 10
    var f float32 = float32(num)
    fmt.Println(f) // 10
    
    /*
    不支持隐式类型转换，比如下例想隐式讲num这个int类型转换为float32就会编译报错:
     cannot use num (type int) as type float32 in assignment
     
    var f float32 = num
    */
}