package main

import "fmt"

const SIZE = 4

func main() {
    /*
    注意：数组的大小不能用变量，比如下面的SIZE必须是常量，如果是变量就会编译报错
     non-constant array bound size
    */
    array := [SIZE]int{1, 2, 3} 
    
    // 方法1：只拿到数组的下标索引
    for index := range array {
        fmt.Printf("index=%d value=%d ", index, array[index])
    }
    fmt.Println()
    
    // 方法2：同时拿到数组的下标索引和对应的值
    for index, value:= range array {
        fmt.Printf("index=%d value=%d ", index, value)
    }
    fmt.Println()
}