# Go十大常见错误第4篇：for/switch和for/select做break操作的注意事项

## 前言

这是Go十大常见错误系列的第4篇：for/switch和for/select做break操作退出的注意事项。素材来源于Go布道者，现Docker公司资深工程师[Teiva Harsanyi](https://teivah.medium.com/)。

本文涉及的源代码全部开源在：[Go十大常见错误源代码](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p28)，欢迎大家关注公众号，及时获取本系列最新更新。



## 场景

### 案例1

大家看看下面这段代码：

```go
for {
  switch f() {
  case true:
    break
  case false:
    // Do something
  }
}
```

如果函数调用`f()`返回的结果是`true`，进入到`case true`分支，会发生什么？会退出for循环么？

答案是：只退出了switch语句，并不会退出for循环，所以break后又继续执行for循环里的代码。

### 案例2

再看下面这段代码

```go
for {
  select {
  case <-ch:
  // Do something
  case <-ctx.Done():
    break
  }
}
```

同样地，如果执行了break语句，退出的只是select语句块，并不会退出for循环。

那在上面2种场景里，如何退出for循环呢？

可以结合label和break进行实现。

```go
loop:
	for {
		select {
		case <-ch:
		// Do something
		case <-ctx.Done():
			break loop
		}
	}
```

对于上面的代码，loop是一个label，`break loop`如果执行了就会退出for循环。



## 推荐阅读

* [Go十大常见错误第1篇：未知枚举值](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484146&idx=1&sn=10fb12b643a2e37c090e5aa3bc583152&chksm=ce124d9df965c48bb954aeddabdff3db12738ded3875542250c5d0ef6cfd4417fc56580288b1&token=1912894792&lang=zh_CN#rd)

* [Go十大常见错误第2篇：benchmark性能测试的坑](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484163&idx=1&sn=b28d61c1f3ec9d914e698dce105ba5d1&chksm=ce124c6cf965c57a90bc85a5295ed9375103de20607b509f845583ff6686385df0ed96653d00&token=1912894792&lang=zh_CN#rd)

* [Go十大常见错误第3篇：go指针的性能问题和内存逃逸](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484247&idx=1&sn=faf716627afb00df646cecff023fb63c&chksm=ce124c38f965c52efd009a4c98691d56b5765dc7dce98aa49b226ad9274bd062d8d01e702e91&token=1899277735&lang=zh_CN#rd)

* [Go switch使用说明](https://github.com/jincheng9/go-tutorial/tree/main/workspace/lesson6)

* [Go for/break使用说明](https://github.com/jincheng9/go-tutorial/tree/main/workspace/lesson7)

* [Go select语义](https://github.com/jincheng9/go-tutorial/tree/main/workspace/lesson29)

  



## 开源地址

文章和示例代码开源在GitHub: [Go语言初级、中级和高级教程](https://github.com/jincheng9/go-tutorial)。

公众号：coding进阶。关注公众号可以获取最新Go面试题和技术栈。

个人网站：[Jincheng's Blog](https://jincheng9.github.io/)。

知乎：[无忌](https://www.zhihu.com/people/thucuhkwuji)。



## 福利

我为大家整理了一份后端开发学习资料礼包，包含编程语言入门到进阶知识(Go、C++、Python)、后端开发技术栈、面试题等。

关注公众号「coding进阶」，发送消息 **backend** 领取资料礼包，这份资料会不定期更新，加入我觉得有价值的资料。还可以发送消息「**进群**」，和同行一起交流学习，答疑解惑。



## References

* https://itnext.io/the-top-10-most-common-mistakes-ive-seen-in-go-projects-4b79d4f6cd65
* https://github.com/jincheng9/go-tutorial/tree/main/workspace/lesson6
* https://github.com/jincheng9/go-tutorial/tree/main/workspace/lesson7