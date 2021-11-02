package main

import (
    "fmt"
    "sync"
    "time"
)


func main() {
    var wg sync.WaitGroup
    /**/
    var mutex sync.Mutex
    cond := sync.NewCond(&mutex)
    size := 10
    wg.Add(size+1)
    
    for i:=0; i<size; i++ {
        i := i
        go func() {
            defer wg.Done()
            /*调用Wait方法时，要对L加锁*/
            cond.L.Lock()
            fmt.Printf("%d ready\n", i)
            /*Wait实际上是会先解锁cond.L，再阻塞当前goroutine
            这样其它goroutine调用上面的cond.L.Lock()才能加锁成功，才能进一步执行到Wait方法，
            等待被Broadcast或者signal唤醒。
            Wait被Broadcast或者Signal唤醒的时候，会再次对cond.L加锁，加锁后Wait才会return
            */
            cond.Wait()
            fmt.Printf("%d done\n", i)
            cond.L.Unlock()
        }()
    }
    
    /*这里sleep 2秒，确保目标goroutine都处于Wait阻塞状态
    如果调用Broadcast之前，目标goroutine不是处于Wait状态，会死锁
    */
    time.Sleep(2*time.Second)
    go func() {
        defer wg.Done()
        cond.Broadcast()
    }()
    wg.Wait()
}