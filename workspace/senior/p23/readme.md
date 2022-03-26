# Fuzzing: 一文读懂Go Fuzzing设计和原理

## 背景

Go 1.18除了引入泛型(generics)这个重大设计之外，Go官方团队在Go 1.18工具链里还引入了fuzzing模糊测试。

Go fuzzing的主要开发者是Katie Hockman, Jay Conrod和Roland Shoemaker。

**编者注**：Katie Hockman已于2022.02.19从Google离职，Jay Conrod也于2021年10月离开Google



##  什么是Fuzzing

Fuzzing中文含义是模糊测试，是一种自动化测试技术，可以随机生成测试数据集，然后调用要测试的功能代码来检查功能是否符合预期。

模糊测试(fuzz test)是对单元测试(unit test)的补充，并不是要替代单元测试。

单元测试是检查指定的输入得到的结果是否和预期的输出结果一致，测试数据集比较有限。

模糊测试可以生成随机测试数据，找出单元测试覆盖不到的场景，进而发现程序的潜在bug和安全漏洞。



## Go Fuzzing怎么使用

Fuzzing在Go语言里并不是一个全新的概念，在Go官方团队发布Fuzzing工具之前，GitHub上已经有了类似的模糊测试工具[go-fuzz](https://github.com/dvyukov/go-fuzz)。

Go官方团队的Fuzzing实现借鉴了go-fuzz的设计思想。

Go 1.18把Fuzzing整合到了`go test`工具链和`testing`包里。

下面举个例子说明下Fuzzing如何使用。



## Go Fuzzing的实现

Coordinator, workers

```bash
MacBook-Air:go-tutorial $ ps aux | grep fuzz
xxx    13913  84.3  1.0  5219184  85124 s001  R+   10:12下午   0:03.90 /var/folders/pv/_x849j6n22x37xxd9cstgwkr0000gn/T/go-build1953131131/b001/fuzz.test -test.fuzzworker -test.paniconexit0 -test.fuzzcachedir=/Users/xxx/Library/Caches/go-build/fuzz/example/fuzz -test.timeout=10m0s -test.fuzz=Fuzz
xxx    13910  81.9  1.0  5221180  86200 s001  R+   10:12下午   0:03.94 /var/folders/pv/_x849j6n22x37xxd9cstgwkr0000gn/T/go-build1953131131/b001/fuzz.test -test.fuzzworker -test.paniconexit0 -test.fuzzcachedir=/Users/xxx/Library/Caches/go-build/fuzz/example/fuzz -test.timeout=10m0s -test.fuzz=Fuzz
xxx    13912  78.3  1.0  5219964  84984 s001  R+   10:12下午   0:03.86 /var/folders/pv/_x849j6n22x37xxd9cstgwkr0000gn/T/go-build1953131131/b001/fuzz.test -test.fuzzworker -test.paniconexit0 -test.fuzzcachedir=/Users/xxx/Library/Caches/go-build/fuzz/example/fuzz -test.timeout=10m0s -test.fuzz=Fuzz
xxx    13911  74.5  1.0  5219184  85132 s001  R+   10:12下午   0:03.76 /var/folders/pv/_x849j6n22x37xxd9cstgwkr0000gn/T/go-build1953131131/b001/fuzz.test -test.fuzzworker -test.paniconexit0 -test.fuzzcachedir=/Users/xxx/Library/Caches/go-build/fuzz/example/fuzz -test.timeout=10m0s -test.fuzz=Fuzz
xxx    13907  43.3  2.3  5944576 191172 s001  R+   10:12下午   0:01.90 /var/folders/pv/_x849j6n22x37xxd9cstgwkr0000gn/T/go-build1953131131/b001/fuzz.test -test.paniconexit0 -test.fuzzcachedir=/Users/xxx/Library/Caches/go-build/fuzz/example/fuzz -test.timeout=10m0s -test.fuzz=Fuzz
xxx    13923   0.0  0.0  4268176    420 s000  R+   10:12下午   0:00.00 grep fuzz
xxx    13891   0.0  0.2  5014396  16868 s001  S+   10:12下午   0:00.52 /Users/xxx/sdk/go1.18beta2/bin/go test -fuzz=Fuzz
xxx    13890   0.0  0.0  4989312   4008 s001  S+   10:12下午   0:00.01 go1.18beta2 test -fuzz=Fuzz
```





## 开源地址

文章和示例代码开源在GitHub: [Go语言初级、中级和高级教程](https://github.com/jincheng9/go-tutorial)。

公众号：coding进阶。关注公众号可以获取最新Go面试题和技术栈。

个人网站：[Jincheng's Blog](https://jincheng9.github.io/)。

知乎：[无忌](https://www.zhihu.com/people/thucuhkwuji)。



## References

* Internals of Go's New Fuzzing System: https://jayconrod.com/posts/123/internals-of-go-s-new-fuzzing-system
* Fuzzing介绍：https://go.dev/doc/fuzz/
* Fuzzing Design Draft: https://go.googlesource.com/proposal/+/master/design/draft-fuzzing.md
* Fuzzing提案：https://github.com/golang/go/issues/44551
* Fuzzing教程：https://go.dev/doc/tutorial/fuzz
* Fuzzing Tesing in Go in 8 Minutes: https://www.youtube.com/watch?v=w8STTZWdG9Y
* GitHub开源工具go-fuzz: https://github.com/dvyukov/go-fuzz

