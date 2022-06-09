# Go十大常见错误第1篇：未知枚举值

## 前言

这是Go十大常见错误系列的第一篇：未知枚举值。素材来源于Go布道者，现Docker公司资深工程师[**Teiva Harsanyi**](https://teivah.medium.com/)。

本文涉及的源代码全部开源在：[Go十大常见错误源代码](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p28)，欢迎大家关注公众号，及时获取本系列最新更新。



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

假设我们业务代码里的数据结构包含了枚举类型，比如下例：

```go
type Request struct {
	ID        int    `json:"Id"`
	Timestamp int    `json:"Timestamp"`
	Status    Status `json:"Status"`
}
```

我们要把接收到的JSON请求反序列化为`Request`结构体类型。

```go
{
  "Id": 1234,
  "Timestamp": 1563362390,
  "Status": 0
}
```

对于上面这个JSON请求数据，`Request`结构体里的`Status`字段会被解析为0，对应的是`StatusOpen`，符合预期。

但是如果由于各种原因没有传`Status`字段，对于如下的JSON请求

```go
{
  "Id": 1235,
  "Timestamp": 1563362390
}
```

在将这个JSON请求反序列化为`Request`结构体类型的时候，因为JSON串里没有`Status`字段，因此`Request`结构体里的`Status`字段的值会是零值，也就是`uint32`的零值0。这个时候`Status`字段的值还是`StatusOpen`，而不是我们预期的`StatusUnknown`。



## 最佳实践

因此对于枚举值的最佳实践，是把枚举的未知值设置为0。

```go
type Status uint32

const (
	StatusUnknown Status = iota
	StatusOpen
	StatusClosed
)
```

这样设计后，如果JSON请求里没有传`Status`字段，那反序列化后的`Request`结构体里的`Status`字段的值就是`StatusUnknown`，符合预期。



## 开源地址

文章和示例代码开源在GitHub: [Go语言初级、中级和高级教程](https://github.com/jincheng9/go-tutorial)。

公众号：coding进阶。关注公众号可以获取最新Go面试题和技术栈。

个人网站：[Jincheng's Blog](https://jincheng9.github.io/)。

知乎：[无忌](https://www.zhihu.com/people/thucuhkwuji)。



## References

* https://itnext.io/the-top-10-most-common-mistakes-ive-seen-in-go-projects-4b79d4f6cd65
* https://gobyexample.com/json