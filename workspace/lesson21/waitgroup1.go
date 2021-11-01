package main

import (
    "fmt"
    "sync"
    "time"
)

var wg sync.WaitGroup

func worker(id int) {
    /*worker执行完成后，会调用Done将wg计数器减1*/
    defer wg.Done()
    fmt.Printf("worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("worker %d done\n", id)
}

func main() {
    /* wg跟踪10个goroutine */
    size := 10
    wg.Add(size)
    /* 开启10个goroutine并发执行 */
    for i:=0; i<size; i++ {
        go worker(i)
    }
    /* Wait一直阻塞，直到wg的计数器变为0 */
    wg.Wait()
    fmt.Println("end")
}