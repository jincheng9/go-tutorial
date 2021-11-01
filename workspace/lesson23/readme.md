# sync.Mutex

## 定义

Mutex是sync包里的一个结构体类型，含义就是互斥锁。Mutex变量的默认值或者说零值是一个没有加锁的mutex，也就是当前mutex的状态是unlocked。

不要对Mutex使用值传递方式进行函数调用。

Mutex允许一个goroutine对其加锁，其它goroutine对其解锁，不一定要加解锁在同一个goroutine里。

Mutex结构体类型有2个方法

* Lock()加锁。Lock()方法会把Mutex变量m锁住，如果m已经锁住了，如果再次调用Lock()就会阻塞，知道锁释放。

  ```go
  func (m *Mutex) Lock()
  ```

* Unlock()解锁。Unlock()方法会把Mutex变量m解锁，如果m没有被锁，还去调用Unlock，会遇到runtime error。

  ```go
  func (m *Mutex) Unlock()
  ```

## 使用

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

  