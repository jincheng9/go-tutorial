# Go Quiz: 从Go面试题看锁的注意事项

##  背景

Google工程师Valentin Deleplace出了2道关于锁的题目，拿出来和大家分享下。

### 题目1

```go
// quiz_lock1.go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Mutex
	fmt.Print("1, ")
	m.Lock()
	m.Lock()
	m.Unlock()
	fmt.Println("2")
}
```

* A: `1, 2`

* B: `1, `

* C: `1, fatal error:......`

* D: 编译报错

  

### 题目2

```go
// quiz_lock2.go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m sync.Mutex
	fmt.Print("1, ")
	m.Lock()

	go func() {
		time.Sleep(200 * time.Millisecond)
		m.Unlock()
	}()

	m.Lock()
	fmt.Println("2")
}
```

* A: `1, 2`

* B: `1, `

* C: `1, fatal error:......`

* D: 编译报错

  

## 解析

Go语言里的sync.Mutex和sync.RWMutex都是不可重入的，Go语言里没有可重入锁(也叫递归锁)。

如果mutex没有释放，在同一个goroutine就不能对这个mutex加2次锁，否则第2次加锁就会阻塞，如果没有其它goroutine去释放这个mutex，就会导致死锁，出现runtime error: `fatal error: all goroutines are asleep - deadlock!`。

同时，sync.Mutex和sync.RWMutex允许一个goroutine对其加锁，其它goroutine对其解锁，不要求加锁和解锁在同一个goroutine里。

所以第一道题目的答案是`C`，第二道题目的答案是`A`。



##  思考题

``` go
// quiz_lock3.go

package main

import (
	"fmt"
	"sync"
)

var a sync.Mutex

func main() {
	a.Lock()
	fmt.Print("1, ")
	a.Unlock()
	fmt.Print("2, ")
	a.Unlock()
	fmt.Println("3")
}
```

* A: `1, 2, 3 `
* B: `1, 2, fatal error:......`
* C: `1, 2`
* D: 编译报错

想知道答案的可以给公众号发送消息`mutex`获取答案。



## 总结

Go语言里的锁和C++，Java里的不太一样，给大家总结了以下注意事项

* Go的锁是不可重入的，没有递归锁
* 允许一个goroutine加锁，另一个goroutine解锁，不要求加锁和解锁在同一个goroutine里
* sync.Mutex的零值是没有加锁的Mutex，sync.RWMutex的零值是没有加锁的RWMutex
* 更多细节可以参考References里Mutex和RWMutex的官方说明



## 开源地址

文章和示例代码开源在GitHub: [Go语言初级、中级和高级教程](https://github.com/jincheng9/go-tutorial)。

公众号：coding进阶。关注公众号可以获取最新Go面试题和技术栈。

个人网站：[Jincheng's Blog](https://jincheng9.github.io/)。

知乎：[无忌](https://www.zhihu.com/people/thucuhkwuji)。



## References

* https://github.com/jincheng9/go-tutorial/tree/main/workspace/lesson23
* https://twitter.com/val_deleplace/status/1501922023467106307
* https://twitter.com/val_deleplace/status/1501930369779519491
* https://pkg.go.dev/sync@go1.17.2#Mutex
* https://pkg.go.dev/sync@go1.17.2#RWMutex