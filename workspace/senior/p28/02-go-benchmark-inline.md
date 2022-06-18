# Go十大常见错误第2篇：benchmark性能测试的坑

## 前言

这是Go十大常见错误系列的第二篇：benchmark性能测试的坑。素材来源于Go布道者，现Docker公司资深工程师[Teiva Harsanyi](https://teivah.medium.com/)。

本文涉及的源代码全部开源在：[Go十大常见错误源代码](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p28)，欢迎大家关注公众号，及时获取本系列最新更新。



## 场景

`go test`支持benchmark性能测试，但是你知道这里可能有坑么？

一个常见的坑是编译器内联优化，我们来看一个具体的例子：

```go
func add(a int, b int) int {
	return a + b
}
```

现在我们要对`add`函数做性能测试，可能会编写如下测试代码：

```go
func BenchmarkWrong(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		add(1000000000, 1000000001)
	}
}
```

这里可能有什么坑呢？对于编译器而言，`add`函数是一个叶子函数(leaf function)，即`add`函数本身没有调用其它函数，所以编译器会对`add`函数的调用做内联(inline)优化，这会导致性能测试的结果不准确。**因为我们通常要测试的是自己程序本身的执行效率，而不是编译器做了优化后的执行效率，这样才方便我们对程序的性能有一个正确的认知，而且你做`go test`测试时编译器的优化效果和实际生产环境运行时编译器的优化效果可能也不一样**。



那怎么知道执行`go test`的时候编译器是否做了内联优化呢？很简单，给`go test`增加`-gcflags="-m"`参数，`-m`表示打印编译器做出的优化决定。

```bash
$ go test -gcflags="-m" -v -bench=BenchmarkWrong -count 1
# example.com/benchmark [example.com/benchmark.test]
./go_util.go:3:6: can inline add
./go_bench_test.go:19:6: inlining call to add
./go_bench_test.go:16:21: b does not escape
# example.com/benchmark.test
/var/folders/pv/_x849j6n22x37xxd9cstgwkr0000gn/T/go-build2365344599/b001/_testmain.go:33:6: can inline init.0
/var/folders/pv/_x849j6n22x37xxd9cstgwkr0000gn/T/go-build2365344599/b001/_testmain.go:41:24: inlining call to testing.MainStart
/var/folders/pv/_x849j6n22x37xxd9cstgwkr0000gn/T/go-build2365344599/b001/_testmain.go:41:42: testdeps.TestDeps{} escapes to heap
/var/folders/pv/_x849j6n22x37xxd9cstgwkr0000gn/T/go-build2365344599/b001/_testmain.go:41:24: &testing.M{...} escapes to heap
goos: darwin
goarch: amd64
pkg: example.com/benchmark
cpu: Intel(R) Core(TM) i5-5250U CPU @ 1.60GHz
BenchmarkWrong
BenchmarkWrong-4        1000000000               0.4601 ns/op
PASS
ok      example.com/benchmark   0.605s
```

上面的执行结果的`./go_bench_test.go:19:6: inlining call to add`就表示编译器对`BenchmarkWrong`里的`add`函数调用做了内联优化。

**备注**: `-gcflags` 的所有参数值可以执行`go tool compile --help`进行查看。



## 最佳实践

那在性能测试的时候怎么禁用编译期的内联优化呢？有2个方案：

### -gcflags="-l"

第一种方案，执行`go test`的时候，增加`-gcfloags="-l"`参数，`-l`表示禁用编译器的内联优化。

```go
$ go test -gcflags="-m -l" -v -bench=BenchmarkWrong -count 3
# example.com/benchmark [example.com/benchmark.test]
./go_bench_test.go:16:21: b does not escape
# example.com/benchmark.test
/var/folders/pv/_x849j6n22x37xxd9cstgwkr0000gn/T/go-build2785655381/b001/_testmain.go:41:42: testdeps.TestDeps{} escapes to heap
goos: darwin
goarch: amd64
pkg: example.com/benchmark
cpu: Intel(R) Core(TM) i5-5250U CPU @ 1.60GHz
BenchmarkWrong
BenchmarkWrong-4        476215998                2.447 ns/op
BenchmarkWrong-4        492860170                2.404 ns/op
BenchmarkWrong-4        483547294                2.388 ns/op
PASS
ok      example.com/benchmark   4.568s
```

通过上面的输出结果可以看出，并没有`inlining call`字样，这就证明了使用`-gcflags="-l"`参数后，编译器没有做内联优化了。

对比下编译期内联优化禁用前后的结果，性能差了将近5倍。

* 开启内联优化，耗时：0.4601 ns/op
* `-gcflags="-l"`关闭内联优化，耗时大概：2.4 ns/op



### go:noinline

第二种方案，使用`//go:noinline`编译器指令(compiler directive)，编译器在编译时会识别到这个指令，不做内联优化。

```go
//go:noinline
func add(a int, b int) int {
	return a + b
}
```

通过这种方式修改代码后，我们就不需要使用`-gcflags="-l"`参数了，我们来看看性能测试结果：

```bash
$ go test -gcflags="-m" -v -bench=BenchmarkWrong -count 3
# example.com/benchmark [example.com/benchmark.test]
./go_bench_test.go:16:21: b does not escape
# example.com/benchmark.test
/var/folders/pv/_x849j6n22x37xxd9cstgwkr0000gn/T/go-build1050705055/b001/_testmain.go:33:6: can inline init.0
/var/folders/pv/_x849j6n22x37xxd9cstgwkr0000gn/T/go-build1050705055/b001/_testmain.go:41:24: inlining call to testing.MainStart
/var/folders/pv/_x849j6n22x37xxd9cstgwkr0000gn/T/go-build1050705055/b001/_testmain.go:41:42: testdeps.TestDeps{} escapes to heap
/var/folders/pv/_x849j6n22x37xxd9cstgwkr0000gn/T/go-build1050705055/b001/_testmain.go:41:24: &testing.M{...} escapes to heap
goos: darwin
goarch: amd64
pkg: example.com/benchmark
cpu: Intel(R) Core(TM) i5-5250U CPU @ 1.60GHz
BenchmarkWrong
BenchmarkWrong-4        482026485                2.422 ns/op
BenchmarkWrong-4        495307399                2.413 ns/op
BenchmarkWrong-4        407674614                2.613 ns/op
PASS
ok      example.com/benchmark   4.439s
```

通过上面的输出结果，同样可以看出编译器没有做内联优化了，最终的执行效率和第一种方案基本一致。

测试源代码地址：[benchmark性能测试源代码](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p28/benchmark)，大家可以下载到本地进行测试。



**备注**: 网上有些文章的说法是把函数调用的结果赋值给一个局部变量，然后使用一个全局变量来承接这个局部变量的值就可以避免编译器的内联优化。这个说法实际上是错误的，原作者Teiva Harsanyi在这方面也犯了错误。要判断编译器是否做了内联优化，参考本文写的方式验证即可。



## 开源地址

文章和示例代码开源在GitHub: [Go语言初级、中级和高级教程](https://github.com/jincheng9/go-tutorial)。

公众号：coding进阶。关注公众号可以获取最新Go面试题和技术栈。

个人网站：[Jincheng's Blog](https://jincheng9.github.io/)。

知乎：[无忌](https://www.zhihu.com/people/thucuhkwuji)。



## References

* https://itnext.io/the-top-10-most-common-mistakes-ive-seen-in-go-projects-4b79d4f6cd65

* https://codeantenna.com/a/xxYYUev2YI#BenchMarking_60

* gcflag参数说明：https://pkg.go.dev/cmd/compile

* https://dave.cheney.net/2018/01/08/gos-hidden-pragmas

  