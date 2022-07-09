# Go 1.19要来了，看看都有哪些变化-第1篇

## 前言

Go官方团队在2022.06.11发布了Go 1.19 Beta 1版本，Go 1.19的正式release版本预计会在今年8月份发布。

让我们先睹为快，看看Go 1.19给我们带来了哪些变化。

这是Go 1.19版本更新内容详解的第1篇，欢迎大家关注公众号，及时获取本系列最新更新。

## Go 1.19发布清单

和Go 1.18相比，改动相对较小，主要涉及语言(Language)、内存模型(Memory Model)、可移植性(Ports)、Go Tool工具链、运行时(Runtime)、编译器(Compiler)、汇编器(Assembler)、链接器(Linker)和核心库(Core library)等方面的优化。

我们逐个看看具体都有哪些变化。

### 语言变化

语言层面的变化很小，只有一个和泛型相关的优化，调整了泛型函数和方法声明里的类型参数(type parameter)的作用域，不影响现有代码。

为什么会有这个调整呢？我们来看看如下的代码

```go
type T[T any] struct {
	m1 T
}
```

这段代码定义了泛型类型`T` ，类型参数列表(type parameter list)里的类型参数(type parameter)命名也是`T`，代码可以正常编译通过。

但是如下代码呢？

```go
type T[T any] struct {
	m1 T
}

func(t T[T]) print() {
	fmt.Println(t.m1)
}
```

我们定义了泛型类型`T` 的方法`print`，编译上面的代码，编译器会提示：

```bash
./main.go:11:8: T is not a generic type
```

也就是说编译器认为方法`print`里的第一个`T`不是泛型类型，为什么会这样呢？我们来看看[官方说明](https://go.dev/ref/spec#Declarations_and_scope)：

> The scope of an identifier denoting a type parameter of a function or declared by a method receiver is the function body and all parameter lists of the function.

这段说明对应到上面的代码，编译器会认为`print`方法里的第2个`T`(类型参数`T`)的作用域是函数体以及函数的参数列表(这里的参数列表包括方法receiver parameter list和函数名后面的参数列表)。因此第2个`T`就覆盖了第1个`T` ，所以编译器会提示`T is not a generic type`。

[这个问题](https://github.com/golang/go/issues/52038)是由Go101作者提出来的，Go泛型的设计者[Robert Griesemer](https://github.com/griesemer)认领了这个问题，把类型参数的作用域按照如下说明进行调整：

> The scope of an identifier denoting a type parameter of a function or declared by a method receiver *starts after the function name and ends at the end of the function body*.

也就是类型参数的作用域是从函数名后面开始一直到函数体，这样上面的代码里的2个`T` 的就不会覆盖第1个`T` ，编译器也就不会编译报错了。

**想了解Go泛型的使用方法、设计思路和最佳实践，推荐大家阅读**：

* [官方教程：Go泛型入门](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483720&idx=1&sn=57ec4877dfd364a59deacf1e74a4fb66&chksm=ce124e27f965c731432dcc89d1e0563cf84baaef482eaa068a91bee61f10cf85b433923b83b4&token=1782465473&lang=zh_CN#rd)
* [一文读懂Go泛型设计和使用场景](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483731&idx=1&sn=b2258b28e2f3c16b065a5a1b22c15b0d&chksm=ce124e3cf965c72a6a22e0ed15deda8238567407bbd7157a79753fc8b605727ab2153009493c&token=1782465473&lang=zh_CN#rd)
* [重磅：Go 1.18将移除用于泛型的constraints包](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483855&idx=1&sn=6ab4aeb140a1a08268dc8a0284a6f375&chksm=ce124ea0f965c7b6776061960d71e4ffb30484a82041f5b1d4786c4b49c4ffabc07a28b1cd48&token=1782465473&lang=zh_CN#rd)
* [泛型最佳实践：Go泛型设计者教你如何用泛型](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484015&idx=1&sn=576b2d8b84b3a8ce5bdd6952c2b84062&chksm=ce124d00f965c416b07dcb81c4dcb9cf75859b2787d4f00ec8c80b37ca42e58cc651420a3b33&token=1782465473&lang=zh_CN#rd)

### 内存模型

[Go内存模型](https://tip.golang.org/ref/mem) 在Go 1.19版本做了修改，和C, C++, Java, JavaScript, Rust和Swift对齐。Go只提供顺序一致原子操作(sequentially consistent atomics)，不提供其它语言里更宽松的内存模型，比如因果一致性(casual consistency)、最终一致性(eventual consistency)。

伴随着内存模型的修改，Go 1.19版本在`sync/atomic`包里引入了新的类型，例如[atomic.Int64](https://pkg.go.dev/sync/atomic@master#Int64)和[atomic.Pointer[T]](https://pkg.go.dev/sync/atomic@master#Pointer)，这些新的类型可以让开发者使用原子操作时更容易。

**想了解Go原子操作和使用方法，推荐大家阅读**：

* [Go并发编程之原子操作sync/atomic](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484082&idx=1&sn=934787c9829391ba743bd611818ad0e2&chksm=ce124dddf965c4cb7d0f2d9d001ab4b7d949fbe87c4c8b7ee8d7498946824ec9aa6581cfe986&token=1782465473&lang=zh_CN#rd)

### 可移植性

Go 1.19版本支持Linux操作系统上的中国龙芯64位CPU架构。

```bash
GOOS=linux
GOARCH=loong64
```

此外，Go现在支持在`riscv64`架构上使用寄存器来传递函数参数和函数执行结果。性能测试表明，`riscv64`架构上的Go语言性能提升了大概10%。

## 总结

下一篇会介绍Go 1.19在Go Tool工具链、运行时、编译器、汇编器、链接器和核心库的优化工作，有一些内容值得学习，欢迎大家保持关注。



## 开源地址

文章和示例代码开源在GitHub: [Go语言初级、中级和高级教程](https://github.com/jincheng9/go-tutorial)。

公众号：coding进阶。关注公众号可以获取最新Go面试题和技术栈。

个人网站：[Jincheng's Blog](https://jincheng9.github.io/)。

知乎：[无忌](https://www.zhihu.com/people/thucuhkwuji)。



## 福利

我为大家整理了一份后端开发学习资料礼包，包含编程语言入门到进阶知识(Go、C++、Python)、后端开发技术栈、面试题等。

关注公众号「coding进阶」，发送消息 **backend** 领取资料礼包，这份资料会不定期更新，加入我觉得有价值的资料。还可以发送消息「**进群**」，和同行一起交流学习，答疑解惑。



## References

* https://tip.golang.org/doc/go1.19
* https://int64.me/2020/%E4%B8%80%E8%87%B4%E6%80%A7%E6%A8%A1%E5%9E%8B%E7%AC%94%E8%AE%B0.html