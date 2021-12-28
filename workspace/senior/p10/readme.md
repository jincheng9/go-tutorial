# Go Quiz: 从Go面试题看switch的注意事项

## 面试题

这是Go Quiz系列的第3篇，关于`switch`的特性。

这道题比较tricky，通过这道题可以加深我们对`switch`和Go语言换行的理解。

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

* `switch`后面的`{`换行后编译器会在背后做什么？

   

## 解析

```go
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

```go
switch f();{
	case true:
		println(1)
	case false:
		println(0)
	default:
		println(-1)
}
```

```go
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





## 思考题

思考下面这道题的输出是什么？大家可以在评论区留下你们的答案。

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
* https://go101.org/article/line-break-rules.html
* https://yourbasic.org/golang/switch-statement/