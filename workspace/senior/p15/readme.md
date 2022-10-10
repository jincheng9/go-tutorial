# Go Quiz: 从Go面试题看defer的注意事项第3篇

## 面试题

这是Go Quiz系列中关于defer的第3篇，这道题目来源于真实的互联网项目里，也是Go初学者容易犯的错误之一。

```go
package main

import "fmt"

func test1() int {
	var i = 0
	defer func() {
		i = 10
	}()
	return i
}

func test2() (result int) {
	defer func() {
		result *= 10
	}()
	return 6
}

func test3() (result int) {
	result = 8
	defer func() {
		result *= 10
	}()
	return
}

func main() {
	result1 := test1()
	result2 := test2()
	result3 := test3()
	fmt.Println(result1, result2, result3)
}
```

- A: 0 6 8 
- B: 0 60 80
- C: 10 6 80
- D: 10 60 8
- E: 编译报错

这道题主要考察以下知识点：

* 在被defer的函数里对返回值做修改在什么情况下会生效？

   

## 解析

我们来看下官方文档里的规定：

> Each time a "defer" statement executes, the function value and parameters to
> the call are evaluated as usual and saved anew but the actual function is not 
> invoked. Instead, deferred functions are invoked immediately before the 
> surrounding function returns, in the reverse order they were deferred. That
> is, if the surrounding function returns through an explicit return statement, 
> deferred functions are executed after any result parameters are set by that 
> return statement but before the function returns to its caller. If a deferred
> function value evaluates to nil, execution panics when the function is 
> invoked, not when the "defer" statement is executed.

重点是这句：

> That is, if the surrounding function returns through an explicit return statement, 
> deferred functions are executed after any result parameters are set by that 
> return statement but before the function returns to its caller.

Go的规定是：如果在函数A里执行了 defer B(xx)，函数A显式地通过return语句来返回时，会先把返回值赋值给A的返回值参数，然后执行被defer的函数B，最后才真正地返回给函数A的调用者。

对于`test1`函数，执行`return i`时，先把`i`的值赋值给`test1`的返回值，defer语句里对`i`的赋值并不会改变函数`test1`的返回值，`test1`函数返回0。

对于`test2`函数，执行`return i`时，先把`i`的值赋值给`test2`的命名返回值result，defer语句里对`result`的修改会改变函数`test2`的返回值，`test2`函数返回60。

对于`test3`函数，虽然`return`后面没有具体的值，但是编译器不会报错，执行`return`时，先执行被defer的函数，在被defer的函数里对result做了修改，result的结果变为80，最后`test3`函数return返回的时候返回80。

所以答案是B。

**所以想要对return返回的值做修改，必须使用命名返回值(Named Return Value)**。



## 加餐

可以回顾Go quiz系列中关于defer的另外2道题目，加深对defer的理解。

题目1：[Go Quiz: 从Go面试题看defer语义的底层原理和注意事项](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483756&idx=1&sn=d536fa3340e1d5f91d72eaa8b67c8123&chksm=ce124e03f965c715e26f5943948e17d8e0ebb3c4a3a180a149219a610f83fc6eb77b3b166b6a&token=531427802&lang=zh_CN#rd)

题目2：[Go Quiz: 从Go面试题看defer的注意事项第2篇](http://link.zhihu.com/?target=https%3A//mp.weixin.qq.com/s%3F__biz%3DMzg2MTcwNjc1Mg%3D%3D%26mid%3D2247483762%26idx%3D1%26sn%3Dca4235d28d513267aa082dc12cb37fda%26chksm%3Dce124e1df965c70b06be48bc537bd628f3caf81e2837ebc2fbd0edddc6eb4f2b2c52e4d5c5d5%26token%3D531427802%26lang%3Dzh_CN%23rd)



## 开源地址

文章和示例代码开源地址在GitHub: https://github.com/jincheng9/go-tutorial

公众号：coding进阶。关注公众号可以获取最新Go面试题和技术栈。

个人网站：https://jincheng9.github.io/

知乎：https://www.zhihu.com/people/thucuhkwuji



## References

* [Go Quiz: 从Go面试题看defer语义的底层原理和注意事项](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483756&idx=1&sn=d536fa3340e1d5f91d72eaa8b67c8123&chksm=ce124e03f965c715e26f5943948e17d8e0ebb3c4a3a180a149219a610f83fc6eb77b3b166b6a&token=531427802&lang=zh_CN#rd)
* [Go Quiz: 从Go面试题看defer的注意事项第2篇](http://link.zhihu.com/?target=https%3A//mp.weixin.qq.com/s%3F__biz%3DMzg2MTcwNjc1Mg%3D%3D%26mid%3D2247483762%26idx%3D1%26sn%3Dca4235d28d513267aa082dc12cb37fda%26chksm%3Dce124e1df965c70b06be48bc537bd628f3caf81e2837ebc2fbd0edddc6eb4f2b2c52e4d5c5d5%26token%3D531427802%26lang%3Dzh_CN%23rd)