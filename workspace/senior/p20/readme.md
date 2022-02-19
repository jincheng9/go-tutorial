# Go Quiz: 从Go面试题看函数命名返回值的注意事项

## 题目

Redhat的首席工程师、Prometheus开源项目Maintainer [Bartłomiej Płotka](https://twitter.com/bwplotka) 在Twitter上出了一道Go编程题，结果超过80%的人都回答错了。

题目如下所示，回答下面这段程序的输出结果。

```go
// named_return.go
package main

import "fmt"

func aaa() (done func(), err error) {
	return func() { print("aaa: done") }, nil
}

func bbb() (done func(), _ error) {
	done, err := aaa()
	return func() { print("bbb: surprise!"); done() }, err
}

func main() {
	done, _ := bbb()
	done()
}
```

* A: `bbb: surprise!`
* B: `bbb: surprise!aaa: done`
* C: 编译报错
* D: 陷入死循环

大家可以先思考下这段代码的输出结果是什么。



## 解析

这里其实涉及Go的2个知识点

1. 命名返回值闭包调用自己陷入递归死循环

所以本题的答案是`D` ，会陷入死循环，不断打印`bbb: surprise!`。



##  总结

这个题目其实很tricky，在实际编程中，要避免对命名返回值采用这种写法，非常容易出错。

想了解国外Go开发者对这个题目的讨论详情可以参考[Go Named Return Parameters Discussion](https://twitter.com/bwplotka/status/1494362886738780165)。

另外题目作者也给了如下所示的[详细解释](https://go.dev/play/p/ELPEi2AK0DP)：

```go
package main

func aaa() (done func(), err error) {
	return func() { print("aaa: done") }, nil
}

func bbb() (done func(), _ error) {
	// NOTE(bwplotka): Here is the problem. We already defined special "return argument" variable called "done".
	// By using `:=` and not `=` we define a totally new variable with the same name in
	// new, local function scope.
	done, err := aaa()

	// NOTE(bwplotka): In this closure (anonymous function), we might think we use `done` from the local scope,
	// but we don't! This is because Go "return" as a side effect ASSIGNS returned values to
	// our special "return arguments". If they are named, this means that after return we can refer
	// to those values with those names during any execution after the main body of function finishes
	// (e.g in defer or closures we created).
	//
	// What is happening here is that no matter what we do in the local "done" variable, the special "return named"
	// variable `done` will get assigned with whatever was returned. Which in bbb case is this closure with
	// "bbb:surprise" print. This means that anyone who runs this closure AFTER `return` did the assignment
	// will start infinite recursive execution.
	//
	// Note that it's a feature, not a bug. We use this often to capture
	// errors (e.g https://github.com/efficientgo/tools/blob/main/core/pkg/errcapture/doc.go)
	//
	// Go compiler actually detects that `done` variable defined above is NOT USED. But we also have `err`
	// variable which is actually used. This makes compiler to satisfy that unused variable check,
	// which is wrong in this context..
	return func() { print("bbb: surprise!"); done() }, err
}

func main() {
	done, _ := bbb()
	done()
}
```

不过这个解释是有瑕疵的，主要是这句描述：

> By using `:=` and not `=` we define a totally new variable with the same name in
>  new, local function scope.

对于`done, err := aaa()`，变量`done`并不是一个新的变量，而是和函数`bbb`的返回变量`done`是同一个变量。

**本人把这个错误反馈给了原作者，原作者同意了我的意见，删除了这块解释**。

![](../../../workspace/img/named_return_paramater.png) 

最新版的英文解释可以参考[修正版解释](https://go.dev/play/p/9J5a3ZtIPnL)。



## 思考题

下面这段代码同样使用了命名返回值，大家可以看看这个道题的输出结果是什么。可以发送消息`nrv`获取答案。

```go
package main

func bar() (r int) {
	defer func() {
		r += 4
		if recover() != nil {
			r += 8
		}
	}()
	
	var f func()
	defer f()
	f = func() {
		r += 2
	}

	return 1
}

func main() {
	println(bar())
}
```



## 开源地址

文章和示例代码开源在GitHub: [Go语言初级、中级和高级教程](https://github.com/jincheng9/go-tutorial)。

公众号：coding进阶。关注公众号可以获取最新Go面试题和技术栈。

个人网站：[Jincheng's Blog](https://jincheng9.github.io/)。

知乎：[无忌](https://www.zhihu.com/people/thucuhkwuji)。



## References

* https://twitter.com/bwplotka/status/1494362886738780165
* https://go.dev/play/p/ELPEi2AK0DP
* https://go.dev/play/p/9J5a3ZtIPnL