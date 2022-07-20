# receive-only channel和send-only channel有用么？

## 背景

Go内置的数据结构channel我们都很熟悉，

### 题目1

```go
ch := make(<-chan int)
close(ch)
fmt.Println("ok")
```

* A: 正常打印ok
* B: 运行时报错：fatal error - deadlock
* C: 运行时报错：painic
* D: 编译失败

### 题目2

```go
c := make(chan<- int)
close(c)
fmt.Println("ok")
```

* A: 正常打印ok
* B: 运行时报错：fatal error - deadlock
* C: 运行时报错：painic
* D: 编译失败

## 解析

题目1的答案是D，题目1创建了一个receive-only channel，只能从channel里接收值，不能往channel里发送值。

对于receive-only channel不能close。

题目2的答案是A，题目2创建了一个send-only channel，只能往channel里发送值，不能从channel里接收值。

对于send-only channel可以正常close。



## 总结



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