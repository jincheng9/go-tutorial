# Go Quiz: 从Go面试题看slice的底层原理和注意事项

## 面试题

最近Go 101的作者发布了11道Go面试题，非常有趣，打算写一个系列对每道题做详细解析。欢迎大家关注。

大家可以看下面这道关于`slice`的题目，通过这道题我们可以对`slice`的特性和注意事项有一个深入理解。

```go
package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3}
	x := a[:1]
	y := a[2:]
	x = append(x, y...)
	x = append(x, y...)
	fmt.Println(a, x)
}
```

- A: [0 1 2 3] [0 2 3 3 3]
- B: [0 2 3 3] [0 2 3 3 3]
- C: [0 1 2 3] [0 2 3 2 3]
- D: [0 2 3 3] [0 2 3 2 3]

大家可以在评论区留下你们的答案。这道题有几个考点：

1. `slice`的底层数据结构是什么？给`slice`赋值，到底赋了什么内容？
2. 通过`:`操作得到的新`slice`和原`slice`是什么关系？新`slice`的长度和容量是多少？
3. `append`会在背后做哪些事情？
4. `slice`的扩容机制是什么？



## 解析

我们先逐个解答上面的问题。

## slice的底层数据结构

`talk is cheap, show me the code`.  直接上`slice`的源码：

`slice`定义在`src/runtime/slice.go`第15行，源码地址：https://github.com/golang/go/blob/master/src/runtime/slice.go#L15。

`Pointer`定义在`src/unsafe/unsafe.go`第184行，源码地址：https://github.com/golang/go/blob/master/src/unsafe/unsafe.go#L184。

```go
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

type Pointer *ArbitraryType
```

`slice`实际上是一个结构体类型，包含3个字段，分别是

* array: 是指针，指向一个数组，切片的数据实际都存储在这个数组里。
* len: 切片的长度。
* cap: 切片的容量，表示切片当前最多可以存储多少个元素，如果超过了现有容量会自动扩容。

**因此给`slice`赋值，实际上都是给`slice`里的这3个字段赋值**。



## `:`操作符



## append机制



## slice扩容机制



## 总结

> 对于slice，时刻想着对slice做了一个操作后，新的slice的指针，长度，容量是怎么变的

* slice数据结构

* `:`操作符

* 下标上下限是长度，`:`上下限是容量

* 拷贝赋值

* append赋值

* 给slice复制：指针赋值，是不是指向新地址， 长度赋值，容量赋值

* Go没有传引用这个说法，都是传值，可以参考我之前的文章xxx

  

## 开源地址

文章和代码开源地址在GitHub: https://github.com/jincheng9/go-tutorial

公众号：coding进阶

个人网站：https://jincheng9.github.io/



## References

* https://go101.org/quizzes/slice-1.html
* https://go.dev/blog/slices-intro