# Go指针的性能问题

## 前言

这是Go十大常见错误系列的第3篇：Go指针的性能问题。素材来源于Go布道者，现Docker公司资深工程师[Teiva Harsanyi](https://teivah.medium.com/)。

本文涉及的源代码全部开源在：[Go十大常见错误源代码](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p28)，欢迎大家关注公众号，及时获取本系列最新更新。



## 场景

如果我们有变量要传参，可以传该变量的值或者传指向该变量的指针。

Go初学者容易有一种误解， 认为传变量的值会对整个变量做拷贝，而传指针只需要拷贝内存地址，所以传指针速度更快。

但事实真的是这样么？

If you believe this, please take a look at [this example](https://gist.github.com/teivah/a32a8e9039314a48f03538f3f9535537). This is a benchmark on a 0.3 KB data structure that we pass and receive by pointer and then by value. 0.3 KB is not huge but that should not be far from the type of data structures we see every day (for most of us).

When I execute these benchmarks on my local environment, passing by value is more than **4 times faster** than passing by pointer. This might a bit counterintuitive, right?

The explanation of this result is related to how the memory is managed in Go. I couldn’t explain it as brilliantly as [William Kennedy](https://www.ardanlabs.com/blog/2017/05/language-mechanics-on-stacks-and-pointers.html) but let’s try to summarize it.

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

## 开源地址

文章和示例代码开源在GitHub: [Go语言初级、中级和高级教程](https://github.com/jincheng9/go-tutorial)。

公众号：coding进阶。关注公众号可以获取最新Go面试题和技术栈。

个人网站：[Jincheng's Blog](https://jincheng9.github.io/)。

知乎：[无忌](https://www.zhihu.com/people/thucuhkwuji)。



## References

* https://itnext.io/the-top-10-most-common-mistakes-ive-seen-in-go-projects-4b79d4f6cd65
* https://www.ardanlabs.com/blog/2017/05/language-mechanics-on-stacks-and-pointers.html
* https://www.youtube.com/watch?v=ZMZpH4yT7M0