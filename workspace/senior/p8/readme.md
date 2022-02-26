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

### slice的底层数据结构

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

### `:`分割操作符

`:`分割操作符有几个特点：

1. `:`可以对数组或者`slice`做数据截取，`:`得到的结果是一个新`slice`。

2. **新`slice`结构体里的`array`指针指向原数组或者原`slice`的底层数组，新切片的长度是`：`右边的数值减去左边的数值，新切片的容量是原切片的容量减去`:`左边的数值。**

3. `:`的左边如果没有写数字，左边的默认值是0，右边如果没有写数字，右边的默认值是被分割的数组或被分割的切片的长度，**注意**，是长度不是容量。

   ```go
   a := make([]int, 0, 4) // a的长度是0，容量是4
   b := a[:] // 等价于 b := a[0:0], b的长度是0，容量是4
   c := a[:1] // 等价于 c := a[0:1], c的长度是1，容量是4
   d := a[1:] // 编译报错 panic: runtime error: slice bounds out of range
   e := a[1:4] // e的长度3，容量3
   ```

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
* 因为长度是2，所以只有`s2[0]`和`s2[1]`是有效的下标索引访问。但是，容量为3，`s2[0:3]`是一个有效的分割表达式。

step3: 对切片`s`做分割`s3 := s2[:cap(s2)]`，得到一个新切片`s3`，结构如下：

![](./slice-3.png)

* `s3`指向切片`s2`的底层数组，同样也是`s`的底层数组，指向的起始位置是`s2`的起始位置，对应数组下标索引为2的位置。
* `s3`的长度`len(s3)`是3，因为`s3 := s2[:cap(s2)]`截取了切片`s2` 下标索引为0，1，2的3个元素。
* `s3`的容量`cap(s3)`是3，因为从`s3`指向的数组位置到底层数组末尾，可以存3个元素。

**因此，对数组或者切片做`:`分割操作产生的新切片还是指向原来的底层数组，并不会把原底层数组的元素拷贝一份到新的内存空间里。**

正是因为他们指向同一块内存空间，所以对原数组或者原切片的修改会影响分割后的新切片的值，反之亦然。

### append机制

要了解append的机制，直接看源码说明。

```go
// The append built-in function appends elements to the end of a slice. If
// it has sufficient capacity, the destination is resliced to accommodate the
// new elements. If it does not, a new underlying array will be allocated.
// Append returns the updated slice. It is therefore necessary to store the
// result of append, often in the variable holding the slice itself:
//	slice = append(slice, elem1, elem2)
//	slice = append(slice, anotherSlice...)
// As a special case, it is legal to append a string to a byte slice, like this:
//	slice = append([]byte("hello "), "world"...)
func append(slice []Type, elems ...Type) []Type
```

* append函数返回的是一个切片，append在原切片的末尾添加新元素，这个末尾是切片长度的末尾，不是切片容量的末尾。

  ```go
  func test() {
  	a := make([]int, 0, 4)
  	b := append(a, 1) // b=[1], a指向的底层数组的首元素为1，但是a的长度和容量不变
  	c := append(a, 2) // a的长度还是0，c=[2], a指向的底层数组的首元素变为2
  	fmt.Println(a, b, c) // [] [2] [2]
  }
  ```

* 如果原切片的容量足以包含新增加的元素，那append函数返回的切片结构里3个字段的值是：

  * array指针字段的值不变，和原切片的array指针的值相同，也就是append是在原切片的底层数组添加元素，返回的切片还是指向原切片的底层数组
  * len长度字段的值做相应增加，增加了N个元素，长度就增加N
  * cap容量不变

* 如果原切片的容量不够存储append新增加的元素，Go会先分配一块容量更大的新内存，然后把原切片里的所有元素拷贝过来，最后在新的内存里添加新元素。append函数返回的切片结构里的3个字段的值是：

  * array指针字段的值变了，不再指向原切片的底层数组了，会指向一块新的内存空间
  * len长度字段的值做相应增加，增加了N个元素，长度就增加N
  * cap容量会增加

**注意**：append不会改变原切片的值，原切片的长度和容量都不变，除非把append的返回值赋值给原切片。

那么问题来了，新切片的容量是按照什么规则计算得出来的呢？

### slice扩容机制

`slice`的扩容机制随着`Go`的版本迭代，是有变化的。目前网上大部分的说法是下面这个：

> 当原 slice 容量小于 `1024` 的时候，新 slice 容量变成原来的 `2` 倍；原 slice 容量超过 `1024`，新 slice 容量变成原来的`1.25`倍。

这里明确告诉大家，这个结论是**错误**的。

`slice`扩容的源码实现在`src/runtime/slice.go`里的`growslice`函数，源码地址：https://github.com/golang/go/blob/master/src/runtime/slice.go。

Go 1.18的扩容实现代码如下，`growslice`的参数et是切片里的元素类型，old是原切片，cap等于原切片的长度+append新增的元素个数。(**注意第3个参数cap的值是原切片的长度+append新增元素个数，不是原切片容量+新增元素个数，可以在growslice里打印cap的值来验证**)

```go
func growslice(et *_type, old slice, cap int) slice {
	// ...
	newcap := old.cap
	doublecap := newcap + newcap
	if cap > doublecap {
		newcap = cap
	} else {
		const threshold = 256
		if old.cap < threshold {
			newcap = doublecap
		} else {
			// Check 0 < newcap to detect overflow
			// and prevent an infinite loop.
			for 0 < newcap && newcap < cap {
				// Transition from growing 2x for small slices
				// to growing 1.25x for large slices. This formula
				// gives a smooth-ish transition between the two.
				newcap += (newcap + 3*threshold) / 4
			}
			// Set newcap to the requested cap when
			// the newcap calculation overflowed.
			if newcap <= 0 {
				newcap = cap
			}
		}
	}

	var overflow bool
	var lenmem, newlenmem, capmem uintptr
	// Specialize for common values of et.size.
	// For 1 we don't need any division/multiplication.
	// For sys.PtrSize, compiler will optimize division/multiplication into a shift by a constant.
	// For powers of 2, use a variable shift.
	switch {
	case et.size == 1:
		lenmem = uintptr(old.len)
		newlenmem = uintptr(cap)
		capmem = roundupsize(uintptr(newcap))
		overflow = uintptr(newcap) > maxAlloc
		newcap = int(capmem)
	case et.size == goarch.PtrSize:
		lenmem = uintptr(old.len) * goarch.PtrSize
		newlenmem = uintptr(cap) * goarch.PtrSize
		capmem = roundupsize(uintptr(newcap) * goarch.PtrSize)
		overflow = uintptr(newcap) > maxAlloc/goarch.PtrSize
		newcap = int(capmem / goarch.PtrSize)
	case isPowerOfTwo(et.size):
		var shift uintptr
		if goarch.PtrSize == 8 {
			// Mask shift for better code generation.
			shift = uintptr(sys.Ctz64(uint64(et.size))) & 63
		} else {
			shift = uintptr(sys.Ctz32(uint32(et.size))) & 31
		}
		lenmem = uintptr(old.len) << shift
		newlenmem = uintptr(cap) << shift
		capmem = roundupsize(uintptr(newcap) << shift)
		overflow = uintptr(newcap) > (maxAlloc >> shift)
		newcap = int(capmem >> shift)
	default:
		lenmem = uintptr(old.len) * et.size
		newlenmem = uintptr(cap) * et.size
		capmem, overflow = math.MulUintptr(et.size, uintptr(newcap))
		capmem = roundupsize(capmem)
		newcap = int(capmem / et.size)
	}
	// ...
	return slice{p, old.len, newcap}
}
```

newcap是扩容后的容量，先根据原切片的长度、容量和要添加的元素个数确定newcap大小，最后再对newcap做内存对齐得到最后的newcap。



## 答案

我们回到本文最开始的题目，逐行解析每行代码的执行结果。

| 代码                      | 切片对应结果                                                 |
| ------------------------- | ------------------------------------------------------------ |
| a := [...]int{0, 1, 2, 3} | a是一个数组，长度是4，值是[ 0 1 2 3]                         |
| x := a[:1]                | x是一个切片，切片里的指针指向数组a的首元素，值是[0]，长度1，容量4 |
| y := a[2:]                | y是一个切片，切片里的指针指向数组a的第2个元素，值是[2 3]，长度2，容量2 |
| x = append(x, y...)       | x的剩余容量还有3个，足以存储y里的2个元素，所以x不会扩容，x的值是[0 2 3]，长度3，容量4。因为x, a, y都指向同一块内存空间，所以x的修改影响了a和y。<br>a的值变为[0 2 3 3]，长度4，容量4<br>y的值变为[3 3]，长度2，容量2 |
| x = append(x, y...)       | x的剩余容量只有1个，不足以存储y里的2个元素，所以要扩容。append(x, y)的结果是得到一个新切片，值是[0 2 3 3 3]，长度5，容量8。<br>append的返回值赋值给x，所以切片x会指向扩容后的新内存。 |
| fmt.Println(a, x)         | a的值还是[0 2 3 3]没有变化，所以打印结果是[0 2 3 3] [0 2 3 3 3 ]，答案是B |



## 加餐：copy机制

Go的内置函数`copy`可以把一个切片里的元素拷贝到另一个切片，源码定义在`src/builtin/builtin.go`，代码如下：

```go
// The copy built-in function copies elements from a source slice into a
// destination slice. (As a special case, it also will copy bytes from a
// string to a slice of bytes.) The source and destination may overlap. Copy
// returns the number of elements copied, which will be the minimum of
// len(src) and len(dst).
func copy(dst, src []Type) int
```

`copy`会从原切片`src`拷贝 `min(len(dst), len(src))`个元素到目标切片`dst`，

因为拷贝的元素个数`min(len(dst), len(src))`不会超过目标切片的长度`len(dst)`，所以`copy`执行后，目标切片的长度不会变，容量不会变。

**注意**：原切片和目标切片的内存空间可能会有重合，`copy`后可能会改变原切片的值，参考下例。

```go
package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := a[1:] // [2 3]
	copy(a, b) // a和b内存空间有重叠
	fmt.Println(a, b) // [2 3 3] [3 3]
}
```



## slice打印和底层数组地址

打印要弄清楚3个问题：

1. `fmt.Println(slice)`打印到切片底层数组的哪个元素截止？

   根据切片的长度len，打印到下标索引为`len-1`的元素截止。比如下例里，虽然切片a的底层数组下标索引`len(a)-1`后面还有个值1，但是因为a的长度为1，就只打印`[0]`，切片b的长度为2，所以会打印`[0 1]`。

   ```go
   a := make([]int, 1, 4) // a的长度是1，容量是4
   b := append(a, 1) // 往a的末尾添加元素1，b=[0 1], a的长度还是1，a和b指向同一个底层数组
   fmt.Println(a, b) // [0] [0 1]
   ```

2. 如何打印`slice`结构体变量的地址？

   ```go
   s := []int{1, 2}
   fmt.Printf("%p\n", &s)
   ```

3. 如何打印`slice`底层数组的地址？有2种方法

   ```go
   s = make([]int, 2, 3)
   fmt.Printf("%p %p\n", s, &s[0])
   ```



## 总结

> 对于slice，时刻想着对slice做了修改后，slice里的3个字段：指针，长度，容量是怎么变的。

* `slice`是一个结构体类型，里面包含3个字段：指向数组的`array`指针，长度`len`和容量`cap`。给slice赋值是对`slice`里的指针，长度和容量3个字段分别赋值。

* `:`分割操作符的结果是一个新切片，**新`slice`结构体里的`array`指针指向原数组或者原`slice`的底层数组，新切片的长度是`：`右边的数值减去左边的数值，新切片的容量是原切片的容量减去`:`左边的数值。**

* `:`分割操作符右边的数值上限有2种情况：

  * 如果分割的是数组，那上限是是被分割的数组的长度。
  * 如果分割的是切片，那上限是被分割的切片的容量。**注意**，这个和下标操作不一样，如果使用下标索引访问切片，下标索引的最大值是(切片的长度-1)，而不是切片的容量。

* 对于`append`操作和`copy`操作，要清楚背后的执行逻辑。

* 打印`slice`时，是根据`slice`的长度来打印的

  ```go
  a := make([]int, 1, 4) // a的长度是1，容量是4
  b := append(a, 1) // 往a的末尾添加元素1，b=[0 1], a的长度还是1，a和b指向同一个底层数组
  fmt.Println(a, b) // [0] [0 1]
  fmt.Printf("%p %p\n", a, b) // 切片a和b的底层数组地址相同
  ```

* Go在函数传参时，没有传引用这个说法，只有传值。网上有些文章写Go的`slice`，`map`，`channel`作为函数参数是传引用，这是错误的，可以参考我之前的文章[Go有引用变量和引用传递么？](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p3)

  

## 开源地址

文章和示例代码开源地址在GitHub: https://github.com/jincheng9/go-tutorial

公众号：coding进阶

个人网站：https://jincheng9.github.io/



## 思考题

留下2道思考题，欢迎大家在评论区留下你们的答案。

* 题目1：

  ```go
  package main
  
  import "fmt"
  
  func main() {
  	a := []int{1, 2}
  	b := append(a, 3)
  
  	c := append(b, 4)
  	d := append(b, 5)
  
  	fmt.Println(a, b, c[3], d[3])
  }
  ```

  

* 题目2

  ```go
  package main
  
  import "fmt"
  
  func main() {
  	s := []int{1, 2}
  	s = append(s, 4, 5, 6)
  	fmt.Println(len(s), cap(s))
  }
  ```

  

## References

* https://go101.org/quizzes/slice-1.html
* https://go.dev/blog/slices-intro
* https://github.com/golang/go/blob/master/src/runtime/slice.go#L156
* https://qcrao91.gitbook.io/go/shu-zu-he-qie-pian/qie-pian-de-rong-liang-shi-zen-yang-zeng-chang-de