# Go常见错误第14篇：过度使用getter和setter方法

## 前言

这是Go常见错误系列的第14篇：过度使用getter和setter方法。

素材来源于Go布道者，现Docker公司资深工程师[Teiva Harsanyi](https://teivah.github.io/)。

本文涉及的源代码全部开源在：[Go常见错误源代码](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p28)，欢迎大家关注公众号，及时获取本系列最新更新。

 

## 常见错误和最佳实践

### 现状

写Java或者C++的人，可能会习惯下面的编程模式：

* 将不希望外部直接访问的类成员变量设置为private私有成员。
* 在类里定义public的get和set方法，用于外部获取和修改这个成员变量的值。get方法我们叫做getter，set方法叫做setter。

这是一种数据封装模式，在Java和C++里被广泛使用。

但是在Go语言里，官方从来没有建议使用getter和setter，我们可以直接访问结构体里的成员变量。

成员变量的可见性通过结构体标识符首字母大小写以及成员变量首字母大小写来控制到package这个层面。

* 如果结构体要被其它package使用，那结构体的标识符或者说结构体的名称首字母要大写。
* 如果结构体的成员要被其它package使用，那结构体和结构体的成员标识符首字母都要大写，否则只能在当前包里使用。

举个Go标准库里的time.Timer结构体的例子：

```go
// The Timer type represents a single event.
// When the Timer expires, the current time will be sent on C,
// unless the Timer was created by AfterFunc.
// A Timer must be created with NewTimer or AfterFunc.
type Timer struct {
	C <-chan Time
	r runtimeTimer
}
```

Timer结构体定义如上所示，里面有一个成员变量`C`用于接收Timer到点后的当前时间。

Timer和C都是大写，所以我们可以直接在下面的代码里访问Timer里的成员变量C拿到当前时间。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// print current time
	fmt.Println(time.Now())

	// NewTimer creates a new Timer that will send
	// the current time on its channel after at least duration d.
	timer := time.NewTimer(5 * time.Second)

	// print current time
	fmt.Println(<-timer.C)
}
```

上面程序执行结果是：

```bash
2022-11-06 13:07:07.706011 +0800 CST m=+0.000174256
2022-11-06 13:07:12.709128 +0800 CST m=+5.003141645
```

这种写法当然不是Go官方所预期的，因为成员变量`C`一般来说是不直接对外访问。

如果`C`暴露了可以对外访问，那我们甚至修改`C`的值，导致程序出错。

尽管不推荐这种写法，但是通过这个例子，我们可以知道如下事实：

**Go标准库里对于结构体里不应该修改的字段，也没有使用getter和setter方法**。



### 辩证来看

尽管Go官方没有使用getter和setter，但是从另一方面来说，在一些特定场景下使用getter和setter是有好处的。

- getter和setter隐藏了内部实现，我们可以自己灵活控制该暴露哪些东西。
- 如果成员变量的值发生了预期之外的变化，那通过getter和setter，我们可以方便做一些调试，更快发现问题。

Go语言里如果要使用getter和setter方法，有一些命名规范需要遵循。

假设我们要对结构体里的成员变量balance增加getter和setter方法，那么规范如下：

- getter方法应该被命名为Balance(而不是GetBalance)。
- setter方法应该被命名为SetBalance。
- 首字母大写是因为要被外部package使用，要大写来保证可见性。

示例如下：

```go
currentBalance := customer.Balance()
if currentBalance < 0 {
    customer.SetBalance(0)
}
```



## 总结

* Java/C++等语言里常用的getter和setter，在Go语言里并不是惯例和规范。
* 但是如果发现有上面讲到的需要使用到getter和setter的场景，那还是应该使用的，而不是完全不用。
* getter和setter方法命名参考上面提到的命名规范。



## 推荐阅读

* [Go面试题系列，看看你会几题？](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=Mzg2MTcwNjc1Mg==&action=getalbum&album_id=2199553588283179010#wechat_redirect)

* [Go常见错误第1篇：未知枚举值](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484146&idx=1&sn=10fb12b643a2e37c090e5aa3bc583152&chksm=ce124d9df965c48bb954aeddabdff3db12738ded3875542250c5d0ef6cfd4417fc56580288b1&token=1912894792&lang=zh_CN#rd)

* [Go常见错误第2篇：benchmark性能测试的坑](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484163&idx=1&sn=b28d61c1f3ec9d914e698dce105ba5d1&chksm=ce124c6cf965c57a90bc85a5295ed9375103de20607b509f845583ff6686385df0ed96653d00&token=1912894792&lang=zh_CN#rd)

* [Go常见错误第3篇：go指针的性能问题和内存逃逸](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484247&idx=1&sn=faf716627afb00df646cecff023fb63c&chksm=ce124c38f965c52efd009a4c98691d56b5765dc7dce98aa49b226ad9274bd062d8d01e702e91&token=1899277735&lang=zh_CN#rd)

* [Go常见错误第4篇：break操作的注意事项](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484262&idx=1&sn=c1bea8af60444a4ef73c4d4d7a09d16d&chksm=ce124c09f965c51f3663ac9089a792d36c3685850e12695dd26d15a1a50f393b2d7c92b9983a&token=461369035&lang=zh_CN#rd)

* [Go常见错误第5篇：Go语言Error管理](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484274&idx=1&sn=711abea3c6fd5d15341ee1b34da8a160&chksm=ce124c1df965c50b3af84965f7ed30b574cd0b247ea6f77b944ec858bd43ee37f4c1554a5bce&token=1846351524&lang=zh_CN#rd)

* [Go常见错误第6篇：slice初始化常犯的错误](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484289&idx=1&sn=2b8171458cde4425b28fdf8f51df8d7c&chksm=ce124ceef965c5f8a14f5951457ce2ac0ecc4612cf2013957f1d818b6e74da7c803b9df1d394&token=1477304797&lang=zh_CN#rd)

* [Go常见错误第7篇：不使用-race选项做并发竞争检测](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484299&idx=1&sn=583c3470a76e93b0af0d5fc04fe29b55&chksm=ce124ce4f965c5f20de5887b113eab91f7c2654a941491a789e4ac53c298fbadb4367acee9bb&token=1918756920&lang=zh_CN#rd)

* [Go常见错误第8篇：并发编程中Context使用常见错误](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484317&idx=1&sn=474dad373684979fc96ba59182f08cf5&chksm=ce124cf2f965c5e4a29e313b4654faacef03e78da7aaf2ba6912d7b490a1df851a1bcbfec1c9&token=1918756920&lang=zh_CN#rd)

* [Go常见错误第9篇：使用文件名称作为函数输入](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484325&idx=1&sn=689c1b3823697cc583e1e818c4c76ee5&chksm=ce124ccaf965c5dce4e497f6251c5f0a8473b8e2ae3824bd72fe8c532d6dd84e6375c3990b3e&token=1266762504&lang=zh_CN#rd)

* [Go常见错误第10篇：Goroutine和循环变量一起使用的坑](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484335&idx=1&sn=cc8c6ceae72b30ec6f4d4e7b4367baca&chksm=ce124cc0f965c5d60410f977cdf31f127694fd0d49c35e2061ce8fb5fb9387bfa321196db438&token=1656737387&lang=zh_CN#rd)

* [Go常见错误第11篇：意外的变量遮蔽(variable shadowing)](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484519&idx=1&sn=00f13bdb95bd8f7c4eb1582a4c981991&chksm=ce124b08f965c21e28ea3a1c67b3b501fe45ba1a2c4af5048cad3a2d5d8303b6fc0192235b6d&token=1762934632&lang=zh_CN#rd)

* [Go常见错误第12篇：如何破解箭头型代码](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484539&idx=1&sn=189ff2e8fdb4a7d18620f9367128d6c4&chksm=ce124b14f965c20269c766a0f18f98cc8b9e340317669034a074d39843f2f9a94f3c9caf18da&token=329552886&lang=zh_CN#rd)

* [Go常见错误第13篇：init函数的常见错误和最佳实践](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484553&idx=1&sn=a4de11c452157193ae4381ab3555c42c&chksm=ce124be6f965c2f0885b8acf21c867e82d09807eba7be7890c54c01acf2b6fc413267e90d371&token=2029492652&lang=zh_CN#rd)

  

## 开源地址

文章和示例代码开源在GitHub: [Go语言初级、中级和高级教程](https://github.com/jincheng9/go-tutorial)。

公众号：coding进阶。关注公众号可以获取最新Go面试题和技术栈。

个人网站：[Jincheng's Blog](https://jincheng9.github.io/)。

知乎：[无忌](https://www.zhihu.com/people/thucuhkwuji)。



## 福利

我为大家整理了一份后端开发学习资料礼包，包含编程语言入门到进阶知识(Go、C++、Python)、后端开发技术栈、面试题等。

关注公众号「coding进阶」，发送消息 **backend** 领取资料礼包，这份资料会不定期更新，加入我觉得有价值的资料。

发送消息「**进群**」，和同行一起交流学习，答疑解惑。



## References

* https://livebook.manning.com/book/100-go-mistakes-how-to-avoid-them/chapter-2/
* https://github.com/jincheng9/go-tutorial/tree/main/workspace/lesson12