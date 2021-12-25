# Go Quiz: 从Go面试题看slice的底层原理和注意事项

## 面试题

最近Go 101的作者发布了11道Go面试题，非常有趣，打算写一个系列对每道题做详细解析。欢迎大家关注。

大家可以看下面这道关于`slice`的题目，通过这道题我们可以对`slice`的特性和注意事项有一个深入理解。

```go
package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3}
	x := a[:1]
	y := a[2:]
	x = append(x, y...)
	x = append(x, y...)
	fmt.Println(a, x)
}
```

- A: [0 1 2 3] [0 2 3 3 3]
- B: [0 2 3 3] [0 2 3 3 3]
- C: [0 1 2 3] [0 2 3 2 3]
- D: [0 2 3 3] [0 2 3 2 3]

大家可以在评论区留下你们的答案。这道题有几个考点：

1. `slice`的底层数据结构是什么？给`slice`赋值，到底赋了什么内容？
2. 通过`:`操作得到的新`slice`和原`slice`是什么关系？新`slice`的长度和容量是多少？
3. `append`在背后到底做了哪些事情？
4. `slice`的扩容机制是什么？



## 解析

我们先逐个解答上面的问题。

## slice的底层数据结构

`talk is cheap, show me the code`.  直接上`slice`的源码：

`slice`定义在`src/runtime/slice.go`第15行，源码地址：https://github.com/golang/go/blob/master/src/runtime/slice.go#L15。

`Pointer`定义在`src/unsafe/unsafe.go`第184行，源码地址：https://github.com/golang/go/blob/master/src/unsafe/unsafe.go#L184。

```go
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

type Pointer *ArbitraryType
```

`slice`实际上是一个结构体类型，包含3个字段，分别是

* array: 是指针，指向一个数组，切片的数据实际都存储在这个数组里。
* len: 切片的长度。
* cap: 切片的容量，表示切片当前最多可以存储多少个元素，如果超过了现有容量会自动扩容。

**因此给`slice`赋值，实际上都是给`slice`里的这3个字段赋值**。看起来这像是一句正确的废话，但是相信我，记住这句话可以帮助你非常清晰地理解对`slice`做修改后`slice`里3个字段的值是怎么变的，`slice` 指向的底层数组的数据是怎么变的。

## `:`分割操作符

`:`分割操作符有几个特点：

1. `:`可以对数组或者`slice`做数据截取

2. `:`得到的结果是一个新`slice`
3. 新`slice`的`array`指针指向原数组或者原`slice`的底层数组。
4. `:`分割操作符右边的数值有上限，上限有2种情况
   * 如果分割的是数组，那上限是是被分割的数组的长度。
   * 如果分割的是切片，那上限是被分割的切片的容量。**注意**，这个和下标操作不一样，如果使用下标索引访问切片，下标索引的最大值是(切片的长度-1)，而不是切片的容量。

一图胜千言，我们通过下面的示例来讲解下切片分割的机制。

下图表示`slice`结构，`ptr`表示`array`指针，指向底层数组，`len`和`cap`分别是切片的长度和容量。

![](./slice-struct.png)

step1: 我们通过代码`s := make([]byte, 5, 5)`来创建一个切片`s`，长度和容量都是5，结构示意如下：

![](./slice-1.png)

step2: 现在对切片`s`做分割`s2 := s[2:4]`，得到一个新切片`s2`，结构如下。

![](./slice-2.png)

* `s2`还是指向原切片`s`的底层数组，只不过指向的起始位置是下标索引为2的位置。
* `s2`的长度`len(s2)`是2，因为`s2 := s[2:4]`只是截取了切片`s`下标索引为2和3的2个元素。
* `s2`的容量`cap(s2)`是3，因为从`s2`指向的数组位置到底层数组末尾，可以存3个元素。

step3: 对切片`s`做分割`s3 := s2[:cap(s2)]`，得到一个新切片`s3`，结构如下：

![](./slice-3.png)

* `s3`指向切片`s2`的底层数组，同样也是`s`的底层数组，指向的起始位置是`s2`的起始位置，对应数组下标索引为2的位置。
* `s3`的长度`len(s3)`是3，因为`s3 := s2[:cap(s2)]`截取了切片`s2` 下标索引为0，1，2的3个元素。
* `s3`的容量`cap(s3)`是3，因为从`s3`指向的数组位置到底层数组末尾，可以存3个元素。

**因此，对数组或者切片做`:`分割操作产生的新切片还是指向原来的底层数组，并不会把原底层数组的元素拷贝一份到新的内存空间里。**



## append机制



## slice扩容机制



## 答案



## 加餐：copy机制



## 总结

> 对于slice，时刻想着对slice做了一个操作后，新的slice的指针，长度，容量是怎么变的

* slice数据结构

* `:`操作符

* 下标上下限是长度，`:`上下限是容量

* 拷贝赋值

* append赋值

* 给slice复制：指针赋值，是不是指向新地址， 长度赋值，容量赋值

* copy机制

* Go没有传引用这个说法，都是传值，可以参考我之前的文章xxx

  

## 开源地址

文章和代码开源地址在GitHub: https://github.com/jincheng9/go-tutorial

公众号：coding进阶

个人网站：https://jincheng9.github.io/



## References

* https://go101.org/quizzes/slice-1.html
* https://go.dev/blog/slices-intro