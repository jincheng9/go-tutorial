# Go 1.19要来了，看看都有哪些变化

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

运行时现在支持软内存限制(soft memory limit)。这个内存限制包括了堆里的内存以及所有其它被运行时管理的内存

The runtime now includes support for a soft memory limit. This memory limit includes the Go heap and all other memory managed by the runtime, and excludes external memory sources such as mappings of the binary itself, memory managed in other languages, and memory held by the operating system on behalf of the Go program. This limit may be managed via [`runtime/debug.SetMemoryLimit`](https://tip.golang.org/pkg/runtime/debug/#SetMemoryLimit) or the equivalent [`GOMEMLIMIT`](https://tip.golang.org/pkg/runtime/#hdr-Environment_Variables) environment variable. The limit works in conjunction with [`runtime/debug.SetGCPercent`](https://tip.golang.org/pkg/runtime/debug/#SetGCPercent) / [`GOGC`](https://tip.golang.org/pkg/runtime/#hdr-Environment_Variables), and will be respected even if `GOGC=off`, allowing Go programs to always make maximal use of their memory limit, improving resource efficiency in some cases. See [the GC guide](https://tip.golang.org/doc/gc-guide) for a detailed guide explaining the soft memory limit in more detail, as well as a variety of common use-cases and scenarios. Please note that small memory limits, on the order of tens of megabytes or less, are less likely to be respected due to external latency factors, such as OS scheduling. See [issue 52433](https://go.dev/issue/52433) for more details. Larger memory limits, on the order of hundreds of megabytes or more, are stable and production-ready.

In order to limit the effects of GC thrashing when the program's live heap size approaches the soft memory limit, the Go runtime also attempts to limit total GC CPU utilization to 50%, excluding idle time, choosing to use more memory over preventing application progress. In practice, we expect this limit to only play a role in exceptional cases, and the new [runtime metric](https://tip.golang.org/pkg/runtime/metrics/#hdr-Supported_metrics) `/gc/limiter/last-enabled:gc-cycle` reports when this last occurred.

The runtime now schedules many fewer GC worker goroutines on idle operating system threads when the application is idle enough to force a periodic GC cycle.

The runtime will now allocate initial goroutine stacks based on the historic average stack usage of goroutines. This avoids some of the early stack growth and copying needed in the average case in exchange for at most 2x wasted space on below-average goroutines.

On Unix operating systems, Go programs that import package [os](https://tip.golang.org/pkg/os/) now automatically increase the open file limit (`RLIMIT_NOFILE`) to the maximum allowed value; that is, they change the soft limit to match the hard limit. This corrects artificially low limits set on some systems for compatibility with very old C programs using the [*select*](https://en.wikipedia.org/wiki/Select_(Unix)) system call. Go programs are not helped by that limit, and instead even simple programs like `gofmt` often ran out of file descriptors on such systems when processing many files in parallel. One impact of this change is that Go programs that in turn execute very old C programs in child processes may run those programs with too high a limit. This can be corrected by setting the hard limit before invoking the Go program.

Unrecoverable fatal errors (such as concurrent map writes, or unlock of unlocked mutexes) now print a simpler traceback excluding runtime metadata (equivalent to a fatal panic) unless `GOTRACEBACK=system` or `crash`. Runtime-internal fatal error tracebacks always include full metadata regardless of the value of `GOTRACEBACK`

Support for debugger-injected function calls has been added on ARM64, enabling users to call functions from their binary in an interactive debugging session when using a debugger that is updated to make use of this functionality.

The [address sanitizer support added in Go 1.18](https://tip.golang.org/doc/go1.18#go-build-asan) now handles function arguments and global variables more precisely.

## 编译器

The compiler now uses a [jump table](https://en.wikipedia.org/wiki/Branch_table) to implement large integer and string switch statements. Performance improvements for the switch statement vary but can be on the order of 20% faster. (`GOARCH=amd64` and `GOARCH=arm64` only)

The Go compiler now requires the `-p=importpath` flag to build a linkable object file. This is already supplied by the `go` command and by Bazel. Any other build systems that invoke the Go compiler directly will need to make sure they pass this flag as well.

## 汇编器

Like the compiler, the assembler now requires the `-p=importpath` flag to build a linkable object file. This is already supplied by the `go` command. Any other build systems that invoke the Go assembler directly will need to make sure they pass this flag as well.

## 链接器

On ELF platforms, the linker now emits compressed DWARF sections in the standard gABI format (`SHF_COMPRESSED`), instead of the legacy `.zdebug` format.

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