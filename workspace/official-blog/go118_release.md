# 官方博文：Go 1.18发布啦！

## 前言

2022年3月15日，Go官方团队在官网上正式宣布了Go 1.18版本的发布。

Go 1.18是Go语言诞生以来变化最大的版本，本人针对官方原文做了一个翻译，以飨读者。

**同时在本文最后，附上《Go泛型官方教程中文版本》以及本人整理的《一文读懂Go泛型设计和最佳实践》**。



## 原文翻译

谷歌Go团队

2022.3.15

今天我们很激动向大家宣布Go 1.18终于发布啦，大家可以去[下载页面](https://go.dev/dl/)进行下载。

Go 1.18版本发布了非常多的新功能、性能优化以及Go语言有史以来最重大的修改(编者注：最重大的修改指的是泛型)。

毫不夸张地说，Go 1.18的部分设计从十多年前我们第一次发布Go就开始了。

## 泛型

Go 1.18引入了对泛型的支持。泛型是Go社区最常被要求支持的功能，我们也很自豪地提供了可以满足大多数用户要求的泛型设计。

Go 1.18的后续版本会对一些更复杂的泛型使用场景提供支持。我们鼓励大家通过我们[官方泛型教程](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483720&idx=1&sn=57ec4877dfd364a59deacf1e74a4fb66&chksm=ce124e27f965c731432dcc89d1e0563cf84baaef482eaa068a91bee61f10cf85b433923b83b4&token=1183396486&lang=zh_CN#rd)来了解泛型以及探索使用泛型来优化代码的最佳实践。

Go 1.18[发布清单](https://go.dev/doc/go1.18)里有更多使用泛型的细节。

## Fuzzing模糊测试

Go 1.18版本还引入了fuzzing模糊测试。Go是第一个在语言的标准工具链中引入fuzzing模糊测试的主流编程语言。

和泛型类似，fuzzing的设计也经历了很多时间，我们很高兴终于在Go 1.18里发布了fuzzing。大家可以查看[官方fuzzing教程](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483931&idx=1&sn=41fc064855d6f2cba24944c0378fee24&chksm=ce124d74f965c46273746ed9054d7149b2eccabcac627ce56388fb5431b0b14b06e284a857b3&token=1183396486&lang=zh_CN#rd)来学习这个新功能。

## 工作区

Go module已经被广泛使用，在我们的年度调查中Go语言开发者对Go module也都表现了高度的满意。

在我们2021的年度用户调查中，大家使用module最常见的挑战就是跨多个Go module工作。

在Go 1.18版本里发布了[工作区模式](https://go.dev/doc/tutorial/workspaces)来解决这个问题，让大家跨多个module工作更容易。

## 20%性能提升

苹果M1，ARM64和PowerPC64用户会欣喜若狂。Go 1.18通过把Go 1.17的寄存器ABI调用约定扩展到这些新的CPU架构，实现了对CPU性能接近20%的提升。

为了强调这个版本的性能提升幅度，我们把`20%的性能提升`作为本文第4个最重要的标题。

关于Go 1.18发布的所有内容更为详细的描述，大家可以参考[Go 1.18发布清单](https://go.dev/doc/go1.18)。

Go 1.18对于整个Go社区而言是一个巨大的里程碑。我们想感谢所有提bug、提交修改、编写使用教程以及对Go 1.18的发布提供过帮助的所有人。没有你们我们无法到达现在这个阶段，感谢你们！

尽情享受Go 1.18吧！



## 开源地址

GitHub: [https://github.com/jincheng9/go-tutorial](https://github.com/jincheng9/go-tutorial)，涵盖Go语言初级、中级和高级实战教程。

公众号：coding进阶，获取最新Go面试题和技术栈

个人网站：[https://jincheng9.github.io/](https://jincheng9.github.io/)

知乎：[https://www.zhihu.com/people/thucuhkwuji](https://www.zhihu.com/people/thucuhkwuji)



## 推荐阅读

* [Go泛型官方教程](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483720&idx=1&sn=57ec4877dfd364a59deacf1e74a4fb66&chksm=ce124e27f965c731432dcc89d1e0563cf84baaef482eaa068a91bee61f10cf85b433923b83b4&token=1183396486&lang=zh_CN#rd)

* [Go泛型设计原理和最佳实践](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483731&idx=1&sn=b2258b28e2f3c16b065a5a1b22c15b0d&chksm=ce124e3cf965c72a6a22e0ed15deda8238567407bbd7157a79753fc8b605727ab2153009493c&token=1183396486&lang=zh_CN#rd)

* [Go Fuzzing官方教程](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483931&idx=1&sn=41fc064855d6f2cba24944c0378fee24&chksm=ce124d74f965c46273746ed9054d7149b2eccabcac627ce56388fb5431b0b14b06e284a857b3&token=1183396486&lang=zh_CN#rd)

  

## References

* https://go.dev/blog/go1.18
* https://go.dev/blog/why-generics
* https://go.dev/doc/go1.18
* https://go.dev/doc/tutorial/workspaces