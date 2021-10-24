package main

import "fmt"

/*
3 ways to define variable
*/
var global_int int = 10
var global_str = "hello"
// global_bool := true // this way can only be used inside a function, refer to definition of stock_code3 below


func main() {
    fmt.Println(global_int)
    fmt.Println(global_str)
    
    var stock_code string = "000001.SZ"
    var stock_code2 = "600000.SH"
    stock_code3 := "600570.SH"
    fmt.Println(stock_code)
    fmt.Println(stock_code2)
    fmt.Println(stock_code3)
    
    /*
    Within a function, variable can be declared first, and then assigned value
    */
    var local_var bool
    local_var = true
    fmt.Println(local_var)

    /*
    浮点数只有float32 and float64这2种类型，没有float类型
    */
    var data float32 = 0
    fmt.Println(data)


    var be byte = 255
    fmt.Println(be)

    var ru rune = -2147483648
    fmt.Println("[func|main]ru=", ru)

    var up uintptr = 100
    fmt.Println("[func|main]up=", up)
}
