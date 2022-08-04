# 详解Go 1.19的新变化

## 前言

Go官方团队在2022.06.11发布了Go 1.19 Beta 1版本，Go 1.19的正式release版本在2022.08.02正式发布。

让我们先睹为快，看看Go 1.19给我们带来了哪些变化。



## Go 1.19发布清单

和Go 1.18相比，改动相对较小，主要涉及语言(Language)、内存模型(Memory Model)、可移植性(Ports)、Go Tool工具链、运行时(Runtime)、编译器(Compiler)、汇编器(Assembler)、链接器(Linker)和核心库(Core library)等方面的优化。

* [Go 1.19要来了，看看都有哪些变化-第1篇](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484179&idx=1&sn=215ea3f092460118b2bc975935015874&chksm=ce124c7cf965c56a7c310b1059683d065810bd18368669d3d42a6cbbb0370d1593979a63620c&token=1899277735&lang=zh_CN#rd)
* [Go 1.19要来了，看看都有哪些变化-第2篇](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484188&idx=1&sn=c14bafb1f89b3b3f988452c5a5f32884&chksm=ce124c73f965c5651a688c42561b02e38253b60943c77a0a6ad7b45621b4296d9e1acd47de7a&token=1899277735&lang=zh_CN#rd)
* [Go 1.19要来了，看看都有哪些变化-第3篇](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484203&idx=1&sn=fcf95ce045a54c2f6a9414d1e9fa732d&chksm=ce124c44f965c5525fc820f99978cf0996d5f3653e4e580302986f14c9ae740d499ba9fb18c4&token=1899277735&lang=zh_CN#rd)
* [Go 1.19要来了，看看都有哪些变化-第4篇](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484210&idx=1&sn=916491404f8a4acd0057f299b77e47de&chksm=ce124c5df965c54b5d21bb00e74fd256d163597d6a0f5cf8c40891e77dafcf719b2e7b7beee4&token=1899277735&lang=zh_CN#rd)



## 开源地址

文章和示例代码开源在GitHub: [Go语言初级、中级和高级教程](https://github.com/jincheng9/go-tutorial)。

公众号：coding进阶。关注公众号可以获取最新Go面试题和技术栈。

个人网站：[Jincheng's Blog](https://jincheng9.github.io/)。

知乎：[无忌](https://www.zhihu.com/people/thucuhkwuji)。



## 福利

我为大家整理了一份后端开发学习资料礼包，包含编程语言入门到进阶知识(Go、C++、Python)、后端开发技术栈、面试题等。

关注公众号「coding进阶」，发送消息 **backend** 领取资料礼包，这份资料会不定期更新，加入我觉得有价值的资料。还可以发送消息「**进群**」，和同行一起交流学习，答疑解惑。



## References

* https://go.dev/blog/go1.19