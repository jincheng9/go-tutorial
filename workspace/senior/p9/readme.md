# Go Quiz: 从Go面试题看channel的注意事项

## 面试题

这是Go Quiz系列的第2篇，关于`channel`和`select`的特性。

这道题比较简单，但是通过这道题可以加深我们对`channel`和`select`的理解。

```go
package main

func main() {
	c := make(chan int, 1)
	for done := false; !done; {
		select {
		default:
			print(1)
			done = true
		case <-c:
			print(2)
			c = nil
		case c <- 1:
			print(3)
		}
	}
}
```

- A: 321
- B: 21
- C: 1
- D: 31

这道题主要考察以下知识点：

* `channel`的数据收发在什么情况会阻塞？

* `select`的运行机制是怎样的？

* `nil channel`收发数据是什么结果？

   

## 解析

1. 对于无缓冲区的`channel`，往`channel`发送数据和从`channel`接收数据都会阻塞。

2. 对于`nil channel`和有缓冲区的`channel`，收发数据的机制如下表所示：

   | channel           | nil   | 空的     | 非空非满 | 满了     |
   | ----------------- | ----- | -------- | -------- | -------- |
   | 往channel发送数据 | 阻塞  | 发送成功 | 发送成功 | 阻塞     |
   | 从channel接收数据 | 阻塞  | 阻塞     | 接收成功 | 接收成功 |
   | 关闭channel       | panic | 关闭成功 | 关闭成功 | 关闭成功 |

3. `channel`被关闭后：

   * 往被关闭的`channel`发送数据会触发panic。

   * 从被关闭的`channel`接收数据，会先读完`channel`里的数据。如果数据读完了，继续从`channel`读数据会拿到`channel`里存储的元素类型的零值。

     ```go
     data, ok := <- c 
     ```

     对于上面的代码，如果channel `c`关闭了，继续从`c`里读数据，当`c`里还有数据时，`data`就是对应读到的值，`ok`的值是`true`。如果`c`的数据已经读完了，那`data`就是零值，`ok`的值是`false`。

   * `channel`被关闭后，如果再次关闭，会引发panic。

4. `select`的运行机制如下：
   * 选取一个可执行不阻塞的`case`分支，如果多个`case`分支都不阻塞，会随机选一个`case`分支执行，和`case`分支在代码里写的顺序没关系。
   * 如果所有`case`分支都阻塞，会进入`default`分支执行。
   * 如果没有`default`分支，那`select`会阻塞，直到有一个`case`分支不阻塞。

根据以上规则，本文最开始的题目，在运行的时候

* 第1次for循环，只有`c <- 1`是不阻塞的，所以最后一个`case`分支执行，打印3。
* 第2次for循环，只有`<-c`是不阻塞的，所以第1个`case`分支执行，打印2，同时`channel`被赋值为`nil`。
* 第3次for循环，因为`channel`是`nil`，`对nil channel`的读和写都阻塞，所以进入`default`分支，打印1，`done`设置为`true`，for循环退出，程序运行结束。

因此打印结果是`321`,答案是`A`。



## 加餐：channel函数传参是引用传递？

网上有些文章写Go的`slice`，`map`，`channel`作为函数参数是传引用，这是**错误**的，Go语言里只有值传递，没有引用传递。

既然`channel`是值传递，那为什么`channel`作为函数形参时，函数内部对`channel`的读写对外部的实参`channel`是可见的呢？

对于这个问题，看Go的源码就一目了然了。`channel`定义在`src/runtime/chan.go`第33行，源码地址：https://github.com/golang/go/blob/master/src/runtime/chan.go#L33

```go
func makechan(t *chantype, size int) *hchan

// channel结构体
type hchan struct {
	qcount   uint           // total data in the queue
	dataqsiz uint           // size of the circular queue
	buf      unsafe.Pointer // points to an array of dataqsiz elements
	elemsize uint16
	closed   uint32
	elemtype *_type // element type
	sendx    uint   // send index
	recvx    uint   // receive index
	recvq    waitq  // list of recv waiters
	sendq    waitq  // list of send waiters

	// lock protects all fields in hchan, as well as several
	// fields in sudogs blocked on this channel.
	//
	// Do not change another G's status while holding this lock
	// (in particular, do not ready a G), as this can deadlock
	// with stack shrinking.
	lock mutex
}
```

我们通过`make`函数来创建`channel`时，Go会调用运行时的`makechan`函数。

从上面的代码可以看出`makechan`返回的是指向`channel`的指针。

因此`channel`作为函数参数时，实参`channel`和形参`channel`都指向同一个`channel`结构体的内存空间，所以在函数内部对`channel`形参的修改对外部`channel`实参是可见的，反之亦然。



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



## References

* https://go101.org/quizzes/channel-1.html
* https://jincheng9.github.io/
* https://github.com/jincheng9/go-tutorial/tree/main/workspace/lesson19
* https://github.com/jincheng9/go-tutorial/tree/main/workspace/lesson29