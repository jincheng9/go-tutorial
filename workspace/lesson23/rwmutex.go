package main

import (
    "fmt"
    "sync"
)

type Counter struct {
    /*
    成员count:计数器
    成员rw: 读写锁，用于实现count的读写并发安全
    */
    count int
    rw sync.RWMutex
}

func (c *Counter) getCounter() int{
    /*
    读数据的时候加读锁
    */
    c.rw.RLock()
    defer c.rw.RUnlock()
    return c.count
}

func (c *Counter) add() {
    /*
    写数据的时候加写锁
    */
    c.rw.Lock()
    defer c.rw.Unlock()
    c.count++
}

func main() {
    var wg sync.WaitGroup
    size := 100
    wg.Add(size)
    
    var c Counter
    /*
    开启size个goroutine对变量c的数据成员count同时进行读写操作
    */
    for i:=0; i<size; i++ {
        go func() {
            defer wg.Done()
            c.getCounter()
            c.add()
        }()
    }
    wg.Wait()
    fmt.Println("count=", c.count)
}