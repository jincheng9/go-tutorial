# Go十大常见错误第6篇：slice初始化常犯的错误

## 前言

这是Go十大常见错误系列的第6篇：slice初始化常犯的错误。素材来源于Go布道者，现Docker公司资深工程师[Teiva Harsanyi](https://teivah.medium.com/)。

本文涉及的源代码全部开源在：[Go十大常见错误源代码](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p28)，欢迎大家关注公众号，及时获取本系列最新更新。



## 场景

假设我们知道要创建的slice的长度，你会怎么创建和初始化这个slice？

比如我们定义了一个结构体叫`Bar`，现在要创建一个slice，里面的元素就是`Bar`类型，而且该slice的长度是已知的。

### 方法1

有的人可能这么来做，先定义slice

```go
var bars []Bar
bars := make([]Bar, 0)
```

每次要往`bars`这个slice插入元素的时候，通过append来操作

```go
bars = append(bars, barElement)
```

`slice`实际上是一个结构体类型，包含3个字段，分别是

- array: 是指针，指向一个数组，切片的数据实际都存储在这个数组里。
- len: 切片的长度。
- cap: 切片的容量，表示切片当前最多可以存储多少个元素，如果超过了现有容量会自动扩容。

slice底层的数据结构定义如下：

```go
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

type Pointer *ArbitraryType
```

如果按照上面的示例先创建一个长度为0的slice，那在append插入元素的过程中，`bars`这个slice会做自动扩容。

如果`bars`的长度比较大，可能会发生多次扩容。每次扩容都要创建一个新的内存空间，然后把旧内存空间的数据拷贝过来，效率比较低。

### 方法2

在定义slice的时候指定长度，代码示例如下：

```go
func convert(foos []Foo) []Bar {
	bars := make([]Bar, len(foos))
	for i, foo := range foos {
		bars[i] = fooToBar(foo)
	}
	return bars
}
```

这行代码`	bars := make([]Bar, len(foos))`直接指定了slice的长度，无需扩容。



###  方法3

在定义slice的时候提前指定容量，长度设置为0，代码示例如下：

```go
func convert(foos []Foo) []Bar {
	bars := make([]Bar, 0, len(foos))
	for _, foo := range foos {
		bars = append(bars, fooToBar(foo))
	}
	return bars
}
```

这种方法也可以，也无需扩容。

那方法2和方法3哪种好一点呢？其实各有优缺点：

* 从效率上来说，方法2比方法3要高一点，因为方法3里调用了append函数，再对bars赋值，效率比直接通过`		bars[i] = fooToBar(foo)`要低一点。
* 从代码的可维护性来说，方法2不能通过append函数来插入元素，因为方法2里的slice定义的时候指定了长度，如果调用append，会扩容，往现有元素后面追加元素。方法3不能通过`bars[i] = `的方式来赋值，因为方法3里的slice定义的时候长度为0，如果使用`bars[1]=`，会触发`panic: runtime error: index out of range [1] with length 0`。



最后强烈推荐大家看看这道关于Go slice的面试题：[Go Quiz: 从Go面试题看slice的底层原理和注意事项](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483741&idx=1&sn=486066a3a582faf457f91b8397178f64&chksm=ce124e32f965c72411e2f083c22531aa70bb7fa0946c505dc886fb054b2a644abde3ad7ea6a0&token=1846351524&lang=zh_CN#rd)。

看完后你会彻底了解Go slice的原理和注意事项。



## 推荐阅读

* [Go十大常见错误第1篇：未知枚举值](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484146&idx=1&sn=10fb12b643a2e37c090e5aa3bc583152&chksm=ce124d9df965c48bb954aeddabdff3db12738ded3875542250c5d0ef6cfd4417fc56580288b1&token=1912894792&lang=zh_CN#rd)

* [Go十大常见错误第2篇：benchmark性能测试的坑](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484163&idx=1&sn=b28d61c1f3ec9d914e698dce105ba5d1&chksm=ce124c6cf965c57a90bc85a5295ed9375103de20607b509f845583ff6686385df0ed96653d00&token=1912894792&lang=zh_CN#rd)

* [Go十大常见错误第3篇：go指针的性能问题和内存逃逸](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484247&idx=1&sn=faf716627afb00df646cecff023fb63c&chksm=ce124c38f965c52efd009a4c98691d56b5765dc7dce98aa49b226ad9274bd062d8d01e702e91&token=1899277735&lang=zh_CN#rd)

* [Go十大常见错误第4篇：break操作的注意事项](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484262&idx=1&sn=c1bea8af60444a4ef73c4d4d7a09d16d&chksm=ce124c09f965c51f3663ac9089a792d36c3685850e12695dd26d15a1a50f393b2d7c92b9983a&token=461369035&lang=zh_CN#rd)

* [Go十大常见错误第5篇：Go语言Error管理](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484274&idx=1&sn=711abea3c6fd5d15341ee1b34da8a160&chksm=ce124c1df965c50b3af84965f7ed30b574cd0b247ea6f77b944ec858bd43ee37f4c1554a5bce&token=1846351524&lang=zh_CN#rd)

* [Go面试题系列，看看你会几题？](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=Mzg2MTcwNjc1Mg==&action=getalbum&album_id=2199553588283179010#wechat_redirect)

  

## 开源地址

文章和示例代码开源在GitHub: [Go语言初级、中级和高级教程](https://github.com/jincheng9/go-tutorial)。

公众号：coding进阶。关注公众号可以获取最新Go面试题和技术栈。

个人网站：[Jincheng's Blog](https://jincheng9.github.io/)。

知乎：[无忌](https://www.zhihu.com/people/thucuhkwuji)。



## 福利

我为大家整理了一份后端开发学习资料礼包，包含编程语言入门到进阶知识(Go、C++、Python)、后端开发技术栈、面试题等。

关注公众号「coding进阶」，发送消息 **backend** 领取资料礼包，这份资料会不定期更新，加入我觉得有价值的资料。还可以发送消息「**进群**」，和同行一起交流学习，答疑解惑。



## References

* https://itnext.io/the-top-10-most-common-mistakes-ive-seen-in-go-projects-4b79d4f6cd65