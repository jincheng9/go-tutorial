# Google重磅发布Go语言编码规范

## 前言

Google官方在2022.11.23重磅发布了Go语言编码规范。

这个编码规范源自于Google内部的Go项目，是Google的开发人员要遵守的代码规范。

在Go语言诞生后，全世界的Go开发者其实一直期盼着能有官方的编码规范，但迟迟未能如愿。

有些技术团队根据自己原来的编程语言背景，直接照搬过来，用于公司内部的Go语言编码规范。

尤其是写Java的，把Java的编程语言规范用于Go语言是非常不合适的。

为了让Go开发者可以知道如何写出更简洁、更地道的Go代码，官方也做出了一些努力，推出了Effective Go和Go Code Review Comments。

* Go官方的Effective Go: https://go.dev/doc/effective_go。

  > Note added January, 2022: This document was written for Go's release in 2009, and has not been updated significantly since. Although it is a good guide to understand how to use the language itself, thanks to the stability of the language, it says little about the libraries and nothing about significant changes to the Go ecosystem since it was written, such as the build system, testing, modules, and polymorphism. There are no plans to update it, as so much has happened and a large and growing set of documents, blogs, and books do a fine job of describing modern Go usage. Effective Go continues to be useful, but the reader should understand it is far from a complete guide. See [issue 28782](https://github.com/golang/go/issues/28782) for context.

* Go官方的Code Review Comments: https://github.com/golang/go/wiki/CodeReviewComments

Effective Go主要讲解的是Go语言的语法细节以及一些最佳实践。Code Review Comments包含了一些Code Review过程中经常出现的问题。这2个指引可以拿来作为参考，但不足以成为一个非常完善齐全的Go语言编码规范。



## 社区的Go语言编码规范

这些年整个Go社区陆续诞生了一些有影响力的Go语言编码规范，主要有以下这些：

* Uber的编码规范：https://github.com/uber-go/guide
* CockroachDB的编码规范：https://wiki.crdb.io/wiki/spaces/CRDB/pages/181371303/Go+style+guide

Uber的编码规范开源在GitHub，业界认可度最高。

接下来我们看看这次Google推出的Go语言编码规范包含哪些内容。



## Google的Go语言编码规范

需要申明的是，这是Google推出的Go语言编码规范，并不是Go团队自己单独推出的Go语言编码规范。

这次发布的Go语言编码规范主要包含Style Guide，Style Decisions和Best Practices这三部分内容。

| Document            | Link                                                  | Primary Audience    | [Normative](https://google.github.io/styleguide/go/index#normative) | [Canonical](https://google.github.io/styleguide/go/index#canonical) |
| ------------------- | ----------------------------------------------------- | ------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| **Style Guide**     | https://google.github.io/styleguide/go/guide          | Everyone            | Yes                                                          | Yes                                                          |
| **Style Decisions** | https://google.github.io/styleguide/go/decisions      | Readability Mentors | Yes                                                          | No                                                           |
| **Best Practices**  | https://google.github.io/styleguide/go/best-practices | Anyone interested   | No                                                           | No                                                           |

1. Style Guide是Go编码规范的基础，这里描述的规则是通用的，所有Go开发者都必须遵守。Style Decisions和Best Practices都遵循了Style Guide里的规则。
2. Style Decisions讲述了部分具体的编码规范以及背后的原因。这里的内容会随着新的语言特性、新的库或者新的编程模式而发生变化。
3. Best Practices讲述了具体写代码过程中的最佳实践。参考这个最佳实践，可以让代码的可读性更好、代码更健壮，有利于代码的可持续维护。



## 编码规范举例

### package命名

比如在style decisions里，对于package的命名规范，Google给出的建议是：

Go package的命名要短小，且只有小写字母。如果package的名字由多个单词组成，需要全部小写，且中间不要用任何符号做分隔。

例如，推荐用`tabwriter`，不推荐用`tabWriter`，`TabWriter`，`tab_writer`。

### Receiver命名

又比如对于Receiver的变量命名，Google的编码风格是：

- 短，通常只有1个或者2个字母。
- 是Receiver变量的类型名称的缩写。
- 对于该类型的所有Receiver变量命名都保持一致。

| 坏的命名风格                | 好的命名风格              |
| --------------------------- | ------------------------- |
| `func (tray Tray)`          | `func (t Tray)`           |
| `func (info *ResearchInfo)` | `func (ri *ResearchInfo)` |
| `func (this *ReportWriter)` | `func (w *ReportWriter)`  |
| `func (self *Scanner)`      | `func (s *Scanner)`       |

还有更多非常好的编码风格就不在本文里全部列出了。

本人正在对该编码风格做翻译，开源在GitHub：https://github.com/jincheng9/google-go-style-guide，欢迎大家关注。



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

* https://google.github.io/styleguide/go/index