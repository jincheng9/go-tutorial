# receive-only channel和send-only channel有用么？

## 背景

Go内置的数据结构channel我们都很熟悉，channel

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

## 解析

* 题目1的答案是D，题目1创建了一个receive-only channel，只能从channel里接收值，不能往channel里发送值。

​	对于receive-only channel不能close，如果做close操作，会有如下报错：

```go
./main.go:9:7: invalid operation: close(ch) (cannot close receive-only channel)
```

* 题目2的答案是A，题目2创建了一个send-only channel，只能往channel里发送值，不能从channel里接收值。

​	对于send-only channel可以正常close。



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