# sync.Once

## 定义

Once是sync包里的一个结构体类型，Once可以在并发场景下让某个操作只执行一次，比如设计模式里的单例只创建一个实例，比如只加载一次配置文件，比如对同一个channel只关闭一次（对一个已经close的channel再次close会引发panic）等。

定义如下：

```go
type Once struct {
    // some fields
}
```

这个结构体只有1个方法Do，参数是要执行的函数。（**注意**：参数是函数类型，而不是函数的返回值，所以只需要把函数名作为参数给到Do即可）

可以看到Do方法的参数**f**这个函数类型没有参数，所以如果要执行的函数f需要传递参数就要结合Go的闭包来使用。

```go
func(o *Once) Do(f func())
```

参考下面的例子，print函数通过Once执行，只会执行1次

```go
package main

import (
    "fmt"
    "sync"
)

func print() {
    fmt.Println("test once")
}

func main() {
    var wg sync.WaitGroup
    var once sync.Once
    size := 10
    wg.Add(size)
    
    /*启用size个goroutine，每个goroutine都调用once.Do(print)
    最终print只会执行一次
    */
    for i:=0; i<size; i++ {
        go func() {
            defer wg.Done()
            once.Do(print)
        }()
    }
    /*等待所有goroutine执行完成*/
    wg.Wait()
    fmt.Println("end")
}
```



## sync.Once实现并发安全的单例

```go
package main

import (
    "fmt"
    "sync"
)

type Singleton struct {
    member int
}

var instance *Singleton

var once sync.Once

func getInstance() *Singleton {
    /*
    通过sync.Once实现单例，只会生成一个Singleton实例
    */
    once.Do(func() {
        fmt.Println("once")
        instance = &Singleton{}
        instance.member = 100
    })
    fmt.Println(instance.member)
    return instance
}

func main() {
    var wg sync.WaitGroup
    size := 10
    wg.Add(size)
    /*
    多个goroutine同时去获取Singelton实例
    */
    for i:=0; i<size; i++ {
        go func() {
            defer wg.Done()
            instance = getInstance()
        }()
    }
    wg.Wait()
    fmt.Println("end")
}
```



## 注意事项

* Once变量作为函数参数传递时，只能传指针，不能传值。传值给函数A的话，对于函数A而言，参数列表里的once形参会是一个新生成的once局部变量，和外部传入的once实参不一样。

  ```go
  package main
  
  import (
      "fmt"
      "sync"
  )
  
  func test() {
      fmt.Println("test once")
  }
  
  func print(once *sync.Once) {
      once.Do(test)
  }
  
  func main() {
      var wg sync.WaitGroup
      var once sync.Once
      size := 10
      wg.Add(size)
      
      /*启用size个goroutine，每个goroutine都调用once.Do(print)
      最终print只会执行一次
      */
      for i:=0; i<size; i++ {
          go func() {
              defer wg.Done()
              print(&once)
          }()
      }
      /*等待所有goroutine执行完成*/
      wg.Wait()
      fmt.Println("end")
  }
  ```

  

* 如果once.Do(f)方法调用的函数**f**发生了panic，那Do也会认为函数**f**已经return了。

* 如果多个goroutine执行了都去调用once.Do(f)，只有某次的函数**f**调用返回了，所有Do方法调用才会返回，否则Do方法会一直阻塞等待。如果在f里继续调用同一个once变量的Do方法，就会死锁了，因为Do在等待**f**返回，**f**又在等待Do返回。

