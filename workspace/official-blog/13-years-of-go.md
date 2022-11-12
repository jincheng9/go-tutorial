# Go语言开源13周年啦，看看负责人说了啥

## 前言

 不知不觉，Go语言已经开源13年了。

日前，Go团队负责人Russ Cox在官方博客上发表了Go开源13年的感想。

Russ首先回顾了2022年3月份发布的Go 1.18版本引入的工作区模式(Go Workspace)、模糊测试(Fuzzing)和泛型设计。

展望Go语言的第14年，

本人针对Russ的原文做了一个翻译，以飨读者。



## 原文翻译

谷歌Go团队Russ Cox

2022.11.10

今天我们可以开心地庆祝Go语言开源13年啦。

### Go 1.18

对Go语言而言，今年有非常多的重要事项。最重要的当然是在今年3月份我们发布了Go 1.18版本，这个版本引入了非常多的新功能，包括大家熟知的Go workspace工作区模式、Go Fuzzing模糊测试和Go泛型。

* Go workspace工作区模式让开发者可以同时开发多个互相有依赖的Module。想了解Go工作区模式的，可以参考Beth Brown's的文章[Get familiar with workspaces](https://go.dev/blog/get-familiar-with-workspaces)以及 [workspace reference](https://go.dev/ref/mod#workspaces)。
* Fuzzing模糊测试是go test的一个新功能，可以生成随机的输入用于检测代码的正确性以及漏洞。想了解Go Fuzzing的可以参考 [Getting started with fuzzing](https://go.dev/doc/tutorial/fuzz)和 [fuzzing reference](https://go.dev/security/fuzz/)。另外，Fuzzing作者katie Hockman在GopherCon 2022会议上分享的主题“Fuzz Testing Made Easy”也即将上线，敬请关注。
* 泛型可能是Go语言里被开发者提及最多的feature。想了解Go泛型入门的，可以参考教程[Getting started with generics](https://go.dev/doc/tutorial/generics)。想了解Go泛型更多设计和使用细节的，可以参考官方博客[An Introduction to Generics](https://go.dev/blog/intro-generics) 和[When to Use Generics](https://go.dev/blog/when-generics)，以及Google 2021开源日的技术分享[Using Generics in Go](https://www.youtube.com/watch?v=nr8EpUO9jhw) 和GopherCon 2021上Go泛型作者Robert Griesemer和Ian Lance Taylor的技术分享[Generics!](https://www.youtube.com/watch?v=Pa_e9EeCdy8)。

### Go 1.19

和Go 1.18相比，今年8月份发布的Go 1.19版本相对修改少一些，吸引的关注也少一些。Go 1.19版本专注在继续优化Go 1.18引入的新特性。此外在Go 1.19版本中，我们支持在Go doc注释中添加超链接、列表以及标题，用于生成Go package的说明文档。另外我们针对Go的垃圾回收期(garbage collector)引入了软内存限制(soft memory limit)，这对于管理容器负载(container workloads)非常有用。关于GC更多的改进和优化，可以参考Michael Knyszek最近的博文[Go runtime: 4 years later](https://go.dev/blog/go119runtime)、视频[Respecting Memory Limits in Go](https://www.youtube.com/watch?v=07wduWyWx8M&list=PLtoVuM73AmsJjj5tnZ7BodjN_zIvpULSx) 和 [Guide to the Go Garbage Collector](https://go.dev/doc/gc-guide)。

### Go开发工具

Go团队还在维护2个Go语言开发工具VS Code Go扩展插件以及Gopls语言服务器。

今年Gopls专注于提升稳定性和性能、支持泛型以及提供更丰富的Go语言分析功能。

想了解VS Code Go和Gopls的最新改进，可以参考Suzy Mueller的技术分享[Building Better Projects with the Go Editor](https://www.youtube.com/watch?v=jMyzsp2E_0U)、[Debugging Treasure Hunt](https://www.youtube.com/watch?v=ZPIPPRjwg7Q)和官方技术博客[Debugging Go in VS Code](https://go.dev/s/vscode-go-debug)。

### Go供应链安全

Another part of development scale is the number of dependencies in a project. A month or so after Go’s 12th birthday, the [Log4shell vulnerability](https://en.wikipedia.org/wiki/Log4Shell) served as a wake-up call for the industry about the importance of supply chain security.

 Go’s module system was designed specifically for this purpose, to help you understand and track your dependencies, identify which specific ones you are using, and determine whether any of them have known vulnerabilities. Filippo Valsorda’s blog post “[How Go Mitigates Supply Chain Attacks](https://go.dev/blog/supply-chain)” gives an overview of our approach. In September, we previewed Go’s approach to vulnerability management in Julie Qiu’s blog post “[Vulnerability Management for Go](https://go.dev/blog/vuln)”. The core of that work is a new, curated vulnerability database and a new [govulncheck command](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck), which uses advanced static analysis to eliminate most of the false positives that would result from using module requirements alone.

### Go开发者调查

Part of our effort to understand Go users is our annual end-of-year Go survey. This year, our user experience researchers added a lightweight mid-year Go survey as well. We aim to gather enough responses to be statistically significant without being a burden on the Go community as a whole. For the results, see Alice Merrick’s blog post “[Go Developer Survey 2021 Results](https://go.dev/blog/survey2021-results)” and Todd Kulesza’s post “[Go Developer Survey 2022 Q2 Results](https://go.dev/blog/survey2022-q2-results)”.

### Go年度会议活动

As the world starts traveling more, we’ve also been happy to meet many of you in person at Go conferences in 2022, particularly at GopherCon Europe in Berlin in July and at GopherCon in Chicago in October. Last week we held our annual virtual event, [Go Day on Google Open Source Live](https://opensourcelive.withgoogle.com/events/go-day-2022). Here are some of the talks we’ve given at those events:

- “[How Go Became its Best Self](https://www.youtube.com/watch?v=vQm_whJZelc)”, by Cameron Balahan, at GopherCon Europe.
- “[Go team Q&A](https://www.youtube.com/watch?v=KbOTTU9yEpI)”, with Cameron Balahan, Michael Knyszek, and Than McIntosh, at GopherCon Europe.
- “[Compatibility: How Go Programs Keep Working](https://www.youtube.com/watch?v=v24wrd3RwGo)”, by Russ Cox at GopherCon.
- “[A Holistic Go Experience](https://www.gophercon.com/agenda/session/998660)”, by Cameron Balahan at GopherCon (video not yet posted)
- “[Structured Logging for Go](https://opensourcelive.withgoogle.com/events/go-day-2022/watch?talk=talk2)”, by Jonathan Amsterdam at Go Day on Google Open Source Live
- “[Writing your Applications Faster and More Securely with Go](https://opensourcelive.withgoogle.com/events/go-day-2022/watch?talk=talk3)”, by Cody Oss at Go Day on Google Open Source Live
- “[Respecting Memory Limits in Go](https://opensourcelive.withgoogle.com/events/go-day-2022/watch?talk=talk4), by Michael Knyszek at Go Day on Google Open Source Live

### Go文章

One other milestone for this year was the publication of “[The Go Programming Language and Environment](https://cacm.acm.org/magazines/2022/5/260357-the-go-programming-language-and-environment/fulltext)”, by Russ Cox, Robert Griesemer, Rob Pike, Ian Lance Taylor, and Ken Thompson, in *Communications of the ACM*. The article, by the original designers and implementers of Go, explains what we believe makes Go so popular and productive. In short, it is that Go effort focuses on delivering a full development environment targeting the entire software development process, with a focus on scaling both to large software engineering efforts and large deployments.

### 展望

In Go’s 14th year, we’ll keep working to make Go the best environment for software engineering at scale. We plan to focus particularly on supply chain security, improved compatibility, and structured logging, all of which have been linked already in this post. And there will be plenty of other improvements as well, including profile-guided optimization.



## 推荐阅读

* [Go面试题系列，看看你会几题？](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=Mzg2MTcwNjc1Mg==&action=getalbum&album_id=2199553588283179010#wechat_redirect)
* [Go泛型]
* [Go工作区模式]
* [Go Fuzzing]
* [Go常见错误和最佳实践系列]
* [Go进阶系列]



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