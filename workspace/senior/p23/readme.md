# Fuzzing: 一文读懂Go Fuzzing设计和原理

## 背景

Go 1.18除了引入泛型(generics)这个重大设计之外，Go官方团队在Go 1.18工具链里还引入了fuzzing模糊测试。

Go fuzzing的主要开发者是Katie Hockman, Jay Conrod和Roland Shoemaker。

**编者注**：Katie Hockman已于2022.02.19从Google离职，Jay Conrod也于2021年10月离开Google。



##  什么是Fuzzing

Fuzzing中文含义是模糊测试，是一种自动化测试技术，可以随机生成测试数据集，然后调用要测试的功能代码来检查功能是否符合预期。

模糊测试(fuzz test)是对单元测试(unit test)的补充，并不是要替代单元测试。

单元测试是检查指定的输入得到的结果是否和预期的输出结果一致，测试数据集比较有限。

模糊测试可以生成随机测试数据，找出单元测试覆盖不到的场景，进而发现程序的潜在bug和安全漏洞。



## Go Fuzzing怎么使用

Fuzzing在Go语言里并不是一个全新的概念，在Go官方团队发布Go Fuzzing之前，GitHub上已经有了类似的模糊测试工具[go-fuzz](https://github.com/dvyukov/go-fuzz)。

Go官方团队的Fuzzing实现借鉴了go-fuzz的设计思想。

Go 1.18把Fuzzing整合到了`go test`工具链和`testing`包里。

### 示例

下面举个例子说明下Fuzzing如何使用。

对于如下的字符串反转函数`Reverse`，大家可以思考下这段代码有什么潜在问题？

```go
// main.go
package fuzz

func Reverse(s string) string {
	bs := []byte(s)
	length := len(bs)
	for i := 0; i < length/2; i++ {
		bs[i], bs[length-i-1] = bs[length-i-1], bs[i]
	}
	return string(bs)
}
```

### 编写Fuzzing模糊测试函数

如果没有发现上面代码的bug，我们不妨来写一个Fuzzing模糊测试函数，来发现上面代码的潜在问题

```go
// fuzz_test.go
package fuzz

import (
	"testing"
	"unicode/utf8"
)

func FuzzReverse(f *testing.F) {
	str_slice := []string{"abc", "bb"}
	for _, v := range str_slice {
		f.Add(v)
	}
	f.Fuzz(func(t *testing.T, str string) {
		rev_str1 := Reverse(str)
		rev_str2 := Reverse(rev_str1)
		if str != rev_str2 {
			t.Errorf("fuzz test failed. str:%s, rev_str1:%s, rev_str2:%s", str, rev_str1, rev_str2)
		}
		if utf8.ValidString(str) && !utf8.ValidString(rev_str1) {
			t.Errorf("reverse result is not utf8. str:%s, len: %d, rev_str1:%s", str, len(str), rev_str1)
		}
	})
}
```

### 执行Fuzzing测试

执行如下命令可以进行Fuzzing测试

```sh
go1.18beta1 test -fuzz=Fuzz
```

使用的Go版本要求是`go 1.18beta 1`或以上版本，结果如下所示：

```sh
xxxMacBook-Air:fuzz$ go1.18beta1 test -fuzz=Fuzz
fuzz: elapsed: 0s, gathering baseline coverage: 0/113 completed
failure while testing seed corpus entry: FuzzReverse/ce9e8c80e2c2de2c96ab9e63b1a8cf18cea932b7d8c6c9c207d5978e0f19027a
fuzz: elapsed: 0s, gathering baseline coverage: 3/113 completed
--- FAIL: FuzzReverse (0.04s)
    --- FAIL: FuzzReverse (0.00s)
        fuzz_test.go:20: reverse result is not utf8. str:æ, len: 2, rev_str1:��
    
FAIL
exit status 1
FAIL    example/fuzz    0.253s
```

Fuzzing帮助我们发现了字符串反转函数`Reverse`的bug。



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
* GitHub开源工具go-fuzz: https://github.com/dvyukov/go-fuz
* Go fuzzing找bug示例：https://julien.ponge.org/blog/playing-with-test-fuzzing-in-go/

