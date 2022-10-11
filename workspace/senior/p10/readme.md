# Go Quiz: 从Go面试题看分号规则和switch的注意事项

## 面试题

这是Go Quiz系列的第3篇，关于Go语言的分号规则和`switch`的特性。

这道题比较tricky，通过这道题可以加深我们对Go语言里的分号`;`规则和`switch`特性的理解。

```go
package main

func f() bool {
	return false
}

func main() {
	switch f() 
  {
	case true:
		println(1)
	case false:
		println(0)
	default:
		println(-1)
	}
}
```

- A: 1
- B: 0
- C: -1

这道题主要考察以下知识点：

* Go语言里的分号`;`规则

* `switch`后面的`{`换行后编译器会在背后做什么？

   

## 解析

**Go语言和`C++`一样，在每行语句(statement)的末尾是以分号`;`结尾的。**

看到这里，你可能会有点懵，是不是在想：我写Go代码的时候也没有在语句末尾加分号啊。。。

那是因为Go编译器的词法解析程序自动帮你做了这个事情，在需要加分号的地方给你加上了分号。

如果你在代码里显式地加上分号，编译器是不会报错的，只是Go不需要、也不建议显式加分号，一切交给编译器去自动完成。

**那编译器是怎么往我们代码里插入分号`;`的呢？规则是什么**？我们看看官方文档的说法：

> 1. When the input is broken into tokens, a semicolon is automatically inserted into the token stream immediately after a line's final token if that token is
>    - an [identifier](https://go.dev/ref/spec#Identifiers)
>    - an [integer](https://go.dev/ref/spec#Integer_literals), [floating-point](https://go.dev/ref/spec#Floating-point_literals), [imaginary](https://go.dev/ref/spec#Imaginary_literals), [rune](https://go.dev/ref/spec#Rune_literals), or [string](https://go.dev/ref/spec#String_literals) literal
>    - one of the [keywords](https://go.dev/ref/spec#Keywords) `break`, `continue`, `fallthrough`, or `return`
>    - one of the [operators and punctuation](https://go.dev/ref/spec#Operators_and_punctuation) `++`, `--`, `)`, `]`, or `}`
> 2. To allow complex statements to occupy a single line, a semicolon may be omitted before a closing `")"` or `"}"`.

根据这2个规则，我们来分析下本文最开始的题目，`switch`代码如下所示：

```go
// 示例1
switch f() 
{
	case true:
		println(1)
	case false:
		println(0)
	default:
		println(-1)
}
```

上面的代码对于`switch f()`满足规则1，会在`)`后面自动加上分号，等价于示例2

```go
// 示例2
switch f();
{
	case true:
		println(1)
	case false:
		println(0)
	default:
		println(-1)
}
```

示例2的代码等价于示例3，程序运行会进入到`case true`这个分支

```go
// 示例3
switch f();true {
	case true:
		println(1)
	case false:
		println(0)
	default:
		println(-1)
}
```

所以本题的答案是1，选择`A`。



##  总结

1. `Effective Go`官方建议：除了`for`循环之外，不要在代码里显式地加上分号`;`。

   你可能在`if`和`switch`关键字后面看到过分号`;`，比如下面的例子。

   ```go
   if i := 0; i < 1 {
     println(i)
   }
   
   switch os := runtime.GOOS; os {
   	case "darwin":
   		fmt.Println("OS X.")
   	case "linux":
   		fmt.Println("Linux.")
   	default:
   		// freebsd, openbsd,
   		// plan9, windows...
   		fmt.Printf("%s.\n", os)
   }
   ```

   虽然代码可以正常工作，但是官方不建议这样做。我们可以把分号`;`前面的变量赋值提前，比如修改为下面的版本：

   ```go
   i := 0
   if i < 1 {
     println(i)
   }
   
   os := runtime.GOOS
   switch os {
   	case "darwin":
   		fmt.Println("OS X.")
   	case "linux":
   		fmt.Println("Linux.")
   	default:
   		// freebsd, openbsd,
   		// plan9, windows...
   		fmt.Printf("%s.\n", os)
   }
   ```

2. `{`不要换行，如果对于`if`, `for`， `switch`和`select`把`{`换行了，Go编译器会自动添加分号`;`，会导致预期之外的结果，比如本文最开头的题目和下面的例子。

   ```go
   // good
   if i < f() {
       g()
   }
   
   // wrong: compile error
   if i < f()  // wrong!
   {           // wrong!
       g()
   }
   ```

3. 养成使用`go fmt`,` go vet`做代码检查的习惯，可以帮我们提前发现和规避潜在隐患。

   

## 思考题

思考下面这2道题的运行结果是什么？大家可以在评论区留下你们的答案。

题目1：

```go
// Foo prints and returns n.
func Foo(n int) int {
    fmt.Println(n)
    return n
}

func main() {
    switch Foo(2) {
    case Foo(1), Foo(2), Foo(3):
        fmt.Println("First case")
        fallthrough
    case Foo(4):
        fmt.Println("Second case")
    }
}
```

题目2：

```go
a := 1
fmt.Println(a++)
```



## 开源地址

文章和示例代码开源地址在GitHub: https://github.com/jincheng9/go-tutorial

公众号：coding进阶

个人网站：https://jincheng9.github.io/

知乎：https://www.zhihu.com/people/thucuhkwuji



## 好文推荐

1. [被defer的函数一定会执行么？](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p2)
2. [Go有引用变量和引用传递么？map,channel和slice作为函数参数是引用传递么？](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p3)
3. [new和make的使用区别是什么？](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p4)
4. [一文读懂Go匿名结构体的使用场景](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p5)
5. [官方教程：Go泛型入门](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p6)
6. [一文读懂Go泛型设计和使用场景](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p7)
7. [Go Quiz: 从Go面试题看slice的底层原理和注意事项](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p8)
8. [Go Quiz: 从Go面试题看channel的注意事项](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p9)



## References

* https://go101.org/quizzes/switch-1.html
* https://go.dev/doc/effective_go
* https://go101.org/article/line-break-rules.html
* https://yourbasic.org/golang/switch-statement/