# Go官方推出了Go 1.18的2个新教程

## 前言

2022年1月14日，Go官方团队的Katie Hockman在Go官方博客网站上发表了一篇新文章，主要介绍了Go 1.18的2个新教程，涉及Go泛型和Go Fuzzing。

本人针对Katie Hockman的原文做了一个翻译，以飨读者。

**同时在本文最后，附上了对Go泛型官方教程的中文翻译，以及针对Go泛型的设计思想和最佳实践。**



## 原文翻译

谷歌Go团队Katie Hockman

2022.1.14

我们很快就会发布Go 1.18版本，这个版本会引入一些新的概念。我们已经发布了2个官方教程来帮助大家熟悉这些新的feature。

第一篇教程是帮助大家熟悉Go泛型，地址[https://go.dev/doc/tutorial/generics](https://go.dev/doc/tutorial/generics)。这个教程会带着大家一步一步实现一个能处理多个类型的泛型函数，并且在代码里调用泛型函数。一旦你实现了泛型函数，你就会学到关于类型约束(type constraint)的知识，并且在你的函数里用到它。同时，也建议大家查阅最近GopherCon上Robert Greisemer和Ian Lance Taylor关于泛型的[技术分享](https://www.youtube.com/watch?v=35eIxI_n5ZM&t=1755s)，可以学到更多关于Go泛型的知识。

第二篇教程是关于Go fuzzing的介绍，地址[https://go.dev/doc/tutorial/fuzz](https://go.dev/doc/tutorial/fuzz)。这个教程展示了如何利用Fuzzing来帮助查找代码里的bug，带你一起利用Fuzzing来诊断和修复代码问题。同时，你也会在这个教程里写一些有bug的代码，利用Fuzzing来发现，修复和验证这些bug。特别感谢Beth Brown写了这篇教程。

Go 1.18 Beta 1版本上个月已经发布了，大家可以从[官方下载地址](https://go.dev/dl/#go1.18beta1)进行下载。

大家也可以查看Go 1.18完整的[发布清单](https://tip.golang.org/doc/go1.18)。

和以前一样，如果你发现了任何问题，请在[GitHub](https://github.com/golang/go/issues/new/choose)上提issue。

希望大家能喜欢这2个教程，我们期待2022年有更多美好事情的发生。



## 后记

本人针对Go泛型写了2篇通俗易懂的入门文章，一个是官方英文教程的中文翻译，一个是本人整理的Go泛型设计思想和使用场景解析，建议感兴趣的可以重点参考。

* Go 泛型官方教程中文版本：[官方教程：Go泛型入门](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483720&idx=1&sn=57ec4877dfd364a59deacf1e74a4fb66&chksm=ce124e27f965c731432dcc89d1e0563cf84baaef482eaa068a91bee61f10cf85b433923b83b4&token=802267677&lang=zh_CN#rd)
* Go 泛型设计思想和最佳实践解析：[一文读懂Go泛型设计和使用场景](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483731&idx=1&sn=b2258b28e2f3c16b065a5a1b22c15b0d&chksm=ce124e3cf965c72a6a22e0ed15deda8238567407bbd7157a79753fc8b605727ab2153009493c&token=802267677&lang=zh_CN#rd)

近期，我也会针对Go Fuzzing写一篇技术分享文章，欢迎大家关注。



## 开源地址

GitHub: https://github.com/jincheng9/go-tutorial

公众号：coding进阶

个人网站：https://jincheng9.github.io/



## References

* https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p6

* https://go.dev/doc/tutorial/generics
* https://jincheng9.github.io/post/go-generics-best-practice/

* https://www.youtube.com/watch?v=35eIxI_n5ZM&t=1755s