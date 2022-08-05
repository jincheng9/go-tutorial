# Go十大常见错误第5篇：go语言Error管理

## 前言

这是Go十大常见错误系列的第5篇：go语言Error管理。素材来源于Go布道者，现Docker公司资深工程师[Teiva Harsanyi](https://teivah.medium.com/)。

本文涉及的源代码全部开源在：[Go十大常见错误源代码](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p28)，欢迎大家关注公众号，及时获取本系列最新更新。



## 场景

Go is still a bit young in its way to deal with errors. It’s not a coincidence if this is one of the most expected features of Go 2.

The current standard library (before Go 1.13) only offers functions to construct errors so you probably want to take a look at [*pkg/errors*](https://github.com/pkg/errors) (if this is not already done).

This library is a good way to respect the following rule of thumb which is not always respected:

*An error should be handled only* ***once\****. Logging an error* ***is\*** *handling an error. So an error should* ***either\*** *be logged or propagated.*

With the current standard library, it is difficult to respect this as we may want to add some context to an error and have some form of hierarchy.

Let’s see an example of what we would expect with a REST call leading to a DB issue:

```
unable to serve HTTP POST request for customer 1234
 |_ unable to insert customer contract abcd
     |_ unable to commit transaction
```

If we use *pkg/errors*, we could do it this way:

<iframe src="https://itnext.io/media/f9de41040d88d7e4cf2dffd98d22f8d8" allowfullscreen="" frameborder="0" height="505" width="692" title="errors.go" class="fp aq as ag cf" scrolling="auto" style="box-sizing: inherit; height: 504.984px; top: 0px; left: 0px; width: 692px; position: absolute;"></iframe>

The initial error (if not returned by an external library) could be created with `errors.New`. The middle layer, `insert`, wraps this error by adding more context to it. Then, the parent handles the error by logging it. Each level either return or handle the error.

We may also want to check at the error cause itself to implement a retry for example. Let’s say we have a `db` package from an external library dealing with the database accesses. This library may return a transient (temporary) error called `db.DBError`. To determine whether we need to retry, we have to check at the error cause:

<iframe src="https://itnext.io/media/f62acbb037a06d9580e40d0e11889c32" allowfullscreen="" frameborder="0" height="527" width="692" title="errors.go" class="fp aq as ag cf" scrolling="auto" style="box-sizing: inherit; height: 527px; top: 0px; left: 0px; width: 692px; position: absolute;"></iframe>

This is done using `errors.Cause` which also comes from *pkg/errors*:

One common mistake I’ve seen was to use *pkg/errors* partially. Checking an error was for example done this way:

<iframe src="https://itnext.io/media/db4ccb35df705f8b09e38fc15d4d5534" allowfullscreen="" frameborder="0" height="197" width="692" title="errors.go" class="fp aq as ag cf" scrolling="auto" style="box-sizing: inherit; height: 197px; top: 0px; left: 0px; width: 692px; position: absolute;"></iframe>

In this example, if the `db.DBError` is wrapped, it will never trigger the retry.

[Don't just check errors, handle them gracefullyThis post is an extract from my presentation at the recent GoCon spring conference in Tokyo, Japan. I've spent a lot of…dave.cheney.net](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully)



## 推荐阅读

* [Go十大常见错误第1篇：未知枚举值](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484146&idx=1&sn=10fb12b643a2e37c090e5aa3bc583152&chksm=ce124d9df965c48bb954aeddabdff3db12738ded3875542250c5d0ef6cfd4417fc56580288b1&token=1912894792&lang=zh_CN#rd)

* [Go十大常见错误第2篇：benchmark性能测试的坑](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484163&idx=1&sn=b28d61c1f3ec9d914e698dce105ba5d1&chksm=ce124c6cf965c57a90bc85a5295ed9375103de20607b509f845583ff6686385df0ed96653d00&token=1912894792&lang=zh_CN#rd)

* [Go十大常见错误第3篇：go指针的性能问题和内存逃逸](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484247&idx=1&sn=faf716627afb00df646cecff023fb63c&chksm=ce124c38f965c52efd009a4c98691d56b5765dc7dce98aa49b226ad9274bd062d8d01e702e91&token=1899277735&lang=zh_CN#rd)

* [Go十大常见错误第4篇：break操作的注意事项](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484262&idx=1&sn=c1bea8af60444a4ef73c4d4d7a09d16d&chksm=ce124c09f965c51f3663ac9089a792d36c3685850e12695dd26d15a1a50f393b2d7c92b9983a&token=461369035&lang=zh_CN#rd)

  

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