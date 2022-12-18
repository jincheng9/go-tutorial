# Go 1.20要来了，看看都有哪些变化-第2篇

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

第3篇主要涉及Go在运行时、编译器、汇编器、链接器等方面的优化，原文链接：[Go 1.20版本升级内容第3篇](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484638&idx=1&sn=459a22d4a9bf5d9715e70d3c25b05b93&chksm=ce124bb1f965c2a76bacc1135799ab268be66a861e99391b354a9f2dfd8c22a60853cc1d689d&token=1342188569&lang=zh_CN#rd)。

本文重点介绍Go 1.20在核心库方面的优化。

### New crypto/ecdh package

Go 1.20 adds a new [`crypto/ecdh`](https://tip.golang.org/pkg/crypto/ecdh/) package to provide direct support for Elliptic Curve Diffie-Hellman key exchange over NIST curves and Curve25519.

Programs should prefer to use `crypto/ecdh` or [`crypto/ecdsa`](https://tip.golang.org/pkg/crypto/ecdsa/) instead of the lower-level functionality in [`crypto/elliptic`](https://tip.golang.org/pkg/crypto/elliptic/).

### Wrapping multiple errors

Go 1.20 expands support for error wrapping to permit an error to wrap multiple other errors.

An error `e` can wrap more than one error by providing an `Unwrap` method that returns a `[]error`.

The [`errors.Is`](https://tip.golang.org/pkg/errors/#Is) and [`errors.As`](https://tip.golang.org/pkg/errors/#As) functions have been updated to inspect multiply wrapped errors.

The [`fmt.Errorf`](https://tip.golang.org/pkg/fmt/#Errorf) function now supports multiple occurrances of the `%w` format verb, which will cause it to return an error that wraps all of those error operands.

The new function [`errors.Join`](https://tip.golang.org/pkg/errors/#Join) returns an error wrapping a list of errors.

### HTTP ResponseController

The new [`"net/http".ResponseController`](https://tip.golang.org/pkg/net/http/#ResponseController) type provides access to extended per-request functionality not handled by the [`"net/http".ResponseWriter`](https://tip.golang.org/pkg/net/http/#ResponseWriter) interface.

Previously, we have added new per-request functionality by defining optional interfaces which a `ResponseWriter` can implement, such as [`Flusher`](https://tip.golang.org/pkg/net/http/#Flusher). These interfaces are not discoverable and clumsy to use.

The `ResponseController` type provides a clearer, more discoverable way to add per-handler controls. Two such controls also added in Go 1.20 are `SetReadDeadline` and `SetWriteDeadline`, which allow setting per-request read and write deadlines. For example:

```
func RequestHandler(w ResponseWriter, r *Request) {
  rc := http.NewResponseController(w)
  rc.SetWriteDeadline(0) // disable Server.WriteTimeout when sending a large response
  io.Copy(w, bigData)
}
```

### New ReverseProxy Rewrite hook

The [`httputil.ReverseProxy`](https://tip.golang.org/pkg/net/http/httputil/#ReverseProxy) forwarding proxy includes a new [`Rewrite`](https://tip.golang.org/pkg/net/http/httputil/#ReverseProxy.Rewrite) hook function, superseding the previous `Director` hook.

The `Rewrite` hook accepts a [`ProxyRequest`](https://tip.golang.org/pkg/net/http/httputil/#ProxyRequest) parameter, which includes both the inbound request received by the proxy and the outbound request that it will send. Unlike `Director` hooks, which only operate on the outbound request, this permits `Rewrite` hooks to avoid certain scenarios where a malicious inbound request may cause headers added by the hook to be removed before forwarding. See [issue #50580](https://go.dev/issue/50580).

The [`ProxyRequest.SetURL`](https://tip.golang.org/pkg/net/http/httputil/#ProxyRequest.SetURL) method routes the outbound request to a provided destination and supersedes the `NewSingleHostReverseProxy` function. Unlike `NewSingleHostReverseProxy`, `SetURL` also sets the `Host` header of the outbound request.

The [`ProxyRequest.SetXForwarded`](https://tip.golang.org/pkg/net/http/httputil/#ProxyRequest.SetXForwarded) method sets the `X-Forwarded-For`, `X-Forwarded-Host`, and `X-Forwarded-Proto` headers of the outbound request. When using a `Rewrite`, these headers are not added by default.

An example of a `Rewrite` hook using these features is:

```
proxyHandler := &httputil.ReverseProxy{
  Rewrite: func(r *httputil.ProxyRequest) {
    r.SetURL(outboundURL) // Forward request to outboundURL.
    r.SetXForwarded()     // Set X-Forwarded-* headers.
    r.Out.Header.Set("X-Additional-Header", "header set by the proxy")
  },
}
```

[`ReverseProxy`](https://tip.golang.org/pkg/net/http/httputil/#ReverseProxy) no longer adds a `User-Agent` header to forwarded requests when the incoming request does not have one.

### Minor changes to the library

As always, there are various minor changes and updates to the library, made with the Go 1 [promise of compatibility](https://tip.golang.org/doc/go1compat) in mind. There are also various performance improvements, not enumerated here.

- [archive/tar](https://tip.golang.org/pkg/archive/tar/)

  When the `GODEBUG=tarinsecurepath=0` environment variable is set, [`Reader.Next`](https://tip.golang.org/pkg/archive/tar/#Reader.Next) method will now return the error [`ErrInsecurePath`](https://tip.golang.org/pkg/archive/tar/#ErrInsecurePath) for an entry with a file name that is an absolute path, refers to a location outside the current directory, contains invalid characters, or (on Windows) is a reserved name such as `NUL`. A future version of Go may disable insecure paths by default.

- [archive/zip](https://tip.golang.org/pkg/archive/zip/)

  When the `GODEBUG=zipinsecurepath=0` environment variable is set, [`NewReader`](https://tip.golang.org/pkg/archive/zip/#NewReader) will now return the error [`ErrInsecurePath`](https://tip.golang.org/pkg/archive/zip/#ErrInsecurePath) when opening an archive which contains any file name that is an absolute path, refers to a location outside the current directory, contains invalid characters, or (on Windows) is a reserved names such as `NUL`. A future version of Go may disable insecure paths by default.Reading from a directory file that contains file data will now return an error. The zip specification does not permit directory files to contain file data, so this change only affects reading from invalid archives.

- [bytes](https://tip.golang.org/pkg/bytes/)

  The new [`CutPrefix`](https://tip.golang.org/pkg/bytes/#CutPrefix) and [`CutSuffix`](https://tip.golang.org/pkg/bytes/#CutSuffix) functions are like [`TrimPrefix`](https://tip.golang.org/pkg/bytes/#TrimPrefix) and [`TrimSuffix`](https://tip.golang.org/pkg/bytes/#TrimSuffix) but also report whether the string was trimmed.The new [`Clone`](https://tip.golang.org/pkg/bytes/#Clone) function allocates a copy of a byte slice.

- [context](https://tip.golang.org/pkg/context/)

  The new [`WithCancelCause`](https://tip.golang.org/pkg/context/#WithCancelCause) function provides a way to cancel a context with a given error. That error can be retrieved by calling the new [`Cause`](https://tip.golang.org/pkg/context/#Cause) function.

- [crypto/ecdsa](https://tip.golang.org/pkg/crypto/ecdsa/)

  The new [`PrivateKey.ECDH`](https://tip.golang.org/pkg/crypto/ecdsa/#PrivateKey.ECDH) method converts an `ecdsa.PrivateKey` to an `ecdh.PrivateKey`.

- [crypto/ed25519](https://tip.golang.org/pkg/crypto/ed25519/)

  The [`PrivateKey.Sign`](https://tip.golang.org/pkg/crypto/ed25519/#PrivateKey.Sign) method and the [`VerifyWithOptions`](https://tip.golang.org/pkg/crypto/ed25519/#VerifyWithOptions) function now support signing pre-hashed messages with Ed25519ph, indicated by an [`Options.HashFunc`](https://tip.golang.org/pkg/crypto/ed25519/#Options.HashFunc) that returns [`crypto.SHA512`](https://tip.golang.org/pkg/crypto/#SHA512). They also now support Ed25519ctx and Ed25519ph with context, indicated by setting the new [`Options.Context`](https://tip.golang.org/pkg/crypto/ed25519/#Options.Context) field.

- [crypto/elliptic](https://tip.golang.org/pkg/crypto/elliptic/)

  Use of custom [`Curve`](https://tip.golang.org/pkg/crypto/elliptic/#Curve) implementations not provided by this package (that is, curves other than [`P224`](https://tip.golang.org/pkg/crypto/elliptic/#P224), [`P256`](https://tip.golang.org/pkg/crypto/elliptic/#P256), [`P384`](https://tip.golang.org/pkg/crypto/elliptic/#P384), and [`P521`](https://tip.golang.org/pkg/crypto/elliptic/#P521)) is deprecated.

- [crypto/rsa](https://tip.golang.org/pkg/crypto/rsa/)

  The new field [`OAEPOptions.MGFHash`](https://tip.golang.org/pkg/crypto/rsa/#OAEPOptions.MGFHash) allows configuring the MGF1 hash separately for OAEP encryption.

- [crypto/subtle](https://tip.golang.org/pkg/crypto/subtle/)

  The new function [`XORBytes`](https://tip.golang.org/pkg/crypto/subtle/#XORBytes) XORs two byte slices together.

- [crypto/tls](https://tip.golang.org/pkg/crypto/tls/)

  The TLS client now shares parsed certificates across all clients actively using that certificate. The savings can be significant in programs that make many concurrent connections to a server or collection of servers sharing any part of their certificate chains.For a handshake failure due to a certificate verification failure, the TLS client and server now return an error of the new type [`CertificateVerificationError`](https://tip.golang.org/pkg/crypto/tls/#CertificateVerificationError), which includes the presented certificates.

- [crypto/x509](https://tip.golang.org/pkg/crypto/x509/)

  [`CreateCertificateRequest`](https://tip.golang.org/pkg/crypto/x509/#CreateCertificateRequest) and [`MarshalPKCS8PrivateKey`](https://tip.golang.org/pkg/crypto/x509/#MarshalPKCS8PrivateKey) now support keys of type [`*crypto/ecdh.PrivateKey`](https://tip.golang.org/pkg/crypto/ecdh.PrivateKey). [`CreateCertificate`](https://tip.golang.org/pkg/crypto/x509/#CreateCertificate) and [`MarshalPKIXPublicKey`](https://tip.golang.org/pkg/crypto/x509/#MarshalPKIXPublicKey) now support keys of type [`*crypto/ecdh.PublicKey`](https://tip.golang.org/pkg/crypto/ecdh.PublicKey). X.509 unmarshaling continues to unmarshal elliptic curve keys into `*ecdsa.PublicKey` and `*ecdsa.PrivateKey`. Use their new `ECDH` methods to convert to the `crypto/ecdh` form.The new [`SetFallbackRoots`](https://tip.golang.org/pkg/crypto/x509/#SetFallbackRoots) function allows a program to define a set of fallback root certificates in case the operating system verifier or standard platform root bundle is unavailable at runtime. It will most commonly be used with a new package, [golang.org/x/crypto/x509roots/fallback](https://tip.golang.org/pkg/golang.org/x/crypto/x509roots/fallback), which will provide an up to date root bundle.

- [debug/elf](https://tip.golang.org/pkg/debug/elf/)

  Attempts to read from a `SHT_NOBITS` section using [`Section.Data`](https://tip.golang.org/pkg/debug/elf/#Section.Data) or the reader returned by [`Section.Open`](https://tip.golang.org/pkg/debug/elf/#Section.Open) now return an error.Additional [`R_LARCH_*`](https://tip.golang.org/pkg/debug/elf/#R_LARCH) constants are defined for use with LoongArch systems.Additional [`R_PPC64_*`](https://tip.golang.org/pkg/debug/elf/#R_PPC64) constants are defined for use with PPC64 ELFv2 relocations.The constant value for [`R_PPC64_SECTOFF_LO_DS`](https://tip.golang.org/pkg/debug/elf/#R_PPC64_SECTOFF_LO_DS) is corrected, from 61 to 62.

- [debug/gosym](https://tip.golang.org/pkg/debug/gosym/)

  Due to a change of [Go's symbol naming conventions](https://tip.golang.org/doc/go1.20#linker), tools that process Go binaries should use Go 1.20's `debug/gosym` package to transparently handle both old and new binaries.

- [debug/pe](https://tip.golang.org/pkg/debug/pe/)

  Additional [`IMAGE_FILE_MACHINE_RISCV*`](https://tip.golang.org/pkg/debug/pe/#IMAGE_FILE_MACHINE_RISCV128) constants are defined for use with RISC-V systems.

- [encoding/binary](https://tip.golang.org/pkg/encoding/binary/)

  The [`ReadVarint`](https://tip.golang.org/pkg/encoding/binary/#ReadVarint) and [`ReadUvarint`](https://tip.golang.org/pkg/encoding/binary/#ReadUvarint) functions will now return `io.ErrUnexpectedEOF` after reading a partial value, rather than `io.EOF`.

- [encoding/xml](https://tip.golang.org/pkg/encoding/xml/)

  The new [`Encoder.Close`](https://tip.golang.org/pkg/encoding/xml/#Encoder.Close) method can be used to check for unclosed elements when finished encoding.The decoder now rejects element and attribute names with more than one colon, such as `<a:b:c>`, as well as namespaces that resolve to an empty string, such as `xmlns:a=""`.The decoder now rejects elements that use different namespace prefixes in the opening and closing tag, even if those prefixes both denote the same namespace.

- [errors](https://tip.golang.org/pkg/errors/)

  The new [`Join`](https://tip.golang.org/pkg/errors/#Join) function returns an error wrapping a list of errors.

- [fmt](https://tip.golang.org/pkg/fmt/)

  The [`Errorf`](https://tip.golang.org/pkg/fmt/#Errorf) function supports multiple occurrences of the `%w` format verb, returning an error that unwraps to the list of all arguments to `%w`.The new [`FormatString`](https://tip.golang.org/pkg/fmt/#FormatString) function recovers the formatting directive corresponding to a [`State`](https://tip.golang.org/pkg/fmt/#State), which can be useful in [`Formatter`](https://tip.golang.org/pkg/fmt/#Formatter). implementations.

- [go/ast](https://tip.golang.org/pkg/go/ast/)

  The new [`RangeStmt.Range`](https://tip.golang.org/pkg/go/ast/#RangeStmt.Range) field records the position of the `range` keyword in a range statement.The new [`File.FileStart`](https://tip.golang.org/pkg/go/ast/#File.FileStart) and [`File.FileEnd`](https://tip.golang.org/pkg/go/ast/#File.FileEnd) fields record the position of the start and end of the entire source file.

- [go/token](https://tip.golang.org/pkg/go/token/)

  The new [`FileSet.RemoveFile`](https://tip.golang.org/pkg/go/token/#FileSet.RemoveFile) method removes a file from a `FileSet`. Long-running programs can use this to release memory associated with files they no longer need.

- [go/types](https://tip.golang.org/pkg/go/types/)

  The new [`Satisfies`](https://tip.golang.org/pkg/go/types/#Satisfies) function reports whether a type satisfies a constraint. This change aligns with the [new language semantics](https://tip.golang.org/doc/go1.20#language) that distinguish satisfying a constraint from implementing an interface.

- [io](https://tip.golang.org/pkg/io/)

  The new [`OffsetWriter`](https://tip.golang.org/pkg/io/#OffsetWriter) wraps an underlying [`WriterAt`](https://tip.golang.org/pkg/io/#WriterAt) and provides `Seek`, `Write`, and `WriteAt` methods that adjust their effective file offset position by a fixed amount.

- [io/fs](https://tip.golang.org/pkg/io/fs/)

  The new error [`SkipAll`](https://tip.golang.org/pkg/io/fs/#SkipAll) terminates a [`WalkDir`](https://tip.golang.org/pkg/io/fs/#WalkDir) immediately but successfully.

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