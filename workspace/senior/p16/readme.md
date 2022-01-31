# Go Quiz: Google工程师的Go语言面试题

## 面试题

Google工程师和Go语言布道者Valentin Deleplace会不定期分享一些Go语言的趣味题目。

我选取了最近的几道题目，给大家春节加加餐。

### 题目1：

```go
// quiz0.go
package main

import "fmt"

func main() 
	s := []string{"a", "b", "c"}

	copy(s[1:], s)

	fmt.Println(s)
}
```

- A: [a a a]
- B: [a a b]
- C: [a a b c]
- D: [b c c]
- E: panic

这道题主要考察Go语言里对切片`slice`使用copy函数的注意事项。

### 题目2：

```go
// quiz1.go
package main

import "fmt"

func main() {
	a := make([]int, 20)

	b := a[18:]
	b = append(b, 2022)

	fmt.Println(len(b), cap(b))
}
```

- A: 1 2
- B: 3 4
- C: 3 33
- D: 3 40

这道题主要考察以下Go语言里对切片`slice`使用冒号`:`截取操作的底层机制以及`slice`的扩容原理。

### 题目3：

```go
// quiz2.go
package main

import "fmt"

func main() {
	c := make(chan int, 1)
	c <- 1
	close(c)
	close(c)
	fmt.Println("OK")
}
```

- A: 死锁
- B: panic
- C: 打印"OK"
- D: 编译报错

这道题主要考察以下Go语言里管道`channel`的注意事项。



## 解析

这几道题对于看过我写的Go面试系列文章的同学来说，应该非常简单，可以给本人vx公众号发送消息`2022`获取答案。

对答案有疑惑的可以看看我之前的2篇Go面试题解析：

* [Go Quiz: 从Go面试题看slice的底层原理和注意事项](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483741&idx=1&sn=486066a3a582faf457f91b8397178f64&chksm=ce124e32f965c72411e2f083c22531aa70bb7fa0946c505dc886fb054b2a644abde3ad7ea6a0&token=1073108956&lang=zh_CN#rd)
* [Go Quiz: 从Go面试题看channel的注意事项](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483746&idx=1&sn=c3ec0e3f67fa7b1cb82e61450d10c7fd&chksm=ce124e0df965c71b7e148ac3ce05c82ffde4137cb901b16c2c9567f3f6ed03e4ff738866ad53&token=1073108956&lang=zh_CN#rd)



## 春节加餐

* [Go Quiz: 从Go面试题看slice的底层原理和注意事项](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483741&idx=1&sn=486066a3a582faf457f91b8397178f64&chksm=ce124e32f965c72411e2f083c22531aa70bb7fa0946c505dc886fb054b2a644abde3ad7ea6a0&token=1073108956&lang=zh_CN#rd)

* [Go Quiz: 从Go面试题搞懂slice range遍历的坑](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483810&idx=1&sn=1f6ab90f481ef340cf48c2458a2a8682&chksm=ce124ecdf965c7dbbf26b331f3e316b9d376f8cd7d9190bfce0e9695593c8bb8b7e8ed06ed8c&token=1073108956&lang=zh_CN#rd)

* [Go Quiz: 从Go面试题看channel的注意事项](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483746&idx=1&sn=c3ec0e3f67fa7b1cb82e61450d10c7fd&chksm=ce124e0df965c71b7e148ac3ce05c82ffde4137cb901b16c2c9567f3f6ed03e4ff738866ad53&token=1073108956&lang=zh_CN#rd)

* [Go Quiz: 从Go面试题看channel在select场景下的注意事项](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483816&idx=1&sn=44e5cf4900b44f9a0cde491df5dd6e51&chksm=ce124ec7f965c7d1edd9ccffe80520981970ad6000cfea3b1a4099a4627f0f24cc33272ec996&token=1073108956&lang=zh_CN#rd)

* [Go Quiz: 从Go面试题看defer语义的底层原理和注意事项](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483756&idx=1&sn=d536fa3340e1d5f91d72eaa8b67c8123&chksm=ce124e03f965c715e26f5943948e17d8e0ebb3c4a3a180a149219a610f83fc6eb77b3b166b6a&token=1073108956&lang=zh_CN#rd)

* [Go Quiz: 从Go面试题看defer的注意事项第2篇](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483762&idx=1&sn=ca4235d28d513267aa082dc12cb37fda&chksm=ce124e1df965c70b06be48bc537bd628f3caf81e2837ebc2fbd0edddc6eb4f2b2c52e4d5c5d5&token=1073108956&lang=zh_CN#rd)

* [Go Quiz: 从Go面试题看defer的注意事项第3篇](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483821&idx=1&sn=2ebfb63b78f5fa3666ca6801985a5462&chksm=ce124ec2f965c7d441efbc0d40d0dd8b4d62c255f8ca0b093d106944adbca9a903e94eb92b19&token=1073108956&lang=zh_CN#rd)

* [Go Quiz: 从Go面试题看分号规则和switch的注意事项](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483750&idx=1&sn=235d959cd0401c2c4299f2ec1bbbfec9&chksm=ce124e09f965c71f1989ac9fe691af6a7697ba12a084d8cbdfe3966da1372f787d8e07c231a7&token=1073108956&lang=zh_CN#rd)

* [官方教程：Go泛型入门](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483720&idx=1&sn=57ec4877dfd364a59deacf1e74a4fb66&chksm=ce124e27f965c731432dcc89d1e0563cf84baaef482eaa068a91bee61f10cf85b433923b83b4&token=1073108956&lang=zh_CN#rd)

* [一文读懂Go泛型设计和使用场景](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483731&idx=1&sn=b2258b28e2f3c16b065a5a1b22c15b0d&chksm=ce124e3cf965c72a6a22e0ed15deda8238567407bbd7157a79753fc8b605727ab2153009493c&token=1073108956&lang=zh_CN#rd)

  

## 开源地址

文章和示例代码开源地址在GitHub: https://github.com/jincheng9/go-tutorial

公众号：coding进阶。关注公众号可以获取最新Go面试题和技术栈。

个人网站：https://jincheng9.github.io/

知乎：https://www.zhihu.com/people/thucuhkwuji



## References

* https://twitter.com/val_deleplace