# 官方博文：Go 1.18 Beta 2发布

## 前言

2022年1月31日，Go官方团队的Jeremy Faller和Steve Francia在Go官方博客网站上发表了一篇新文章，宣布Go官方正式推出了Go 1.18 Beta 2版本。

同时，支持Go泛型的正式release版本也从原计划的2022年2月份推迟到2022年3月份。

本人针对官方原文做了一个翻译，以飨读者。

**同时在本文最后，附上了对Go泛型官方教程的中文翻译，以及针对Go泛型的设计思想和最佳实践。**



## 原文翻译

谷歌Go团队Jeremy Faller和Steve Francia

2022.1.31

Go 1.18 release版本会新增对泛型、fuzzing自动化测试和新的工作区模式(workspace mode)的支持，整个Go社区对此都感到兴奋，我们作为Go官方团队的成员也深受鼓舞。

2个月前，我们发布了Go 1.18 beta 1版本，这是有史以来下载量最多的Go beta版本，下载量是之前任何版本的2倍多。Beta 1版本被证明非常稳定可靠，实际上我们已经在Google的生产环境上正式使用了Go 1.18 Beta 1版本。

你们对Beta 1版本的反馈帮助我们发现了Go泛型里隐藏的bug，确保了一个更加稳定的最终版本。

我们已经在Go 1.18 Beta 2里解决了这些问题，我们希望大家都可以去尝试使用。

安装Go 1.18 Beta 2最简单的方式是运行如下命令：

```bash
go install golang.org/dl/go1.18beta2@latest
go1.18beta2 download
```

执行上面的命令后，你可以使用`go1.18beta2`来代替`go`命令。

更多的下载选项，可以访问[https://go.dev/dl/#go1.18beta2](https://go.dev/dl/#go1.18beta2)。

因为我们花了一些时间去发布Go 1.18的第2个beta版本，因此我们现在规划是在2月份发布Go 1.18的候选版本，并在3月份发布最终的Go 1.18正式版本。

此外，Go语言服务器`gopls`和VS Code的Go扩展插件现在也支持Go泛型了。

安装支持泛型的`gopls`，可以参考这篇文档：https://github.com/golang/tools/blob/master/gopls/doc/advanced.md#working-with-generic-code。

配置VS Code的Go插件，可以参考这篇说明：https://github.com/golang/vscode-go/blob/master/docs/advanced.md#using-go118

和以前一样，如果大家发现任何问题，尤其是beta版本的问题，请到GitHub上提交issue，提交地址：https://github.com/golang/go/issues/new/choose



## 后记

本人针对Go泛型写了2篇通俗易懂的入门文章，一个是官方英文教程的中文翻译，一个是本人整理的Go泛型设计思想和使用场景解析，建议感兴趣的可以重点参考。

* Go 泛型官方教程中文版本：[官方教程：Go泛型入门](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483720&idx=1&sn=57ec4877dfd364a59deacf1e74a4fb66&chksm=ce124e27f965c731432dcc89d1e0563cf84baaef482eaa068a91bee61f10cf85b433923b83b4&token=802267677&lang=zh_CN#rd)
* Go 泛型设计思想和最佳实践解析：[一文读懂Go泛型设计和使用场景](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483731&idx=1&sn=b2258b28e2f3c16b065a5a1b22c15b0d&chksm=ce124e3cf965c72a6a22e0ed15deda8238567407bbd7157a79753fc8b605727ab2153009493c&token=802267677&lang=zh_CN#rd)

近期，我也会针对Go Fuzzing写一篇技术分享文章，欢迎大家关注。



## 开源地址

GitHub: [https://github.com/jincheng9/go-tutorial](https://github.com/jincheng9/go-tutorial)

公众号：coding进阶，获取最新Go面试题和技术栈

个人网站：[https://jincheng9.github.io/](https://jincheng9.github.io/)

知乎：[https://www.zhihu.com/people/thucuhkwuji](https://www.zhihu.com/people/thucuhkwuji)



## References

* https://go.dev/blog/go1.18beta2
* https://go.dev/blog/fuzz-beta
* https://go.googlesource.com/proposal/+/master/design/45713-workspace.md