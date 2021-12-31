# Go Quiz: 从Go面试题看defer的注意事项第2篇

## 面试题

这是Go Quiz系列的第5篇，是考察Go语言的`defer`语义，也是`defer`语义的第2篇。

没有看过`defer`第1篇的可以先回顾下：[Go defer语义第1篇](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483756&idx=1&sn=d536fa3340e1d5f91d72eaa8b67c8123&chksm=ce124e03f965c715e26f5943948e17d8e0ebb3c4a3a180a149219a610f83fc6eb77b3b166b6a&token=1521159887&lang=zh_CN#rd)。

本文的题目如下：

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

- A: 1
- B: 7
- C: 12
- D: 13
- E: 15

这道题相比于`defer`第1篇的题目，主要考察以下知识点：

* 被`defer`的函数的值是什么时候确定的？**注意**：这里说的是函数值，不是函数的参数。
* 如果函数的值是`nil`，那`defer`一个nil函数是什么结果？
* 多个被`defer`的函数的执行先后顺序，遵循LIFO原则。
* `defer` 和`recover`结合可以捕获`panic`。
* `defer`如果对函数的命名返回值参数做修改，有什么影响？

## 解析

### 函数值

上面提到了函数值的概念，对应的英文是`function value`。可能刚学习Go的同学还不太了解，我先讲解下这个概念。

首先，函数和`struct`，`int`等一样，也是一个类型(type)。

我们可以先定义一个函数类型，再声明一个该函数类型的变量，并且给该变量赋值。该函数类型变量的值我们就可以叫做函数值。看下面的代码示例就一目了然了。

```go
package main

import "fmt"

// 函数类型FuncType
type FuncType func(int) int

// 定义变量f1, 类型是FuncType
var f1 FuncType = func(a int) int { return a }

// 定义变量f, 类型是一个函数类型，函数签名是func(int) int
// 在main函数里给f赋值，零值是nil
var f func(int) int

func main() {
	fmt.Println(f1(1))
	// 定义变量f2，类型是FuncType
	var f2 FuncType
	f2 = func(a int) int {
		a++
		return a
	}
	fmt.Println(f2(1))

	// 给函数类型变量f赋值
	f = func(a int) int {
		return 10
	}
	fmt.Println(f(1))
}
```

我们平时实现函数的时候，通常是把函数的类型定义，函数变量和变量赋值一起做了。

函数类型的变量如果只声明不初始化，零值是nil，也就是默认值是nil。

通过上面的例子我们知道可以先定义函数类型，后面再定义函数类型的变量。



### 原理解析

我们再回顾下官方文档怎么说的：

>Each time a "defer" statement executes, the function value and parameters to
>the call are evaluated as usual and saved anew but the actual function is not 
>invoked. 
>
>Instead, deferred functions are invoked immediately before the 
>surrounding function returns, in the reverse order they were deferred. 
>
>That is, if the surrounding function returns through an explicit return statement, 
>deferred functions are executed after any result parameters are set by that 
>return statement but before the function returns to its caller. 
>
>If a deferred function value evaluates to nil, execution panics when the function is 
>invoked, not when the "defer" statement is executed.

本题的代码里的`bar()`函数在return之前按照如下顺序执行

| 执行                                | 执行结果                                                     |
| ----------------------------------- | ------------------------------------------------------------ |
| 执行return 1                        | 把1赋值给函数返回值参数`r`                                   |
| 执行f()                             | 因为`f`的值在`defer f()`这句执行的时候就已经确定下来是nil了，所以会引发panic |
| 执行`bar`函数里第1个被`defer`的函数 | r先加4，值变为5，然后recover捕获上一步的panic，r的值再加8，结果就是13 |
| `bar()`返回`r`的值                  | r的值是13，返回13。`main`里打印13                            |

因此本题的运行结果是13，答案是`D`。



##  总结

在`defer`第1篇文章里我们列出了`defer`六大原则，参考[Go defer语义6大原则](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483756&idx=1&sn=d536fa3340e1d5f91d72eaa8b67c8123&chksm=ce124e03f965c715e26f5943948e17d8e0ebb3c4a3a180a149219a610f83fc6eb77b3b166b6a&token=1521159887&lang=zh_CN#rd)。

本文再总结下本题目涉及的`defer`语义另外2大注意事项：

* 被defer的函数值(function value)在执行到defer语句的时候就被确定下来了。

  ```go
  package main
  
  import "fmt"
  
  func main() {
  	defer func() {
  		r := recover()
  		fmt.Println(r)
  	}()
  	var f func(int) // f没有初始化赋值，默认值是nil
  	defer f(1) // 函数变量f的值已经确定下来是nil了
  	f = func(a int) {}
  }
  ```

* 如果被`defer`的函数或方法的值是nil，在执行`defer`这条语句的时候不会报错，但是最后调用nil函数或方法的时候就引发`panic: runtime error: invalid memory address or nil pointer dereference`。



## 思考题

思考下面这道题的运行结果是什么？大家可以在评论区留下你们的答案。也可以在我的wx公号发送消息 `defer2` 获取答案和原因。

题目1：程序运行结果是什么？

```go
package main

import "fmt"

func main() {
	defer func() {
		fmt.Print(recover())
	}()
	defer func() {
		defer fmt.Print(recover())
		defer panic(1)
		recover()
	}()
	defer recover()
	panic(2)
}
```



## 开源地址

文章和示例代码开源地址在GitHub: [https://github.com/jincheng9/go-tutorial](https://github.com/jincheng9/go-tutorial)

公众号：coding进阶

个人网站：[https://jincheng9.github.io/](https://jincheng9.github.io/)

知乎：[https://www.zhihu.com/people/thucuhkwuji](https://www.zhihu.com/people/thucuhkwuji)



## 好文推荐

1. [被defer的函数一定会执行么？](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p2)
2. [Go Quiz: 从Go面试题看slice的底层原理和注意事项](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p8)
3. [Go Quiz: 从Go面试题看channel的注意事项](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p9)
4. [Go Quiz: 从Go面试题看分号规则和switch的注意事项](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p10)
5. [Go Quiz: 从Go面试题看defer语义的底层原理和注意事项](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p11)
6. [Go有引用变量和引用传递么？map,channel和slice作为函数参数是引用传递么？](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p3)
7. [new和make的使用区别是什么？](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p4)
8. [官方教程：Go泛型入门](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p6)
9. [一文读懂Go泛型设计和使用场景](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p7)
10. [一文读懂Go匿名结构体的使用场景](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p5)



## References

* https://go101.org/quizzes/defer-2.html
* https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p11
* https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p2
* https://golang.google.cn/ref/spec#Defer_statements
* https://chai2010.gitbooks.io/advanced-go-programming-book/content/appendix/appendix-a-trap.html