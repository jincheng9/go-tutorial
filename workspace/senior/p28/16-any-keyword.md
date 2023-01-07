# Go常见错误第16篇：any的常见错误和最佳实践

## 前言

这是Go常见错误系列的第16篇：any的常见错误和最佳实践。

素材来源于Go布道者，现Docker公司资深工程师[Teiva Harsanyi](https://teivah.github.io/)。

本文涉及的源代码全部开源在：[Go常见错误源代码](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p28)，欢迎大家关注公众号，及时获取本系列最新更新。

 

## 常见错误和最佳实践

Go语言中，没有方法的接口类型是空接口，也就是大家熟知的`interface{}`。

从Go 1.18开始，定义了一个预声明标识符(Predeclared identifiers)：any。

any实际上是空接口的别名，所以任何用了`interface{}`的地方都可以把`interface{}` 替换为any。

```go
func main() {
    var i any
 
    i = 42
    i = "foo"
    i = struct {
        s string
    }{
        s: "bar",
    }
    i = f
 
    _ = i
}
 
func f() {}
```

很多场景里，如果直接用any，会带来代码的过度抽象。

Rob Pike在[Gopherfest 2015](https://www.youtube.com/watch?v=PAAkCSZUG1c&t=7m36s)上，曾经分享过他的观点：

> interface{} says nothing.



### 常见错误

给any类型的变量赋值的时候，我们其实失去了所有类型信息，需要依赖类型断言(type assertion)来获取类型信息。

[类型断言](https://go.dev/tour/methods/15)即`t, ok := i.(T)`，代码示例如下所示：

```go
package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}
```

我们看看下面这个例子，体会下使用any带来的问题。

```go
package store
 
type Customer struct{
    // Some fields
}
type Contract struct{
    // Some fields
}
 
type Store struct{}
 
func (s *Store) Get(id string) (any, error) {
    // ...
}
 
func (s *Store) Set(id string, v any) error {
    // ...
}
```

这段代码里，我们定义了一个Store结构体，这个结构体有2个方法Get和Set，可以用来设置和获取Customer和Contract这2个结构体类型的变量。

示例代码计划用Get和Set方法来设置和查询Customer结构体和Contract结构体。

Get和Set方法虽然只存储Customer和Contrac这2个结构体类型，但是使用了any作为方法参数和方法返回值类型。

如果一个开发者只是看到函数签名里参数和返回值为any，会很容易误以为可以存储和查询任何类型的变量，比如int。

但实际上Get和Set方法的实现只是为了服务于Customer和Contract结构体。

这就是any类型带来的问题，因为隐藏了类型信息，开发者看到any要特别留意，仔细阅读代码和文档，才能避免出错。

比如有的开发者可能出现以下误用的情况：

```go
s := store.Store{}
s.Set("foo", 42)
```

Set的第2个参数虽然是any，但实际是要存储Customer和Contract类型的结构体，但由于参数为any，有的开发者可能就会直接存一个int类型，那就和代码设计的预期不符。

使用any会丢失类型信息，Go作为静态类型语言的优势就被影响了。

如下代码，其实是更好的设计。看代码一目了然，很方便知道每个方法存储和查询的是什么类型的结构体。

```go
func (s *Store) GetContract(id string) (Contract, error) {
    // ...
}
 
func (s *Store) SetContract(id string, contract Contract) error {
    // ...
}
 
func (s *Store) GetCustomer(id string) (Customer, error) {
    // ...
}
 
func (s *Store) SetCustomer(id string, customer Customer) error {
    // ...
}
```

### 最佳实践

any当然也不是一无是处，我们来看看Go标准库里的以下几个关于any的使用场景。

* 第一个例子是encoding/json这个package里的Marshal函数。

  因为Marshal函数可以操作任何类型，所以参数类型为any。

  ```go
  func Marshal(v any) ([]byte, error) {
      // ...
  }
  ```

* 第二个例子是database/sql包里的QueryContext方法。

  如果这个方法的第2个query参数是格式化字符串，比如`SELECT * FROM FOO WHERE id = ?`，由于查询语句里`?`的值其实可以是任何类型，所以后面的args参数是any类型。

  ```go
  func (c *Conn) QueryContext(ctx context.Context, query string,
      args ...any) (*Rows, error) {
      // ...
  }
  ```

## 总结

如果具体的使用场景的确是适合任意类型，那可以使用any。

但通常而言，为了代码的易读性和可维护性，我们应该避免过度抽象我们的代码



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

* [Go常见错误第14篇：过度使用getter和setter方法 ](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484568&idx=1&sn=2f078aa6561093691b4aeae58de44830&chksm=ce124bf7f965c2e17b99896393dfb684d5868156ea4e0b324bc86e2dcc596ea55d5e6c562c29&token=1629431500&lang=zh_CN#rd)

* [Go常见错误第15篇：interface使用的常见错误和最佳实践](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484589&idx=1&sn=0e0e71b71e6f236349aff55fe3b32b9e&chksm=ce124bc2f965c2d44b6f5bf74fd704db1417f8af2a584ce009a683f7ddae5678a95b09c8388f&token=1258925621&lang=zh_CN#rd)

  

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
* https://go.dev/ref/spec#Predeclared_identifiers
* https://go101.org/article/keywords-and-identifiers.html#:~:text=Keywords%20are%20the%20special%20words,understand%20and%20parse%20user%20code.&text=They%20can%20be%20categorized%20as,code%20elements%20in%20Go%20programs.