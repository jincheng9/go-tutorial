# Go 1.21的2个语言变化

## 语言行为变化

Go 1.20已经于今年2月份发布，Go 1.21也不远了，我们来先睹为快，看看Go 1.21版本里几个有趣的变化。

文末附送2道面试题。

### panic(nil)

```go
func main() {
  defer func() {
    print(recover() == nil)
  }()
  panic(nil)
}
```

大家先想一想这段代码会输出什么？是true还是false。

在Go 1.20版本及以前会输出true。

但是在Go 1.21版本开始会输出false。这是因为Go 1.21定义了一个新的类型`*runtime.PanicNilError`。

`panic(nil)`后，`recover()`会返回一个类型为`*runtime.PanicNilError`，值为`panic called with nil argument`的变量，具体可以参考如下代码：

```go
func main() {
	defer func() {
		r := recover()
		fmt.Printf("%T\n", r) // *runtime.PanicNilError
		fmt.Println(r) // panic called with nil argument
	}()
	panic(nil)
}
```



### clear函数

Go 1.21会新增一个clear函数，用于清理map和slice里的元素。示例代码如下：

```go
package main

import "fmt"

var x = 0.0
var nan = x / x

func main() {
	s := []int{1, 2, 3}
	clear(s)
	fmt.Println(s) // [0 0 0]

	m := map[float64]int{0.1: 9}
	m[nan] = 5
	clear(m)
	fmt.Println(len(m)) // 0
}
```

官方源码说明如下：

> // The clear built-in function clears maps and slices.
>
> // For maps, clear deletes all entries, resulting in an empty map.
>
> // For slices, clear sets all elements up to the length of the slice
>
> // to the zero value of the respective element type. If the argument
>
> // type is a type parameter, the type parameter's type set must
>
> // contain only map or slice types, and clear performs the operation
>
> // implied by the type argument.
>
> func clear[T ~[]Type | ~map[Type]Type1](t T)

对于map，调用clear函数，会直接把map里的元素清空，成为一个empty map。

对于slice，调用clear函数，会保持原slice的长度不变，把里面元素的值修改为slice元素类型的零值。



## 面试题

defer语义是Go开发人员经常使用到的，也是最容易理解错误的地方。

大家看看下面2道关于defer的程序会输出什么结果。

```go
package main

import "fmt"

func f() {
	defer func() {
		defer func() { recover() }()
		defer recover()
		panic(2)
	}()
	panic(1)
}

func main() {
	defer func() { fmt.Print(recover()) }()
	f()
}
```

* A: 2
* B: 1
* C: nil
* D: 抛panic异常



```go
package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		defer func() { print(i) }()
	}
	for i := range [3]int{} {
		defer func() { print(i) }()
	}
}

```

* A: 222333
* B: 210333
* C: 333333
* D: 210210

想知道答案的发送消息`121`到公众号。



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

* https://twitter.com/go100and1
* https://twitter.com/go100and1/status/1623546829773361152