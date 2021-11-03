package main

import "fmt"

func main() {
    /*counter没有初始化，给counter赋值会在运行时报错
    panic: assignment to entry in nil map
    */
    var counter map[string]int
    counter["a"] = 1
    fmt.Println(counter)
}