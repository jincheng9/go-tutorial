# Go十大常见错误第一篇：未知枚举值

## 前言

这是Go十大常见错误的第一篇：未知枚举值。



## 场景

让我们来看下面的代码示例：

```go
type Status uint32

const (
	StatusOpen Status = iota
	StatusClosed
	StatusUnknown
)
```

这里我们使用`iota`定义了一个枚举，对应的枚举值分别是：

```go
StatusOpen = 0
StatusClosed = 1
StatusUnknown = 2
```





## 开源地址

文章和示例代码开源在GitHub: [Go语言初级、中级和高级教程](https://github.com/jincheng9/go-tutorial)。

公众号：coding进阶。关注公众号可以获取最新Go面试题和技术栈。

个人网站：[Jincheng's Blog](https://jincheng9.github.io/)。

知乎：[无忌](https://www.zhihu.com/people/thucuhkwuji)。



## References

* https://itnext.io/the-top-10-most-common-mistakes-ive-seen-in-go-projects-4b79d4f6cd65
* https://gobyexample.com/json