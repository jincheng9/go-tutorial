# Go 1.20要来了，看看都有哪些变化-第2篇

## 前言

Go官方团队在2022.12.08发布了Go 1.20 rc1(release candidate)版本，Go 1.20的正式release版本预计会在2023年2月份发布。

让我们先睹为快，看看Go 1.20给我们带来了哪些变化。

安装方法：

```bash
$ go install golang.org/dl/go1.20rc1@latest
$ go1.20rc1 download
```

这是Go 1.20版本更新内容详解的第2篇，欢迎大家关注公众号，及时获取本系列最新更新。

## Go 1.20发布清单

和Go 1.19相比，改动内容适中，主要涉及语言(Language)、可移植性(Ports)、工具链(Go Tools)、运行时(Runtime)、编译器(Compiler)、汇编器(Assembler)、链接器(Linker)和核心库(Core library)等方面的优化。

第1篇主要涉及Go 1.20在语言、可移植性方面的优化，原文链接：[Go 1.20版本升级内容第1篇](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484629&idx=1&sn=60a01d3cc85ef2462156f0565c30738d&chksm=ce124bbaf965c2ac351cd9c602e8b67d5119b2a89a7f2de0289bdeb7608ae589c329eb8f7275&token=1619842941&lang=zh_CN#rd)。

本文重点介绍Go 1.20在Go工具链方面的优化。

### Go command

`$GOROOT/pkg`路径不再存储标准库源代码编译后生成的文件，包括以下几点：

* `go install`不再往`$GOROOT/pkg`目录写文件。
* `go build`不再检查`$GOROOT/pkg`下的文件。
* Go发布包不再带有这些编译文件。

以macOS环境来演示：在Go 1.16版本里，`$GOROOT/pkg`目录下的内容如下：

```bash
$ go version
go version go1.16.5 darwin/amd64
$ go env GOROOT
/usr/local/opt/go/libexec
$ ls /usr/local/opt/go/libexec/pkg
darwin_amd64		darwin_amd64_race	include			tool
```

但是在Go 1.20rc1版本里，`$GOROOT/pkg`目录下的内容如下：

```bash
$ go1.20rc1 version
go version go1.20rc1 darwin/amd64
$ go1.20rc1 env GOROOT
/Users/xxx/sdk/go1.20rc1
$ ls /Users/xxx/sdk/go1.20rc1/pkg
include	tool
```

少了`darwin_amd64`和`darwin_amd64_race`这2个文件夹：

```bash
$ ls /usr/local/opt/go/libexec/pkg/darwin_amd64
archive		database	go		io		net.a		runtime		testing.a
bufio.a		debug		hash		io.a		os		runtime.a	text
bytes.a		embed.a		hash.a		log		os.a		sort.a		time
cmd		encoding	html		log.a		path		strconv.a	time.a
compress	encoding.a	html.a		math		path.a		strings.a	unicode
container	errors.a	image		math.a		plugin.a	sync		unicode.a
context.a	expvar.a	image.a		mime		reflect.a	sync.a		vendor
crypto		flag.a		index		mime.a		regexp		syscall.a
crypto.a	fmt.a		internal	net		regexp.a	testing
$ ls /usr/local/opt/go/libexec/pkg/darwin_amd64_race/
archive		debug		hash		io.a		os		runtime.a	text
bufio.a		embed.a		hash.a		log		os.a		sort.a		time
bytes.a		encoding	html		log.a		path		strconv.a	time.a
compress	encoding.a	html.a		math		path.a		strings.a	unicode
container	errors.a	image		math.a		plugin.a	sync		unicode.a
context.a	expvar.a	image.a		mime		reflect.a	sync.a		vendor
crypto		flag.a		index		mime.a		regexp		syscall.a
crypto.a	fmt.a		internal	net		regexp.a	testing
database	go		io		net.a		runtime		testing.a
```

从Go 1.20开始，标准库的package会按需编译，编译后生成的文件会缓存在编译缓存(build cache)里，就像非`GOROOT`下的package一样。这个修改减小了Go安装包的大小。



`go test -json`的实现在Go 1.20版本更加鲁棒，对使用`go test -json`的开发者来说，不用做任何改变。

但是，直接调用Go工具`test2json`的开发者，需要给测试的可执行程序增加`-v=test2json`参数，例如`go test -v=test2json`或者`./pkg.test -test.v=test2json`，而不是仅仅加一个`-v`标记。

`go test -json`的另外一个修改是在每个测试程序执行的开始，增加了一个`Action`事件。当使用go命令同时运行多个测试程序时，这些`Action`事件的执行会按照命令行里的package的顺序按序执行。

`go`命令现在新增了和CPU架构相关的编译标记参数，例如`amd64.v2`。有了这个参数后，在业务代码实现的时候就可以根据CPU架构的不同而做不同的处理。更多细节可以参考：[`go help buildconstraint`](https://tip.golang.org/cmd/go#hdr-Build_constraints) 。

`go`子命令现在支持`-C <dir>`参数，可以在执行命令前改变目录到`<dir>`下。如果一个脚本要在多个不同的Go module下执行，这个特性会带来方便。

`go` `build` 和 `go` `test` 命令不再支持`-i`参数，这个参数从 [Go 1.16版本开始弃用](https://go.dev/issue/41696)。

`go generate` 命令接受 `-skip <pattern>` 参数，可以跳过匹配 `<pattern>`格式的 `//go:generate` 指令。

`go test`命令接受`-skip <pattern>` 参数，可以跳过匹配 `<pattern>`格式的测试用例。

`go build`, `go install`和其它编译相关的命令新增了一个`-pgo`标记参数，可以辅助开发者做程序优化。`-pgo`指定的是profile文件的路径。如果`-pgo=auto`，那go命令会在main这个包的路径下去找名为`default.pgo`的文件。`-pgo=off`可以关闭优化。详情可以参考：[PGO Proposal](https://github.com/golang/go/issues/55022)。

`go build`, `go install`和其他编译相关的命令新增了一个`-cover`标记参数，可以用来对编译出来的可执行程序做代码覆盖率收集，详情可以参考本文后面介绍。

#### go version

`go version -m`命令支持读取和解析更多类型的Go二进制文件。

比如通过`go build -buildmode=c-share`编译出来的Windows DLL文件以及没有可执行权限的Linux二进制文件，现在都可以被`go version -m`解析和识别到。

### Cgo

早期Go语言的一些标准库是用C语言实现的，需要依赖cgo来作为go语言和C语言的桥梁。

现在Go语言开始去除对C语言的的依赖。

从Go 1.20版本开始，如果机器上没有C语言的工具链，go命令会默认禁用`cgo`。

具体而言：如果没有设置`CGO_ENABLED`和`CC`环境变量，而且默认的C语言编译器(例如`clang`和`gcc`)也找不到，那`CGO_ENABLED`会默认为0。当然开发者可以通过设置`CGO_ENABLED`环境变量的值来改变`CGO_ENABLED`的值。

这个修改会让Go语言减少对C语言工具链的依赖，适配更多的环境，尤其是最小化的容器环境以及macOS环境。

Go标准库里使用了cgo的package有： [`net`](https://tip.golang.org/pkg/net/), [`os/user`](https://tip.golang.org/pkg/os/user/) 和 [`plugin`](https://tip.golang.org/pkg/plugin/)。

在macOS环境，`net`和`os/user`包已经被重写了，不再依赖cgo。现在`net`和`os/user`的代码实现既可以用cgo编译，也可以不用cgo编译。同时，在macOS环境，race detector已经被重写了，不再依赖cgo。

在Windows环境，`net`和`os/user`没有使用过cgo。

在其它操作系统上，如果编译的时候禁用了cgo，那会使用这些包的纯go语言实现。

### Cover(代码覆盖率检测)

Go 1.20版本之前只支持对单元测试(unit test)收集代码覆盖率，从Go 1.20版本开始支持对任何Go程序做代码覆盖率收集。

那如何收集呢？需要做如下操作：

* 给`go build`编译命令增加`-cover`标记
* 给环境变量`GOCOVERDIR`赋值为某个路径
* 运行`go build`编译出来的可执行程序时，会把代码覆盖率文件输出到`GOCOVERDIR`指定的路径下。

详细的介绍文档和使用说明可以参考： [coverage for integration tests' landing page](https://go.dev/testing/coverage) 。

想了解设计原理和实现的可以参考： [proposal](https://golang.org/issue/51430)。

### Vet

#### 检测循环变量被嵌套子函数错误使用的场景

```go
func TestTLog(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		value int
	}{
		{name: "test 1", value: 1},
		{name: "test 2", value: 2},
		{name: "test 3", value: 3},
		{name: "test 4", value: 4},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
      t.Parallel()
			// Here you test tc.value against a test function.
			// Let's use t.Log as our test function :-)
			t.Log(tc.value)
		})
	}
}
```

大家可以猜一下这段程序里t.Log打印的结果是什么？结果是为1,2,3,4还是4,4,4,4？是否和自己的预期相符。

想了解详情的可以参考：[Be Careful with Table Driven Tests and t.Parallel()](https://gist.github.com/posener/92a55c4cd441fc5e5e85f27bca008721)。

这个本质上和goroutine与闭包函数一起使用时，遇到的循环变量问题一样。

Go 1.20版本开始通过`go vet`可以检测出单元测试里的这类问题。

#### 检查错误时间格式

对于 [`Time.Format`](https://tip.golang.org/pkg/time/#Time.Format) 和[`time.Parse`](https://tip.golang.org/pkg/time/#Parse)，如果代码里要转成yyyy-dd-mm的格式，会给出提示。

因为yyyy-dd-mm不符合常用的日期格式标准，ISO 8601日期格式是yyyy-mm-dd格式。

## 总结

下一篇会介绍Go 1.20在运行时、编译器、汇编器、链接器和核心库的优化工作，有一些内容值得学习，欢迎大家保持关注。



## 推荐阅读

* [Go 1.20要来了，看看都有哪些变化-第1篇](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484629&idx=1&sn=60a01d3cc85ef2462156f0565c30738d&chksm=ce124bbaf965c2ac351cd9c602e8b67d5119b2a89a7f2de0289bdeb7608ae589c329eb8f7275&token=1342188569&lang=zh_CN#rd)

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