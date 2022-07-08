# Go 1.19要来了，看看都有哪些变化-第3篇

## 前言

Go官方团队在2022.06.11发布了Go 1.19 Beta 1版本，Go 1.19的正式release版本预计会在今年8月份发布。

让我们先睹为快，看看Go 1.19给我们带来了哪些变化。

这是Go 1.19版本更新内容详解的第3篇，欢迎大家关注公众号，及时获取本系列最新更新。

第1篇主要涉及Go泛型的改动、Go内存模型和原子操作的优化，原文链接：[Go 1.19版本变更内容第1篇](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484179&idx=1&sn=215ea3f092460118b2bc975935015874&chksm=ce124c7cf965c56a7c310b1059683d065810bd18368669d3d42a6cbbb0370d1593979a63620c#rd)。

第2篇主要涉及Go文档注释(doc comments)、编译约束(build constraint)以及Go命令的修改，原文链接：[Go 1.19版本变更内容第2篇](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484188&idx=1&sn=c14bafb1f89b3b3f988452c5a5f32884&chksm=ce124c73f965c5651a688c42561b02e38253b60943c77a0a6ad7b45621b4296d9e1acd47de7a#rd)。

## Go 1.19发布清单

和Go 1.18相比，改动相对较小，主要涉及语言(Language)、内存模型(Memory Model)、可移植性(Ports)、Go Tool工具链、运行时(Runtime)、编译器(Compiler)、汇编器(Assembler)、链接器(Linker)和核心库(Core library)等方面的优化。

本文重点介绍Go 1.19版本在运行时、编译器、汇编器和链接器方面的变化。

## 运行时

###  软内存限制(soft memory limit)

运行时现在支持软内存限制(soft memory limit)。这个内存限制包括了Go heap里的内存以及所有其它被Go运行时管理的内存。如果不是被Go运行时管理的内存，比如二进制程序本身映射的内存、其它语言管理的内存，是不在这个软内存限制里的。

这个限制可以通过[`runtime/debug.SetMemoryLimit`](https://tip.golang.org/pkg/runtime/debug/#SetMemoryLimit) 函数或者 [`GOMEMLIMIT`](https://tip.golang.org/pkg/runtime/#hdr-Environment_Variables) 环境变量进行设置。

软内存限制和[`runtime/debug.SetGCPercent`](https://tip.golang.org/pkg/runtime/debug/#SetGCPercent) 函数以及 [`GOGC`](https://tip.golang.org/pkg/runtime/#hdr-Environment_Variables)环境变量是可以结合起来工作的，而且即使在`GOGC=off`模式下，软内存限制也会生效。设计目的是为了让Go程序可以最大化内存使用，提升某些场景下的内存资源使用效率。

可以参考[the GC guide](https://tip.golang.org/doc/gc-guide)查看更多软内存限制的设计实现细节，以及一些常见使用场景和最佳实践。

需要注意的是，对于数十MB或者更小的内存限制，由于考虑到一些性能问题，软内存限制是有可能不会生效的，可以参考[issue 52433](https://go.dev/issue/52433)查看更多细节。对于数百MB或者更大的内存限制，目前是可以稳定运行在生产环境上的。

当Go程序堆内存接近软内存限制时，为了减少GC抖动的影响，Go运行时会尝试限制GC CPU利用率不超过50%(不包括CPU空闲时间)。

在实际使用中，一般只在一些特殊场景才建议使用软内存限制，当Go堆内存占用真的超过软内存限制时，新的运行时度量( [runtime metric](https://tip.golang.org/pkg/runtime/metrics/#hdr-Supported_metrics))工具`/gc/limiter/last-enabled:gc-cycle`会报告这个事件。

如果应用程序空闲到足以执行一次GC时，Go运行时(Runtime)现在会在空闲的操作系统线程里调度更少的GC工作协程(goroutine)。

Go运行时现在会根据goroutine过去栈空间的平均使用大小来分配一个初始的goroutine栈大小，提升栈空间的分配和使用效率。

### 最大文件描述符

在Unix操作系统上，import了 [os](https://tip.golang.org/pkg/os/) 包的Go程序现在会自动增加 **可打开的最大文件描述符数量** 到允许的最大值。

修改这个的原因是有些系统为了兼容很老的使用select系统调用的C语言程序，会把最大文件描述符数量设置一个比较小的值，通常是1024。这个设置对Go程序没有用处，相反，还会带来一些问题。比如`gofmt`在并行处理很多文件时，很容易遇到文件描述符用尽的问题。另外这个修改需要注意的一个点是，如果Go程序和C语言混合编程，执行了很老的C代码，那在运行这个Go程序时，需要先设置最大文件描述符的hard limit，将其调小。

### 堆栈信息

不可恢复的致命错误(例如并发的map写入，解锁一个未加锁的mutex)现在会打印更简单的堆栈信息，不再打印Go runtime的元数据，除非设置了`GOTRACEBACK=system` or `crash`。但是对于runtime本身的致命错误，不管`GOTRACEBACK` 的值是什么，还是会打印包含元数据的所有堆栈信息。

### 调试

ARM64架构现在支持注入了调试器的函数调用，允许用户在交互式的调试会话里调用函数。

 [Go 1.18新增的address sanitizer](https://tip.golang.org/doc/go1.18#go-build-asan) 现在可以更精确地处理函数参数和全局变量。

## 编译器

针对`GOARCH=amd64` 和 `GOARCH=arm64` 架构，编译器现在使用跳表(jump table)来实现大整数和字符串的switch语句。

带来的优化效果是switch语句的性能提升了大概20%左右。

Go编译器现在需要 `-p=importpath` 标记来编译出一个可链接的目标文件。`go`命令和Bazel现在已经支持`-p=importpath`标记。

任何其它直接调用Go编译器的编译系统也需要确保传递了这个标记参数。

## 汇编器

和编译器一样，汇编器现在也需要`-p=importpath`标记来编译出一个可链接的目标文件。`go`命令已经支持该标记参数。

任何其它直接调用Go汇编器的编译系统也需要确保传递了这个标记参数。

## 链接器

在ELF(Executable and Linkable Format)平台上，链接器现在会以gABI格式(`SHF_COMPRESSED`压缩方式)压缩DWARF章节，而不是传统的 `.zdebug` 格式。

## 推荐阅读

第1篇主要涉及Go泛型的改动、Go内存模型和原子操作的优化，原文链接：[Go 1.19版本变更内容第1篇](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484179&idx=1&sn=215ea3f092460118b2bc975935015874&chksm=ce124c7cf965c56a7c310b1059683d065810bd18368669d3d42a6cbbb0370d1593979a63620c#rd)。

第2篇主要涉及Go文档注释(doc comments)、编译约束(build constraint)以及Go命令的修改，原文链接：[Go 1.19版本变更内容第2篇](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484188&idx=1&sn=c14bafb1f89b3b3f988452c5a5f32884&chksm=ce124c73f965c5651a688c42561b02e38253b60943c77a0a6ad7b45621b4296d9e1acd47de7a#rd)。

**想了解Go泛型的使用方法、设计思路和最佳实践，推荐大家阅读**：

* [官方教程：Go泛型入门](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483720&idx=1&sn=57ec4877dfd364a59deacf1e74a4fb66&chksm=ce124e27f965c731432dcc89d1e0563cf84baaef482eaa068a91bee61f10cf85b433923b83b4&token=1782465473&lang=zh_CN#rd)
* [一文读懂Go泛型设计和使用场景](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483731&idx=1&sn=b2258b28e2f3c16b065a5a1b22c15b0d&chksm=ce124e3cf965c72a6a22e0ed15deda8238567407bbd7157a79753fc8b605727ab2153009493c&token=1782465473&lang=zh_CN#rd)
* [重磅：Go 1.18将移除用于泛型的constraints包](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483855&idx=1&sn=6ab4aeb140a1a08268dc8a0284a6f375&chksm=ce124ea0f965c7b6776061960d71e4ffb30484a82041f5b1d4786c4b49c4ffabc07a28b1cd48&token=1782465473&lang=zh_CN#rd)
* [泛型最佳实践：Go泛型设计者教你如何用泛型](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484015&idx=1&sn=576b2d8b84b3a8ce5bdd6952c2b84062&chksm=ce124d00f965c416b07dcb81c4dcb9cf75859b2787d4f00ec8c80b37ca42e58cc651420a3b33&token=1782465473&lang=zh_CN#rd)

**想了解Go原子操作和使用方法，推荐大家阅读**：

* [Go并发编程之原子操作sync/atomic](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484082&idx=1&sn=934787c9829391ba743bd611818ad0e2&chksm=ce124dddf965c4cb7d0f2d9d001ab4b7d949fbe87c4c8b7ee8d7498946824ec9aa6581cfe986&token=1782465473&lang=zh_CN#rd)



## 总结

下一篇会介绍Go 1.19对核心库的优化工作，有一些内容值得学习，欢迎大家保持关注。



## 开源地址

文章和示例代码开源在GitHub: [Go语言初级、中级和高级教程](https://github.com/jincheng9/go-tutorial)。

公众号：coding进阶。关注公众号可以获取最新Go面试题和技术栈。

个人网站：[Jincheng's Blog](https://jincheng9.github.io/)。

知乎：[无忌](https://www.zhihu.com/people/thucuhkwuji)。



## 福利

我为大家整理了一份后端开发学习资料礼包，包含编程语言入门到进阶知识(Go、C++、Python)、后端开发技术栈、面试题等。

关注公众号「coding进阶」，发送消息 **backend** 领取资料礼包，这份资料会不定期更新，加入我觉得有价值的资料。还可以发送消息「**进群**」，和同行一起交流学习，答疑解惑。



## References

* https://tip.golang.org/doc/go1.19