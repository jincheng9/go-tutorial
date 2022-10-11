# 并发编程之sync包介绍和sync.WaitGroup

## sync包介绍

sync包提供了基本的并发编程同步原语(concurrency primitives or synchronization primitives)，例如互斥锁sync.Mutex。sync包囊括了以下数据类型：

* sync.Cond
* sync.Locker
* sync.Map
* sync.Mutex
* sync.Once
* sync.Pool
* sync.RWMutex
* sync.WaitGroup

除了sync.Once和sync.WaitGroup这2个类型之外，其它类型主要给一些偏底层的库程序用。业务代码里的goroutine同步，Go设计者是建议通过channel通信来实现。

## sync.WaitGroup

WaitGroup是sync包里的一个结构体类型，定义如下

```go
type WaitGroup struct {
    // some fields
}
```

这个结构体有如下3个方法

* Add：

  ```go
  func (wg *WaitGroup) Add(delta int)
  ```

  

* Done：Done调用会将WiatGroup的计数器减1

  ```go
  func (wg *WaitGroup) Done()
  ```

  

* Wait：Wait调用会阻塞，直到WaitGroup的计数器为0

  ```go
  func (wg *WaitGroup) Wait()
  ```

定义一个WaitGroup变量的目的是为了等待若干个goroutine执行完成，主goroutine调用Add方法，指明要等待的子goroutine数量，这些子goroutine执行完成后调用Done方法。同时，主goroutine要调用Wait方法阻塞程序，等WaitGroup的计数器减小到0时，Wait方法不再阻塞。

示例如下：

```go
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
```

* **注意事项**

  * WaitGroup不要拷贝传值，如果要显式地把WaitGroup作为函数参数，**一定要传指针**。

    WaitGroup给函数A传值，在函数A内部这个WaitGroup会是一个局部变量，对WaitGroup的操作只会在函数内部生效。示例如下：

    ```go
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
    ```

    程序运行时wg.Wait()会报错：fatal error: all goroutines are asleep - deadlock!

    改为下面的传指针就正常了：

    ```go
    package main
    
    import (
        "fmt"
        "sync"
        "time"
    )
    
    
    func worker(id int, wg *sync.WaitGroup) {
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
            /*wg传指针给worker*/
            go worker(i, &wg)
        }
        /* Wait会一直阻塞，直到wg的计数器为0*/
        wg.Wait()
        fmt.Println("end")
    }
    ```


## references

* https://pkg.go.dev/sync@go1.17.2

