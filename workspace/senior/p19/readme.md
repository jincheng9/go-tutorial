# Go语言中fmt.Println(true)的结果一定是true么？

##  背景

[Honeycomb](https://www.honeycomb.io/)的首席开发布道者[Jessica](https://twitter.com/jessitron)在Twitter上发了一条状态：

> fmt.Println("What is truth?", true)
>
> can output:
>
> What is truth? false

意思是下面代码的执行结果可能是`What is truth? false`

```go
fmt.Println("What is truth?", true)
```

大家可以先想想什么情况下会出现这样的结果。



## 解析

我们来看下面这段代码：

```go
// identifier.go
package main

import "fmt"

func main() {
	true := false
	fmt.Println(true)
}
```

大家觉得这段代码是会编译报错呢，还是正常执行？

实际执行结果是打印`false`，并不会编译报错。

因此本文开头的代码`fmt.Println("What is truth?", true)`是可能打印`What is truth? false`的。

**有的同学可能好奇了？为什么会这样呢？`true`不是Go语言的关键字么，为什么还可以定义标识符为`true` 的变量？**

答案是：`true` 并不是Go语言的关键字，Go语言目前只有25个关键字，如下所示：

> ```
> break        default      func         interface    select
> case         defer        go           map          struct
> chan         else         goto         package      switch
> const        fallthrough  if           range        type
> continue     for          import       return       var
> ```

这些[关键字](https://go.dev/ref/spec#Keywords)是不可以用来作为Go语言的标识符的。`true`是[预声明标识符](https://go.dev/ref/spec#Predeclared_identifiers)，是可以作为Go语言的标识符的，官方说明如下：

> ### Predeclared identifiers
>
> The following identifiers are implicitly declared in the [universe block](https://go.dev/ref/spec#Blocks):
>
> ```
> Types:
> 	bool byte complex64 complex128 error float32 float64
> 	int int8 int16 int32 int64 rune string
> 	uint uint8 uint16 uint32 uint64 uintptr
> 
> Constants:
> 	true false iota
> 
> Zero value:
> 	nil
> 
> Functions:
> 	append cap close complex copy delete imag len
> 	make new panic print println real recover
> ```

因此`true := false`这样的代码在Go语言是可以正常编译通过的，并且`go vet`也不会检测出任何潜在错误。

不仅仅是`true`，预声明标识符里的所有标识符都可以用来作为全局变量和局部变量的标识符，比如下面的代码：

```go
// identifier2.go
package main

import "fmt"

var nil = 100

var false = true

func main() {
	true := false
	fmt.Println(true, false, nil)
}
```

大家可以想想输出结果是什么？

* A: `true false nil`
* B: `true false 100`
* C: `true true 100`
* D: `false true 100`
* E: `false true nil`

想知道答案的可以给公众号发送消息`bsf`获取答案。



## 总结

Go语言的这个特性也引起了很多[争议](https://twitter.com/jessitron/status/1492940889043443723)，正如Go的错误处理一样。

我们作为使用者，需要注意：**Go不允许把关键字作为标识符，其它都可以作为标识符**。



## 开源地址

文章和示例代码开源在GitHub: [Go语言初级、中级和高级教程](https://github.com/jincheng9/go-tutorial)。

公众号：coding进阶。关注公众号可以获取最新Go面试题和技术栈。

个人网站：[Jincheng's Blog](https://jincheng9.github.io/)。

知乎：[无忌](https://www.zhihu.com/people/thucuhkwuji)。



## References

* https://twitter.com/jessitron/status/1492940889043443723
* https://go.dev/ref/spec#Keywords
* https://go.dev/ref/spec#Predeclared_identifiers