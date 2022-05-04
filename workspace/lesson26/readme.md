# 并发编程之原子操作sync/atomic

## sync/atomic定义

官方文档地址：https://pkg.go.dev/sync/atomic@go1.18.1

Go语言标准库中的`sync/atomic`包提供了偏底层的原子内存原语(atomic memory primitives)，用于实现同步算法，其本质是将底层CPU提供的原子操作指令封装成了Go函数。

使用`sync/atomic`提供的原子操作可以确保在任意时刻只有一个goroutine对变量进行操作，避免并发冲突。

使用`sync/atomic`需要特别小心，Go官方建议只有在一些偏底层的应用场景里才去使用`sync/atomic`，其它场景建议使用`channel`或者`sync`包里的锁。

> Share memory by communicating; don't communicate by sharing memory.

`sync/atomic`提供了5种类型的原子操作和1个`Value`类型。

### 5种类型的原子操作

* swap操作：`SwapXXX`
* compare-and-swap操作：`CompareAndSwapXXX`
* add操作：`AddXXX`
* load操作：`LoadXXX`
* store操作：`StoreXXX`

这几种类型的原子操作只支持几个基本的数据类型。

add操作的`Addxxx`函数只支持`int32`, `int64`, `uint32`, `uint64`, `uintptr`这5种基本数据类型。

其它类型的操作函数只支持`int32`, `int64`, `uint32`, `uint64`, `uintptr`, `unsafe.Pointer`这6种基本数据类型。

### Value类型

由于上面5种类型的原子操作只支持几种基本的数据类型，因此为了扩大原子操作的使用范围，Go团队在1.4版本的`sync/atomic`包中引入了一个新的类型`Value`。`Value`类型可以用来读取(Load)和修改(Store)**任意类型**的值。

Go 1.4版本的`Value`类型只有`Load`和`Store`2个方法，Go 1.17版本又给`Value`类型新增了`CompareAndSwap`和`Swap`这2个新方法。

## sync/atomic实践

### swap操作

swap操作支持`int32`, `int64`, `uint32`, `uint64`, `uintptr`, `unsafe.Pointer`这6种基本数据类型，对应有6个swap操作函数。

```go
func SwapInt32(addr *int32, new int32) (old int32)
func SwapInt64(addr *int64, new int64) (old int64)
func SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer) (old unsafe.Pointer)
func SwapUint32(addr *uint32, new uint32) (old uint32)
func SwapUint64(addr *uint64, new uint64) (old uint64)
func SwapUintptr(addr *uintptr, new uintptr) (old uintptr)
```

swap操作实现的功能是把`addr` 指针指向的内存里的值替换为新值`new`，然后返回旧值`old`，是如下伪代码的原子实现：

```go
old = *addr
*addr = new
return old
```

我们拿`SwapInt32`举个例子：

```go
// swap.go
package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var newValue int32 = 200
	var dst int32 = 100
	// 把dst的值替换为newValue
	old := atomic.SwapInt32(&dst, newValue)
	// 打印结果
	fmt.Println("old value: ", old, " new value:", dst)
}
```

上面程序的执行结果如下：

```bash
old value:  100  new value: 200
```

### compare-and-swap操作

compare-and-swap(CAS)操作支持`int32`, `int64`, `uint32`, `uint64`, `uintptr`, `unsafe.Pointer`这6种基本数据类型，对应有6个compare-and-swap操作函数。

```go
func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)
func CompareAndSwapInt64(addr *int64, old, new int64) (swapped bool)
func CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer) (swapped bool)
func CompareAndSwapUint32(addr *uint32, old, new uint32) (swapped bool)
func CompareAndSwapUint64(addr *uint64, old, new uint64) (swapped bool)
func CompareAndSwapUintptr(addr *uintptr, old, new uintptr) (swapped bool)
```

compare-and-swap操作实现的功能是先比较`addr` 指针指向的内存里的值是否为旧值`old`相等。

* 如果相等，就把`addr`指针指向的内存里的值替换为新值`new`，并返回`true`，表示操作成功。
* 如果不相等，直接返回`false`，表示操作失败。

compare-and-swap操作是如下伪代码的原子实现：

```go
if *addr == old {
	*addr = new
	return true
}
return false
```

我们拿`CompareAndSwapInt32`举个例子：

```go
// compare-and-swap.go
package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var dst int32 = 100
	oldValue := atomic.LoadInt32(&dst)
	var newValue int32 = 200
	// 先比较dst的值和oldValue的值，如果相等，就把dst的值替换为newValue
	swapped := atomic.CompareAndSwapInt32(&dst, oldValue, newValue)
	// 打印结果
	fmt.Printf("old value: %d, swapped value: %d, swapped success: %v\n", oldValue, dst, swapped)
}
```

上面程序的执行结果如下：

```bash
old value: 100, swapped value: 200, swapped success: true
```

### add操作

add操作支持`int32`, `int64`, `uint32`, `uint64`, `uintptr`这5种基本数据类型，对应有5个add操作函数。

```go
func AddInt32(addr *int32, delta int32) (new int32)
func AddInt64(addr *int64, delta int64) (new int64)
func AddUint32(addr *uint32, delta uint32) (new uint32)
func AddUint64(addr *uint64, delta uint64) (new uint64)
func AddUintptr(addr *uintptr, delta uintptr) (new uintptr)
```

add操作实现的功能是把`addr` 指针指向的内存里的值和`delta`做加法，然后返回新值，是如下伪代码的原子实现：

```go
*addr += delta
return *addr
```

我们拿`AddInt32`举个例子：

```go
// add.go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

// 多个goroutine并发读写sum，有并发冲突，最终计算得到的sum值是不准确的
func test1() {
	var sum int32 = 0
	N := 100
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func(i int32) {
			sum += i
			wg.Done()
		}(int32(i))
	}
	wg.Wait()
	fmt.Println("func test1, sum=", sum)
}

// 使用原子操作计算sum，没有并发冲突，最终计算得到sum的值是准确的
func test2() {
	var sum int32 = 0
	N := 100
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func(i int32) {
			atomic.AddInt32(&sum, i)
			wg.Done()
		}(int32(i))
	}
	wg.Wait()
	fmt.Println("func test2, sum=", sum)
}

func main() {
	test1()
	test2()
}
```

上面程序的执行结果如下：

```bash
func test1, sum= 4857
func test2, sum= 4950
```

**注意**：对于test1函数，你本地运行得到的结果可能和我的不一样，这个值并不是一个固定值。

### load操作

load操作支持`int32`, `int64`, `uint32`, `uint64`, `uintptr`, `unsafe.Pointer`这6种基本数据类型，对应有6个load操作函数。

```go
func LoadInt32(addr *int32) (val int32)
func LoadInt64(addr *int64) (val int64)
func LoadPointer(addr *unsafe.Pointer) (val unsafe.Pointer)
func LoadUint32(addr *uint32) (val uint32)
func LoadUint64(addr *uint64) (val uint64)
func LoadUintptr(addr *uintptr) (val uintptr)
```

load操作实现的功能是返回`addr` 指针指向的内存里的值，是如下伪代码的原子实现：

```go
return *addr
```

我们拿`LoadInt32`举个例子：

```go
// load.go
package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var sum int32 = 100
	result := atomic.LoadInt32(&sum)
	fmt.Println("result=", result)
}
```

上面程序的执行结果如下：

```bash
result= 100
```

### store操作

store操作支持`int32`, `int64`, `uint32`, `uint64`, `uintptr`, `unsafe.Pointer`这6种基本数据类型，对应有6个store操作函数。

```go
func StoreInt32(addr *int32, val int32)
func StoreInt64(addr *int64, val int64)
func StorePointer(addr *unsafe.Pointer, val unsafe.Pointer)
func StoreUint32(addr *uint32, val uint32)
func StoreUint64(addr *uint64, val uint64)
func StoreUintptr(addr *uintptr, val uintptr)
```

store操作实现的功能是把`addr` 指针指向的内存里的值修改为`val`，是如下伪代码的原子实现：

```go
*addr = val
```

我们拿`StoreInt32`举个例子：

```go
// store.go
package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var sum int32 = 100
	var newValue int32 = 200
	// 将sum的值修改为newValue
	atomic.StoreInt32(&sum, newValue)
	// 读取修改后的sum值
	result := atomic.LoadInt32(&sum)
	// 打印结果
	fmt.Println("result=", result)
}
```

上面程序的执行结果如下：

```bash
result= 200
```

#### Value类型

Go标准库里的`sync/atomic`包提供了`Value`类型，可以用来并发读取和修改任何类型的值。

`Value`类型的定义如下：

```go
// A Value provides an atomic load and store of a consistently typed value.
// The zero value for a Value returns nil from Load.
// Once Store has been called, a Value must not be copied.
//
// A Value must not be copied after first use.
type Value struct {
	v any
}
```

`Value`类型有4个方法：`CompareAndSwap`,  `Load`, `Store`, `Swap`，定义如下：

```go
func (v *Value) CompareAndSwap(old, new any) (swapped bool)
func (v *Value) Load() (val any)
func (v *Value) Store(val any)
func (v *Value) Swap(new any) (old any)
```

源码实现：https://cs.opensource.google/go/go/+/refs/tags/go1.18.1:src/sync/atomic/value.go

下面是一个具体的示例：对`map[string][string]`类型做并发读写，为了避免加锁，使用`value`类型来读取和修改`map[string][string]`。

```go
package main

import (
	"sync/atomic"
	"time"
)

func loadConfig() map[string]string {
	// 从数据库或者文件系统中读取配置信息，然后以map的形式存放在内存里
	return make(map[string]string)
}

func requests() chan int {
	// 将从外界中接收到的请求放入到channel里
	return make(chan int)
}

func main() {
	// config变量用来存放该服务的配置信息
	var config atomic.Value
	// 初始化时从别的地方加载配置文件，并存到config变量里
	config.Store(loadConfig())
	go func() {
		// 每10秒钟定时拉取最新的配置信息，并且更新到config变量里
		for {
			time.Sleep(10 * time.Second)
			// 对应于赋值操作 config = loadConfig()
			config.Store(loadConfig())
		}
	}()
	// 创建协程，每个工作协程都会根据它所读取到的最新的配置信息来处理请求
	for i := 0; i < 10; i++ {
		go func() {
			for r := range requests() {
				// 对应于取值操作 c := config
				// 由于Load()返回的是一个interface{}类型，所以我们要先强制转换一下
				c := config.Load().(map[string]string)
				// 这里是根据配置信息处理请求的逻辑...
				_, _ = r, c
			}
		}()
	}
}
```



## 总结和注意事项

* 原子操作由底层CPU的原子操作指令支持。

* 5种原子操作和`Value`类型的官方文档地址：https://pkg.go.dev/sync/atomic@go1.18.1

* CAS操作会有ABA问题

* 对于386处理器架构，64-bit原子操作函数使用了奔腾MMX或更新处理器型号才支持的CPU指令。对于非Linux的ARM处理器架构，64-bit原子操作函数使用了ARMv6k core或更新处理器型号才支持的CPU指令。对于ARM, 386和32-bit MIPS处理器架构，原子操作的调用者要对进行原子访问的64bit字(word)按照64-bit进行内存对齐。变量或者分配的结构体、数组和切片的第1个字可以认为是64-bit对齐的。(这块涉及到内存对齐，后面抽个专题详解)

  

## 开源地址

文章和示例代码开源在GitHub: [Go语言初级、中级和高级教程](https://github.com/jincheng9/go-tutorial)。

公众号：coding进阶。关注公众号可以获取最新Go面试题和技术栈。

个人网站：[Jincheng's Blog](https://jincheng9.github.io/)。

知乎：[无忌](https://www.zhihu.com/people/thucuhkwuji)。



## References

* https://pkg.go.dev/sync/atomic@go1.18.1
* https://blog.betacat.io/post/golang-atomic-value-exploration/
* https://gfw.go101.org/article/concurrent-atomic-operation.html