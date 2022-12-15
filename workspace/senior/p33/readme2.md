# Go 1.20要来了，看看都有哪些变化-第2篇

## 前言

Go官方团队在2022.12.08发布了Go 1.20 rc1(release candidate)版本，Go 1.20的正式release版本预计会在2023年2月份发布。

让我们先睹为快，看看Go 1.20给我们带来了哪些变化。(文末有彩蛋！)

安装方法：

```bash
$ go install golang.org/dl/go1.20rc1@latest
$ go1.20rc1 download
```

这是Go 1.20版本更新内容详解的第2篇，欢迎大家关注公众号，及时获取本系列最新更新。

第1篇主要涉及Go 1.20在语言、可移植性方面的优化，原文链接：[Go 1.20版本升级内容第1篇](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484629&idx=1&sn=60a01d3cc85ef2462156f0565c30738d&chksm=ce124bbaf965c2ac351cd9c602e8b67d5119b2a89a7f2de0289bdeb7608ae589c329eb8f7275&token=1619842941&lang=zh_CN#rd)。



## Go 1.20发布清单

和Go 1.19相比，改动内容适中，主要涉及语言(Language)、可移植性(Ports)、工具链(Go Tools)、运行时(Runtime)、编译器(Compiler)、汇编器(Assembler)、链接器(Linker)和核心库(Core library)等方面的优化。

本文重点介绍Go 1.20在Go工具链方面的优化。

### Go command

`$GOROOT/pkg`路径不再存储标准库预先编译好的

The directory `$GOROOT/pkg` no longer stores pre-compiled package archives for the standard library: `go` `install` no longer writes them, the `go` build no longer checks for them, and the Go distribution no longer ships them. 

Instead, packages in the standard library are built as needed and cached in the build cache, just like packages outside `GOROOT`. This change reduces the size of the Go distribution and also avoids C toolchain skew for packages that use cgo.

The implementation of `go` `test` `-json` has been improved to make it more robust. Programs that run `go` `test` `-json` do not need any updates. Programs that invoke `go` `tool` `test2json` directly should now run the test binary with `-v=test2json` (for example, `go` `test` `-v=test2json` or `./pkg.test` `-test.v=test2json`) instead of plain `-v`.

A related change to `go` `test` `-json` is the addition of an event with `Action` set to `start` at the beginning of each test program's execution. When running multiple tests using the `go` command, these start events are guaranteed to be emitted in the same order as the packages named on the command line.

The `go` command now defines architecture feature build tags, such as `amd64.v2`, to allow selecting a package implementation file based on the presence or absence of a particular architecture feature. See [`go` `help` `buildconstraint`](https://tip.golang.org/cmd/go#hdr-Build_constraints) for details.

The `go` subcommands now accept `-C` `<dir>` to change directory to <dir> before performing the command, which may be useful for scripts that need to execute commands in multiple different modules.

The `go` `build` and `go` `test` commands no longer accept the `-i` flag, which has been [deprecated since Go 1.16](https://go.dev/issue/41696).

The `go` `generate` command now accepts `-skip` `<pattern>` to skip `//go:generate` directives matching `<pattern>`.

The `go` `test` command now accepts `-skip` `<pattern>` to skip tests, subtests, or examples matching `<pattern>`.

When the main module is located within `GOPATH/src`, `go` `install` no longer installs libraries for non-`main` packages to `GOPATH/pkg`, and `go` `list` no longer reports a `Target` field for such packages. (In module mode, compiled packages are stored in the [build cache](https://pkg.go.dev/cmd/go#hdr-Build_and_test_caching) only, but [a bug](https://go.dev/issue/37015) had caused the `GOPATH` install targets to unexpectedly remain in effect.)

The `go` `build`, `go` `install`, and other build-related commands now support a `-pgo` flag that enables profile-guided optimization, which is described in more detail in the [Compiler](https://tip.golang.org/doc/go1.20#compiler) section below. The `-pgo` flag specifies the file path of the profile. Specifying `-pgo=auto` causes the `go` command to search for a file named `default.pgo` in the main package's directory and use it if present. This mode currently requires a single main package to be specified on the command line, but we plan to lift this restriction in a future release. Specifying `-pgo=off` turns off profile-guided optimization.

The `go` `build`, `go` `install`, and other build-related commands now support a `-cover` flag that builds the specified target with code coverage instrumentation. This is described in more detail in the [Cover](https://tip.golang.org/doc/go1.20#cover) section below.

#### go version

`go version -m`命令支持读取和解析更多类型的Go二进制文件。

比如通过`go build -buildmode=c-share`编译出来的Windows DLL文件以及没有可执行权限的Linux二进制文件，现在都可以被`go version -m`解析和识别到。

### Cgo

The `go` command now disables `cgo` by default on systems without a C toolchain. More specifically, when the `CGO_ENABLED` environment variable is unset, the `CC` environment variable is unset, and the default C compiler (typically `clang` or `gcc`) is not found in the path, `CGO_ENABLED` defaults to `0`. As always, you can override the default by setting `CGO_ENABLED` explicitly.

The most important effect of the default change is that when Go is installed on a system without a C compiler, it will now use pure Go builds for packages in the standard library that use cgo, instead of using pre-distributed package archives (which have been removed, as [noted above](https://tip.golang.org/doc/go1.20#go-command)) or attempting to use cgo and failing. This makes Go work better in some minimal container environments as well as on macOS, where pre-distributed package archives have not been used for cgo-based packages since Go 1.16.

The packages in the standard library that use cgo are [`net`](https://tip.golang.org/pkg/net/), [`os/user`](https://tip.golang.org/pkg/os/user/), and [`plugin`](https://tip.golang.org/pkg/plugin/). On macOS, the `net` and `os/user` packages have been rewritten not to use cgo: the same code is now used for cgo and non-cgo builds as well as cross-compiled builds. On Windows, the `net` and `os/user` packages have never used cgo. On other systems, builds with cgo disabled will use a pure Go version of these packages.

On macOS, the race detector has been rewritten not to use cgo: race-detector-enabled programs can be built and run without Xcode. On Linux and other Unix systems, and on Windows, a host C toolchain is required to use the race detector.

### Cover(代码覆盖率检测)

Go 1.20版本之前只支持对单元测试场景收集代码覆盖率，从Go 1.20版本开始支持对任何Go程序做代码覆盖率收集。

那如何收集呢？需要做如下操作：

* 给`go build`编译命令增加`-cover`标记
* 给环境变量`GOCOVERDIR`赋值为某个路径
* 运行`go build`编译出来的可执行程序时，会把代码覆盖率文件输出到`GOCOVERDIR`指定的路径下。

详细的介绍文档和使用说明可以参考： [coverage for integration tests' landing page](https://go.dev/testing/coverage) 。

想了解设计原理和实现的可以参考： [proposal](https://golang.org/issue/51430)。

### Vet

#### Improved detection of loop variable capture by nested functions

The `vet` tool now reports references to loop variables following a call to [`T.Parallel()`](https://tip.golang.org/pkg/testing/#T.Parallel) within subtest function bodies. Such references may observe the value of the variable from a different iteration (typically causing test cases to be skipped) or an invalid state due to unsynchronized concurrent access.

The tool also detects reference mistakes in more places. Previously it would only consider the last statement of the loop body, but now it recursively inspects the last statements within if, switch, and select statements.

#### New diagnostic for incorrect time formats

The vet tool now reports use of the time format 2006-02-01 (yyyy-dd-mm) with [`Time.Format`](https://tip.golang.org/pkg/time/#Time.Format) and [`time.Parse`](https://tip.golang.org/pkg/time/#Parse). This format does not appear in common date standards, but is frequently used by mistake when attempting to use the ISO 8601 date format (yyyy-mm-dd).

## 总结

下一篇会介绍Go 1.20在运行时、编译器、汇编器、链接器和核心库的优化工作，有一些内容值得学习，欢迎大家保持关注。



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