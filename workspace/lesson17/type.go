package main

import "fmt"

func main() {
    total_weight := 100
    num := 12
    // total_weight和num都是整数，相除结果还是整数
    fmt.Println("average=", total_weight/num) //  average= 8
    
    // 转成float32再相除，结果就是准确值了
    fmt.Println("average=", float32(total_weight)/float32(num)) // average= 8.333333
    
    /* 注意，float32只能和float32做运算，否则会报错，比如下例里float32和int相加，编译报错:
    invalid operation: float32(total_weight) + num (mismatched types float32 and int)
   
    res := float32(total_weight) + num
    fmt.Println(res)
    */
}