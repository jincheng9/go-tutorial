# Go Quiz: 从Go面试题看defer语义的底层原理和注意事项

## 面试题

这是Go Quiz系列的第4篇，关于Go语言的`defer`语义。

这道题稍微有点迷惑性，通过这道题可以加深我们对Go语言`defer`关键字底层运行机制的理解。

```go
package main

type Foo struct {
	v int
}

func NewFoo(n *int) Foo {
	print(*n)
	return Foo{}
}

func (Foo) Bar(n *int) {
	print(*n)
}

func main() {
	var x = 1
	var p = &x
	defer NewFoo(p).Bar(p)
	x = 2
	p = new(int)
	NewFoo(p)
}
```

- A: 100
- B: 102
- C: 022
- D: 011

这道题主要考察以下知识点：

* 被`defer`的函数或方法什么时候执行？

* 被`defer`的函数或方法的参数的值是什么时候确定的？

* 被`defer`的函数或方法如果存在多级调用是什么机制？比如本题的`NewFoo(p).Bar(p)`就存在二级调用，先调用了`NewFoo`函数，再调用了`Bar`方法。

   

## 解析

我们再看看官方文档怎么说的：

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

官方文档的前两句话对我们求解本题至关重要，用中文来表述就是：

**假设我们在函数`A`的函数体内运行了`defer B(params)`，那被`defer`的函数`B`的参数会像普通函数调用一样被即时计算，但是被`defer`的函数`B`的调用会延迟到函数`A` return或者panic之前执行。**

如果`defer`后面跟的是多级函数的调用，只有最后一个函数会被延迟执行。

比如上例里的`defer NewFoo(p).Bar(p)`中，`NewFoo(p)`是会被即时执行的，不会延后，只有最后一个方法`Bar`的调用会被延后执行。

同时，函数的传参也是被即时计算的，也就是`defer NewFoo(p).Bar(p)`里涉及的参数`p`的值也是被即时计算保存好的，延后执行的时候就用事先计算好的值。

这道题求解过程如下：

| 代码                   | 执行结果                                                     |
| ---------------------- | ------------------------------------------------------------ |
| var x = 1              | 定义一个整型变量x，值为1                                     |
| var p = &x             | 定义一个指针p，指向变量x                                     |
| defer NewFoo(p).Bar(p) | NewFoo(p)和参数p会被即时计算。NewFoo(p)的执行结果会打印1，指针p的值是变量x的内存地址，*p也就是变量x的值，Bar的调用会延后到main函数return之前执行 |
| x = 2                  | 修改x的值为2                                                 |
| p = new(int)           | 修改指针p的值，不再指向x，而是指向新的内存地址，该内存地址存的值是int的零值，也就是0 |
| NewFoo(p)              | 调用函数NewFoo，打印0。<br>main函数return之前，执行Bar方法，Bar的参数p的值在defer语句执行的时候就已经确定下来了，是变量x的内存地址，因此Bar方法打印的是变量x的值，也就是2 |

因此本题的运行结果是102，答案是B。

##  defer六大原则

最后总结下`defer`语义要注意的六大关键点：

1. defer后面跟的必须是函数或者方法调用，defer后面的表达式不能加括号。

   ```go
   defer (fmt.Println(1)) // 编译报错，因为defer后面跟的表达式不能加括号
   ```

2. 被defer的函数或方法的参数的值在执行到defer语句的时候就被确定下来了。

   ```go
   func a() {
       i := 0
       defer fmt.Println(i) // 最终打印0
       i++
       return
   }
   ```

   上例中，被defer的函数fmt.Println的参数`i`在执行到defer这一行的时候，`i`的值是0，fmt.Println的参数就被确定下来是0了，因此最终打印的结果是0，而不是1。

3. 被`defer`的函数或者方法如果存在多级调用，只有最后一个函数或方法会被`defer`到函数return或者panic之前执行，参见上面的说明。

4. 被defer的函数执行顺序满足LIFO原则，后defer的先执行。

   ```go
   func b() {
       for i := 0; i < 4; i++ {
           defer fmt.Print(i)
       }
   }
   ```

   上例中，输出的结果是3210，后defer的先执行。

5. 被defer的函数可以对defer语句所在的函数的命名返回值(named return value)做读取和修改操作。

   ```go
   // f returns 42
   func f() (result int) {
   	defer func() {
   		// result is accessed after it was set to 6 by the return statement
   		result *= 7
   	}()
   	return 6
   }
   ```

   上例中，被defer的函数func对defer语句所在的函数**f**的命名返回值result做了修改操作。

   调用函数`f`，返回的结果是42。

   执行顺序是函数`f`先把要返回的值6赋值给result，然后执行被defer的函数func，result被修改为42，然后函数`f`返回result给调用方，也就返回了42。

6. 即使`defer`语句执行了，被`defer`的函数不一定会执行。对这句话不理解的可以参考下面的思考题。

## 思考题

思考下面这3道题的运行结果是什么？大家可以在评论区留下你们的答案。

题目1：程序运行结果是什么？

```go
package main

type T int

func (t T) M(n int) T {
	print(n)
	return t
}

func main() {
	var t T
	defer t.M(1).M(2)
	t.M(3).M(4)
}
```

题目2："end"会被打印么？`f(1, 2)`会不会编译报错？

```go
package main

import "fmt"

type Add func(int, int) int

func main() {
	var f Add
	defer f(1, 2)
	fmt.Println("end")
}
```

题目3：“test”会被打印么？

```go
package main

import (
	"fmt"
	"os"
)

func test1() {
	fmt.Println("test")
}

func main() {
	defer test1()
	os.Exit(0)
}
```



## 开源地址

文章和示例代码开源地址在GitHub: [https://github.com/jincheng9/go-tutorial](https://github.com/jincheng9/go-tutorial)

公众号：coding进阶

个人网站：[https://jincheng9.github.io/](https://jincheng9.github.io/)

知乎：[https://www.zhihu.com/people/thucuhkwuji](https://www.zhihu.com/people/thucuhkwuji)



## 好文推荐

1. [被defer的函数一定会执行么？](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p2)
2. [Go有引用变量和引用传递么？map,channel和slice作为函数参数是引用传递么？](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p3)
3. [new和make的使用区别是什么？](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p4)
4. [一文读懂Go匿名结构体的使用场景](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p5)
5. [官方教程：Go泛型入门](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p6)
6. [一文读懂Go泛型设计和使用场景](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p7)
7. [Go Quiz: 从Go面试题看slice的底层原理和注意事项](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p8)
8. [Go Quiz: 从Go面试题看channel的注意事项](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p9)
9. [Go Quiz: 从Go面试题看分号规则和switch的注意事项](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p10)



## References

* https://go101.org/quizzes/defer-1.html
* https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p2
* https://golang.google.cn/ref/spec#Defer_statements