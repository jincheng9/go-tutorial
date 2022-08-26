# Go十大常见错误第7篇：不使用-race选项做并发竞争检测

## 前言

这是Go十大常见错误系列的第7篇：不使用`-race`选项做并发竞争检测。素材来源于Go布道者，现Docker公司资深工程师[Teiva Harsanyi](https://teivah.medium.com/)。

本文涉及的源代码全部开源在：[Go十大常见错误源代码](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p28)，欢迎大家关注公众号，及时获取本系列最新更新。



## 背景

并发编程里很容易遇到并发访问冲突的问题。

Go语言里多个goroutine同时操作某个共享变量的时候，如果一个goroutine对该变量做写操作，其它goroutine做读操作，假设没有做好并发访问控制，就容易出现并发访问冲突，导致程序crash。

大家可以看如下的代码示例：

```go
package main

import (
	"fmt"
)

func main() {
	c := make(chan bool)
	m := make(map[string]string)
	go func() {
		m["1"] = "a" // First conflicting access.
		c <- true
	}()
	m["2"] = "b" // Second conflicting access.
	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}
}
```

上面的代码会出现并发访问冲突，2个goroutine同时对共享变量`m`做写操作。

通过`-race`选项，我们就可以利用编译器帮我们快速发现问题。

```bash
$ go run -race main.go 
==================
WARNING: DATA RACE
Write at 0x00c000074180 by goroutine 7:
  runtime.mapassign_faststr()
      /usr/local/opt/go/libexec/src/runtime/map_faststr.go:202 +0x0
  main.main.func1()
      /Users/xxx/github/go-tutorial/workspace/senior/p28/data-race/main.go:11 +0x5d

Previous write at 0x00c000074180 by main goroutine:
  runtime.mapassign_faststr()
      /usr/local/opt/go/libexec/src/runtime/map_faststr.go:202 +0x0
  main.main()
      /Users/xxx/github/go-tutorial/workspace/senior/p28/data-race/main.go:14 +0xcb

Goroutine 7 (running) created at:
  main.main()
      /Users/xxx/github/go-tutorial/workspace/senior/p28/data-race/main.go:10 +0x9c
==================
1 a
2 b
Found 1 data race(s)
exit status 66
```



## 常见问题

一个常见的错误是开发者测试Go程序的时候，不使用`-race`选项。

尽管Go语言设计的目的之一是为了让并发编程更简单、更不容易出错，但Go语言开发者还是会遇到并发问题。

因此，大家在测试Go程序的时候，应该开启`-race`选项，及时发现代码里的并发访问冲突问题。

```bash
$ go test -race mypkg    // to test the package
$ go run -race mysrc.go  // to run the source file
$ go build -race mycmd   // to build the command
$ go install -race mypkg // to install the package
```



## 推荐阅读

* [Go十大常见错误第1篇：未知枚举值](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484146&idx=1&sn=10fb12b643a2e37c090e5aa3bc583152&chksm=ce124d9df965c48bb954aeddabdff3db12738ded3875542250c5d0ef6cfd4417fc56580288b1&token=1912894792&lang=zh_CN#rd)

* [Go十大常见错误第2篇：benchmark性能测试的坑](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484163&idx=1&sn=b28d61c1f3ec9d914e698dce105ba5d1&chksm=ce124c6cf965c57a90bc85a5295ed9375103de20607b509f845583ff6686385df0ed96653d00&token=1912894792&lang=zh_CN#rd)

* [Go十大常见错误第3篇：go指针的性能问题和内存逃逸](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484247&idx=1&sn=faf716627afb00df646cecff023fb63c&chksm=ce124c38f965c52efd009a4c98691d56b5765dc7dce98aa49b226ad9274bd062d8d01e702e91&token=1899277735&lang=zh_CN#rd)

* [Go十大常见错误第4篇：break操作的注意事项](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484262&idx=1&sn=c1bea8af60444a4ef73c4d4d7a09d16d&chksm=ce124c09f965c51f3663ac9089a792d36c3685850e12695dd26d15a1a50f393b2d7c92b9983a&token=461369035&lang=zh_CN#rd)

* [Go十大常见错误第5篇：Go语言Error管理](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484274&idx=1&sn=711abea3c6fd5d15341ee1b34da8a160&chksm=ce124c1df965c50b3af84965f7ed30b574cd0b247ea6f77b944ec858bd43ee37f4c1554a5bce&token=1846351524&lang=zh_CN#rd)

* [Go十大常见错误第6篇：slice初始化常犯的错误](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484289&idx=1&sn=2b8171458cde4425b28fdf8f51df8d7c&chksm=ce124ceef965c5f8a14f5951457ce2ac0ecc4612cf2013957f1d818b6e74da7c803b9df1d394&token=1477304797&lang=zh_CN#rd)

* [Go面试题系列，看看你会几题？](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=Mzg2MTcwNjc1Mg==&action=getalbum&album_id=2199553588283179010#wechat_redirect)

* [Go编译器的race detector可以发现所有的并发冲突么？](https://medium.com/@val_deleplace/does-the-race-detector-catch-all-data-races-1afed51d57fb)

  

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
* https://go.dev/doc/articles/race_detector
* https://medium.com/@val_deleplace/does-the-race-detector-catch-all-data-races-1afed51d57fb