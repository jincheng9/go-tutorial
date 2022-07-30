# Go十大常见错误第3篇：Go指针的性能问题和内存逃逸

## 前言

这是Go十大常见错误系列的第3篇：Go指针的性能问题和内存逃逸。素材来源于Go布道者，现Docker公司资深工程师[Teiva Harsanyi](https://teivah.medium.com/)。

本文涉及的源代码全部开源在：[Go十大常见错误源代码](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p28)，欢迎大家关注公众号，及时获取本系列最新更新。



## 场景

我们知道，函数参数和返回值可以使用变量或者指向变量的指针。

Go初学者容易有一种误解：

* 认为函数参数和返回值如果使用变量的值会对整个变量做拷贝，速度慢
* 认为函数参数和返回值如果使用指针类型，只需要拷贝内存地址，速度更快

但事实真的是这样么？我们可以看下这段代码 [this example](https://github.com/jincheng9/go-tutorial/blob/main/workspace/senior/p28/pointer/pointer_test.go)，做性能测试的结果如下：

```bash
$ go test -bench .
goos: darwin
goarch: amd64
pkg: pointer
cpu: Intel(R) Core(TM) i5-5250U CPU @ 1.60GHz
BenchmarkByPointer-4     6473781               178.2 ns/op
BenchmarkByValue-4      21760696                47.11 ns/op
PASS
ok      pointer 2.894s
```

可以看出，参数和返回值都用指针的函数比参数和返回值都用变量值的函数慢很多，前者的耗时是后者的4倍。

初学者看到这个，可能会觉得有点反直觉，为什么会这样呢？

这和Go对stack(栈)和heap(堆)的内存管理有关系，变量分配在stack上还是heap上，对性能是会有影响的。

* stack上分配内存效率比heap更高，而且stack上分配的内存不用做GC，超出了作用域，就自动回收内存。

* 放在heap上的内存，需要由GC来做内存回收，而且容易产生内存碎片。

* 编译器在编译期决定变量分配在stack还是heap上，需要做逃逸分析(escape analysis)，逃逸分析在编译阶段就完成了。

什么是逃逸分析呢？

Go编译器解析源代码，决定哪些变量分配在stack内存空间，哪些变量分配在heap内存空间的过程就叫做逃逸分析，属于Go代码编译过程中的一个分析环节。

通过逃逸分析，编译器会尽可能把能分配在栈上的对象分配在栈上，避免堆内存频繁GC垃圾回收带来的系统开销，影响程序性能(只有heap内存空间才会发生GC)。



### 案例1

我们看下面的代码：其中结构体`foo`的可以参考 [this example](https://github.com/jincheng9/go-tutorial/blob/main/workspace/senior/p28/pointer/pointer_test.go)。

```go
func getFooValue() foo {
	var result foo
	// Do something
	return result
}
```

变量`result`定义的时候会在这个goroutine的stack上分配`result`的内存空间。

当函数返回时，`getFooValue`的调用方如果有接收返回值，那`result`的值会被拷贝给对应的接收变量。

stack上变量`result`的内存空间会被释放(标记为不可用，不能再被访问，除非这块空间再次被分配给其它变量)。

**注意**：本案例的结构体`foo`占用的内存空间比较小，约0.3KB，goroutine的stack空间足够存储，如果`foo`占用的空间过大，在stack里存储不了，就会分配内存到heap上。



### 案例2

我们看下面的代码：

```go
func getFooPointer() *foo {
	var result foo
	// Do something
	return &result
}
```

函数`getFooPointer`因为返回的是一个指针，如果变量`result`分配在stack上，那函数返回后，`result`的内存空间会被释放，就会导致接受函数返回值的变量无法访问原本`result`的内存空间，成为一个悬浮指针(dangling pointer)。

所以这种情况会发生内存逃逸，`result`会分配在heap上，而不是stack上。



### 案例3

我们看下面的代码：

```go
func main()  {
	p := &foo{}
	f(p)
}
```

指针变量`p`是函数`f`的实参，因为我们是在main所在的goroutine里调用函数`f`，并没有跨goroutine，所以指针变量`p`分配在stack上就可以，不需要分配在heap上。



## 总结

那我们怎么知道到底变量是分配在stack上还是head上呢？

Go官方给的说法是：

* 从程序正确性的角度而言，你不需要关心变量是分配在stack上还是heap上。变量分配在哪块内存空间不改变Go语言的语义。
* 从程序性能的角度而言，你可以关心变量到底是分配在stack上还是heap上，因为正如上文所言，变量存储的位置是对性能有影响的。

> ### How do I know whether a variable is allocated on the heap or the stack?
>
> From a correctness standpoint, you don't need to know. Each variable in Go exists as long as there are references to it. The storage location chosen by the implementation is irrelevant to the semantics of the language.
>
> The storage location does have an effect on writing efficient programs. When possible, the Go compilers will allocate variables that are local to a function in that function's stack frame. However, if the compiler cannot prove that the variable is not referenced after the function returns, then the compiler must allocate the variable on the garbage-collected heap to avoid dangling pointer errors. Also, if a local variable is very large, it might make more sense to store it on the heap rather than the stack.
>
> In the current compilers, if a variable has its address taken, that variable is a candidate for allocation on the heap. However, a basic *escape analysis* recognizes some cases when such variables will not live past the return from the function and can reside on the stack.

一般而言，遇到以下情况会发生逃逸行为，Go编译器会将变量存储在heap上

* 函数内局部变量在函数外部被引用
* 接口(interface)类型的变量
* size未知或者动态变化的变量，如slice，map，channel，[]byte等
* size过大的局部变量，因为stack内存空间比较小。

此外，我们还可以借助内存逃逸分析工具来帮助我们。

因为内存逃逸分析是编译器在编译期就完成的，可以使用以编译下命令来做内存逃逸分析：

* `go build -gcflags="-m"`，可以展示逃逸分析、内联优化等各种优化结果。
* `go build -gcflags="-m -l"`，`-l`会禁用内联优化，这样可以过滤掉内联优化的结果展示，让我们可以关注逃逸分析的结果。
* `go build -gcflags="-m -m"`，多一个`-m`会展示更详细的分析结果。



## 推荐阅读

* [Go十大常见错误第1篇：未知枚举值](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484146&idx=1&sn=10fb12b643a2e37c090e5aa3bc583152&chksm=ce124d9df965c48bb954aeddabdff3db12738ded3875542250c5d0ef6cfd4417fc56580288b1&token=1912894792&lang=zh_CN#rd)
* [Go十大常见错误第2篇：benchmark性能测试的坑](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484163&idx=1&sn=b28d61c1f3ec9d914e698dce105ba5d1&chksm=ce124c6cf965c57a90bc85a5295ed9375103de20607b509f845583ff6686385df0ed96653d00&token=1912894792&lang=zh_CN#rd)
* [Go栈和指针的语法机制](https://www.ardanlabs.com/blog/2017/05/language-mechanics-on-stacks-and-pointers.html)
* [逃逸分析原理 by ArdanLabs](https://www.ardanlabs.com/blog/2017/05/language-mechanics-on-escape-analysis.html)
* [逃逸分析原理 by Gopher Con](https://www.youtube.com/watch?v=ZMZpH4yT7M0)



## 开源地址

文章和示例代码开源在GitHub: [Go语言初级、中级和高级教程](https://github.com/jincheng9/go-tutorial)。

公众号：coding进阶。关注公众号可以获取最新Go面试题和技术栈。

个人网站：[Jincheng's Blog](https://jincheng9.github.io/)。

知乎：[无忌](https://www.zhihu.com/people/thucuhkwuji)。



## 福利

我为大家整理了一份后端开发学习资料礼包，包含编程语言入门到进阶知识(Go、C++、Python)、后端开发技术栈、面试题等。

关注公众号「coding进阶」，发送消息 **backend** 领取资料礼包，这份资料会不定期更新，加入我觉得有价值的资料。还可以发送消息「**进群**」，和同行一起交流学习，答疑解惑。



## References

* https://itnext.io/the-top-10-most-common-mistakes-ive-seen-in-go-projects-4b79d4f6cd65
* https://www.ardanlabs.com/blog/2017/05/language-mechanics-on-stacks-and-pointers.html
* https://www.ardanlabs.com/blog/2017/05/language-mechanics-on-escape-analysis.html
* https://www.youtube.com/watch?v=ZMZpH4yT7M0