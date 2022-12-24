# Go 1.20要来了，看看都有哪些变化-第3篇

## 前言

Go官方团队在2022.12.08发布了Go 1.20 rc1(release candidate)版本，Go 1.20的正式release版本预计会在2023年2月份发布。

让我们先睹为快，看看Go 1.20给我们带来了哪些变化。

安装方法：

```bash
$ go install golang.org/dl/go1.20rc1@latest
$ go1.20rc1 download
```

这是Go 1.20版本更新内容详解的第3篇，欢迎大家关注公众号，及时获取本系列最新更新。

## Go 1.20发布清单

和Go 1.19相比，改动内容适中，主要涉及语言(Language)、可移植性(Ports)、工具链(Go Tools)、运行时(Runtime)、编译器(Compiler)、汇编器(Assembler)、链接器(Linker)和核心库(Core library)等方面的优化。

第1篇主要涉及Go 1.20在语言、可移植性方面的优化，原文链接：[Go 1.20版本升级内容第1篇](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484629&idx=1&sn=60a01d3cc85ef2462156f0565c30738d&chksm=ce124bbaf965c2ac351cd9c602e8b67d5119b2a89a7f2de0289bdeb7608ae589c329eb8f7275&token=1619842941&lang=zh_CN#rd)。

第2篇主要涉及Go命令和工具链方面的优化，原文链接：[Go 1.20版本升级内容第2篇](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484638&idx=1&sn=459a22d4a9bf5d9715e70d3c25b05b93&chksm=ce124bb1f965c2a76bacc1135799ab268be66a861e99391b354a9f2dfd8c22a60853cc1d689d&token=1342188569&lang=zh_CN#rd)。

本文重点介绍Go 1.20在运行时、编译器、汇编器和链接器等方面的优化。

## Runtime

Go 1.20版本的运行时新增了arena内存分配这个新功能的实验性支持，可以让Go程序释放更多内存空间，节省内存占用。

想了解什么是arena内存分配的，可以参考:https://github.com/golang/go/issues/51317。

如果area内存分配使用恰当，对于需要频繁内存分配的应用，可以提升多达15%的CPU性能。

使用方式为编译Go程序时，添加 `GOEXPERIMENT=arenas`参数。代码里如果有`import arena`，也需要添加这个编译参数。

此外，垃圾回收器的一些内部数据结构的设计做了优化，在时间和空间上更高效，可以节省内存开销，提升2%左右的CPU总体性能。

Go 1.20还新增了一个`runtime/coverage`包，调用这个包的API可以把程序运行的代码覆盖率数据输出到指定文件。

## Compiler

Go 1.20新增了PGO(profile-guided optimization)特性，可以帮助开发者做程序性能优化。

目前，编译器支持pprof CPU profile，这种类型的profile可以通过例如`runtime/pprof`或`net/http/pprof`收集得到。

如果要开启PGO，在使用`go build`编译程序的时候，要增加`-pgo`参数。`-pgo`指定的是profile文件的路径。如果`-pgo=auto`，那go命令会在main这个包的路径下去找名为`default.pgo`的文件。`-pgo=off`可以关闭优化。详情可以参考：[PGO Proposal](https://github.com/golang/go/issues/55022)。

如果使用了PGO，编译器会对被调用比较多的函数，更多地使用inline function的方式去做性能优化。

性能测试表明，如果开启了profile-guided inlining optimization，可以提升3%-4%的性能，后期Go会加入更多的PGO优化支持。

注意，由于PGO并不是稳定版本，生产环境使用需要小心。

此外，从Go 1.20开始，编译器禁止匿名interface嵌套，如下代码会编译失败。

```go
type I interface {
  m() interface {
    I 
  }
}
```

## Linker

在Linux环境，链接器会在进行link操作的时候，为`glibc`或者`musl`选择动态解释器。

在Windows环境，链接器现在支持基于LLVM的C语言工具链。

Go 1.20版本开始，使用`go:`和`type:`作为前缀，用于编译器生成的符号，而抛弃使用`go.`和`type.`作为前缀。

这可以避免歧义，因为有的go package的命名是以`go.`开始的。

## Bootstrap(自举)

如果要编译Go语言本身的源代码，必须要设置`GOROOT_BOOTSTRAP`环境变量。

在 1.4 版本开始，Go语言实现了自举，即可以通过1.4版本来编译安装之后版本的编译器。

假设`GOROOT_BOOTSTRAP`环境变量没有设置，那Go 1.17版本及之前的版本，会尝试在 `$HOME/go1.4` (`%HOMEDRIVE%%HOMEPATH%\go1.4` on Windows)寻找Go 1.4或者更新的bootstrap工具链。

Go 1.18和Go 1.19首先会找 `$HOME/go1.17` 或 `$HOME/sdk/go1.17`，找不到的话，就回退去找 `$HOME/go1.4`。

Go 1.20要实现自举，需要依赖Go 1.17的最新子版本，即Go 1.17.13版本。

Go 1.20会先找 `$HOME/go1.17.13` 或 `$HOME/sdk/go1.17.13`，如果找不到，就回退找`$HOME/go1.4` 。

在接下来，Go官方会把自举的工具链每年向前推进一步。预期是Go 1.22会要求依赖Go 1.20的最新子版本用于Go 1.22的自举。

## 总结

下一篇会介绍Go 1.20在核心库方面的优化工作，有一些内容值得学习，欢迎大家保持关注。



## 推荐阅读

* [Go 1.20要来了，看看都有哪些变化-第1篇](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484629&idx=1&sn=60a01d3cc85ef2462156f0565c30738d&chksm=ce124bbaf965c2ac351cd9c602e8b67d5119b2a89a7f2de0289bdeb7608ae589c329eb8f7275&token=1342188569&lang=zh_CN#rd)

* [Go 1.20要来了，看看都有哪些变化-第2篇](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484638&idx=1&sn=459a22d4a9bf5d9715e70d3c25b05b93&chksm=ce124bb1f965c2a76bacc1135799ab268be66a861e99391b354a9f2dfd8c22a60853cc1d689d&token=1342188569&lang=zh_CN#rd)

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