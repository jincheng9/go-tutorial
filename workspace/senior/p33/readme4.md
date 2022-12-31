# Go 1.20要来了，看看都有哪些变化-第4篇

## 前言

Go官方团队在2022.12.08发布了Go 1.20 rc1(release candidate)版本，Go 1.20的正式release版本预计会在2023年2月份发布。

让我们先睹为快，看看Go 1.20给我们带来了哪些变化。

安装方法：

```bash
$ go install golang.org/dl/go1.20rc1@latest
$ go1.20rc1 download
```

这是Go 1.20版本更新内容详解的第4篇，欢迎大家关注公众号，及时获取本系列最新更新。

## Go 1.20发布清单

和Go 1.19相比，改动内容适中，主要涉及语言(Language)、可移植性(Ports)、工具链(Go Tools)、运行时(Runtime)、编译器(Compiler)、汇编器(Assembler)、链接器(Linker)和核心库(Core library)等方面的优化。

第1篇主要涉及Go 1.20在语言、可移植性方面的优化，原文链接：[Go 1.20版本升级内容第1篇](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484629&idx=1&sn=60a01d3cc85ef2462156f0565c30738d&chksm=ce124bbaf965c2ac351cd9c602e8b67d5119b2a89a7f2de0289bdeb7608ae589c329eb8f7275&token=1619842941&lang=zh_CN#rd)。

第2篇主要涉及Go命令和工具链方面的优化，原文链接：[Go 1.20版本升级内容第2篇](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484638&idx=1&sn=459a22d4a9bf5d9715e70d3c25b05b93&chksm=ce124bb1f965c2a76bacc1135799ab268be66a861e99391b354a9f2dfd8c22a60853cc1d689d&token=1342188569&lang=zh_CN#rd)。

第3篇主要涉及Go在运行时、编译器、汇编器、链接器等方面的优化，原文链接：[Go 1.20版本升级内容第3篇](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484644&idx=1&sn=3c1c4d852b220595ef633f30084f3a11&chksm=ce124b8bf965c29d38c8f17702003c3531b58be15470f7b5c13f67784806a532850b79f464cc&token=1794942092&lang=zh_CN#rd)。

本文重点介绍Go 1.20在核心库方面的优化。

### crypto/ecdh

Go 1.20新增了 [`crypto/ecdh`](https://tip.golang.org/pkg/crypto/ecdh/) 这个package，`ecdh`实现了Elliptic Curve Diffie-Hellman这个新的加密算法。

### 封装多个error

Go 1.20允许一个error变量里封装多个error。

```go
package main

import (
	"errors"
	"fmt"
)

func main() {
	err1 := errors.New("err1")
	err2 := errors.New("err2")
	err := errors.Join(err1, err2)
	fmt.Printf("%T, %v\n", err, err)
	if errors.Is(err, err1) {
		fmt.Println("err is err1")
	}
	if errors.Is(err, err2) {
		fmt.Println("err is err2")
	}
	err3 := fmt.Errorf("error3: %w", err)
	fmt.Printf("%T, %v\n", err3, errors.Unwrap(err3))
	if errors.Is(err3, err1) {
		fmt.Println("err3 is err1")
	}
	if errors.Is(err3, err2) {
		fmt.Println("err3 is err2")
	}
}
```

这段程序的输出结果为：

```bash
*errors.joinError, err1
err2
err is err1
err is err2
*fmt.wrapError, err1
err2
err3 is err1
err3 is err2
```

详情可以参考：https://pkg.go.dev/errors@master#pkg-overview

`fmt.Errorf`里带有`%w`参数，就会返回一个实现了Unwrap方法的error类型的变量。

### HTTP ResponseController

`net/http`这个package新增了名为`ResponseController`的新类型。

```go
func RequestHandler(w ResponseWriter, r *Request) {
  rc := http.NewResponseController(w)
  rc.SetWriteDeadline(0) // disable Server.WriteTimeout when sending a large response
  io.Copy(w, bigData)
}
```

> A ResponseController is used by an HTTP handler to control the response.
>
> A ResponseController may not be used after the Handler.ServeHTTP method has returned.

详情可以参考：https://pkg.go.dev/net/http@master#ResponseController。

### Rewrite钩子函数

 [`httputil.ReverseProxy`](https://tip.golang.org/pkg/net/http/httputil/#ReverseProxy) 类型新增了一个 [`Rewrite`](https://tip.golang.org/pkg/net/http/httputil/#ReverseProxy.Rewrite) 方法，这是一个钩子函数，用来取代之前的`Director`钩子函数。

```go
proxyHandler := &httputil.ReverseProxy{
  Rewrite: func(r *httputil.ProxyRequest) {
    r.SetURL(outboundURL) // Forward request to outboundURL.
    r.SetXForwarded()     // Set X-Forwarded-* headers.
    r.Out.Header.Set("X-Additional-Header", "header set by the proxy")
  },
}
```

详情可以参考：https://pkg.go.dev/net/http/httputil@master#ReverseProxy.Rewrite。

### 标准库的修改

- [bytes](https://tip.golang.org/pkg/bytes/)

  新增了 [`CutPrefix`](https://tip.golang.org/pkg/bytes/#CutPrefix) 和 [`CutSuffix`](https://tip.golang.org/pkg/bytes/#CutSuffix) 函数，这2个函数功能上类似 [`TrimPrefix`](https://tip.golang.org/pkg/bytes/#TrimPrefix) 和 [`TrimSuffix`](https://tip.golang.org/pkg/bytes/#TrimSuffix) ，但是还会返回一个bool类型的变量，表示这个string是否被修改了。

  新增了 [`Clone`](https://tip.golang.org/pkg/bytes/#Clone) 函数，会创建一个byte slice的拷贝。

- [encoding/binary](https://tip.golang.org/pkg/encoding/binary/)

   [`ReadVarint`](https://tip.golang.org/pkg/encoding/binary/#ReadVarint) 和 [`ReadUvarint`](https://tip.golang.org/pkg/encoding/binary/#ReadUvarint)函数如果读的数据的值被损坏，比如只写了一部分内容，会返回 `io.ErrUnexpectedEOF`，而不是像之前返回`io.EOF`。

- [errors](https://tip.golang.org/pkg/errors/)

  新的 `Join`函数可以把多个error变量的值组合在一起，封装为一个新的error变量。

- [fmt](https://tip.golang.org/pkg/fmt/)

  `Errorf`支持`%w`格式化字符串，可以返回一个实现了Unwrap方法的error类型变量。

- [strings](https://tip.golang.org/pkg/strings/)

  新增了`CutPrefix`和 `CutSuffix`函数，这2个函数功能上类似 `TrimPrefix`和`TrimSuffix` ，但是还会返回一个bool类型的变量，表示这个string是否被修改了。

  新增了 `Clone`函数，会创建一个string的拷贝。

- [sync](https://tip.golang.org/pkg/sync/)

   [`Map`](https://tip.golang.org/pkg/sync/#Map) 类型新增了3个新方法： [`Swap`](https://tip.golang.org/pkg/sync/#Map.Swap)， [`CompareAndSwap`](https://tip.golang.org/pkg/sync/#Map.CompareAndSwap) 和[`CompareAndDelete`](https://tip.golang.org/pkg/sync/#Map.CompareAndDelete) ，允许对已有的map做原子更新。

- [testing](https://tip.golang.org/pkg/testing/)

  新增了[`B.Elapsed`](https://tip.golang.org/pkg/testing/#B.Elapsed) 方法，可以返回当前的benchmark性能测试耗时了多久。

- [time](https://tip.golang.org/pkg/time/)

  新增了3个常量`DateTime`， `DateOnly`和`TimeOnly`，方便开发者做格式转换，不用在代码里写死"2006-01-02 15:04:05"。

## 总结

Go 1.20版本即将在2023年2月份发布。

总体而言，Go 1.20版本没有大的修改，主要在细节方面做了各种优化。



## 推荐阅读

* [Go 1.20要来了，看看都有哪些变化-第1篇](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484629&idx=1&sn=60a01d3cc85ef2462156f0565c30738d&chksm=ce124bbaf965c2ac351cd9c602e8b67d5119b2a89a7f2de0289bdeb7608ae589c329eb8f7275&token=1342188569&lang=zh_CN#rd)

* [Go 1.20要来了，看看都有哪些变化-第2篇](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484638&idx=1&sn=459a22d4a9bf5d9715e70d3c25b05b93&chksm=ce124bb1f965c2a76bacc1135799ab268be66a861e99391b354a9f2dfd8c22a60853cc1d689d&token=1342188569&lang=zh_CN#rd)

* [Go 1.20要来了，看看都有哪些变化-第3篇](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484644&idx=1&sn=3c1c4d852b220595ef633f30084f3a11&chksm=ce124b8bf965c29d38c8f17702003c3531b58be15470f7b5c13f67784806a532850b79f464cc&token=1794942092&lang=zh_CN#rd)

* [Go面试题系列，看看你会几题](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=Mzg2MTcwNjc1Mg==&action=getalbum&album_id=2199553588283179010#wechat_redirect)

* [Go常见错误和最佳实践系列](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=Mzg2MTcwNjc1Mg==&action=getalbum&album_id=2549657749539028992#wechat_redirect)

  

## 开源地址

文章和示例代码开源在GitHub: [Go语言初级、中级和高级教程](https://github.com/jincheng9/go-tutorial)。

公众号：coding进阶。关注公众号可以获取最新Go面试题和技术栈。

个人网站：[Jincheng's Blog](https://jincheng9.github.io/)。

知乎：[无忌](https://www.zhihu.com/people/thucuhkwuji)。



## 福利

我为大家整理了一份后端开发学习资料礼包，包含编程语言入门到进阶知识(Go、C++、Python)、后端开发技术栈、面试题等。

关注公众号「coding进阶」，发送消息 **backend** 领取资料礼包，这份资料会不定期更新，加入我觉得有价值的资料。还可以发送消息「**进群**」，和同行一起交流学习，答疑解惑。



## References

* https://tip.golang.org/doc/go1.20