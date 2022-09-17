# Go Quiz: 从Go面试题看recover注意事项第1篇

## 面试题

这是Go Quiz系列里关于`recover`的第1篇，主要考察recover函数在什么情况下可以捕获到panic。

```go
func main() {
  fmt.Print("1 ")
  defer recover()
  fmt.Print("2 ")
  var a []int
  _ = a[0]
  fmt.Print("3 ")
}
```

- A: 1 2
- B: 1 2 3
- C: 1 2 panic
- D: 1 2 3 panic
- E: 编译报错

这道题有以下几个考点：

* defer语句在什么时候执行
* recover在什么时候可以捕捉到panic

大家可以先暂停，思考下答案。

## 题目解析

我们先来回顾下recover, panic的基础知识点：

recover是Go的内置函数，可以捕获panic异常，但是recover必须结合defer一起使用才能生效。

如果当前goroutine触发了panic，可以在代码的适当位置调用recover函数捕获异常，让程序继续正常执行，而不是异常终止。

如果程序正常执行过程中，没有panic产生，这时调用recover函数会返回nil，除此之外，没有其它任何效果。

recover在以下几种情况返回nil

- panic的参数是nil。这种情况recover捕获后，拿到的返回值也是nil。
- goroutine没有panic产生。没有panic，那当然recover拿到的也就是nil了。
- recover不是**在被defer的函数里面**被**直接调用**执行，这个时候捕获不到panic，recover函数的返回值是nil。

官方文档的如法如下：

> The recover built-in function allows a program to manage behavior of a panicking goroutine. Executing a call to recover inside a deferred function (but not any function called by it) stops the panicking sequence by restoring normal execution and retrieves the error value passed to the call of panic. If recover is called outside the deferred function it will not stop a panicking sequence. In this case, or when the goroutine is not panicking, or if the argument supplied to panic was nil, recover returns nil. Thus the return value from recover reports whether the goroutine is panicking.

recover必须是在被defer的函数里面执行，如果直接在被defer的函数外面执行revcover是不会捕捉到panic的。

对于本题，`defer recover()`这行代码的`recover`并不是在被defer的函数里执行，而是直接跟着defer。

所以这道题目的recover捕获不到panic。因此本题的答案是C，程序输出的结果如下：

```bash
1 2 panic: runtime error: index out of range [0] with length 0

goroutine 1 [running]:
main.main()
        /Users/xxx/quiz2.go:11 +0x106
exit status 2
```



## 思考题

留几道相关的思考题给大家，想知道答案的可以给本人vx公众号发送消息`recover`获取答案和题目解析。

### 题目1

```go
package main

import "fmt"

func main() {
	defer func() { fmt.Println(recover()) }()
	defer panic(1)
	panic(2)
}
```

* A: 1
* B: 2
* C: 先打印1，然后panic
* D: 先打印2，然后panic

### 题目2

```go
func main() {
  	fmt.Println(recover())
    panic(1)
}
```

* A: 抛panic
* B: 打印1

### 题目3

```go
func main() {
    defer fmt.Println(recover())
    panic(1)
}
```

* A: 抛panic
* B: 打印1

### 题目4

```go
func main() {
    defer func() {
      func() { fmt.Println(recover()) }()
    }()
    panic(1)
}
```

* A: 抛panic
* B: 打印1

### 题目5

```go
func main() {
    defer func() {
      fmt.Println(recover())
    }()
    panic(1)
}
```

* A: 抛panic
* B: 打印1



## 推荐阅读

* [Go面试题集锦](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=Mzg2MTcwNjc1Mg==&action=getalbum&album_id=2199553588283179010#wechat_redirect)

* [Go Quiz: Google工程师的Go语言面试题](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483826&idx=1&sn=867f05f3de482259a16369d5e7dff84f&chksm=ce124eddf965c7cb6fee82f567ac86bcf48aaf6bc7c2dc4261c0c9f8f13a2d6f6e060ccb9d16&token=258755563&lang=zh_CN#rd)

* [Go常见错误和最佳实践](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=Mzg2MTcwNjc1Mg==&action=getalbum&album_id=2549657749539028992#wechat_redirect)

  


## 开源地址

文章和示例代码开源在GitHub: [Go语言初级、中级和高级教程](https://github.com/jincheng9/go-tutorial)。

公众号：coding进阶。关注公众号可以获取最新Go面试题和技术栈。

个人网站：[Jincheng's Blog](https://jincheng9.github.io/)。

知乎：[无忌](https://www.zhihu.com/people/thucuhkwuji)。



## 福利

我为大家整理了一份后端开发学习资料礼包，包含编程语言入门到进阶知识(Go、C++、Python)、后端开发技术栈、面试题等。

关注公众号「coding进阶」，发送消息 **backend** 领取资料礼包，这份资料会不定期更新，加入我觉得有价值的资料。还可以发送消息「**进群**」，和同行一起交流学习，答疑解惑。



## References

* https://twitter.com/val_deleplace/status/1564972735540989953
* https://chai2010.gitbooks.io/advanced-go-programming-book/content/appendix/appendix-a-trap.html