# Go语言new和make的使用区别和最佳实践

## 前言

在Go语言中，初始化数据结构的时候，可能会用到2个内置函数：new和make。

new和make都可以用来分配内存，那他们有什么区别呢？在写代码过程中，对于new和make的最佳实践又是什么呢？

## new是什么

我们先看看new的官方定义：

> func new(Type) *Type
>
> The new built-in function allocates memory. The first argument is a type, not a value, and the value returned is a pointer to a newly allocated zero value of that type.

从官方定义里可以看到，new有以下几个特点：

1. 分配内存。内存里存的值是对应类型的[零值](../../lesson3)。
2. 只有一个参数。参数是分配的内存空间所存储的变量类型，Go语言里的任何类型都可以是new的参数，比如int， 数组，结构体，甚至函数类型都可以。
3. 返回的是指针。

以下代码是等价的，可以认为new(T)是 var t T; &t 的语法糖。

```go
// 方式1
ptr := new(T)

// 方式2
var t T
ptr := &t
```

注意：Go里的new和C++的new是不一样的：

* Go的new分配的内存可能在栈(stack)上，可能在堆(heap)上。C++ new分配的内存一定在堆上。
* Go的new分配的内存里的值是对应类型的零值，不能显式初始化指定要分配的值。C++ new分配内存的时候可以显式指定要存储的值。
* Go里没有构造函数，Go的new不会去调用构造函数。C++的new是会调用对应类型的构造函数。



## make是什么

我们再看看make的官方定义：

> func make(t Type, size ...IntegerType) Type
>
> The make built-in function allocates and initializes an object of type slice, map, or chan (only). Like new, the first argument is a type, not a value. Unlike new, make's return type is the same as the type of its argument, not a pointer to it. The specification of the result depends on the type:
>
>  Slice: The size specifies the length. The capacity of the slice is
>  equal to its length. A second integer argument may be provided to
>  specify a different capacity; it must be no smaller than the
>  length. For example, make([]int, 0, 10) allocates an underlying array
>  of size 10 and returns a slice of length 0 and capacity 10 that is
>  backed by this underlying array.
>
>  Map: An empty map is allocated with enough space to hold the
>  specified number of elements. The size may be omitted, in which case
>  a small starting size is allocated.
>
>  Channel: The channel's buffer is initialized with the specified
>  buffer capacity. If zero, or the size is omitted, the channel is
>  unbuffered.

从官方定义里可以看到，make有如下几个特点：

1. 分配和初始化内存。

2. 只能用于slice, map和chan这3个类型，不能用于其它类型。

   如果是用于slice类型，make函数的第2个参数表示slice的长度，这个参数必须给值。

3. 返回的是原始类型，也就是slice, map和chan，不是返回指向slice, map和chan的指针。



## 几个问题

这里来回答几个使用new和make过程中的常见问题，从简单到复杂：

* 为什么针对slice, map和chan类型专门定义一个make函数？

  答案：这是因为slice, map和chan的底层结构上要求在使用slice，map和chan的时候必须初始化，如果不初始化，那slice，map和chan的值就是零值，也就是nil。我们知道：

  * map如果是nil，是不能往map插入元素的，插入元素会引发panic

  * chan如果是nil，往chan发送数据或者从chan接收数据都会阻塞

  * slice会有点特殊，理论上slice如果是nil，也是没法用的。但是append函数处理了nil slice的情况，可以调用append函数对nil slice做扩容。但是我们使用slice，总是会希望可以自定义长度或者容量，这个时候就需要用到make。

    

* 可以用new来创建slice, map和chan么？

  答案：可以。代码可以参考如下示例：

  ```go
  // example1.go
  package main
  
  import "fmt"
  
  func main() {
  	a := *new([]int)
  	fmt.Printf("%T, %v\n", a, a==nil)
  
  	b := *new(map[string]int)
  	fmt.Printf("%T, %v\n", b, b==nil)
  
  	c := *new(chan int)
  	fmt.Printf("%T, %v\n", c, c==nil)
  }
  ```

  输出结果是：

  ```go
  []int, true
  map[string]int, true
  chan int, true
  ```

  虽然new可以用来创建slice, map和chan，但实际上并没有卵用，因为new创建的slice, map和chan的值都是零值，也就是nil。这3种类型如果是nil，那遇到的问题我们在上面第一个问题已经解答过了，这里不再赘述。

  

* 为什么slice是nil也可以直接append？

  答案：对于nil slice，append会对slice的底层数组做扩容，通过调用mallocgc向Go的内存管理器申请内存空间，再赋值给原来的nil slice。

* slice用make创建的时候，如果指定的长度len>0，则make创建的slice下标索引从0到len-1的值都是对应slice里元素类型的零值。参考下例：

  ```go
  package main
  
  import "fmt"
  
  func main() {
  	s := make([]int, 2, 3)
  	fmt.Println(s) // [0 0]
  }
  ```

  

## 最佳实践

1. 尽量不使用new
2. 对于slice, map和chan的定义和初始化，优先使用make函数



## References

* https://draveness.me/golang/docs/part2-foundation/ch05-keyword/golang-make-and-new/
* https://dave.cheney.net/2014/08/17/go-has-both-make-and-new-functions-what-gives
* https://mp.weixin.qq.com/s/MTZ0C9zYsNrb8wyIm2D8BA
* https://juejin.cn/post/6859145664316571661
* https://www.godesignpatterns.com/2014/04/new-vs-make.html