# sync.Cond

## 定义

Cond是sync包里的一个结构体类型，表示条件变量。我们知道sync.WaitGroup可以用于等待所有goroutine都执行完成，**sync.Cond可以用于控制goroutine什么时候开始执行**。

Cond结构体类型定义如下：

```go
type Cond struct {
    L Locker
    // some other fields
}
```

Cond结构体类型以下几个方法与其紧密相关：

* NewCond函数，用于创建条件变量，条件变量的成员`L`是NewCond函数的参数`l`

  ```go
  func NewCond(l Locker) *Cond
  ```

  

* Broadcast，发出广播，唤醒所有等待条件变量c的goroutine开始执行。**注意**：在调用Broadcast方法之前，要确保目标goroutine处于Wait阻塞状态，不然会出现死锁问题。

  ```go
  func (c *Cond) Broadcast()
  ```

  

* Signal，发出信号，唤醒某一个等待条件变量c的goroutine开始执行。**注意**：在调用Signal方法之前，要确保目标goroutine处于Wait阻塞状态，不然会出现死锁问题。

  ```go
  func (c *Cond) Signal()
  ```

  

* Wait，这个方法会解锁c.L以及阻塞当前goroutine往下执行，解锁和阻塞组合在一起构成原子操作。Wait被Broadcast或者Signal唤醒时，会先对c.L加锁，然后Wait才return返回。

  ```go
  func (c *Cond) Wait()
  ```

  

每个Cond变量都有一个Locker类型的成员L，L通常是\*Mutex或者\*RWMutex类型，**调用Wait方法时要对L加锁**。

不要对Cond变量使用值传递进行函数调用。

## 示例

下面这个示例，先开启了10个goroutine，这10个goroutine都进入Wait阻塞状态，等待被唤醒。

```go
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
```



## References

https://pkg.go.dev/sync@go1.17.2#Cond

