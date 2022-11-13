# Go语言开源13周年啦，看看负责人说了啥

## 前言

 不知不觉，Go语言已经开源13年了。

日前，Go团队负责人Russ Cox在官方博客上发表了Go开源13年的感想。

Russ首先回顾了2022年3月份发布的Go 1.18版本引入的工作区模式(Go Workspace)、模糊测试(Fuzzing)和泛型设计。

接着介绍了Go 1.19版本引入的新优化、Go开发工具链、Go供应链安全以及其它重要的里程碑事件。

本人针对Russ的原文做了一个归纳整理，以飨读者。



## 原文归纳整理内容

谷歌Go团队Russ Cox

2022.11.10

今天我们可以开心地庆祝Go语言开源13年啦。

### Go 1.18

对Go语言而言，今年有非常多的重要事项。最重要的当然是在今年3月份我们发布了Go 1.18版本，这个版本引入了非常多的新功能，包括大家熟知的Go workspace工作区模式、Go Fuzzing模糊测试和Go泛型。

* Go workspace工作区模式让开发者可以同时开发多个互相有依赖的Module。想了解Go工作区模式的，可以参考Beth Brown's的文章[Get familiar with workspaces](https://go.dev/blog/get-familiar-with-workspaces)以及 [workspace reference](https://go.dev/ref/mod#workspaces)。
* Fuzzing模糊测试是go test的一个新功能，可以生成随机的输入用于检测代码的正确性以及漏洞。想了解Go Fuzzing的可以参考 [Getting started with fuzzing](https://go.dev/doc/tutorial/fuzz)和 [fuzzing reference](https://go.dev/security/fuzz/)。另外，Fuzzing作者katie Hockman在GopherCon 2022会议上分享的主题“Fuzz Testing Made Easy”也即将上线，敬请关注。
* 泛型可能是Go语言里被开发者提及最多的feature。想了解Go泛型入门的，可以参考教程[Getting started with generics](https://go.dev/doc/tutorial/generics)。想了解Go泛型更多设计和使用细节的，可以参考官方博客[An Introduction to Generics](https://go.dev/blog/intro-generics) 和[When to Use Generics](https://go.dev/blog/when-generics)，以及Google 2021开源日的技术分享[Using Generics in Go](https://www.youtube.com/watch?v=nr8EpUO9jhw) 和GopherCon 2021上Go泛型作者Robert Griesemer和Ian Lance Taylor的技术分享[Generics!](https://www.youtube.com/watch?v=Pa_e9EeCdy8)。

### Go 1.19

和Go 1.18相比，今年8月份发布的Go 1.19版本相对修改少一些，吸引的关注也少一些。Go 1.19版本专注在继续优化Go 1.18引入的新特性。此外在Go 1.19版本中，我们支持在Go doc注释中添加超链接、列表以及标题，用于生成Go package的说明文档。另外我们针对Go的垃圾回收器(garbage collector)引入了软内存限制(soft memory limit)，这对于管理容器负载(container workloads)非常有用。关于GC更多的改进和优化，可以参考Michael Knyszek最近的博文[Go runtime: 4 years later](https://go.dev/blog/go119runtime)、视频[Respecting Memory Limits in Go](https://www.youtube.com/watch?v=07wduWyWx8M&list=PLtoVuM73AmsJjj5tnZ7BodjN_zIvpULSx) 和 [Guide to the Go Garbage Collector](https://go.dev/doc/gc-guide)。

### Go开发工具

Go团队还在维护2个Go语言开发工具VS Code Go扩展插件以及Gopls语言服务器。

今年Gopls专注于提升稳定性和性能、支持泛型以及提供更丰富的Go语言分析功能。

想了解VS Code Go和Gopls的最新改进，可以参考Suzy Mueller的技术分享[Building Better Projects with the Go Editor](https://www.youtube.com/watch?v=jMyzsp2E_0U)、[Debugging Treasure Hunt](https://www.youtube.com/watch?v=ZPIPPRjwg7Q)和官方技术博客[Debugging Go in VS Code](https://go.dev/s/vscode-go-debug)。

### Go供应链安全

Go语言开源12周年后的一个月左右，爆发了全球轰动的[Log4shell vulnerability](https://en.wikipedia.org/wiki/Log4Shell)安全漏洞事件，给大家敲醒了警钟，在做开发过程中需要考虑到依赖库的的安全性。

Go语言在设计之初就考虑到了供应链安全问题，可以帮助Go开发者更好地理解和跟踪依赖Module的安全性。

Filippo Valsorda的技术博客[How Go Mitigates Supply Chain Attacks](https://go.dev/blog/supply-chain)就介绍了如何管理依赖包的安全。

在今年9月的时候，我们在Julie Qiu的技术博客[Vulnerability Management for Go](https://go.dev/blog/vuln)里预发布了Go语言漏洞管理机制。

该漏洞管理机制的核心是维护了一个新的、经过筛选的安全漏洞数据库以及一个新的Go命令[govulncheck](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck)用于找出Go代码里的安全漏洞。

### Go开发者调研

我们以前都是每年年尾向Go开发者做关于Go语言的年度使用调研。

今年我们改变了调研频率，改为一年调研2次，每次调研的内容减少。这样可以保证我们既可以收集到足够多的调研结果，也能减轻被调研者的负担。

2021年调研结果参考Alice Merrick的文章[Go Developer Survey 2021 Results](https://go.dev/blog/survey2021-results)。

2022年上半年的调研结果参考Todd Kulesza的文章[Go Developer Survey 2022 Q2 Results](https://go.dev/blog/survey2022-q2-results)。

### Go年度会议活动

2022年举办了多场Go语言开发者会议。包括7月份在欧洲柏林举办的GopherCon Europe、10月份在芝加哥举行的GopherCon。

上周，我们在Google开源直播活动中举办了Go语言日的线上活动，主要分享了以下内容：

- “[How Go Became its Best Self](https://www.youtube.com/watch?v=vQm_whJZelc)”, by Cameron Balahan, at GopherCon Europe.
- “[Go team Q&A](https://www.youtube.com/watch?v=KbOTTU9yEpI)”, with Cameron Balahan, Michael Knyszek, and Than McIntosh, at GopherCon Europe.
- “[Compatibility: How Go Programs Keep Working](https://www.youtube.com/watch?v=v24wrd3RwGo)”, by Russ Cox at GopherCon.
- “[A Holistic Go Experience](https://www.gophercon.com/agenda/session/998660)”, by Cameron Balahan at GopherCon (video not yet posted)
- “[Structured Logging for Go](https://opensourcelive.withgoogle.com/events/go-day-2022/watch?talk=talk2)”, by Jonathan Amsterdam at Go Day on Google Open Source Live
- “[Writing your Applications Faster and More Securely with Go](https://opensourcelive.withgoogle.com/events/go-day-2022/watch?talk=talk3)”, by Cody Oss at Go Day on Google Open Source Live
- “[Respecting Memory Limits in Go](https://opensourcelive.withgoogle.com/events/go-day-2022/watch?talk=talk4), by Michael Knyszek at Go Day on Google Open Source Live

### Go里程碑

今年还有一个重要的里程碑是我们在权威期刊Communications of the ACM里发表了文章[The Go Programming Language and Environment](https://cacm.acm.org/magazines/2022/5/260357-the-go-programming-language-and-environment/fulltext)。作者都是Go语言的早期设计者和主要实现者，包括Russ Cox, Robert Griesemer, Rob Pike, Ian Lance Taylor和Ken Thompson。

这篇文章里详细阐述了为什么Go语言会如此流行和提高开发效率。

简而言之，Go语言专注为软件研发流程提供一个高效的开发语言和环境。同时提供强大的可扩展性，支持大规模软件项目开发和大规模系统部署。

### 展望

在Go语言的第14个年头，我们会继续让Go语言成为软件工程里最好的开发语言。

我们会重点提升软件供应链安全、程序兼容性和结构化的logging。

同时，还会有其它大量功能改进和优化，比如基于profling结果的优化。大家敬请关注。

## 推荐阅读

* [Go面试题系列，看看你会几题？](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=Mzg2MTcwNjc1Mg==&action=getalbum&album_id=2199553588283179010#wechat_redirect)
* [Go泛型的使用教程和设计原理](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=Mzg2MTcwNjc1Mg==&action=getalbum&album_id=2184751156453834753#wechat_redirect)
* [Go工作区模式](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=Mzg2MTcwNjc1Mg==&action=getalbum&album_id=2339933847347544066#wechat_redirect)
* [Go Fuzzing模糊测试使用详解](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=Mzg2MTcwNjc1Mg==&action=getalbum&album_id=2287546159059566596#wechat_redirect)
* [Go语言常见错误和最佳实践系列](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=Mzg2MTcwNjc1Mg==&action=getalbum&album_id=2549657749539028992#wechat_redirect)
* [Go开发进阶系列](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=Mzg2MTcwNjc1Mg==&action=getalbum&album_id=2549661543605764097#wechat_redirect)



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

* https://go.dev/blog/13years