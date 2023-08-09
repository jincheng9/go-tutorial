# Go语言里被defer的函数一定会执行么？

## 前言

大家都知道Go编程中，假设在函数F里，执行了defer A()，那在函数F正常return之前或者因为panic要结束运行之前，被defer关键字修饰的函数调用A()都会被执行到。

比如下面的2个例子：

### test1()会在main结束之前执行

```go
// defer1.go
package main

import (
	"fmt"
)

func test1() {
	fmt.Println("test")
}

func main() {
	fmt.Println("main start")
	defer test1()
	fmt.Println("main end")
}
```

这个例子输出的结果是：

```go
main start
main end
test
```



### test1()会在panic之前执行

```go
// defer2.go
package main

import (
	"fmt"
)

func test1() {
	fmt.Println("test")
}

func test2() {
	panic(1)
}
func main() {
	fmt.Println("main start")
	defer test1()
	test2()
	fmt.Println("main end")
}
```

这个例子输出的结果是：

```go
main start
test
panic: 1

goroutine 1 [running]:
main.test2(...)
        /path/to/defer2.go:13
main.main()
        /path/to/defer2.go:18 +0xb8
exit status 2
```

## 问题

如果在函数F里，defer A()这个语句执行了，是否意味着A()这个函数调用一定会执行？

这里大家可以先脑补一会。



请看下面的例子：

```go
// defer3.go
package main

import (
	"fmt"
	"os"
)

func test1() {
	fmt.Println("test")
}

func main() {
	fmt.Println("main start")
	defer test1()
	fmt.Println("main end")
	os.Exit(0)
}
```

上面的代码运行结果会是怎么样？



## 结论

上面defer3.go执行的结果是：

```go
main start
main end
```

被defer的test1()并没有在main结束之前执行。这是为什么呢？

查看os.Exit的说明如下：

```go
Exit causes the current program to exit with the given status code. 
Conventionally, code zero indicates success, non-zero an error. 
The program terminates immediately; deferred functions are not run.

For portability, the status code should be in the range [0, 125].
```

如果在函数里是因为执行了os.Exit而退出，而不是正常return退出或者panic退出，那程序会立即停止，被defer的函数调用不会执行。

## defer 4原则回顾

1. defer后面跟的必须是函数或者方法调用，defer后面的表达式不能加括号。

   ```go
   defer (fmt.Println(1)) // 编译报错，因为defer后面跟的表达式不能加括号
   ```

2. 被defer的函数的参数在执行到defer语句的时候就被确定下来了。

   ```go
   func a() {
       i := 0
       defer fmt.Println(i) // 最终打印0
       i++
       return
   }
   ```

   上例中，被defer的函数fmt.Println的参数**i**在执行到defer这一行的时候，**i**的值是0，fmt.Println的参数就被确定下来是0了，因此最终打印的结果是0，而不是1。

3. 被defer的函数执行顺序满足LIFO原则，后defer的先执行。

   ```go
   func b() {
       for i := 0; i < 4; i++ {
           defer fmt.Print(i)
       }
   }
   ```

   上例中，输出的结果是3210，后defer的先执行。

4. 被defer的函数可以对defer语句所在的函数的命名返回值做读取和修改操作。

   ```go
   // f returns 42
   func f() (result int) {
   	defer func() {
   		// result is accessed after it was set to 6 by the return statement
   		result *= 7
   	}()
   	return 6
   }
   ```

   上例中，被defer的函数func对defer语句所在的函数**f**的命名返回值result做了修改操作。

   调用函数**f**，返回的结果是42。

   执行顺序是函数**f**先把要返回的值6赋值给result，然后执行被defer的函数func，result被修改为42，然后函数**f**返回result，也就是返回了42。

   官方说明如下：

   ```go
   Each time a "defer" statement executes, the function value and parameters to
   the call are evaluated as usual and saved anew but the actual function is not 
   invoked. Instead, deferred functions are invoked immediately before the 
   surrounding function returns, in the reverse order they were deferred. That
   is, if the surrounding function returns through an explicit return statement, 
   deferred functions are executed after any result parameters are set by that 
   return statement but before the function returns to its caller. If a deferred
   function value evaluates to nil, execution panics when the function is 
   invoked, not when the "defer" statement is executed.
   ```

## go-defer 的实现原理
    ```go
    type _defer struct {
        siz int32 // 参数和返回值的内存大小 
        started bool 
        heap bool //是否分配在堆上面 
        openDefer bool // 是否经过开放编码优化 
        sp uintptr // sp 计数器值，栈指针 
        pc uintptr // pc 计数器值，程序计数器 
        fn *funcval // defer 传入的函数地址，也就是延后执行的函数 
        _panic *_panic // defer 的 panic 结构体 
        link *_defer // 同一个协程里面的defer 延迟函数，会通过该指针连接在一起 
    }
    ```
   内部维护这一个单链表，在对应的函数体里面执行的是单链表的头插法方式，每次在结束时候都是从头开始遍历，也就是FILO（First In Last Out)（先进后出)

## 代码

相关代码和说明开源在GitHub：[被defer的函数是否一定会执行？](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p2)

也可以搜索公众号：coding进阶，查看更多Go知识。

![df](../../official-blog/qrcode_wechat.jpg) 

