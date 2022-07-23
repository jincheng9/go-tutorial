# send-only channel和receive-only channel的争议

## 背景

Go内置的数据结构channel相信大家都很熟悉，channel在Go语言里扮演了非常重要的角色。

* channel可以帮助实现goroutine之间的通信和同步。
* Go语言通过goroutine和channel实现了CSP(Communicating Sequencial Process)模型。

Go语言的发明人之一Robe Pike说过下面这句话：

>  Don't communicate by sharing memory, share memory by communicating.

这句话里隐藏的含义其实是希望Go开发者使用channel来实现goroutine之间的通信。

channel分为2种类型：

* bi-directional channel(双向channel)：既可以往channel里发送数据，也可以从channel接收数据
* uni-directional channel(单向channel)
  * send-only channel：只能往channel发送数据，不能从channel接收数据，否则会编译报错
  * receive-only channel：只能从channel接收数据，不能往channel发送数据，否则会编译报错

单向channel的一个典型使用场景是作为函数或方法参数，用来控制只能往channel发送数据或者只能从channel接收数据，避免误操作。

```go
package main

import (
	"fmt"
)

// send-only channel
func testSendChan(c chan<- int) {
	c <- 20
}

// receive-only channel
func testRecvChan(c <-chan int) {
	result := <-c
	fmt.Println("result:", result)
}

func main() {
	ch := make(chan int, 3)
	testSendChan(ch)
	testRecvChan(ch)
}
```

比如上面这段程序，`testSendChan`的参数是一个send-only channel，`testRecvChan`的参数是一个receive-only channel。

实参`ch`是一个双向channel，Go runtime会把双向channel转换成单向channel传给对应的函数。



## 问题

上面的例子是单向channel作为函数形参，双向channel作为函数实参，make函数创建的是一个双向channel。

那我们可以用make来创建单向channel么？make创建的单向channel有实际使用场景么？有潜在的坑么？

我们先看看如下2道题目：

### 题目1

```go
ch := make(<-chan int)
close(ch)
fmt.Println("ok")
```

* A: 打印ok
* B: 运行时报错：fatal error - deadlock
* C: 运行时报错：painic
* D: 编译失败

### 题目2

```go
c := make(chan<- int)
close(c)
fmt.Println("ok")
```

* A: 打印ok
* B: 运行时报错：fatal error - deadlock
* C: 运行时报错：painic
* D: 编译失败

大家可以先稍微停顿一下，想想这2道题的答案会是什么。

## 解析

### 答案

* 题目1的答案是D，题目1创建了一个receive-only channel，只能从channel里接收值，不能往channel里发送值。

  对于receive-only channel不能close，如果做close操作，会有如下报错：

```go
./main.go:9:7: invalid operation: close(ch) (cannot close receive-only channel)
```

* 题目2的答案是A，题目2创建了一个send-only channel，只能往channel里发送值，不能从channel里接收值。

  对于send-only channel可以正常close。

为什么receive-only channel不能close，但是send-only channel可以close呢？

这是因为receive-only channel表示只能从这个channel里接收数据，使用方对这个channel只有读权限，对channel没有写权限，即不能往channel发送数据或者对channel做close操作。

所以题目1的答案是D，题目2的答案是A。

### 衍生问题

```go
// send-only channel
func testSendChan(c chan<- int) {
	c <- 20
}

// receive-only channel
func testRecvChan(c <-chan int) {
	result := <-c
	fmt.Println("result:", result)
}

func main() {
	ch := make(chan int, 3)
	testSendChan(ch)
	testRecvChan(ch)
}
```

上面这段代码是单向channel的典型使用场景，函数实参是双向channel，函数参数是单向channel，我们可以通过send-only channel往channel发送数据，通过receive-only channel从channel里接收数据，这个很好理解。**但是要注意的是这里的函数实参是双向channel**。

细心的开发者可能有个疑问，对于上面的2道题目，make创建的是单向channel，要么只能往这个channel里发数据，要么只能从这个channel接收数据，而不像上面代码里make创建的channel是个双向channel。

**问题来了，对于make创建的单向channel，有实际用途么**？

比如`make(chan<- int)`创建的send-only channel，只能往这个channel发数据，没办法从这个channel里读数据，有啥用么？

比如`make(<-chan int)`创建的receive-only channel，只能从这个channel里收数据，没办法往这个channel里发数据，有啥用么？

关于这个问题，其实Go语言的大佬之间也产生过比较大的争议。

Golang团队的[Brad Fitzpatrick](https://github.com/bradfitz)就抱怨过，认为编译器应该直接对make创建单向channel报错，因为make创建的单向channel没啥用。

> ```
> I'm surprised the compiler permits create send- or receive-only channels:
> ------
> package main
> 
> import "fmt"
> 
> func main() {
>     c := make(<-chan int)
>     fmt.Printf("%T\n", c)
> }
> 
> ------
> 
> Is that an accident?
> 
> I can't see what utility they'd have.
> ```

Go语言负责人[Russ Cox](https://github.com/rsc)给的反馈是：认为没有明确理由让编译器禁止make创建单向channel，因为对单向channel如果有非法的操作，编译器还是会报错，而且让编译器直接禁止make创建单向channel会增加Go设计的复杂性，没有必要，得不偿失。

> ```
> I can't see a reason to disallow it either though.
> 
> All I am saying is that explicitly disallowing it adds
> complexity to the spec that seems unnecessary.
> What bugs or problems does it avoid to make this
> special exclusion in the rules for make?
> Russ
> ```



## 推荐阅读

* [Go Quiz: 从Go面试题看channel的注意事项](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483746&idx=1&sn=c3ec0e3f67fa7b1cb82e61450d10c7fd&chksm=ce124e0df965c71b7e148ac3ce05c82ffde4137cb901b16c2c9567f3f6ed03e4ff738866ad53&token=1224150345&lang=zh_CN#rd)

* [Go Quiz: 从Go面试题看channel在select场景下的注意事项](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483816&idx=1&sn=44e5cf4900b44f9a0cde491df5dd6e51&chksm=ce124ec7f965c7d1edd9ccffe80520981970ad6000cfea3b1a4099a4627f0f24cc33272ec996&token=1224150345&lang=zh_CN#rd)

  

## 总结

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



## 开源地址

文章和示例代码开源在GitHub: [Go语言初级、中级和高级教程](https://github.com/jincheng9/go-tutorial)。

公众号：coding进阶。关注公众号可以获取最新Go面试题和技术栈。

个人网站：[Jincheng's Blog](https://jincheng9.github.io/)。

知乎：[无忌](https://www.zhihu.com/people/thucuhkwuji)。



## 福利

我为大家整理了一份后端开发学习资料礼包，包含编程语言入门到进阶知识(Go、C++、Python)、后端开发技术栈、面试题等。

关注公众号「coding进阶」，发送消息 **backend** 领取资料礼包，这份资料会不定期更新，加入我觉得有价值的资料。还可以发送消息「**进群**」，和同行一起交流学习，答疑解惑。



## References

* https://twitter.com/val_deleplace/status/1544951594659287041
* https://github.com/golang/go/issues/2431