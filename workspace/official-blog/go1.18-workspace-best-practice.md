# 官方博文：Go工作区模式最佳实践

## 前言

Go 1.18除了引入泛型(generics)、模糊测试(Fuzzing)之外，另外一个重大功能是引入了工作区模式(workspace mode)。

Go官方团队的*Beth Brown*于2022.04.05在官网上专门写了一篇博文，详细介绍了workspace模式的使用场景和最佳实践。

本人针对官方原文做了一个翻译，以飨读者。**同时在本文最后，附上了对workspace模式的入门介绍。**



## 原文翻译

Go官方团队*Beth Brown*

2022.04.05

Go 1.18新增了工作区模式(workspace mode)，让你可以同时跨多个Go Module进行开发。

你可以从[download](https://go.dev/dl/)地址下载Go 1.18，[release notes](https://go.dev/doc/go1.18)有更多关于Go 1.18的变化细节。



## 后记



## 开源地址

文章和示例代码开源在GitHub: [Go语言初级、中级和高级教程](https://github.com/jincheng9/go-tutorial)。

公众号：coding进阶。关注公众号可以获取最新Go面试题和技术栈。

个人网站：[Jincheng's Blog](https://jincheng9.github.io/)。

知乎：[无忌](https://www.zhihu.com/people/thucuhkwuji)。https://www.zhihu.com/people/thucuhkwuji)



## References

* https://go.dev/blog/get-familiar-with-workspaces