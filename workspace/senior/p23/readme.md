# Fuzzing: 一文读懂Go Fuzzing使用和原理

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

**注意**：如果Go Fuzzing运行过程中发现了你的bug，会把对应的输入数据写到`testdata/fuzz/FuzzXXX`目录下。比如上面的例子里，`go1.18beta1 test -v -fuzz=Fuzz`的输出结果里打印了如下内容：`Failing input written to testdata/fuzz/FuzzReverse/ce9e8c80e2c2de2c96ab9e63b1a8cf18cea932b7d8c6c9c207d5978e0f19027a`，这就表示把这个测试输入写到了`testdata/fuzz/FuzzReverse/xxx`这个语料文件里。



## Go Fuzzing的底层机制

 `go test` 执行的时候，会为每个被测试的package先编译生成一个可执行文件，然后运行这个可执行文件得到对应package的`TestXXX`和`BenchmarkXXX`的测试结果。Go Fuzzing运行的模式和这个类似，但是也有一点区别。

当`go test`执行的时候如果有`-fuzz`标记，`go test`会结合覆盖率工具来编译生成用于模糊测试的可执行文件。大部分的Fuzzing逻辑都实现在[internal/fuzz](https://pkg.go.dev/internal/fuzz)。

当`go test`编译生成了可执行文件后，该可执行文件就会运行起来，这个运行起来的进程叫做协调进程(coordinator process)。协调进程的启动参数里有`go test`命令的大部分标记，包括`-fuzz=pattern`这个标记，`-fuzz=pattern`用来识别对哪个模糊测试函数(fuzz test)进行Fuzzing测试。

目前，对于每一个`go test -fuzz=pattern`调用，只支持匹配一个模糊测试函数。如果`go test -fuzz=pattern`可以匹配多个`FuzzXXX`函数，就会报如下错误：

```sh
$ go1.18beta1 test -v -fuzz=Fuzz
testing: will not fuzz, -fuzz matches more than one fuzz test: [FuzzReverse FuzzReverse2]
FAIL
exit status 1
FAIL    example/fuzz    0.752s
```

协调进程启动后，主要的程序逻辑都在[`fuzz.CoordinateFuzzing`](https://pkg.go.dev/internal/fuzz#CoordinateFuzzing)。[`fuzz.CoordinateFuzzing`](https://pkg.go.dev/internal/fuzz#CoordinateFuzzing)会初始化fuzzing系统，开启coordinator事件循环。

coordinator进程会启动多个worker进程，每个worker进程和coordinator进程运行相同的可执行程序，真正的fuzzing模糊测试由worker进程来完成。worker进程启动时带有一个标记参数`-test.fuzzworker`，表明这是一个worker进程。启动的worker进程数量等于GOMAXPROCS。

这里我给了一个示例，大家可以在执行`go test -fuzz=pattern`的过程中，运行`ps aux | grep fuzz`来查看当前fuzzing相关的进程。

```go
$ ps aux | grep fuzz
xxx    13913  84.3  1.0  5219184  85124 s001  R+   10:12下午   0:03.90 /var/folders/pv/_x849j6n22x37xxd9cstgwkr0000gn/T/go-build1953131131/b001/fuzz.test -test.fuzzworker -test.paniconexit0 -test.fuzzcachedir=/Users/xxx/Library/Caches/go-build/fuzz/example/fuzz -test.timeout=10m0s -test.fuzz=Fuzz
xxx    13910  81.9  1.0  5221180  86200 s001  R+   10:12下午   0:03.94 /var/folders/pv/_x849j6n22x37xxd9cstgwkr0000gn/T/go-build1953131131/b001/fuzz.test -test.fuzzworker -test.paniconexit0 -test.fuzzcachedir=/Users/xxx/Library/Caches/go-build/fuzz/example/fuzz -test.timeout=10m0s -test.fuzz=Fuzz
xxx    13912  78.3  1.0  5219964  84984 s001  R+   10:12下午   0:03.86 /var/folders/pv/_x849j6n22x37xxd9cstgwkr0000gn/T/go-build1953131131/b001/fuzz.test -test.fuzzworker -test.paniconexit0 -test.fuzzcachedir=/Users/xxx/Library/Caches/go-build/fuzz/example/fuzz -test.timeout=10m0s -test.fuzz=Fuzz
xxx    13911  74.5  1.0  5219184  85132 s001  R+   10:12下午   0:03.76 /var/folders/pv/_x849j6n22x37xxd9cstgwkr0000gn/T/go-build1953131131/b001/fuzz.test -test.fuzzworker -test.paniconexit0 -test.fuzzcachedir=/Users/xxx/Library/Caches/go-build/fuzz/example/fuzz -test.timeout=10m0s -test.fuzz=Fuzz
xxx    13907  43.3  2.3  5944576 191172 s001  R+   10:12下午   0:01.90 /var/folders/pv/_x849j6n22x37xxd9cstgwkr0000gn/T/go-build1953131131/b001/fuzz.test -test.paniconexit0 -test.fuzzcachedir=/Users/xxx/Library/Caches/go-build/fuzz/example/fuzz -test.timeout=10m0s -test.fuzz=Fuzz
xxx    13923   0.0  0.0  4268176    420 s000  R+   10:12下午   0:00.00 grep fuzz
xxx    13891   0.0  0.2  5014396  16868 s001  S+   10:12下午   0:00.52 /Users/xxx/sdk/go1.18beta2/bin/go test -fuzz=Fuzz
xxx    13890   0.0  0.0  4989312   4008 s001  S+   10:12下午   0:00.01 go1.18beta2 test -fuzz=Fuzz
```

worker进程在运行模糊测试(fuzzing)的时候如果crash了，coordinator进程可以记录导致worker进程crash的测试数据。如果直接交给coordinator进程执行fuzzing，在遇到了会导致程序crash的输入时，coordinator进程本身就会crash，就没有办法记录导致程序crash的输入了(Failing input)。Go Fuzzing运行的模型如下所示：



![Diagram showing the relationship between fuzzing processes. At the top is a box showing "go test (cmd/go)". An arrow points downward to a box labelled "coordinator (test binary)". From that, three arrows point downward to three boxes labelled "worker (test binary)".](https://jayconrod.com/images/fuzz-processes.svg)

coordinator进程和worker进程通过一对管道进行通信，使用基于JSON的RPC通信协议。这个协议非常精简，因为我们并不需要gRPC一样复杂的RPC协议，我们也不希望给Go标准库引入任何新的依赖。

每个worker进程在mmap文件里保存自己的状态，这个mmap文件和coordinator进程共享。大多数情况下，mmap里记录的只是迭代次数和随机数生成器的状态。如果worker进程crash了，那coordinator进程就可以从共享内存里恢复其状态，而不需要worker进程通过管道发送消息。

整个Fuzzing过程分为3个阶段：

![Diagram showing communication between coordinator and worker. Two arrows point down: the left is labelled "coordinator", the right is labelled "worker". Three pairs of horizontal arrows point from the coordinator to the worker and back. The top pair is labelled "baseline coverage", the middle is labelled "fuzz", the bottom is labelled "minimize".](https://jayconrod.com/images/fuzz-communication.svg)

### 阶段1：Baseline coverage

coordinator进程启动时，会拉起worker进程。coordinator进程会给worker进程发送种子语料(包括`f.Add`里添加的测试数据以及`testdata/fuzz`目录下的测试输入)和fuzzing缓存语料(cache corpus，位于`$GOCACHE`的子目录下)。

每个worker进程运行指定的输入，然后给coordinator进程报告其覆盖率计数器的快照，coordinator会将收集到的worker的覆盖率数据合并为一个覆盖率数组。

这个阶段叫基线覆盖率收集阶段，worker只会运行coordinator发送给它们的指定输入，不会生成随机测试数据。



### 阶段2：Fuzzing模糊测试

这个阶段，coordinator进程会再次发送种子语料(seed corpus)和缓存语料(cache corpus)给worker进程，用于真正的fuzzing。

每个worker进程会收到一个coordinator发送的输入数据和基线覆盖率数组的拷贝。然后worker进程会随机对这个指定的输入做变异来得到新的测试数据。变异的方式有多种，可能是对bit位做反转，0改为1，1改为0，也可能是删除或者新增字节，等等。然后再把变异后的数据作为参数给到fuzz target函数去运行。

为了减少coordinator进程和worker进程的通信开销，每个worker进程可以在100ms内一直变异拿到新的测试数据，然后调用fuzz target函数，而不需要coordinator进程的进一步输入。

每次对生成的随机数据调用fuzz target函数后，worker进程会检查2种场景：

* 和基线覆盖率数组相比，是否找到了新的覆盖率数据。
* 是否有error产生，也就是代码里执行了[`T.Fail`](https://pkg.go.dev/testing@go1.18beta2#T.Fail)或[`T.FailNow`](https://pkg.go.dev/testing@go1.18beta2#T.FailNow)。**注意**：`T.Error`、`T.Errorf`会自动调用`T.Fail`,`T.Fatal`和`T.Fatalf`会自动调用`T.FailNow`。

如果二者满足其一，worker进程就会把输入数据立即发送给coordinator进程。



### 阶段3：Minimization最小化

如果coordinator进程收到了worker进程发送过来的输入数据是场景1，也就是收到了会产生新覆盖率的输入，coordinator会把这个worker的覆盖率数据和当前组合的覆盖率数组做比较。

因为有可能其它worker已经发现了会提供相同覆盖率的输入，如果是这样的话，那coordinator会直接ignore这个输入。如果这个新的输入的确提供了新的覆盖率，那coordinator会把这个输入发送给一个worker(很可能是不同的worker)用于最小化(minimization)。

最小化有点像fuzzing，但是worker会通过随机变异来创建一个仍然会产生新覆盖率的更小输入。更小的输入通常会让fuzzing执行更快，因此minimization是有必要的。worker进程完成最小化后会报告给coordinator最小化的输入，即使它未能找到更小的输入。coordinator进程会把这个最小化的输入添加到缓存语料库(cache corpus)并继续执行Fuzzing。后续，coordinator可能会把这个最小化的输入发送给所有worker用于进一步fuzzing。这就是fuzzing系统如何自动调节找到新的覆盖率。

如果coordinator进程收到了worker进程发送过来的输入数据是场景2：也就是`引发error的输入`，coordinator进程会把这个输入再次发送给worker进行最小化。在这种场景下，worker会试图找到一个会引发error的更小输入，尽管不一定是同一个error。在输入数据被最小化后，coordinator进程会把最小化后的数据存储到`testdata/fuzz/$FuzzTarget`，优雅关闭所有worker进程，然后以非0状态(non-zero status)退出。

如果worker进程在fuzzing过程中crash了，那coordinator进程可以使用发送给worker的输入、worker的RNG状态和迭代次数(留在共享内存中)来恢复导致worker进程crash的输入。crash的输入通常没有被最小化，因为最小化是一个高度状态化的过程，而每次crash都会对状态进行破坏。对导致crash的输入做最小化在[理论](https://github.com/golang/go/issues/48163)上是可行的，但是目前还没能实现。

Fuzzing通常遇到以下场景才会结束运行，否则会一直运行：

* Fuzzing找到了error，也就是触发了你模糊测试函数里的error条件
* 用户按Ctrl-C来中断程序
* 运行时间达到了`-fuzztime`设定的时间

fuzzing引擎会优雅处理中断，不管中断是被发送给了coordinator进程还是worker进程。举个例子，如果worker进程在最小化输入的时候遇到了中断，coordinator进程会保存没有被最小化的输入。



## 注意事项

* `FuzzXXX`的实现也是放在以`_test.go`结尾的go文件里。
* seed corpus(种子语料)：既包含通过`f.Add`指定的输入，也包括`testdata/fuzz/$FuzzTarget`目录下的文件里面的输入。
* `go test` 不带`-fuzz`标记会默认执行`TestXXX`和`FuzzXXX`开头的函数，对于`FuzzXXX`只会使用种子语料库里的输入，而不会生成随机数据。如果需要生成随机输入，要使用`go test -fuzz=pattern`。



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
* 专注于Fuzzing技术的博客网站：https://blog.fuzzbuzz.io/

