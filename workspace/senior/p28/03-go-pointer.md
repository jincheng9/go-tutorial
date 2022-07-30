# Go指针的性能问题和内存逃逸

## 前言

这是Go十大常见错误系列的第3篇：Go指针的性能问题和内存逃逸。素材来源于Go布道者，现Docker公司资深工程师[Teiva Harsanyi](https://teivah.medium.com/)。

本文涉及的源代码全部开源在：[Go十大常见错误源代码](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p28)，欢迎大家关注公众号，及时获取本系列最新更新。



## 场景

我们知道，函数参数和返回值可以使用变量或者指向变量的指针。

Go初学者容易有一种误解：

* 认为函数参数和返回值如果使用变量的值会对整个变量做拷贝，速度慢
* 认为函数参数和返回值如果使用指针类型，只需要拷贝内存地址，速度更快

但事实真的是这样么？

我们可以看下这段代码 [this example](https://github.com/jincheng9/go-tutorial/blob/main/workspace/senior/p28/pointer/pointer_test.go)，做性能测试的结果如下：

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

* 编译器决定变量分配在stack还是heap上，需要做逃逸分析(escape analysis)，逃逸分析在编译阶段就完成了。

通过逃逸分析，编译器会尽可能把能分配在栈上的对象分配在栈上，避免频繁GC垃圾回收带来的系统开销，对程序性能有影响。



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



## 总结

那我们怎么知道到底变量是分配在stack上还是head上呢？通常有以下原则，



此外，我们还可以借助内存逃逸分析工具来帮助我们。

因为内存逃逸分析是编译器在编译期就完成的，可以使用以编译下命令来做内存逃逸分析：

* `go build -gcflags="-m"`
* `go build -gcflags="-m -l"`，`-l`会禁用内联优化。
* `go build -gcflags="-m -m"`，多一个`-m`会展示更详细的分析结果。



A variable can be allocated on the **heap** or the **stack**. As a rough draft:

- The stack contains the **ongoing** variables for a given **goroutine**. Once a function returned, the variables are popped from the stack.
- The heap contains the **shared** variables (global variables, etc.).

Let’s check at a simple example where we return a value:

<iframe src="https://itnext.io/media/e449d978f99c08184cd5a202cdf2a1d4" allowfullscreen="" frameborder="0" height="153" width="692" title="value.go" class="fp aq as ag cf" scrolling="auto" style="box-sizing: inherit; height: 153px; top: 0px; left: 0px; width: 692px; position: absolute;"></iframe>

Here, a `result` variable is created by the current goroutine. This variable is pushed into the current stack. Once the function returns, the client will receive a copy of this variable. The variable itself is popped from the stack. It still exists in memory until it is erased by another variable but it **cannot be accessed anymore**.

Now, the same example but with a pointer:

<iframe src="https://itnext.io/media/c509bec83765998aee524baeaa8b6741" allowfullscreen="" frameborder="0" height="153" width="692" title="pointer.go" class="fp aq as ag cf" scrolling="auto" style="box-sizing: inherit; height: 153px; top: 0px; left: 0px; width: 692px; position: absolute;"></iframe>

The `result` variable is still created by the current goroutine but the client will receive a pointer (a copy of the variable address). If the `result` variable was popped from the stack, the client of this function **could not access** it anymore.

In this scenario, the Go compiler will **escape** the `result` variable to a place where the variables can be shared: **the heap**.

Passing pointers, though, is another scenario. For example:

<iframe src="https://itnext.io/media/6e7878703f82261e74637f928cb92234" allowfullscreen="" frameborder="0" height="131" width="692" title="pointer.go" class="fp aq as ag cf" scrolling="auto" style="box-sizing: inherit; height: 131px; top: 0px; left: 0px; width: 692px; position: absolute;"></iframe>

Because we are calling `f` within the same goroutine, the `p` variable does not need to be escaped. It is simply pushed to the stack and the sub-function can access it.

For example, this is a direct consequence of receiving a slice in the `Read` method of `io.Reader` instead of returning one. Returning a slice (which is a pointer) would have escaped it to the heap.

Why is the stack so **fast** then? There are two main reasons:

- There is no need to have a **garbage collector** for the stack. As we said, a variable is simply pushed once it is created then popped from the stack once the function returns. There is no need to get a complex process reclaiming unused variables, etc.
- A stack belongs to one goroutine so storing a variable does not need to be **synchronized** compared to storing it on the heap. This also results in a performance gain.

As a conclusion, when we create a function, our default behavior should be to use **values instead of pointers**. A pointer should only be used if we want to **share** a variable.

Then, if we suffer from performance issues, one possible optimization could be to check whether pointers would help or not in some specific situations. It is possible to know when the compiler will escape a variable to the heap by using the following command: `go build -gcflags "-m -m"`.

But again, for most of our day-to-day use cases, values are the best fit.

## 推荐阅读



## 开源地址

文章和示例代码开源在GitHub: [Go语言初级、中级和高级教程](https://github.com/jincheng9/go-tutorial)。

公众号：coding进阶。关注公众号可以获取最新Go面试题和技术栈。

个人网站：[Jincheng's Blog](https://jincheng9.github.io/)。

知乎：[无忌](https://www.zhihu.com/people/thucuhkwuji)。



## References

* https://itnext.io/the-top-10-most-common-mistakes-ive-seen-in-go-projects-4b79d4f6cd65
* https://www.ardanlabs.com/blog/2017/05/language-mechanics-on-stacks-and-pointers.html
* https://www.ardanlabs.com/blog/2017/05/language-mechanics-on-escape-analysis.html
* https://www.youtube.com/watch?v=ZMZpH4yT7M0