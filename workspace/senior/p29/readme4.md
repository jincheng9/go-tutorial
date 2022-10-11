# Go 1.19要来了，看看都有哪些变化-第4篇

## 前言

Go官方团队在2022.06.11发布了Go 1.19 Beta 1版本，Go 1.19的正式release版本预计会在今年8月份发布。

让我们先睹为快，看看Go 1.19给我们带来了哪些变化。

这是Go 1.19版本更新内容详解的第4篇，欢迎大家关注公众号，及时获取本系列最新更新。

第1篇主要涉及Go泛型的改动、Go内存模型和原子操作的优化，原文链接：[Go 1.19版本变更内容第1篇](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484179&idx=1&sn=215ea3f092460118b2bc975935015874&chksm=ce124c7cf965c56a7c310b1059683d065810bd18368669d3d42a6cbbb0370d1593979a63620c#rd)。

第2篇主要涉及Go文档注释(doc comments)、编译约束(build constraint)以及Go命令的修改，原文链接：[Go 1.19版本变更内容第2篇](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484188&idx=1&sn=c14bafb1f89b3b3f988452c5a5f32884&chksm=ce124c73f965c5651a688c42561b02e38253b60943c77a0a6ad7b45621b4296d9e1acd47de7a#rd)。

第3篇主要涉及Go运行时、编译器、汇编器和链接器方面的改动和优化，原文链接：[Go 1.19版本变更内容第3篇](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484203&idx=1&sn=fcf95ce045a54c2f6a9414d1e9fa732d&chksm=ce124c44f965c5525fc820f99978cf0996d5f3653e4e580302986f14c9ae740d499ba9fb18c4&token=1117798005&lang=zh_CN#rd)。

## Go 1.19发布清单

和Go 1.18相比，改动相对较小，主要涉及语言(Language)、内存模型(Memory Model)、可移植性(Ports)、Go Tool工具链、运行时(Runtime)、编译器(Compiler)、汇编器(Assembler)、链接器(Linker)和核心库(Core library)等方面的优化。

本文重点介绍Go 1.19版本在核心库(Core library)方面的变化。

### 新的原子类型(New atomic types)

[`sync/atomic`](https://tip.golang.org/pkg/sync/atomic/)包里现在定义了新的类型： [`Bool`](https://tip.golang.org/pkg/sync/atomic/#Bool), [`Int32`](https://tip.golang.org/pkg/sync/atomic/#Int32), [`Int64`](https://tip.golang.org/pkg/sync/atomic/#Int64), [`Uint32`](https://tip.golang.org/pkg/sync/atomic/#Uint32), [`Uint64`](https://tip.golang.org/pkg/sync/atomic/#Uint64), [`Uintptr`](https://tip.golang.org/pkg/sync/atomic/#Uintptr), and [`Pointer`](https://tip.golang.org/pkg/sync/atomic/#Pointer)。

这些新的类型定义了相应的原子方法，要修改或者读取这些类型的变量的值就必须使用该类型的原子方法，这样可以避免误操作。

```go
type Bool struct {
	// contains filtered or unexported fields
}

func (x *Bool) CompareAndSwap(old, new bool) (swapped bool)
func (x *Bool) Load() bool
func (x *Bool) Store(val bool)
func (x *Bool) Swap(new bool) (old bool)
```

比如上面的`sync/atomic`包里的`Bool`类型就有4个原子方法，要读取或者修改`atomic.Bool`类型的变量的值就要使用这4个方法。

`sync/atomic`包有了`Pointer`类型后，开发者不需要先把变量转成`unsafe.Pointer`类型再去调用`sync/atomic`包里的函数，直接使用`Pointer`类型的原子方法即可。

`Int64` 和`Uint64`类型在结构体(structs)和分配的内存里会自动按照64位自动对齐，即使在32位系统上也是按照64位对齐。

### 路径查找(PATH lookups)

[`Command`](https://tip.golang.org/pkg/os/exec/#Command) 和 [`LookPath`](https://tip.golang.org/pkg/os/exec/#LookPath) 不再允许在当前目录查找可执行程序，这个修改解决了一个[常见的安全问题](https://tip.golang.org/blog/path-security)，但是也带来了破坏性更新。

比如以前有段代码是`exec.Command("prog")`，表示要执行当前目录下名为`prog`的可执行文件(在Windows系统上对应的是`prog.exe`)，那使用Go 1.19后就不会生效了。可以参考 [`os/exec`](https://tip.golang.org/pkg/os/exec/) 包的说明来修改代码以适配`Command`和`LookPath`的改动。

在Windows系统上，`Command`和`LookPath`现在会感知 [`NoDefaultCurrentDirectoryInExePath`](https://docs.microsoft.com/en-us/windows/win32/api/processenv/nf-processenv-needcurrentdirectoryforexepatha) 环境变量。我们可以在Windows系统上设置该环境变量来禁止从当前目录`.` 查找可执行程序。

### 核心库的微小改动

Go标准库在Go 1.19版本有很多细微的改动和优化，主要涵盖以下内容：

- [archive/zip](https://tip.golang.org/pkg/archive/zip/)

  [`Reader`](https://tip.golang.org/pkg/archive/zip/#Reader) 现在会忽略掉ZIP文件开头的非ZIP数据部分，这在读一些Java的JAR文件时会很有必要。

- [crypto/rand](https://tip.golang.org/pkg/crypto/rand/)

  [`Read`](https://tip.golang.org/pkg/crypto/rand/#Read) 不再缓存从操作系统里获取的随机数。对于Plan 9操作系统，`Read`被重新实现了，用fast key erasure替换掉了ANSI X9.31算法。

- [crypto/tls](https://tip.golang.org/pkg/crypto/tls/)

   `tls10default` `GODEBUG` 选项在Go 1.19版本已经被移除。 不过，我们还是可以通过设置 [`Config.MinVersion`](https://tip.golang.org/pkg/crypto/tls#Config.MinVersion) 来支持client侧使用TLS 1.0协议。根据RFC 5246中7.4.1.4章节和RFC 8446中4.2章节的要求，TLS server和client现在会拒绝TLS握手里重复的扩展(duplicate extensions)。

- [crypto/x509](https://tip.golang.org/pkg/crypto/x509/)

  [`CreateCertificate`](https://tip.golang.org/pkg/crypto/x509/#CreateCertificate) 不再支持使用`MD5WITHRSA`的签名算法来创建证书。

  `CreateCertificate` 不再接受SerialNumber为负数。

  [`ParseCertificate`](https://tip.golang.org/pkg/crypto/x509/#ParseCertificate) 和 [`ParseCertificateRequest`](https://tip.golang.org/pkg/crypto/x509/#ParseCertificateRequest) 现在会拒绝包含有重复扩展的证书和CSR(Certifcate Signing Request)。

  新方法 [`CertPool.Clone`](https://tip.golang.org/pkg/crypto/x509/#CertPool.Clone) 和 [`CertPool.Equal`](https://tip.golang.org/pkg/crypto/x509/#CertPool.Equal) 可以克隆一个`CertPool`，并且检查2个`CertPool`是否相同。

  新函数 [`ParseRevocationList`](https://tip.golang.org/pkg/crypto/x509/#ParseRevocationList) 提供了一个更快、更安全的方式去使用CRL解析器(parser)。

- [crypto/x509/pkix](https://tip.golang.org/pkg/crypto/x509/pkix)

   [`CertificateList`](https://tip.golang.org/pkg/crypto/x509/pkix#CertificateList) 和 [`TBSCertificateList`](https://tip.golang.org/pkg/crypto/x509/pkix#TBSCertificateList) 现在被废弃了，应该使用新的 [`crypto/x509` CRL functionality](https://tip.golang.org/doc/go1.19#crypto/x509)。 

- [debug](https://tip.golang.org/pkg/debug/)

  新的 `EM_LONGARCH` and `R_LARCH_*` 常量现在支持龙芯loong64架构。

- [debug/pe](https://tip.golang.org/pkg/debug/pe/)

  引入了新方法 [`File.COFFSymbolReadSectionDefAux`](https://tip.golang.org/pkg/debug/pe/#File.COFFSymbolReadSectionDefAux) ，该方法返回 [`COFFSymbolAuxFormat5`](https://tip.golang.org/pkg/debug/pe/#COFFSymbolAuxFormat5)类型，可以让开发者访问PE文件里的COMDAT信息。

- [encoding/binary](https://tip.golang.org/pkg/encoding/binary/)

  新接口 [`AppendByteOrder`](https://tip.golang.org/pkg/encoding/binary/#AppendByteOrder) 提供了高效的方法用于把 `uint16`，`uint32`，或 `uint64` 添加到一个byte切片里。

   [`BigEndian`](https://tip.golang.org/pkg/encoding/binary/#BigEndian) 和 [`LittleEndian`](https://tip.golang.org/pkg/encoding/binary/#LittleEndian) 都实现了该接口。

- [encoding/csv](https://tip.golang.org/pkg/encoding/csv/)

  新方法 [`Reader.InputOffset`](https://tip.golang.org/pkg/encoding/csv/#Reader.InputOffset) 会返回当前读到的位置，以偏移的字节数来表示，类似于 `encoding/json`包里的 [`Decoder.InputOffset`](https://tip.golang.org/pkg/encoding/json/#Decoder.InputOffset)。

- [encoding/xml](https://tip.golang.org/pkg/encoding/xml/)

  新方法 [`Decoder.InputPos`](https://tip.golang.org/pkg/encoding/xml/#Decoder.InputPos) 会返回当前读到的位置，以行和列来表示，类似于 `encoding/csv`包里的 [`Decoder.FieldPos`](https://tip.golang.org/pkg/encoding/csv/#Decoder.FieldPos)方法。

- [flag](https://tip.golang.org/pkg/flag/)

  新函数 [`TextVar`](https://tip.golang.org/pkg/flag/#TextVar) 定义了一个 [`encoding.TextUnmarshaler`](https://tip.golang.org/pkg/encoding/#TextUnmarshaler)参数，允许命令行里传入的flag变量使用 [`big.Int`](https://tip.golang.org/pkg/math/big/#Int), [`netip.Addr`](https://tip.golang.org/pkg/net/netip/#Addr)和[`time.Time`](https://tip.golang.org/pkg/time/#Time)类型。

- [fmt](https://tip.golang.org/pkg/fmt/)

  新函数 [`Append`](https://tip.golang.org/pkg/fmt/#Append), [`Appendf`](https://tip.golang.org/pkg/fmt/#Appendf) 和 [`Appendln`](https://tip.golang.org/pkg/fmt/#Appendln) 可以添加格式化的数据到byte切片中。

- [go/parser](https://tip.golang.org/pkg/go/parser/)

  `go/parser`会把 `~x`解析为一元表达式(unary expression)，其中操作符是`~`，`~`操作符的官方说明参考 [token.TILDE](https://tip.golang.org/pkg/go/token#TILDE)。

  当类型约束(type constraint)用在错误的上下文时，比如`~int`，可以允许更好的错误恢复。

- [go/types](https://tip.golang.org/pkg/go/types/)

  新方法 [`Func.Origin`](https://tip.golang.org/pkg/go/types/#Func.Origin) 和 [`Var.Origin`](https://tip.golang.org/pkg/go/types/#Var.Origin) 会返回 [`Func`](https://tip.golang.org/pkg/go/types/#Func) 和 [`Var`](https://tip.golang.org/pkg/go/types/#Var) 实例化后的对象。

- [hash/maphash](https://tip.golang.org/pkg/hash/maphash/)

  新函数 [`Bytes`](https://tip.golang.org/pkg/hash/maphash/#Bytes) 和 [`String`](https://tip.golang.org/pkg/hash/maphash/#String) 提供了高效的方式用于对一个byte slice或者字符串做hash。

- [html/template](https://tip.golang.org/pkg/html/template/)

   [`FuncMap`](https://tip.golang.org/pkg/html/template/#FuncMap) 类型现在是`text/template`包里 [`FuncMap`](https://tip.golang.org/pkg/text/template/#FuncMap) 类型的别名，本身不再是一个独立的类型。

- [image/draw](https://tip.golang.org/pkg/image/draw/)

  当目标图像和源头像都是[`image.NRGBA`](https://tip.golang.org/pkg/image/#NRGBA) 或者都是 [`image.NRGBA64`](https://tip.golang.org/pkg/image/#NRGBA64) 类型时，operator为 [`Src`](https://tip.golang.org/pkg/image/draw/#Src) 的[`Draw`](https://tip.golang.org/pkg/image/draw/#Draw) 会保留non-premultiplied-alpha颜色。

  其实Go 1.17及更早版本的行为就是如此，但是Go 1.18版本做库优化的时候改变了这个行为，Go 1.19版本将这个行为还原了。

- [io](https://tip.golang.org/pkg/io/)

  [`NopCloser`](https://tip.golang.org/pkg/io/#NopCloser)的结果现在实现了 [`WriterTo`](https://tip.golang.org/pkg/io/#WriterTo) 接口。

  [`MultiReader`](https://tip.golang.org/pkg/io/#MultiReader)的结果现在无条件地实现了 [`WriterTo`](https://tip.golang.org/pkg/io/#WriterTo)。如果任何底层的reader没有实现`WriteTo`，也会模拟`WriteTo`的行为。

- [mime](https://tip.golang.org/pkg/mime/)

  `.js`扩展名的文件本来应该被mime包识别为 `text/plain`类型，但是在Windows系统上有bug，会导致以`.js`为扩展名的文件被mime包识别为`text/javascript; charset=utf-8`类型。

  如果在Windows系统上，想让以`.js`为扩展名的文件被mime包识别为 `text/plain` ，必须显式调用 [`AddExtensionType`](https://tip.golang.org/pkg/mime/#AddExtensionType)。

- [net/http](https://tip.golang.org/pkg/net/http/)

  [`ResponseWriter.WriteHeader`](https://tip.golang.org/pkg/net/http/#ResponseWriter) 现在支持发送用户自定义的1xx信息头(informational header)。

   [`MaxBytesReader`](https://tip.golang.org/pkg/net/http/#MaxBytesReader) 的返回值 `io.ReadCloser`在超过读上限(read limit)后，会返回一个错误类型 [`MaxBytesError`](https://tip.golang.org/pkg/net/http/#MaxBytesError) 。

  HTTP client会把状态码为3xx但是没有`Location` header的Http Response返回给调用者，而不是直接当做错误处理。

- [net/url](https://tip.golang.org/pkg/net/url/)

  新增的 [`JoinPath`](https://tip.golang.org/pkg/net/url/#JoinPath) 函数 和 [`URL.JoinPath`](https://tip.golang.org/pkg/net/url/#URL.JoinPath) 方法可以把一组path元素组合在一起，创建一个新的 `URL`。

   `URL`类型现在会区分没有host的URL和host为空的URL。举个例子， `http:///path` 是有host的，只是host为空，但是 `http:/path` 就没有host。

  当URL的host为空时，[`URL`](https://tip.golang.org/pkg/net/url/#URL) 类型里的字段 `OmitHost` 的值会被设置为`true`。

- [os/exec](https://tip.golang.org/pkg/os/exec/)

  如果 [`Cmd`](https://tip.golang.org/pkg/os/exec/#Cmd) 类型的 `Dir` 字段非空， `Env`字段为nil，会隐式地为子进程设置`PWD`环境变量，值为`Dir`字段的值。

  新方法 [`Cmd.Environ`](https://tip.golang.org/pkg/os/exec/#Cmd.Environ) 可以获取到运行cmd的环境，包括隐式设置的`PWD`环境变量。

- [reflect](https://tip.golang.org/pkg/reflect/)

   [`Value.Bytes`](https://tip.golang.org/pkg/reflect/#Value.Bytes) 方法现在除了接收slice切片，现在还接收可取址的数组(addressable array)。 [`Value.Len`](https://tip.golang.org/pkg/reflect/#Value.Len) 和 [`Value.Cap`](https://tip.golang.org/pkg/reflect/#Value.Cap) 方法现在可以操作指向数组的指针，返回数组的长度。

- [regexp/syntax](https://tip.golang.org/pkg/regexp/syntax/)

  Go 1.18 release candidate 1， Go 1.17.8和 Go 1.16.15 这3个版本包含了对正则表达式解析可能带来的安全问题的修复，会拒绝嵌套很深的正则表达式。由于Go的补丁版本不能引入新的API，对于这种情况，解析器会返回 [`syntax.ErrInternalError`](https://tip.golang.org/pkg/regexp/syntax/#ErrInternalError) 。

  Go 1.19对于上述情况，新增了一个更具体的错误 [`syntax.ErrNestingDepth`](https://tip.golang.org/pkg/regexp/syntax/#ErrNestingDepth)，不再返回 [`syntax.ErrInternalError`](https://tip.golang.org/pkg/regexp/syntax/#ErrInternalError) 。

- [runtime](https://tip.golang.org/pkg/runtime/)

  当Go可执行程序使用了`-trimpath`标记进行编译并且在进程运行环境里没有设置`GOROOT`环境变量， [`GOROOT`](https://tip.golang.org/pkg/runtime/#GOROOT) 函数会返回空串。

- [runtime/metrics](https://tip.golang.org/pkg/runtime/metrics/)

  新的 `/sched/gomaxprocs:threads` [度量指标](https://tip.golang.org/pkg/runtime/metrics/#hdr-Supported_metrics) 会报告 [`runtime.GOMAXPROCS`](https://tip.golang.org/pkg/runtime/#GOMAXPROCS) 的当前值。

  新的 `/cgo/go-to-c-calls:calls` [度量指标](https://tip.golang.org/pkg/runtime/metrics/#hdr-Supported_metrics) 会报告Go调用C的总次数。这个指标等同于 [`runtime.NumCgoCall`](https://tip.golang.org/pkg/runtime/#NumCgoCall) 函数的执行结果。

  新的 `/gc/limiter/last-enabled:gc-cycle` [度量指标](https://tip.golang.org/pkg/runtime/metrics/#hdr-Supported_metrics) 在GC CPU limiter开启时，会报告最新的GC循环(cycle)。可以参考runtime notes](https://tip.golang.org/doc/go1.19#runtime) 了解更多关于GC CPU limiter的细节。

- [runtime/pprof](https://tip.golang.org/pkg/runtime/pprof/)

  `pprof`在收集goroutine profile时做了优化，可以大大减少对应用程序的性能影响。

  所有Unix操作系统上做`pprof` 的heap profile结果都包含了`MaxRSS`，之前只有 `GOOS=android`, `darwin`, `ios` 和 `linux`系统上才会包含有`MaxRSS`结果。

- [runtime/race](https://tip.golang.org/pkg/runtime/race/)

  race detector在Go 1.19版本做了升级，使用v3版本的 thread sanitizer，支持除了 `windows/amd64` 和 `openbsd/amd64` 的所有平台， `windows/amd64`和 `openbsd/amd64`平台仍然使用v2版本的thread sanitizer。

  和v2版本相比，v3版本速度提升了1.5-2倍，并且内存开销减半，还不限制goroutine的数量。

  在Linux操作系统上，race detector现在要求glibc的版本最低是2.17。

  race detector现在支持`GOARCH=s390x`架构。

  新版的thread sanitizer不再支持`openbsd/amd64`平台，因此`openbsd/amd64`平台还是会沿用旧的v2版本的thread sanitizer。

- [runtime/trace](https://tip.golang.org/pkg/runtime/trace/)

  当tracing和 [CPU profiler](https://tip.golang.org/pkg/runtime/pprof#StartCPUProfile) 同时开启时，tracing也会记录CPU Profile采样的结果。

- [sort](https://tip.golang.org/pkg/sort/)

  Go自带的排序算法使用了[pattern-defeating quicksort](https://arxiv.org/pdf/2106.05123.pdf)进行重写，速度更快。

  新的函数 [Find](https://tip.golang.org/pkg/sort/#Find) 类似 函数[Search](https://tip.golang.org/pkg/sort/#Search) ，但是更好用。`Find`函数会额外返回一个bool值，用于表示是否找到了相同的数。

- [strconv](https://tip.golang.org/pkg/strconv/)

  [`Quote`](https://tip.golang.org/pkg/strconv/#Quote) 函数和相关函数为了和其它ASCII码值保持一致，会引用字符U+007F为`\x7f`，而不是`\u007f`。

- [syscall](https://tip.golang.org/pkg/syscall/)

  对于PowerPC (`GOARCH=ppc64`, `ppc64le`)架构，[`Syscall`](https://tip.golang.org/pkg/syscall/#Syscall)， [`Syscall6`](https://tip.golang.org/pkg/syscall/#Syscall6)，[`RawSyscall`](https://tip.golang.org/pkg/syscall/#RawSyscall)，和 [`RawSyscall6`](https://tip.golang.org/pkg/syscall/#RawSyscall6) 函数的第2个返回值`r2` 现在永远返回0，而不是之前的未定义值(undefined value)。

  对于AIX和Solaris系统，可以使用 `Getrusage` 函数了。

- [time](https://tip.golang.org/pkg/time/)

  新方法 [`Duration.Abs`](https://tip.golang.org/pkg/time/#Duration.Abs) 可以得到duration的绝对值，更方便和安全，其中对于边界情况，−2⁶³ 会被转换为 2⁶³−1。

  新方法 [`Time.ZoneBounds`](https://tip.golang.org/pkg/time/#Time.ZoneBounds) 可以返回指定时间所在时区的开始和结束时间。

## 推荐阅读

**想了解Go泛型的使用方法、设计思路和最佳实践，推荐大家阅读**：

* [官方教程：Go泛型入门](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483720&idx=1&sn=57ec4877dfd364a59deacf1e74a4fb66&chksm=ce124e27f965c731432dcc89d1e0563cf84baaef482eaa068a91bee61f10cf85b433923b83b4&token=1782465473&lang=zh_CN#rd)
* [一文读懂Go泛型设计和使用场景](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483731&idx=1&sn=b2258b28e2f3c16b065a5a1b22c15b0d&chksm=ce124e3cf965c72a6a22e0ed15deda8238567407bbd7157a79753fc8b605727ab2153009493c&token=1782465473&lang=zh_CN#rd)
* [重磅：Go 1.18将移除用于泛型的constraints包](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483855&idx=1&sn=6ab4aeb140a1a08268dc8a0284a6f375&chksm=ce124ea0f965c7b6776061960d71e4ffb30484a82041f5b1d4786c4b49c4ffabc07a28b1cd48&token=1782465473&lang=zh_CN#rd)
* [泛型最佳实践：Go泛型设计者教你如何用泛型](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484015&idx=1&sn=576b2d8b84b3a8ce5bdd6952c2b84062&chksm=ce124d00f965c416b07dcb81c4dcb9cf75859b2787d4f00ec8c80b37ca42e58cc651420a3b33&token=1782465473&lang=zh_CN#rd)

**想了解Go原子操作和使用方法，推荐大家阅读**：

* [Go并发编程之原子操作sync/atomic](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484082&idx=1&sn=934787c9829391ba743bd611818ad0e2&chksm=ce124dddf965c4cb7d0f2d9d001ab4b7d949fbe87c4c8b7ee8d7498946824ec9aa6581cfe986&token=1782465473&lang=zh_CN#rd)

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