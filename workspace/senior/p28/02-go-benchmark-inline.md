# Go十大常见错误第2篇：benchmark性能测试的坑

## 前言

这是Go十大常见错误系列的第二篇：benchmark性能测试的坑。素材来源于Go布道者，现Docker公司资深工程师[**Teiva Harsanyi**](https://teivah.medium.com/)。

本文涉及的源代码全部开源在：[Go十大常见错误源代码](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p28)，欢迎大家关注公众号，及时获取本系列最新更新。



## 场景

`go test`支持benchmark性能测试，但是你知道这里可能有坑么？

一个常见的坑是编译器优化，我们来看一个具体的例子：

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

这里可能有什么坑呢？对于编译器而言，`add`函数是一个叶子函数(leaf function)，即`add`函数本身没有调用其它函数，所以编译器会对`add`函数的调用做内联(inline)优化，这会导致性能测试的结果不准确。



## 最佳实践



```go
var result int

func BenchmarkCorrect(b *testing.B) {
	b.ResetTimer()
	var r int
	for i := 0; i < b.N; i++ {
		r = add(1000000000, 1000000001)
	}
	result = r
}
```



我们对比看看性能测试的结果

```bash
$ go test -v -bench . -count=3
goos: darwin
goarch: amd64
pkg: example.com/benchmark
cpu: Intel(R) Core(TM) i5-5250U CPU @ 1.60GHz
BenchmarkWrong
BenchmarkWrong-4        1000000000               0.4009 ns/op
BenchmarkWrong-4        1000000000               0.4139 ns/op
BenchmarkWrong-4        1000000000               0.4065 ns/op
BenchmarkCorrect
BenchmarkCorrect-4      1000000000               0.5590 ns/op
BenchmarkCorrect-4      1000000000               0.5500 ns/op
BenchmarkCorrect-4      1000000000               0.5654 ns/op
PASS
ok      example.com/benchmark   3.320s
```

源码地址：[benchmark性能测试源代码](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p28/benchmark)。

## 

## 开源地址

文章和示例代码开源在GitHub: [Go语言初级、中级和高级教程](https://github.com/jincheng9/go-tutorial)。

公众号：coding进阶。关注公众号可以获取最新Go面试题和技术栈。

个人网站：[Jincheng's Blog](https://jincheng9.github.io/)。

知乎：[无忌](https://www.zhihu.com/people/thucuhkwuji)。



## References

* https://itnext.io/the-top-10-most-common-mistakes-ive-seen-in-go-projects-4b79d4f6cd65