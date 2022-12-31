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

  新的 [`Join`](https://tip.golang.org/pkg/errors/#Join) 函数可以封装多个error变量，返回一个新的error变量。

- [fmt](https://tip.golang.org/pkg/fmt/)

  

  The [`Errorf`](https://tip.golang.org/pkg/fmt/#Errorf) function supports multiple occurrences of the `%w` format verb, returning an error that unwraps to the list of all arguments to `%w`.The new [`FormatString`](https://tip.golang.org/pkg/fmt/#FormatString) function recovers the formatting directive corresponding to a [`State`](https://tip.golang.org/pkg/fmt/#State), which can be useful in [`Formatter`](https://tip.golang.org/pkg/fmt/#Formatter). implementations.

- [math/rand](https://tip.golang.org/pkg/math/rand/)

  The [math/rand](https://tip.golang.org/pkg/math/rand/) package now automatically seeds the global random number generator (used by top-level functions like `Float64` and `Int`) with a random value, and the top-level [`Seed`](https://tip.golang.org/pkg/math/rand/#Seed) function has been deprecated. Programs that need a reproducible sequence of random numbers should prefer to allocate their own random source, using `rand.New(rand.NewSource(seed))`.Programs that need the earlier consistent global seeding behavior can set `GODEBUG=randautoseed=0` in their environment.The top-level [`Read`](https://tip.golang.org/pkg/math/rand/#Read) function has been deprecated. In almost all cases, [`crypto/rand.Read`](https://tip.golang.org/pkg/crypto/rand/#Read) is more appropriate.

- [mime](https://tip.golang.org/pkg/mime/)

  The [`ParseMediaType`](https://tip.golang.org/pkg/mime/#ParseMediaType) function now allows duplicate parameter names, so long as the values of the names are the same.

- [mime/multipart](https://tip.golang.org/pkg/mime/multipart/)

  Methods of the [`Reader`](https://tip.golang.org/pkg/mime/multipart/#Reader) type now wrap errors returned by the underlying `io.Reader`.

- [net](https://tip.golang.org/pkg/net/)

  The [`LookupCNAME`](https://tip.golang.org/pkg/net/#LookupCNAME) function now consistently returns the contents of a `CNAME` record when one exists. Previously on Unix systems and when using the pure Go resolver, `LookupCNAME` would return an error if a `CNAME` record referred to a name that with no `A`, `AAAA`, or `CNAME` record. This change modifies `LookupCNAME` to match the previous behavior on Windows, allowing `LookupCNAME` to succeed whenever a `CNAME` exists.[`Interface.Flags`](https://tip.golang.org/pkg/net/#Interface.Flags) now includes the new flag `FlagRunning`, indicating an operationally active interface. An interface which is administratively configured but not active (for example, because the network cable is not connected) will have `FlagUp` set but not `FlagRunning`.The new [`Dialer.ControlContext`](https://tip.golang.org/pkg/net/#Dialer.ControlContext) field contains a callback function similar to the existing [`Dialer.Control`](https://tip.golang.org/pkg/net/#Dialer.Control) hook, that additionally accepts the dial context as a parameter. `Control` is ignored when `ControlContext` is not nil.The Go DNS resolver recognizes the `trust-ad` resolver option. When `options trust-ad` is set in `resolv.conf`, the Go resolver will set the AD bit in DNS queries. The resolver does not make use of the AD bit in responses.DNS resolution will detect changes to `/etc/nsswitch.conf` and reload the file when it changes. Checks are made at most once every five seconds, matching the previous handling of `/etc/hosts` and `/etc/resolv.conf`.

- [net/http](https://tip.golang.org/pkg/net/http/)

  The [`ResponseWriter.WriteHeader`](https://tip.golang.org/pkg/net/http/#ResponseWriter.WriteHeader) function now supports sending `1xx` status codes.The new [`Server.DisableGeneralOptionsHandler`](https://tip.golang.org/pkg/net/http/#Server.DisableGeneralOptionsHandler) configuration setting allows disabling the default `OPTIONS *` handler.The new [`Transport.OnProxyConnectResponse`](https://tip.golang.org/pkg/net/http/#Transport.OnProxyConnectResponse) hook is called when a `Transport` receives an HTTP response from a proxy for a `CONNECT` request.The HTTP server now accepts HEAD requests containing a body, rather than rejecting them as invalid.HTTP/2 stream errors returned by `net/http` functions may be converted to a [`golang.org/x/net/http2.StreamError`](https://tip.golang.org/pkg/golang.org/x/net/http2/#StreamError) using [`errors.As`](https://tip.golang.org/pkg/errors/#As).Leading and trailing spaces are trimmed from cookie names, rather than being rejected as invalid. For example, a cookie setting of "name =value" is now accepted as setting the cookie "name".

- [net/netip](https://tip.golang.org/pkg/net/netip/)

  The new [`IPv6LinkLocalAllRouters`](https://tip.golang.org/pkg/net/netip/#IPv6LinkLocalAllRouters) and [`IPv6Loopback`](https://tip.golang.org/pkg/net/netip/#IPv6Loopback) functions are the `net/netip` equivalents of [`net.IPv6loopback`](https://tip.golang.org/pkg/net/#IPv6loopback) and [`net.IPv6linklocalallrouters`](https://tip.golang.org/pkg/net/#IPv6linklocalallrouters).

- [os](https://tip.golang.org/pkg/os/)

  On Windows, the name `NUL` is no longer treated as a special case in [`Mkdir`](https://tip.golang.org/pkg/os/#Mkdir) and [`Stat`](https://tip.golang.org/pkg/os/#Stat).On Windows, [`File.Stat`](https://tip.golang.org/pkg/os/#File.Stat) now uses the file handle to retrieve attributes when the file is a directory. Previously it would use the path passed to [`Open`](https://tip.golang.org/pkg/os/#Open), which may no longer be the file represented by the file handle if the file has been moved or replaced. This change modifies `Open` to open directories without the `FILE_SHARE_DELETE` access, which match the behavior of regular files.On Windows, [`File.Seek`](https://tip.golang.org/pkg/os/#File.Seek) now supports seeking to the beginning of a directory.

- [os/exec](https://tip.golang.org/pkg/os/exec/)

  The new [`Cmd`](https://tip.golang.org/pkg/os/exec/#Cmd) fields [`Cancel`](https://tip.golang.org/pkg/os/exec/#Cmd.Cancel) and [`WaitDelay`](https://tip.golang.org/pkg/os/exec/#Cmd.WaitDelay) specify the behavior of the `Cmd` when its associated `Context` is canceled or its process exits with I/O pipes still held open by a child process.

- [path/filepath](https://tip.golang.org/pkg/path/filepath/)

  The new error [`SkipAll`](https://tip.golang.org/pkg/path/filepath/#SkipAll) terminates a [`Walk`](https://tip.golang.org/pkg/path/filepath/#Walk) immediately but successfully.The new [`IsLocal`](https://tip.golang.org/pkg/path/filepath/#IsLocal) function reports whether a path is lexically local to a directory. For example, if `IsLocal(p)` is `true`, then `Open(p)` will refer to a file that is lexically within the subtree rooted at the current directory.

- [reflect](https://tip.golang.org/pkg/reflect/)

  The new [`Value.Comparable`](https://tip.golang.org/pkg/reflect/#Value.Comparable) and [`Value.Equal`](https://tip.golang.org/pkg/reflect/#Value.Equal) methods can be used to compare two `Value`s for equality. `Comparable` reports whether `Equal` is a valid operation for a given `Value` receiver.The new [`Value.Grow`](https://tip.golang.org/pkg/reflect/#Value.Grow) method extends a slice to guarantee space for another `n` elements.The new [`Value.SetZero`](https://tip.golang.org/pkg/reflect/#Value.SetZero) method sets a value to be the zero value for its type.Go 1.18 introduced [`Value.SetIterKey`](https://tip.golang.org/pkg/reflect/#Value.SetIterKey) and [`Value.SetIterValue`](https://tip.golang.org/pkg/reflect/#Value.SetIterValue) methods. These are optimizations: `v.SetIterKey(it)` is meant to be equivalent to `v.Set(it.Key())`. The implementations incorrectly omitted a check for use of unexported fields that was present in the unoptimized forms. Go 1.20 corrects these methods to include the unexported field check.

- [regexp](https://tip.golang.org/pkg/regexp/)

  Go 1.19.2 and Go 1.18.7 included a security fix to the regular expression parser, making it reject very large expressions that would consume too much memory. Because Go patch releases do not introduce new API, the parser returned [`syntax.ErrInternalError`](https://tip.golang.org/pkg/regexp/syntax/#ErrInternalError) in this case. Go 1.20 adds a more specific error, [`syntax.ErrLarge`](https://tip.golang.org/pkg/regexp/syntax/#ErrLarge), which the parser now returns instead.

- [runtime/cgo](https://tip.golang.org/pkg/runtime/cgo/)

  Go 1.20 adds new [`Incomplete`](https://tip.golang.org/pkg/runtime/cgo/#Incomplete) marker type. Code generated by cgo will use `cgo.Incomplete` to mark an incomplete C type.

- [runtime/metrics](https://tip.golang.org/pkg/runtime/metrics/)

  Go 1.20 adds new [supported metrics](https://tip.golang.org/pkg/runtime/metrics/#hdr-Supported_metrics), including the current `GOMAXPROCS` setting (`/sched/gomaxprocs:threads`), the number of cgo calls executed (`/cgo/go-to-c-calls:calls`), total mutex block time (`/sync/mutex/wait/total`), and various measures of time spent in garbage collection.Time-based histogram metrics are now less precise, but take up much less memory.

- [runtime/pprof](https://tip.golang.org/pkg/runtime/pprof/)

  Mutex profile samples are now pre-scaled, fixing an issue where old mutex profile samples would be scaled incorrectly if the sampling rate changed during execution.Profiles collected on Windows now include memory mapping information that fixes symbolization issues for position-independent binaries.

- [runtime/trace](https://tip.golang.org/pkg/runtime/trace/)

  The garbage collector's background sweeper now yields less frequently, resulting in many fewer extraneous events in execution traces.

- [strings](https://tip.golang.org/pkg/strings/)

  The new [`CutPrefix`](https://tip.golang.org/pkg/bytes/#CutPrefix) and [`CutSuffix`](https://tip.golang.org/pkg/bytes/#CutSuffix) functions are like [`TrimPrefix`](https://tip.golang.org/pkg/bytes/#TrimPrefix) and [`TrimSuffix`](https://tip.golang.org/pkg/bytes/#TrimSuffix) but also report whether the string was trimmed.The new [`Clone`](https://tip.golang.org/pkg/strings/#Clone) function allocates a copy of a string.

- [sync](https://tip.golang.org/pkg/sync/)

  The new [`Map`](https://tip.golang.org/pkg/sync/#Map) methods [`Swap`](https://tip.golang.org/pkg/sync/#Map.Swap), [`CompareAndSwap`](https://tip.golang.org/pkg/sync/#Map.CompareAndSwap), and [`CompareAndDelete`](https://tip.golang.org/pkg/sync/#Map.CompareAndDelete) allow existing map entries to be updated atomically.

- [syscall](https://tip.golang.org/pkg/syscall/)

  On FreeBSD, compatibility shims needed for FreeBSD 11 and earlier have been removed.On Linux, additional [`CLONE_*`](https://tip.golang.org/pkg/syscall/#CLONE_CLEAR_SIGHAND) constants are defined for use with the [`SysProcAttr.Cloneflags`](https://tip.golang.org/pkg/syscall/#SysProcAttr.Cloneflags) field.On Linux, the new [`SysProcAttr.CgroupFD`](https://tip.golang.org/pkg/syscall/#SysProcAttr.CgroupFD) and [`SysProcAttr.UseCgroupFD`](https://tip.golang.org/pkg/syscall/#SysProcAttr.UseCgroupFD) fields provide a way to place a child process into a specific cgroup.

- [testing](https://tip.golang.org/pkg/testing/)

  The new method [`B.Elapsed`](https://tip.golang.org/pkg/testing/#B.Elapsed) reports the current elapsed time of the benchmark, which may be useful for calculating rates to report with `ReportMetric`.

- [time](https://tip.golang.org/pkg/time/)

  The new time layout constants [`DateTime`](https://tip.golang.org/pkg/time/#DateTime), [`DateOnly`](https://tip.golang.org/pkg/time/#DateOnly), and [`TimeOnly`](https://tip.golang.org/pkg/time/#TimeOnly) provide names for three of the most common layout strings used in a survey of public Go source code.The new [`Time.Compare`](https://tip.golang.org/pkg/time/#Time.Compare) method compares two times.[`Parse`](https://tip.golang.org/pkg/time/#Parse) now ignores sub-nanosecond precision in its input, instead of reporting those digits as an error.The [`Time.MarshalJSON`](https://tip.golang.org/pkg/time/#Time.MarshalJSON) and [`Time.UnmarshalJSON`](https://tip.golang.org/pkg/time/#Time.UnmarshalJSON) methods are now more strict about adherence to RFC 3339.

- [unicode/utf16](https://tip.golang.org/pkg/unicode/utf16/)

  The new [`AppendRune`](https://tip.golang.org/pkg/unicode/utf16/#AppendRune) function appends the UTF-16 encoding of a given rune to a uint16 slice, analogous to [`utf8.AppendRune`](https://tip.golang.org/pkg/unicode/utf8/#AppendRune).

## 总结



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