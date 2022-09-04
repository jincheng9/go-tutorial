# Go十大常见错误第8篇：并发编程中Context使用常见错误

## 前言

这是Go十大常见错误系列的第8篇：并发编程中Context使用常见错误。素材来源于Go布道者，现Docker公司资深工程师[Teiva Harsanyi](https://teivah.medium.com/)。

本文涉及的源代码全部开源在：[Go十大常见错误源代码](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p28)，欢迎大家关注公众号，及时获取本系列最新更新。



## Context是什么

Go语言标准库里有一个package叫`context`，该package里定义了[context.Context](https://pkg.go.dev/context)类型，在并发编程里非常有用，但是也经常被开发者误解。

官方对Context的表述是：

> Package context defines the Context type, which carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes.

光看这段描述，还是很容易让人迷糊的，我们接下来具体看看Context到底是什么以及可以帮助我们做什么事情。

Context顾名思义，表示的是goroutine的上下文，Context定义如下所示：

```go
// A Context carries a deadline, cancellation signal, and request-scoped values
// across API boundaries. Its methods are safe for simultaneous use by multiple
// goroutines.
type Context interface {
    // Done returns a channel that is closed when this Context is canceled
    // or times out.
    Done() <-chan struct{}

    // Err indicates why this context was canceled, after the Done channel
    // is closed.
    Err() error

    // Deadline returns the time when this Context will be canceled, if any.
    Deadline() (deadline time.Time, ok bool)

    // Value returns the value associated with key or nil if none.
    Value(key interface{}) interface{}
}
```

Context可以通过超时设置、携带取消信号、附加参数信息来方便goroutine里做相应的逻辑控制。

- 超时控制。 通过`context.WithTimeout`函数和`context.WithDeadline`函数可以创建一个有超时时间的Context。通过Context的`Done`函数可以判断是否超时了。

  ```go
  func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
  func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)
  ```

- 取消信号。通过`context.WithCancel`函数可以创建一个接收cancel信号的Context。通过Context的`Done`函数可以判断是否发出了cancel信号。父Context发出的cancel信号，子Context也可以接收到。

  ```go
  func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
  ```

- 附加参数信息。通过`context.WithValue`函数可以给Context添加参数。其中key和value都是空接口类型(`interface{}`)。通过Context的`Value`函数可以获取附加参数信息。

  ```go
  func WithValue(parent Context, key, val any) Context
  ```

在实际开发过程中，Context的使用流程一般是：

* Step 1: 创建Context，给Context指定超时时间，设置取消信号，或者附加参数(链路跟踪里经常使用Context里的附加参数，传递IP等链路跟踪信息)。
* Step 2: goroutine使用Step 1里的Context作为第一个参数，在该goroutine里就可以做如下事情：
  * 使用Context里的`Done`函数判断是否达到了Context设置的超时时间或者Context是否被主动取消了。
  * 使用Context里的`Value`函数获取该Context里的附加参数信息。
  * 使用Context里的`Err`函数获取错误原因，目前原因就2个，要么是超时，要么是主动取消。

有2点要补充：

* 第一，Context是可以组合的。比如，我们可以通过`context.WithTimeout`创建一个有超时时间的Context，再调用`context.WithValue`添加一些附加参数信息。
* 第二，多个goroutine可以共享同一个Context，可以通过该Context的超时设置、携带取消信号以及附加参数来控制多个goroutine的行为。



## 常见错误

在Context使用过程中有以下几个常见错误：

* 第一，不执行`cancel`函数去释放Context资源。

  * 对于`context.WithTimeout`、`context.WithDeadline`、`context.WithCancel`函数返回的cancel函数，需要做执行。官方说明如下：

    > Canceling this context releases resources associated with it, so code should call cancel as soon as the operations running in this Context complete:

    参考代码示例：

    ```go
    func slowOperationWithTimeout(ctx context.Context) (Result, error) {
    	ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
    	defer cancel()  // releases resources if slowOperation completes before timeout elapses
    	return slowOperation(ctx)
    }
    ```

* 第二，不加超时控制，如果执行了非常耗时的rpc操作或者数据库操作，就会阻塞程序。如果rpc调用接口或者数据库操作接口支持传递Context参数，建议加上超时设置。代码示例参考如下：

  ```go
  ctx, cancel := context.WithTimeout(parent, 100 * time.Millisecond)
  response, err := grpcClient.Send(ctx, request)
  ```



## 推荐阅读

* [Go十大常见错误第1篇：未知枚举值](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484146&idx=1&sn=10fb12b643a2e37c090e5aa3bc583152&chksm=ce124d9df965c48bb954aeddabdff3db12738ded3875542250c5d0ef6cfd4417fc56580288b1&token=1912894792&lang=zh_CN#rd)

* [Go十大常见错误第2篇：benchmark性能测试的坑](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484163&idx=1&sn=b28d61c1f3ec9d914e698dce105ba5d1&chksm=ce124c6cf965c57a90bc85a5295ed9375103de20607b509f845583ff6686385df0ed96653d00&token=1912894792&lang=zh_CN#rd)

* [Go十大常见错误第3篇：go指针的性能问题和内存逃逸](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484247&idx=1&sn=faf716627afb00df646cecff023fb63c&chksm=ce124c38f965c52efd009a4c98691d56b5765dc7dce98aa49b226ad9274bd062d8d01e702e91&token=1899277735&lang=zh_CN#rd)

* [Go十大常见错误第4篇：break操作的注意事项](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484262&idx=1&sn=c1bea8af60444a4ef73c4d4d7a09d16d&chksm=ce124c09f965c51f3663ac9089a792d36c3685850e12695dd26d15a1a50f393b2d7c92b9983a&token=461369035&lang=zh_CN#rd)

* [Go十大常见错误第5篇：Go语言Error管理](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484274&idx=1&sn=711abea3c6fd5d15341ee1b34da8a160&chksm=ce124c1df965c50b3af84965f7ed30b574cd0b247ea6f77b944ec858bd43ee37f4c1554a5bce&token=1846351524&lang=zh_CN#rd)

* [Go十大常见错误第6篇：slice初始化常犯的错误](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484289&idx=1&sn=2b8171458cde4425b28fdf8f51df8d7c&chksm=ce124ceef965c5f8a14f5951457ce2ac0ecc4612cf2013957f1d818b6e74da7c803b9df1d394&token=1477304797&lang=zh_CN#rd)

* [Go十大常见错误第7篇：不使用-race选项做并发竞争检测](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484299&idx=1&sn=583c3470a76e93b0af0d5fc04fe29b55&chksm=ce124ce4f965c5f20de5887b113eab91f7c2654a941491a789e4ac53c298fbadb4367acee9bb&token=1918756920&lang=zh_CN#rd)

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

* 参考文章：https://itnext.io/the-top-10-most-common-mistakes-ive-seen-in-go-projects-4b79d4f6cd65
* 官方文档：https://pkg.go.dev/context
* 官方Context入门介绍：https://go.dev/blog/context
* Context使用介绍：https://mp.weixin.qq.com/s/PoXSEDHRyKCyjibFGS0wHw
* https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go