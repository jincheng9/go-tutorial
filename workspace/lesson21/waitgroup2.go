package main

import (
    "fmt"
    "sync"
    "time"
)


func worker(id int, wg sync.WaitGroup) {
    /*worker执行完成后，会调用Done将wg计数器减1*/
    defer wg.Done()
    fmt.Printf("worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup
    /* wg跟踪10个goroutine */
    size := 10
    wg.Add(size)
    /* 开启10个goroutine并发执行 */
    for i:=0; i<size; i++ {
        go worker(i, wg)
    }
    /* 这个例子里Wait会一直阻塞，因为函数worker内部的Done调用对外部的wg其实不生效*/
    wg.Wait()
    fmt.Println("end")
}