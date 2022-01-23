# Go Quiz: 从Go面试题搞懂slice range遍历的坑

## 面试题

最近Go 101的作者发布了11道Go面试题，非常有趣，打算写一个系列对每道题做详细解析，欢迎大家关注。

本题是Go quiz `slice`系列的第2道题目，这道题非常有迷惑性。

通过这道题我们可以知晓对`slice`做range遍历的坑，避免在实际项目中踩坑。

```go
package main

func main() {
	var x = []string{"A", "B", "C"}

	for i, s := range x {
		print(i, s, ",")
		x[i+1] = "M"
		x = append(x, "Z")
		x[i+1] = "Z"
	}
}
```

- 0A,1B,2C,
- 0A,1Z,2Z,
- 0A,1M,2M,
- 0A,1M,2C,
- 0A,1Z,2M,
- 0A,1M,2Z,
- (infinite loop)

大家可以在评论区留下你们的答案。这道题主要有以下几个考点：

1. `slice`做range遍历，Go编译器背后会做哪些事情？
2. `slice`什么时候扩容，扩容后的行为是怎么样的？



## 解析

我们先逐个解答上面的问题。

### range遍历机制

range对`slice`做遍历的时候，实际上是先构造一个原slice的拷贝，再对这个拷贝做遍历。

在for循环里面的逻辑执行之前，这个拷贝的值就确定下来了。因此这个拷贝的长度和容量是不会在for循环的时候发生改变的。

以上面的题目为例：`range x `实际上是会先构造一个原切片`x`的拷贝，我们假设为`y`，然后对`y`做遍历。

```go
for i, s := range x {
		print(i, s, ",")
		x[i+1] = "M"
		x = append(x, "Z")
		x[i+1] = "Z"
}
```

上面这段代码可以等价为：

```go
y := x
for i := 0; i < len(y); i++ {
	print(i, y[i], ",")
	x[i+1] = "M"
	x = append(x, "Z")
	x[i+1] = "Z"
}
```



## slice扩容机制

通过`append`函数给`slice`添加元素的时候，有2种情况：

* 如果切片的容量足够，就会在切片指向的底层数组里追加元素。
* 如果切片的容量不足以承载新添加的元素，就会开辟一个新的底层数组，把原切片里的元素拷贝过来，再追加新的元素。切片结构里的指针会指向新的底层数组。



## 答案

我们回到本文最开始的题目，逐行解析每行代码的执行结果。

| 代码                            | 程序执行结果                                                 |
| ------------------------------- | ------------------------------------------------------------ |
| var x = []string{"A", "B", "C"} | x是一个切片，长度是3，容量是3，x指向的底层数组的值是[ "A" "B" "C"] |
| for i, s := range x             | 编译器先构造一个切片`x`的拷贝，假设为切片`y`，然后对`y`做遍历。`y`的值在for循环执行前就确定下来了，长度为3，容量为3，固定不变。 |
| print(i, s, ",")                | 第一次for循环，`i`的值是0，`s`的值是切片`y`里下标索引为0的元素，值为"A"，打印0A |
| x[i+1] = "M"                    | 执行x[1] = "M"，因为切片`x`和`y`现在指向同一个底层数组，切片`y`里下标索引为1的元素的值也被改成了"M"，`y`指向的底层数组的值为["A", "M", "C"] |
| x = append(x, "Z")              | 给切片`x`添加新元素"Z"，因为当前切片`x`的长度为3，容量为3，容量已满，不足以承载新增加的元素，所以要对`x`的底层数组做扩容，`x`指向新的底层数组，新底层数组的值是["A", "M", "C", "Z"]，`y`还是指向原来的底层数组，`y`指向的底层数组的值是["A", "M", "C"] |
| x[i+1] = "Z"                    | 切片`x`指向的底层数组的值变为["A", "Z", "C", "Z"]，切片`y`指向的底层数组的值不变，还是["A", "M", "C"] |
| 后续for循环                     | 因为从第2次for循环开始，`x`和`y`指向了不同的底层数组，所以对切片`x`的修改不会影响到`y`，因此后面打印的结果依次是1M,2C |

所以本题的答案是 0A, 1M, 2C。



## 总结

> 对于slice，时刻想着对slice做了修改后，slice里的3个字段：指针，长度，容量是怎么变的。
>
> ​																																		zen of go

* 对切片`x`做range遍历，实际上是对`x`的拷贝(假设为`y`)做range遍历，`y`的值(包括`y`结构体里指向底层数组的指针的值，`y`的长度和容量)都在执行`for`循环前确定下来了。
* 切片的底层数据结构和扩容机制，如果有不清楚的，参考我写的[slice底层原理篇](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483741&idx=1&sn=486066a3a582faf457f91b8397178f64&chksm=ce124e32f965c72411e2f083c22531aa70bb7fa0946c505dc886fb054b2a644abde3ad7ea6a0&token=609026015&lang=zh_CN#rd)，包含了slice的所有注意事项。



## 开源地址

文章和示例代码开源地址在GitHub: https://github.com/jincheng9/go-tutorial

公众号：coding进阶

个人网站：https://jincheng9.github.io/



## 思考题

留下2道思考题，欢迎大家在评论区留下你们的答案。也可以在我的wx公号发送消息`slice2`获取答案和原因。

* 题目1：

  ```go
  package main
  
  func main() {
  	var x = []string{"A", "B", "C"}
  
  	for i, s := range x {
  		print(i, s, " ")
  		x = append(x, "Z")
  		x[i+1] = "Z"
  	}
  }
  ```

  

* 题目2

  ```go
  package main
  
  func main() {
  	var y = []string{"A", "B", "C", "D"}
  	var x = y[:3]
  
  	for i, s := range x {
  		print(i, s, ",")
  		x = append(x, "Z")
  		x[i+1] = "Z"
  	}
  }
  ```

  

## References

* https://go101.org/quizzes/slice-2.html
* https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247483741&idx=1&sn=486066a3a582faf457f91b8397178f64&chksm=ce124e32f965c72411e2f083c22531aa70bb7fa0946c505dc886fb054b2a644abde3ad7ea6a0&token=609026015&lang=zh_CN#rd
* https://jincheng9.github.io/post/go-slice-principle/
* "For statements with `range` clause"： https://go.dev/ref/spec

