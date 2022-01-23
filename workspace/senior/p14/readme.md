# Go Quiz: 从Go面试题看channel在select场景下的注意事项

## 面试题

这是Go Quiz系列中关于channel的第2篇，涉及`channel`被close后的特性，以及在`select`和`channel`一起使用时的注意事项。

这道题目来源于Google的工程师Valentin Deleplace。

```go
package main

import "fmt"

func main() {
	data := make(chan int)
	shutdown := make(chan int)
	close(shutdown)
	close(data)

	select {
	case <-shutdown:
		fmt.Print("CLOSED, ")
	case data <- 1:
		fmt.Print("HAS WRITTEN, ")
	default:
		fmt.Print("DEFAULT, ")
	}
}
```

- A: 进入default分支，打印"DEFAULT, "
- B: 进入shutdown分支，打印"CLOSED, "
- C: 进入data分支，打印"HAS WRITTEN, "
- D: 程序会panic
- E: 程序可能panic，也可能打印"CLOSED, "

这道题主要考察以下知识点：

* `channel`被关闭后，从`channel`接收数据和往`channel`发送数据会有什么结果？

* `select`的运行机制是怎样的？

   

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

* data和shutdown这2个channel都被关闭了。
* 对于关闭的channel，从channel里接收数据，拿到的是channel的存储的元素类型的零值，因此`case <-shutdown`这个case分支不会阻塞。
* 对于关闭的channel，向其发送数据会引发panic，因此`case data <- 1`这个case分支不会阻塞，会引发panic。
* 因此这个select语句执行的时候，2个case分支都不会阻塞，都可能执行到。如果执行的是`case <-shutdown`这个case分支，会打印"CLOSED, "。如果执行的是`case data <- 1`这个case分支，会导致程序panic。

因此本题的答案是`E`。



## 加餐

可以回顾Go quiz系列中关于channel的第一道题目，加深对channel的理解。

题目链接地址：[channel面试题和注意事项](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483746&idx=1&sn=c3ec0e3f67fa7b1cb82e61450d10c7fd&chksm=ce124e0df965c71b7e148ac3ce05c82ffde4137cb901b16c2c9567f3f6ed03e4ff738866ad53&token=609026015&lang=zh_CN#rd)



## 开源地址

文章和示例代码开源地址在GitHub: https://github.com/jincheng9/go-tutorial

公众号：coding进阶

个人网站：https://jincheng9.github.io/

知乎：https://www.zhihu.com/people/thucuhkwuji



## References

* https://twitter.com/val_deleplace/status/1484172826974195712
* https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p9
* https://github.com/jincheng9/go-tutorial/tree/main/workspace/lesson19
* https://github.com/jincheng9/go-tutorial/tree/main/workspace/lesson29