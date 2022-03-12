# sync.Mutex和sync.RWMutex

## sync.Mutex

### 定义

Mutex是sync包里的一个结构体类型，含义就是互斥锁。Mutex变量的默认值或者说零值是一个没有加锁的mutex，也就是当前mutex的状态是unlocked。

不要对Mutex使用值传递方式进行函数调用。

Mutex允许一个goroutine对其加锁，其它goroutine对其解锁，不要求加锁和解锁在同一个goroutine里。

Mutex结构体类型有2个方法

* Lock()加锁。Lock()方法会把Mutex变量m锁住，如果m已经锁住了，如果再次调用Lock()就会阻塞，直到锁释放。

  ```go
  func (m *Mutex) Lock()
  ```

* Unlock()解锁。Unlock()方法会把Mutex变量m解锁，如果m没有被锁，还去调用Unlock，会遇到runtime error。

  ```go
  func (m *Mutex) Unlock()
  ```

### 不加锁

* 场景举例：多个 goroutine对共享变量同时执行写操作，并发是不安全的，结果和预期不符。

* 示例代码

  ```go
  package main
  
  import (
  	"fmt"
  	"sync"
  )
  
  var sum int = 0
  
  /*多个goroutine同时访问add
  sum是多个goroutine共享的
  也就是多个goroutine同时对共享变量sum做写操作不是并发安全的
  */
  func add(i int) {
  	sum += i
  }
  
  func main() {
  	var wg sync.WaitGroup
  	size := 100
  	wg.Add(size)
  	for i:=1; i<=size; i++ {
  		i := i
  		go func() {
  			defer wg.Done()
  			add(i)
  		}()
  	}
  	wg.Wait()
  	fmt.Printf("sum of 1 to %d is: %d\n", size, sum)
  }
  ```

### 加锁

* 示例代码，通过对共享变量加互斥锁来保证并发安全，结果和预期相符。

  ```go
  package main
  
  import (
  	"fmt"
  	"sync"
  )
  
  var sum int = 0
  var mutex sync.Mutex
  /*多个goroutine同时访问add
  sum是多个goroutine共享的
  通过加互斥锁来保证并发安全
  */
  func add(i int) {
  	mutex.Lock()
  	defer mutex.Unlock()
  	sum += i
  }
  
  func main() {
  	var wg sync.WaitGroup
  	size := 100
  	wg.Add(size)
  	for i:=1; i<=size; i++ {
  		i := i
  		go func() {
  			defer wg.Done()
  			add(i)
  		}()
  	}
  	wg.Wait()
  	fmt.Printf("sum of 1 to %d is: %d\n", size, sum)
  }
  ```




## sync.RWMutex

### 定义

RWMutex是sync包里的一个结构体类型，含义是读写锁。RWMutex变量的零值是一个没有加锁的mutex。

不要对RWMutex变量使用值传递的方式进行函数调用。

RWMutex允许一个goroutine对其加锁，其它goroutine对其解锁，不要求加锁和解锁在同一个goroutine里。

RWMutex结构体类型的定义如下：

```go
type RWMutex struct {
    // some fields
}
```

RWMutex结构体类型有5个方法：

* Lock()，加写锁。某个goroutine加了写锁后，其它goroutine不能获取读锁，也不能获取写锁

  ```go
  func (rw *RWMutex) Lock()
  ```

* Unlock()，释放写锁。

  ```go
  func (rw *RWMutex) Unlock()
  ```

* RLock()，加读锁。某个goroutine加了读锁后，其它goroutine可以获取读锁，但是不能获取写锁

  ```go
  func (rw *RWMutex) RLock()
  ```

* RUnlock()，释放读锁

  ```go
  func (rw *RWMutex) RUnlock()
  ```

* RLocker()，获取一个类型为Locker的接口，Locker类型定义了Lock()和Unlock()方法

  ```go
  func (rw *RWMutex) RLocker() Locker
  ```

  类型Locker的定义如下

  ```go
  type Locker interface {
      Lock()
      Unlock()
  }
  ```

  Mutex和RWMutex这2个结构体类型实现了Locker这个interface里的所有方法，因此可以把Mutex和RWMutex变量或者指针赋值给Locker实例，然后通过Locker实例来加锁和解锁，这个在条件变量sync.Cond里会用到，可以参考[sync.Cond](./workspace/lesson24)

### 示例

```go
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
```



## 注意事项

* Mutex和RWMutex都不是递归锁，不可重入

## References

* https://pkg.go.dev/sync@go1.17.2#Mutex

* https://pkg.go.dev/sync@go1.17.2#RWMutex

