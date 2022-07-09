# Go 1.19要来了，看看都有哪些变化-第4篇

## 前言

Go官方团队在2022.06.11发布了Go 1.19 Beta 1版本，Go 1.19的正式release版本预计会在今年8月份发布。

让我们先睹为快，看看Go 1.19给我们带来了哪些变化。

这是Go 1.19版本更新内容详解的第4篇，欢迎大家关注公众号，及时获取本系列最新更新。

第1篇主要涉及Go泛型的改动、Go内存模型和原子操作的优化，原文链接：[Go 1.19版本变更内容第1篇](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484179&idx=1&sn=215ea3f092460118b2bc975935015874&chksm=ce124c7cf965c56a7c310b1059683d065810bd18368669d3d42a6cbbb0370d1593979a63620c#rd)。

第2篇主要涉及Go文档注释(doc comments)、编译约束(build constraint)以及Go命令的修改，原文链接：[Go 1.19版本变更内容第2篇](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484188&idx=1&sn=c14bafb1f89b3b3f988452c5a5f32884&chksm=ce124c73f965c5651a688c42561b02e38253b60943c77a0a6ad7b45621b4296d9e1acd47de7a#rd)。

第2篇主要涉及Go运行时、编译器、汇编器和链接器方面的改动和优化，原文链接：[Go 1.19版本变更内容第3篇]()。

## Go 1.19发布清单

和Go 1.18相比，改动相对较小，主要涉及语言(Language)、内存模型(Memory Model)、可移植性(Ports)、Go Tool工具链、运行时(Runtime)、编译器(Compiler)、汇编器(Assembler)、链接器(Linker)和核心库(Core library)等方面的优化。

本文重点介绍Go 1.19版本在核心库(Core library)方面的变化。

### 新的原子类型(New atomic types)

The [`sync/atomic`](https://tip.golang.org/pkg/sync/atomic/) package defines new atomic types [`Bool`](https://tip.golang.org/pkg/sync/atomic/#Bool), [`Int32`](https://tip.golang.org/pkg/sync/atomic/#Int32), [`Int64`](https://tip.golang.org/pkg/sync/atomic/#Int64), [`Uint32`](https://tip.golang.org/pkg/sync/atomic/#Uint32), [`Uint64`](https://tip.golang.org/pkg/sync/atomic/#Uint64), [`Uintptr`](https://tip.golang.org/pkg/sync/atomic/#Uintptr), and [`Pointer`](https://tip.golang.org/pkg/sync/atomic/#Pointer). These types hide the underlying values so that all accesses are forced to use the atomic APIs. [`Pointer`](https://tip.golang.org/pkg/sync/atomic/#Pointer) also avoids the need to convert to [`unsafe.Pointer`](https://tip.golang.org/pkg/unsafe/#Pointer) at call sites. [`Int64`](https://tip.golang.org/pkg/sync/atomic/#Int64) and [`Uint64`](https://tip.golang.org/pkg/sync/atomic/#Uint64) are automatically aligned to 64-bit boundaries in structs and allocated data, even on 32-bit systems.

### 路径查找(PATH lookups)

[`Command`](https://tip.golang.org/pkg/os/exec/#Command) and [`LookPath`](https://tip.golang.org/pkg/os/exec/#LookPath) no longer allow results from a PATH search to be found relative to the current directory. This removes a [common source of security problems](https://tip.golang.org/blog/path-security) but may also break existing programs that depend on using, say, `exec.Command("prog")` to run a binary named `prog` (or, on Windows, `prog.exe`) in the current directory. See the [`os/exec`](https://tip.golang.org/pkg/os/exec/) package documentation for information about how best to update such programs.

On Windows, `Command` and `LookPath` now respect the [`NoDefaultCurrentDirectoryInExePath`](https://docs.microsoft.com/en-us/windows/win32/api/processenv/nf-processenv-needcurrentdirectoryforexepatha) environment variable, making it possible to disable the default implicit search of “`.`” in PATH lookups on Windows systems.

### 核心库的微小改动

As always, there are various minor changes and updates to the library, made with the Go 1 [promise of compatibility](https://tip.golang.org/doc/go1compat) in mind. There are also various performance improvements, not enumerated here.

- [archive/zip](https://tip.golang.org/pkg/archive/zip/)

  [`Reader`](https://tip.golang.org/pkg/archive/zip/#Reader) now ignores non-ZIP data at the start of a ZIP file, matching most other implementations. This is necessary to read some Java JAR files, among other uses.

- [crypto/rand](https://tip.golang.org/pkg/crypto/rand/)

  [`Read`](https://tip.golang.org/pkg/crypto/rand/#Read) no longer buffers random data obtained from the operating system between calls.On Plan 9, `Read` has been reimplemented, replacing the ANSI X9.31 algorithm with fast key erasure.

- [crypto/tls](https://tip.golang.org/pkg/crypto/tls/)

  The `tls10default` `GODEBUG` option has been removed. It is still possible to enable TLS 1.0 client-side by setting [`Config.MinVersion`](https://tip.golang.org/pkg/crypto/tls#Config.MinVersion).The TLS server and client now reject duplicate extensions in TLS handshakes, as required by RFC 5246, Section 7.4.1.4 and RFC 8446, Section 4.2.

- [crypto/x509](https://tip.golang.org/pkg/crypto/x509/)

  [`CreateCertificate`](https://tip.golang.org/pkg/crypto/x509/#CreateCertificate) no longer supports creating certificates with `SignatureAlgorithm` set to `MD5WithRSA`.`CreateCertificate` no longer accepts negative serial numbers.[`ParseCertificate`](https://tip.golang.org/pkg/crypto/x509/#ParseCertificate) and [`ParseCertificateRequest`](https://tip.golang.org/pkg/crypto/x509/#ParseCertificateRequest) now reject certificates and CSRs which contain duplicate extensions.The new [`CertPool.Clone`](https://tip.golang.org/pkg/crypto/x509/#CertPool.Clone) and [`CertPool.Equal`](https://tip.golang.org/pkg/crypto/x509/#CertPool.Equal) methods allow cloning a `CertPool` and checking the equality of two `CertPool`s respectively.The new function [`ParseRevocationList`](https://tip.golang.org/pkg/crypto/x509/#ParseRevocationList) provides a faster, safer to use CRL parser which returns a [`RevocationList`](https://tip.golang.org/pkg/crypto/x509/#RevocationList). To support this addition, `RevocationList` adds new fields `RawIssuer`, `Signature`, `AuthorityKeyId`, and `Extensions`. The new method [`RevocationList.CheckSignatureFrom`](https://tip.golang.org/pkg/crypto/x509/#RevocationList.CheckSignatureFrom) checks that the signature on a CRL is a valid signature from a [`Certificate`](https://tip.golang.org/pkg/crypto/x509/#Certificate). With the new CRL functionality, the existing functions [`ParseCRL`](https://tip.golang.org/pkg/crypto/x509/#ParseCRL) and [`ParseDERCRL`](https://tip.golang.org/pkg/crypto/x509/#ParseDERCRL) are deprecated. Additionally the method [`Certificate.CheckCRLSignature`](https://tip.golang.org/pkg/crypto/x509#Certificate.CheckCRLSignature) is deprecated.When building paths, [`Certificate.Verify`](https://tip.golang.org/pkg/crypto/x509/#Certificate.Verify) now considers certificates to be equal when the subjects, public keys, and SANs are all equal. Before, it required byte-for-byte equality.

- [crypto/x509/pkix](https://tip.golang.org/pkg/crypto/x509/pkix)

  The types [`CertificateList`](https://tip.golang.org/pkg/crypto/x509/pkix#CertificateList) and [`TBSCertificateList`](https://tip.golang.org/pkg/crypto/x509/pkix#TBSCertificateList) have been deprecated. The new [`crypto/x509` CRL functionality](https://tip.golang.org/doc/go1.19#crypto/x509) should be used instead.

- [debug](https://tip.golang.org/pkg/debug/)

  The new `EM_LONGARCH` and `R_LARCH_*` constants support the loong64 port.

- [debug/pe](https://tip.golang.org/pkg/debug/pe/)

  The new [`File.COFFSymbolReadSectionDefAux`](https://tip.golang.org/pkg/debug/pe/#File.COFFSymbolReadSectionDefAux) method, which returns a [`COFFSymbolAuxFormat5`](https://tip.golang.org/pkg/debug/pe/#COFFSymbolAuxFormat5), provides access to COMDAT information in PE file sections. These are supported by new `IMAGE_COMDAT_*` and `IMAGE_SCN_*` constants.

- [encoding/binary](https://tip.golang.org/pkg/encoding/binary/)

  The new interface [`AppendByteOrder`](https://tip.golang.org/pkg/encoding/binary/#AppendByteOrder) provides efficient methods for appending a `uint16`, `uint32`, or `uint64` to a byte slice. [`BigEndian`](https://tip.golang.org/pkg/encoding/binary/#BigEndian) and [`LittleEndian`](https://tip.golang.org/pkg/encoding/binary/#LittleEndian) now implement this interface.Similarly, the new functions [`AppendUvarint`](https://tip.golang.org/pkg/encoding/binary/#AppendUvarint) and [`AppendVarint`](https://tip.golang.org/pkg/encoding/binary/#AppendVarint) are efficient appending versions of [`PutUvarint`](https://tip.golang.org/pkg/encoding/binary/#PutUvarint) and [`PutVarint`](https://tip.golang.org/pkg/encoding/binary/#PutVarint).

- [encoding/csv](https://tip.golang.org/pkg/encoding/csv/)

  The new method [`Reader.InputOffset`](https://tip.golang.org/pkg/encoding/csv/#Reader.InputOffset) reports the reader's current input position as a byte offset, analogous to `encoding/json`'s [`Decoder.InputOffset`](https://tip.golang.org/pkg/encoding/json/#Decoder.InputOffset).

- [encoding/xml](https://tip.golang.org/pkg/encoding/xml/)

  The new method [`Decoder.InputPos`](https://tip.golang.org/pkg/encoding/xml/#Decoder.InputPos) reports the reader's current input position as a line and column, analogous to `encoding/csv`'s [`Decoder.FieldPos`](https://tip.golang.org/pkg/encoding/csv/#Decoder.FieldPos).

- [flag](https://tip.golang.org/pkg/flag/)

  The new function [`TextVar`](https://tip.golang.org/pkg/flag/#TextVar) defines a flag with a value implementing [`encoding.TextUnmarshaler`](https://tip.golang.org/pkg/encoding/#TextUnmarshaler), allowing command-line flag variables to have types such as [`big.Int`](https://tip.golang.org/pkg/math/big/#Int), [`netip.Addr`](https://tip.golang.org/pkg/net/netip/#Addr), and [`time.Time`](https://tip.golang.org/pkg/time/#Time).

- [fmt](https://tip.golang.org/pkg/fmt/)

  The new functions [`Append`](https://tip.golang.org/pkg/fmt/#Append), [`Appendf`](https://tip.golang.org/pkg/fmt/#Appendf), and [`Appendln`](https://tip.golang.org/pkg/fmt/#Appendln) append formatted data to byte slices.

- [go/parser](https://tip.golang.org/pkg/go/parser/)

  The parser now recognizes `~x` as a unary expression with operator [token.TILDE](https://tip.golang.org/pkg/go/token#TILDE), allowing better error recovery when a type constraint such as `~int` is used in an incorrect context.

- [go/types](https://tip.golang.org/pkg/go/types/)

  The new methods [`Func.Origin`](https://tip.golang.org/pkg/go/types/#Func.Origin) and [`Var.Origin`](https://tip.golang.org/pkg/go/types/#Var.Origin) return the corresponding [`Object`](https://tip.golang.org/pkg/go/types/#Object) of the generic type for synthetic [`Func`](https://tip.golang.org/pkg/go/types/#Func) and [`Var`](https://tip.golang.org/pkg/go/types/#Var) objects created during type instantiation.It is no longer possible to produce an infinite number of distinct-but-identical [`Named`](https://tip.golang.org/pkg/go/types/#Named) type instantiations via recursive calls to [`Named.Underlying`](https://tip.golang.org/pkg/go/types/#Named.Underlying) or [`Named.Method`](https://tip.golang.org/pkg/go/types/#Named.Method).

- [hash/maphash](https://tip.golang.org/pkg/hash/maphash/)

  The new functions [`Bytes`](https://tip.golang.org/pkg/hash/maphash/#Bytes) and [`String`](https://tip.golang.org/pkg/hash/maphash/#String) provide an efficient way hash a single byte slice or string. They are equivalent to using the more general [`Hash`](https://tip.golang.org/pkg/hash/maphash/#Hash) with a single write, but they avoid setup overhead for small inputs.

- [html/template](https://tip.golang.org/pkg/html/template/)

  The type [`FuncMap`](https://tip.golang.org/pkg/html/template/#FuncMap) is now an alias for `text/template`'s [`FuncMap`](https://tip.golang.org/pkg/text/template/#FuncMap) instead of its own named type. This allows writing code that operates on a `FuncMap` from either setting.

- [image/draw](https://tip.golang.org/pkg/image/draw/)

  [`Draw`](https://tip.golang.org/pkg/image/draw/#Draw) with the [`Src`](https://tip.golang.org/pkg/image/draw/#Src) operator preserves non-premultiplied-alpha colors when destination and source images are both [`image.NRGBA`](https://tip.golang.org/pkg/image/#NRGBA) or both [`image.NRGBA64`](https://tip.golang.org/pkg/image/#NRGBA64). This reverts a behavior change accidentally introduced by a Go 1.18 library optimization; the code now matches the behavior in Go 1.17 and earlier.

- [io](https://tip.golang.org/pkg/io/)

  [`NopCloser`](https://tip.golang.org/pkg/io/#NopCloser)'s result now implements [`WriterTo`](https://tip.golang.org/pkg/io/#WriterTo) whenever its input does.[`MultiReader`](https://tip.golang.org/pkg/io/#MultiReader)'s result now implements [`WriterTo`](https://tip.golang.org/pkg/io/#WriterTo) unconditionally. If any underlying reader does not implement `WriterTo`, it is simulated appropriately.

- [mime](https://tip.golang.org/pkg/mime/)

  On Windows only, the mime package now ignores a registry entry recording that the extension `.js` should have MIME type `text/plain`. This is a common unintentional misconfiguration on Windows systems. The effect is that `.js` will have the default MIME type `text/javascript; charset=utf-8`. Applications that expect `text/plain` on Windows must now explicitly call [`AddExtensionType`](https://tip.golang.org/pkg/mime/#AddExtensionType).

- [net](https://tip.golang.org/pkg/net/)

  The pure Go resolver will now use EDNS(0) to include a suggested maximum reply packet length, permitting reply packets to contain up to 1232 bytes (the previous maximum was 512). In the unlikely event that this causes problems with a local DNS resolver, setting the environment variable `GODEBUG=netdns=cgo` to use the cgo-based resolver should work. Please report any such problems on [the issue tracker](https://tip.golang.org/issue/new).When a net package function or method returns an "I/O timeout" error, the error will now satisfy `errors.Is(err, context.DeadlineExceeded)`. When a net package function returns an "operation was canceled" error, the error will now satisfy `errors.Is(err, context.Canceled)`. These changes are intended to make it easier for code to test for cases in which a context cancellation or timeout causes a net package function or method to return an error, while preserving backward compatibility for error messages.[`Resolver.PreferGo`](https://tip.golang.org/pkg/net/#Resolver.PreferGo) is now implemented on Windows and Plan 9. It previously only worked on Unix platforms. Combined with [`Dialer.Resolver`](https://tip.golang.org/pkg/net/#Dialer.Resolver) and [`Resolver.Dial`](https://tip.golang.org/pkg/net/#Resolver.Dial), it's now possible to write portable programs and be in control of all DNS name lookups when dialing.The `net` package now has initial support for the `netgo` build tag on Windows. When used, the package uses the Go DNS client (as used by `Resolver.PreferGo`) instead of asking Windows for DNS results. The upstream DNS server it discovers from Windows may not yet be correct with complex system network configurations, however.

- [net/http](https://tip.golang.org/pkg/net/http/)

  [`ResponseWriter.WriteHeader`](https://tip.golang.org/pkg/net/http/#ResponseWriter) now supports sending user-defined 1xx informational headers.The `io.ReadCloser` returned by [`MaxBytesReader`](https://tip.golang.org/pkg/net/http/#MaxBytesReader) will now return the defined error type [`MaxBytesError`](https://tip.golang.org/pkg/net/http/#MaxBytesError) when its read limit is exceeded.The HTTP client will handle a 3xx response without a `Location` header by returning it to the caller, rather than treating it as an error.

- [net/url](https://tip.golang.org/pkg/net/url/)

  The new [`JoinPath`](https://tip.golang.org/pkg/net/url/#JoinPath) function and [`URL.JoinPath`](https://tip.golang.org/pkg/net/url/#URL.JoinPath) method create a new `URL` by joining a list of path elements.The `URL` type now distinguishes between URLs with no authority and URLs with an empty authority. For example, `http:///path` has an empty authority (host), while `http:/path` has none.The new [`URL`](https://tip.golang.org/pkg/net/url/#URL) field `OmitHost` is set to `true` when a `URL` has an empty authority.

- [os/exec](https://tip.golang.org/pkg/os/exec/)

  A [`Cmd`](https://tip.golang.org/pkg/os/exec/#Cmd) with a non-empty `Dir` field and nil `Env` now implicitly sets the `PWD` environment variable for the subprocess to match `Dir`.The new method [`Cmd.Environ`](https://tip.golang.org/pkg/os/exec/#Cmd.Environ) reports the environment that would be used to run the command, including the implicitly set `PWD` variable.

- [reflect](https://tip.golang.org/pkg/reflect/)

  The method [`Value.Bytes`](https://tip.golang.org/pkg/reflect/#Value.Bytes) now accepts addressable arrays in addition to slices.The methods [`Value.Len`](https://tip.golang.org/pkg/reflect/#Value.Len) and [`Value.Cap`](https://tip.golang.org/pkg/reflect/#Value.Cap) now successfully operate on a pointer to an array and return the length of that array, to match what the [builtin `len` and `cap` functions do](https://tip.golang.org/ref/spec#Length_and_capacity).

- [regexp/syntax](https://tip.golang.org/pkg/regexp/syntax/)

  Go 1.18 release candidate 1, Go 1.17.8, and Go 1.16.15 included a security fix to the regular expression parser, making it reject very deeply nested expressions. Because Go patch releases do not introduce new API, the parser returned [`syntax.ErrInternalError`](https://tip.golang.org/pkg/regexp/syntax/#ErrInternalError) in this case. Go 1.19 adds a more specific error, [`syntax.ErrNestingDepth`](https://tip.golang.org/pkg/regexp/syntax/#ErrNestingDepth), which the parser now returns instead.

- [runtime](https://tip.golang.org/pkg/runtime/)

  The [`GOROOT`](https://tip.golang.org/pkg/runtime/#GOROOT) function now returns the empty string (instead of `"go"`) when the binary was built with the `-trimpath` flag set and the `GOROOT` variable is not set in the process environment.

- [runtime/metrics](https://tip.golang.org/pkg/runtime/metrics/)

  The new `/sched/gomaxprocs:threads` [metric](https://tip.golang.org/pkg/runtime/metrics/#hdr-Supported_metrics) reports the current [`runtime.GOMAXPROCS`](https://tip.golang.org/pkg/runtime/#GOMAXPROCS) value.The new `/cgo/go-to-c-calls:calls` [metric](https://tip.golang.org/pkg/runtime/metrics/#hdr-Supported_metrics) reports the total number of calls made from Go to C. This metric is identical to the [`runtime.NumCgoCall`](https://tip.golang.org/pkg/runtime/#NumCgoCall) function.The new `/gc/limiter/last-enabled:gc-cycle` [metric](https://tip.golang.org/pkg/runtime/metrics/#hdr-Supported_metrics) reports the last GC cycle when the GC CPU limiter was enabled. See the [runtime notes](https://tip.golang.org/doc/go1.19#runtime) for details about the GC CPU limiter.

- [runtime/pprof](https://tip.golang.org/pkg/runtime/pprof/)

  Stop-the-world pause times have been significantly reduced when collecting goroutine profiles, reducing the overall latency impact to the application.`MaxRSS` is now reported in heap profiles for all Unix operating systems (it was previously only reported for `GOOS=android`, `darwin`, `ios`, and `linux`).

- [runtime/race](https://tip.golang.org/pkg/runtime/race/)

  The race detector has been upgraded to use thread sanitizer version v3 on all supported platforms except `windows/amd64` and `openbsd/amd64`, which remain on v2. Compared to v2, it is now typically 1.5x to 2x faster, uses half as much memory, and it supports an unlimited number of goroutines. On Linux, the race detector now requires at least glibc version 2.17.The race detector is now supported on `GOARCH=s390x`.Race detector support for `openbsd/amd64` has been removed from thread sanitizer upstream, so it is unlikely to ever be updated from v2.

- [runtime/trace](https://tip.golang.org/pkg/runtime/trace/)

  When tracing and the [CPU profiler](https://tip.golang.org/pkg/runtime/pprof#StartCPUProfile) are enabled simultaneously, the execution trace includes CPU profile samples as instantaneous events.

- [sort](https://tip.golang.org/pkg/sort/)

  The sorting algorithm has been rewritten to use [pattern-defeating quicksort](https://arxiv.org/pdf/2106.05123.pdf), which is faster for several common scenarios.The new function [Find](https://tip.golang.org/pkg/sort/#Find) is like [Search](https://tip.golang.org/pkg/sort/#Search) but often easier to use: it returns an additional boolean reporting whether an equal value was found.

- [strconv](https://tip.golang.org/pkg/strconv/)

  [`Quote`](https://tip.golang.org/pkg/strconv/#Quote) and related functions now quote the rune U+007F as `\x7f`, not `\u007f`, for consistency with other ASCII values.

- [syscall](https://tip.golang.org/pkg/syscall/)

  On PowerPC (`GOARCH=ppc64`, `ppc64le`), [`Syscall`](https://tip.golang.org/pkg/syscall/#Syscall), [`Syscall6`](https://tip.golang.org/pkg/syscall/#Syscall6), [`RawSyscall`](https://tip.golang.org/pkg/syscall/#RawSyscall), and [`RawSyscall6`](https://tip.golang.org/pkg/syscall/#RawSyscall6) now always return 0 for return value `r2` instead of an undefined value.On AIX and Solaris, `Getrusage` is now defined.

- [time](https://tip.golang.org/pkg/time/)

  The new method [`Duration.Abs`](https://tip.golang.org/pkg/time/#Duration.Abs) provides a convenient and safe way to take the absolute value of a duration, converting −2⁶³ to 2⁶³−1. (This boundary case can happen as the result of subtracting a recent time from the zero time.)The new method [`Time.ZoneBounds`](https://tip.golang.org/pkg/time/#Time.ZoneBounds) returns the start and end times of the time zone in effect at a given time. It can be used in a loop to enumerate all the known time zone transitions at a given location.

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