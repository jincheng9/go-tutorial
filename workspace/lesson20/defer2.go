package main

import (
    "fmt"
    "sync"
)

func worker(id int) {
    fmt.Println(id)
}

func main() {
    var wg sync.WaitGroup
    size := 10
    wg.Add(size)
    
    for i:=0; i<size; i++ {
        i := i 
        /*把worker的调用和defer放在一个闭包里
        这样worker函数内部就不用使用WaitGroup了
        */
        go func() {
            defer wg.Done()
            worker(i)
        }()
    }
    
    wg.Wait()
}