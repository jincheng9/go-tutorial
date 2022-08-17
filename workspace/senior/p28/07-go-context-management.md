# Go十大常见错误第7篇：Context管理

## 前言

这是Go十大常见错误系列的第7篇：Context管理。素材来源于Go布道者，现Docker公司资深工程师[Teiva Harsanyi](https://teivah.medium.com/)。

本文涉及的源代码全部开源在：[Go十大常见错误源代码](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p28)，欢迎大家关注公众号，及时获取本系列最新更新。



## 场景

Go语言里的[context.Context](https://pkg.go.dev/context)非常有用，但是也经常被开发者误解。

官方对Context的表述是：

> A Context carries a deadline, a cancelation signal, and other values across API boundaries.

光看这段描述，还是很容易让人迷糊的，我们接下来具体看看Context到底可以帮助我们做什么事情。

Let’s try to detail it. A context can carry:

- A **deadline**. It means either a duration (e.g. 250 ms) or a date-time (e.g. `2019-01-08 01:00:00`) by which we consider that if it is reached, we must cancel an ongoing activity (an I/O request, awaiting a channel input, etc.).
- A **cancelation signal** (basically a `<-chan struct{}`). Here, the behavior is similar. Once we receive a signal, we must stop an ongoing activity. For example, let’s imagine that we receive two requests. One to insert some data and another one to cancel the first request (because it’s not relevant anymore or whatever). This could be achieved by using a cancelable context in the first call that would be then canceled once we get the second request.
- A list of key/value (both based on an `interface{}` type).

Two things to add. First, a context is **composable**. So, we can have a context that carries a deadline and a list of key/value for example. Moreover, multiple goroutines can **share** the same context so a cancelation signal can potentially stop **multiple activities**.

Coming back to our topic, here is a concrete mistake I’ve seen.

A Go application was based on [*urfave/cli*](https://github.com/urfave/cli) (if you don’t know it, that’s a nice library to create command-line applications in Go). Once started, the developer *inherits* from a sort of application context. It means when the application is stopped, the library will use this context to send a cancellation signal.

What I experienced is that this very context was directly passed while calling a gRPC endpoint for example. This is **not** what we want to do.

Instead, we want to indicate to the gRPC library: *Please cancel the request either when the application is being stopped or after 100 ms for example.*

To achieve this, we can simply create a composed context. If `parent` is the name of the application context (created by *urfave/cli*), then we can simply do this:

```go
ctx, cancel := context.WithTimeout(parent, 100 * time.Millisecond)
response, err := grpcClient.Send(ctx, request)
```

Contexts are not that complex to understand and it is one of the best feature of the language in my opinion.



## 推荐阅读

* [Go十大常见错误第1篇：未知枚举值](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484146&idx=1&sn=10fb12b643a2e37c090e5aa3bc583152&chksm=ce124d9df965c48bb954aeddabdff3db12738ded3875542250c5d0ef6cfd4417fc56580288b1&token=1912894792&lang=zh_CN#rd)

* [Go十大常见错误第2篇：benchmark性能测试的坑](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484163&idx=1&sn=b28d61c1f3ec9d914e698dce105ba5d1&chksm=ce124c6cf965c57a90bc85a5295ed9375103de20607b509f845583ff6686385df0ed96653d00&token=1912894792&lang=zh_CN#rd)

* [Go十大常见错误第3篇：go指针的性能问题和内存逃逸](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484247&idx=1&sn=faf716627afb00df646cecff023fb63c&chksm=ce124c38f965c52efd009a4c98691d56b5765dc7dce98aa49b226ad9274bd062d8d01e702e91&token=1899277735&lang=zh_CN#rd)

* [Go十大常见错误第4篇：break操作的注意事项](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484262&idx=1&sn=c1bea8af60444a4ef73c4d4d7a09d16d&chksm=ce124c09f965c51f3663ac9089a792d36c3685850e12695dd26d15a1a50f393b2d7c92b9983a&token=461369035&lang=zh_CN#rd)

* [Go十大常见错误第5篇：slice初始化常犯的错误](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484274&idx=1&sn=711abea3c6fd5d15341ee1b34da8a160&chksm=ce124c1df965c50b3af84965f7ed30b574cd0b247ea6f77b944ec858bd43ee37f4c1554a5bce&token=1846351524&lang=zh_CN#rd)

* [Go面试题系列，看看你会几题？](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=Mzg2MTcwNjc1Mg==&action=getalbum&album_id=2199553588283179010#wechat_redirect)

  

## 开源地址

文章和示例代码开源在GitHub: [Go语言初级、中级和高级教程](https://github.com/jincheng9/go-tutorial)。

公众号：coding进阶。关注公众号可以获取最新Go面试题和技术栈。

个人网站：[Jincheng's Blog](https://jincheng9.github.io/)。

知乎：[无忌](https://www.zhihu.com/people/thucuhkwuji)。



## 福利

我为大家整理了一份后端开发学习资料礼包，包含编程语言入门到进阶知识(Go、C++、Python)、后端开发技术栈、面试题等。

关注公众号「coding进阶」，发送消息 **backend** 领取资料礼包，这份资料会不定期更新，加入我觉得有价值的资料。还可以发送消息「**进群**」，和同行一起交流学习，答疑解惑。



## References

* https://itnext.io/the-top-10-most-common-mistakes-ive-seen-in-go-projects-4b79d4f6cd65