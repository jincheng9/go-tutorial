# Fuzzing: 一文读懂Go Fuzzing设计和原理

## 背景

Go 1.18除了引入泛型(generics)这个重大设计之外，Go官方团队在Go 1.18工具链里还引入了fuzzing模糊测试。

Go fuzzing的主要开发者是Katie Hockman, Jay Conrod和Roland Shoemaker。

**编者注**：Katie Hockman已于2022.02.19从Google离职，Jay Conrod也于2021年10月离开Google。



##  什么是Fuzzing

Fuzzing中文含义是模糊测试，是一种自动化测试技术，可以随机生成测试数据集，然后调用要测试的功能代码来检查功能是否符合预期。

模糊测试(fuzz test)是对单元测试(unit test)的补充，并不是要替代单元测试。

单元测试是检查指定的输入得到的结果是否和预期的输出结果一致，测试数据集比较有限。

模糊测试可以生成随机测试数据，找出单元测试覆盖不到的场景，进而发现程序的潜在bug和安全漏洞。



## Go Fuzzing怎么使用

Fuzzing在Go语言里并不是一个全新的概念，在Go官方团队发布Go Fuzzing之前，GitHub上已经有了类似的模糊测试工具[go-fuzz](https://github.com/dvyukov/go-fuzz)。

Go官方团队的Fuzzing实现借鉴了go-fuzz的设计思想。

Go 1.18把Fuzzing整合到了`go test`工具链和`testing`包里。

### 示例

下面举个例子说明下Fuzzing如何使用。

对于如下的字符串反转函数`Reverse`，大家可以思考下这段代码有什么潜在问题？

```go
// main.go
package fuzz

func Reverse(s string) string {
	bs := []byte(s)
	length := len(bs)
	for i := 0; i < length/2; i++ {
		bs[i], bs[length-i-1] = bs[length-i-1], bs[i]
	}
	return string(bs)
}
```

### 编写Fuzzing模糊测试

如果没有发现上面代码的bug，我们不妨写一个Fuzzing模糊测试函数，来发现上面代码的潜在问题。

Go Fuzzing模糊测试函数的语法如下所示：

* 模糊测试函数定义在`xxx_test.go`文件里，这点和Go已有的单元测试(unit test)和性能测试(benchmark test)一样。

* 函数名以`Fuzz`开头，参数是`* testing.F`类型，`testing.F`类型有2个重要方法`Add`和`Fuzz`。
* `Add`方法是用于添加种子语料(seed corpus)数据，Fuzzing底层可以根据种子语料数据自动生成随机测试数据。
* `Fuzz`方法接收一个函数类型的变量作为参数，该函数类型的第一个参数必须是`*testing.T`类型，其余的参数类型和`Add`方法里传入的实参类型保持一致。比如下面的例子里，`f.Add(5, "hello")`传入的第一个实参是`5`，第二个实参是`hello`，对应的是`i int`和`s string`。

![](../../img/fuzzing.png)

* Go Fuzzing底层会根据`Add`里指定的种子语料，随机生成测试数据，执行模糊测试。比如上图的例子里，会根据`Add`里指定的`5`和`hello`，随机生产新的测试数据，赋值给`i`和`s`，然后不断调用作为`f.Fuzz`方法的实参，也就是`func(t *testing.T, i int, s string){...}`这个函数。

知道了上述规则后，我们来给`Reverse`函数编写一个如下的模糊测试函数。

```go
// fuzz_test.go
package fuzz

import (
	"testing"
	"unicode/utf8"
)

func FuzzReverse(f *testing.F) {
	str_slice := []string{"abc", "bb"}
	for _, v := range str_slice {
		f.Add(v)
	}
	f.Fuzz(func(t *testing.T, str string) {
		rev_str1 := Reverse(str)
		rev_str2 := Reverse(rev_str1)
		if str != rev_str2 {
			t.Errorf("fuzz test failed. str:%s, rev_str1:%s, rev_str2:%s", str, rev_str1, rev_str2)
		}
		if utf8.ValidString(str) && !utf8.ValidString(rev_str1) {
			t.Errorf("reverse result is not utf8. str:%s, len: %d, rev_str1:%s", str, len(str), rev_str1)
		}
	})
}
```

### 运行Fuzzing测试

使用的Go版本要求是`go 1.18beta 1`或以上版本，执行如下命令可以进行Fuzzing测试，结果如下：

```sh
$ go1.18beta1 test -v -fuzz=Fuzz
fuzz: elapsed: 0s, gathering baseline coverage: 0/111 completed
fuzz: minimizing 60-byte failing input file
fuzz: elapsed: 0s, gathering baseline coverage: 5/111 completed
--- FAIL: FuzzReverse (0.04s)
    --- FAIL: FuzzReverse (0.00s)
        fuzz_test.go:20: reverse result is not utf8. str:æ, len: 2, rev_str1:��
    
    Failing input written to testdata/fuzz/FuzzReverse/ce9e8c80e2c2de2c96ab9e63b1a8cf18cea932b7d8c6c9c207d5978e0f19027a
    To re-run:
    go test -run=FuzzReverse/ce9e8c80e2c2de2c96ab9e63b1a8cf18cea932b7d8c6c9c207d5978e0f19027a
FAIL
exit status 1
FAIL    example/fuzz    0.179s
```

重点看`fuzz_test.go:20: reverse result is not utf8. str:æ, len: 2, rev_str1:��`

这个例子里，随机生成了一个字符串`æ`，这是由2个字节组成的一个UTF-8字符串，按照`Reverse`函数进行反转后，得到了一个非UTF-8的字符串`��`。

所以我们之前实现的按照字节进行字符串反转的函数`Reverse`是有bug的，该函数对于ASCII码里的字符组成的字符串是可以正确反转的，但是对于非ASCII码里的字符，如果简单按照字节进行反转，得到的可能是一个非法的字符串。

感兴趣的朋友，可以看看如果对字符串"吃"，调用`Reverse` 函数，会得到怎样的结果。

**注意**：如果Go Fuzzing运行过程中发现了你的bug，会把对应的输入数据写到`go test`所在目录下的`testdata/fuzz/FuzzXXX`子目录下。比如上面的例子里，`go1.18beta1 test -v -fuzz=Fuzz`的输出结果里打印了如下内容：`Failing input written to testdata/fuzz/FuzzReverse/ce9e8c80e2c2de2c96ab9e63b1a8cf18cea932b7d8c6c9c207d5978e0f19027a`，这就表示把这个测试输入写到了`testdata/fuzz/FuzzReverse/xxx`这个语料文件里。



## Go Fuzzing的底层机制

 `go test` 执行的时候，会为每个被测试的package先编译生成一个可执行文件，然后运行这个可执行文件得到对应package的`TestXXX`和`BenchmarkXXX`的测试结果。

You may already know that `go test` builds a test executable for each package being tested, then runs those executables to get test and benchmark results. Fuzzing follows this pattern, though there are some differences.

When `go test` is invoked with the `-fuzz` flag, `go test` compiles the test executable with additional coverage instrumentation. The Go compiler already had instrumentation support for [libFuzzer](https://llvm.org/docs/LibFuzzer.html), so we reused that. The compiler adds an 8-bit counter to each basic block. The counter is fast and approximate: it wraps on overflow, and there's no synchronization across threads. (We had to tell the race detector not to complain about writes to these counters). The counter data is used at run-time by the [internal/fuzz](https://pkg.go.dev/internal/fuzz) package, where most of the fuzzing logic is.

After `go test` builds an instrumented executable, it runs it as usual. This is called the coordinator process. This process is started with most of the flags that were passed to `go test`, including `-fuzz=pattern`, which it uses to identify which target to fuzz; for now, only one target may be fuzzed per `go test` invocation ([#46312](https://github.com/golang/go/issues/46312)). When that target calls [`F.Fuzz`](https://pkg.go.dev/testing@go1.18beta2#F.Fuzz), control is passed to [`fuzz.CoordinateFuzzing`](https://pkg.go.dev/internal/fuzz#CoordinateFuzzing), which initializes the fuzzing system and begins the coordinator event loop.

The coordinator starts several worker processes, which run the same test executable and perform the actual fuzzing. Workers are started with an undocumented command line flag that tells them to be workers. Fuzzing must be done in separate processes so that if a worker process crashes entirely, the coordinator can still find and record the input that caused the crash.

![Diagram showing the relationship between fuzzing processes. At the top is a box showing "go test (cmd/go)". An arrow points downward to a box labelled "coordinator (test binary)". From that, three arrows point downward to three boxes labelled "worker (test binary)".](https://jayconrod.com/images/fuzz-processes.svg)

The coordinator communicates with each worker using an improvised JSON-based RPC protocol over a pair of pipes. The protocol is pretty basic because we didn't need anything sophisticated like gRPC, and we didn't want to introduce anything new into the standard library. Each worker also keeps some state in a memory mapped temporary file, shared with the coordinator. Mostly this is just an iteration count and random number generator state. If the worker crashes entirely, the coordinator can recover its state from shared memory without requiring the worker to politely send a message over the pipes first.

After the coordinator starts the workers, it gathers baseline coverage by sending workers inputs from the seed corpus and the fuzzing cache corpus (in a subdirectory of `$GOCACHE`). Each worker runs its given input, then reports back with a snapshot of its coverage counters. The coordinator coarsens and merges these counters into a combined coverage array.

Next, the coordinator sends out inputs from the seed corpus and cached corpus for fuzzing: each worker is given an input and a copy of the baseline coverage array. Each worker then randomly mutates its input (flipping bits, deleting or inserting bytes, etc.) and calls the fuzz function. In order to reduce communication overhead, each worker can keep mutating and calling for 100 ms without further input from the coordinator. After each call, the worker checks whether an error was reported (with [`T.Fail`](https://pkg.go.dev/testing@go1.18beta2#T.Fail)) or new coverage was found compared with the baseline coverage array. If so, the worker reports the "interesting" input back to the coordinator immediately.

When the coordiantor receives an input that produces new coverage, it compares the worker's coverage to the current combined coverage array: it's possible that another worker already discovered an input that provides the same coverage. If so, the new input is discarded. If the new input actually does provide new coverage, the coordinator sends it back to a worker (perhaps a different worker) for minimization. Minimization is like fuzzing, but the worker performs random mutations to create a smaller input that still provides at least some new coverage. Smaller inputs tend to be faster, so it's worth spending the time to minimize up front to make the fuzzing process faster later. The worker process reports back when it's done minimizing, even if it failed to find anything smaller. The coordinator adds the minimized input to the cached corpus and continues. Later on, the coordinator may send the minimized input out to workers for further fuzzing. This is how the fuzzing system adapts to find new coverage.

When the coordinator receives an input that causes an error, it again sends the input back to workers for minimization. In this case, a worker attempts to find a smaller input that still causes an error, though not necessarily the same error. After the input is minimized, the coordinator saves it into `testdata/corpus/$FuzzTarget`, shuts worker processes down gracefully, then exits with a non-zero status.

![Diagram showing communication between coordinator and worker. Two arrows point down: the left is labelled "coordinator", the right is labelled "worker". Three pairs of horizontal arrows point from the coordinator to the worker and back. The top pair is labelled "baseline coverage", the middle is labelled "fuzz", the bottom is labelled "minimize".](https://jayconrod.com/images/fuzz-communication.svg)

If a worker process crashes while fuzzing, the coordinator can recover the input that caused the crash using the input sent to the worker, and the worker's RNG state and iteration count (both left in shared memory). Crashing inputs are generally not minimized, since minimization is a highly stateful process, and each crash blows that state away. It is [theoretically possible](https://github.com/golang/go/issues/48163) but hasn't been done yet.

Fuzzing usually continues until an error is discovered or the user interrupts the process by pressing Ctrl-C or the deadline set with the `-fuzztime` flag is passed. The fuzzing engine handles interrupts gracefully, whether they are delivered to the coordinator or worker processes. For example, if a worker is interrupted while minimizing an input that caused an error, the coordinator will save the unminimized input.

```bash
MacBook-Air:go-tutorial $ ps aux | grep fuzz
xxx    13913  84.3  1.0  5219184  85124 s001  R+   10:12下午   0:03.90 /var/folders/pv/_x849j6n22x37xxd9cstgwkr0000gn/T/go-build1953131131/b001/fuzz.test -test.fuzzworker -test.paniconexit0 -test.fuzzcachedir=/Users/xxx/Library/Caches/go-build/fuzz/example/fuzz -test.timeout=10m0s -test.fuzz=Fuzz
xxx    13910  81.9  1.0  5221180  86200 s001  R+   10:12下午   0:03.94 /var/folders/pv/_x849j6n22x37xxd9cstgwkr0000gn/T/go-build1953131131/b001/fuzz.test -test.fuzzworker -test.paniconexit0 -test.fuzzcachedir=/Users/xxx/Library/Caches/go-build/fuzz/example/fuzz -test.timeout=10m0s -test.fuzz=Fuzz
xxx    13912  78.3  1.0  5219964  84984 s001  R+   10:12下午   0:03.86 /var/folders/pv/_x849j6n22x37xxd9cstgwkr0000gn/T/go-build1953131131/b001/fuzz.test -test.fuzzworker -test.paniconexit0 -test.fuzzcachedir=/Users/xxx/Library/Caches/go-build/fuzz/example/fuzz -test.timeout=10m0s -test.fuzz=Fuzz
xxx    13911  74.5  1.0  5219184  85132 s001  R+   10:12下午   0:03.76 /var/folders/pv/_x849j6n22x37xxd9cstgwkr0000gn/T/go-build1953131131/b001/fuzz.test -test.fuzzworker -test.paniconexit0 -test.fuzzcachedir=/Users/xxx/Library/Caches/go-build/fuzz/example/fuzz -test.timeout=10m0s -test.fuzz=Fuzz
xxx    13907  43.3  2.3  5944576 191172 s001  R+   10:12下午   0:01.90 /var/folders/pv/_x849j6n22x37xxd9cstgwkr0000gn/T/go-build1953131131/b001/fuzz.test -test.paniconexit0 -test.fuzzcachedir=/Users/xxx/Library/Caches/go-build/fuzz/example/fuzz -test.timeout=10m0s -test.fuzz=Fuzz
xxx    13923   0.0  0.0  4268176    420 s000  R+   10:12下午   0:00.00 grep fuzz
xxx    13891   0.0  0.2  5014396  16868 s001  S+   10:12下午   0:00.52 /Users/xxx/sdk/go1.18beta2/bin/go test -fuzz=Fuzz
xxx    13890   0.0  0.0  4989312   4008 s001  S+   10:12下午   0:00.01 go1.18beta2 test -fuzz=Fuzz
```



## 注意事项

* FuzzXXX也是实现在xxx_test.go里
* go test 不带-fuzz 会默认执行 FuzzXXX，带-fuzz会怎么样
* seed corpus：包括f.Add里新增的，也包括Fuzzing找出来写到testdata下的





## 开源地址

文章和示例代码开源在GitHub: [Go语言初级、中级和高级教程](https://github.com/jincheng9/go-tutorial)。

公众号：coding进阶。关注公众号可以获取最新Go面试题和技术栈。

个人网站：[Jincheng's Blog](https://jincheng9.github.io/)。



## References

* Internals of Go's New Fuzzing System: https://jayconrod.com/posts/123/internals-of-go-s-new-fuzzing-system
* Fuzzing介绍：https://go.dev/doc/fuzz/
* Fuzzing Design Draft: https://go.googlesource.com/proposal/+/master/design/draft-fuzzing.md
* Fuzzing提案：https://github.com/golang/go/issues/44551
* Fuzzing教程：https://go.dev/doc/tutorial/fuzz
* tesing.F说明文档：https://pkg.go.dev/testing@go1.18#F
* Fuzzing Tesing in Go in 8 Minutes: https://www.youtube.com/watch?v=w8STTZWdG9Y
* GitHub开源工具go-fuzz: https://github.com/dvyukov/go-fuz
* Go fuzzing找bug示例：https://julien.ponge.org/blog/playing-with-test-fuzzing-in-go/

