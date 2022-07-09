# Go 1.19要来了，看看都有哪些变化-第2篇

## 前言

Go官方团队在2022.06.11发布了Go 1.19 Beta 1版本，Go 1.19的正式release版本预计会在今年8月份发布。

让我们先睹为快，看看Go 1.19给我们带来了哪些变化。

这是Go 1.19版本更新内容详解的第2篇，欢迎大家关注公众号，及时获取本系列最新更新。

## Go 1.19发布清单

和Go 1.18相比，改动相对较小，主要涉及语言(Language)、内存模型(Memory Model)、可移植性(Ports)、Go Tool工具链、运行时(Runtime)、编译器(Compiler)、汇编器(Assembler)、链接器(Linker)和核心库(Core library)等方面的优化。

第1篇详细介绍了Go 1.19在语言、内存模型、可移植性方面的改进。

本文重点介绍Go 1.19版本在Go Tool工具链方面的变化。

### 文档注释

[文档注释(doc comments)](https://tip.golang.org/doc/comment) 是Go语言里的对包(package), 常量(const), 函数(func), 类型(type)和变量(var)声明的一种注释规范。按照这个规范来注释，就可以使用`go doc`命令生成对应的代码说明文档。

像大家熟知的https://pkg.go.dev/里的说明文档就是通过编写符合doc comments规范的文档注释来生成的。

Go 1.19 在文档注释里新增了对于链接、列表和更清晰的标题的支持，可以参考“[Go Doc Comments](https://tip.golang.org/doc/comment)” 了解语法细节。

作为这个修改的一部分，`gofmt`现在会把文档注释重新格式化，让文档展示的样式更清晰。

同时，新增了一个package:  [go/doc/comment](https://tip.golang.org/pkg/go/doc/comment)，可以用于解析和重新格式化文档注释，并且支持把文档注释渲染为HTML, Markdown和text格式。

### 新的编译约束 `unix` 

Go语言支持使用编译约束(build constraint)进行条件编译。Go 1.19版本新增了编译约束 `unix` ，可以在`//go:build`后面使用`unix`。

```go
//go:build unix
```

`unix`表示编译的目标操作系统是Unix或者类Unix系统。对于Go 1.19版本而言，如果`GOOS`是 `aix`, `android`, `darwin`, `dragonfly`, `freebsd`, `hurd`, `illumos`, `ios`, `linux`, `netbsd`, `openbsd`, 或 `solaris`中的某一个，那就满足`unix`这个编译约束。

未来`unix`约束还会匹配一些新的类Unix操作系统。 

### Go命令

`go build`如果使用`-trimpath`标记，会在生成的可执行文件里打上`trimpath`标签，我们可以使用 [`go` `version` `-m`](https://pkg.go.dev/cmd/go#hdr-Print_Go_version) 或[`debug.ReadBuildInfo`](https://pkg.go.dev/runtime/debug#ReadBuildInfo) 检查可执行文件是否是使用`-trimpath`标记编译生成的。

**备注**：编译的时候带上`trimpath`标记可以去除Go程序运行时打印的堆栈信息里包含的Go程序的编译路径和编译机用户信息，避免信息泄露。

`go` `generate` 现在会在生成器环境里设置 `GOROOT` 环境变量，所以即使使用了`-trimpath`进行编译，生成器也可以精准定位到`GOROOT`的路径。

`go` `test` 和 `go` `generate` 运行时会把 `GOROOT/bin` 放在 `PATH` 环境变量的开头，这样设计后，`go test`和`go generate`执行的时候可以解析到同一个`GOROOT`。解决的是这个[GitHub Issue](https://github.com/golang/go/issues/23635)。

`go` `env`会把环境变量的值中带有空格的加上双引号括起来，包括`CGO_CFLAGS`, `CGO_CPPFLAGS`, `CGO_CXXFLAGS`, `CGO_FFLAGS`, `CGO_LDFLAGS`, and `GOGCCFLAGS` 这些环境变量。解决的是Windows环境下不带空格会报错的bug，详情可以参考[GitHub Issue](https://github.com/golang/go/issues/45637)。

```bash
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -arch x86_64 -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/pv/_x849j6n22x37xxd9cstgwkr0000gn/T/go-build4165054210=/tmp/go-build -gno-record-gcc-switches -fno-common"
```

`go`命令现在会缓存必要的信息用于加载模块(module)，这会带来`go list`调用的加速。

对`-trimpath`和`go generate`不了解的，推荐阅读官方文档：

* [trimpath and go generate](https://pkg.go.dev/cmd/go#hdr-Print_Go_version)
* [go generate介绍](https://mp.weixin.qq.com/s/YBGppDhhBMorqkSzvufHlA)

### Vet

`go vet`新增了一个`errorsas`检查规则，可以对`errors.As`函数调用进行正确性检查。

如果`errors.As`的第2个参数是`*error`类型，`go vet`会提示错误，这也是大家使用`errors.As`常犯的一个错误。



## 推荐阅读

[Go 1.19版本变更内容第1篇](https://mp.weixin.qq.com/s/3xfCgtpEGu5Vm3XSSIze5w)，第1篇主要涉及到Go泛型的细小改动以及Go内存模型和原子操作的优化。

**想了解Go泛型的使用方法、设计思路和最佳实践，推荐大家阅读**：

* [官方教程：Go泛型入门](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483720&idx=1&sn=57ec4877dfd364a59deacf1e74a4fb66&chksm=ce124e27f965c731432dcc89d1e0563cf84baaef482eaa068a91bee61f10cf85b433923b83b4&token=1782465473&lang=zh_CN#rd)
* [一文读懂Go泛型设计和使用场景](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483731&idx=1&sn=b2258b28e2f3c16b065a5a1b22c15b0d&chksm=ce124e3cf965c72a6a22e0ed15deda8238567407bbd7157a79753fc8b605727ab2153009493c&token=1782465473&lang=zh_CN#rd)
* [重磅：Go 1.18将移除用于泛型的constraints包](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483855&idx=1&sn=6ab4aeb140a1a08268dc8a0284a6f375&chksm=ce124ea0f965c7b6776061960d71e4ffb30484a82041f5b1d4786c4b49c4ffabc07a28b1cd48&token=1782465473&lang=zh_CN#rd)
* [泛型最佳实践：Go泛型设计者教你如何用泛型](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484015&idx=1&sn=576b2d8b84b3a8ce5bdd6952c2b84062&chksm=ce124d00f965c416b07dcb81c4dcb9cf75859b2787d4f00ec8c80b37ca42e58cc651420a3b33&token=1782465473&lang=zh_CN#rd)

**想了解Go原子操作和使用方法，推荐大家阅读**：

* [Go并发编程之原子操作sync/atomic](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484082&idx=1&sn=934787c9829391ba743bd611818ad0e2&chksm=ce124dddf965c4cb7d0f2d9d001ab4b7d949fbe87c4c8b7ee8d7498946824ec9aa6581cfe986&token=1782465473&lang=zh_CN#rd)



## 总结

下一篇会介绍Go 1.19在运行时、编译器、汇编器、链接器和核心库的优化工作，有一些内容值得学习，欢迎大家保持关注。



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
* https://tip.golang.org/doc/comment