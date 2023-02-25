# 一文读懂Go 1.21引入的PGO性能优化

## 背景

Go 1.21版本于2023年2月份正式发布，在这个版本里引入了PGO性能优化机制。

PGO的英文全称是Profile Guided Optimization，基本原理分为以下2个步骤：

* 先对程序做profile，得到一个.pgo文件
* 编译程序时启用pgo，编译器会根据.pgo文件里的内容对程序做性能优化

## 实例



## 推荐阅读

* [Go 1.20来了，看看都有哪些变化](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484693&idx=1&sn=9f84d42dfadb7319f8c4e4645893d218&chksm=ce124a7af965c36c63deafc09b9f2bfdae35bc8714aa2f76bca63e233f664b499bad742a8c3f&token=293290824&lang=zh_CN#rd)

* [Go面试题系列，看看你会几题](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=Mzg2MTcwNjc1Mg==&action=getalbum&album_id=2199553588283179010#wechat_redirect)

* [Go常见错误和最佳实践系列](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=Mzg2MTcwNjc1Mg==&action=getalbum&album_id=2549657749539028992#wechat_redirect)

  

## 开源地址

文章和示例代码开源在GitHub: [Go语言初级、中级和高级教程](https://github.com/jincheng9/go-tutorial)。

公众号：coding进阶。关注公众号可以获取最新Go面试题和技术栈。

个人网站：[Jincheng's Blog](https://jincheng9.github.io/)。

知乎：[无忌](https://www.zhihu.com/people/thucuhkwuji)。



## 福利

我为大家整理了一份后端开发学习资料礼包，包含编程语言入门到进阶知识(Go、C++、Python)、后端开发技术栈、面试题等。

关注公众号「coding进阶」，发送消息 **backend** 领取资料礼包，这份资料会不定期更新，加入我觉得有价值的资料。还可以发送消息「**进群**」，和同行一起交流学习，答疑解惑。



## References

* https://go.dev/blog/pgo-preview