# Go 1.20要来了，看看都有哪些变化-第1篇

## 前言

Go官方团队在2022.12.08发布了Go 1.20 rc1版本，Go 1.20的正式release版本预计会在2023年2月份发布。

让我们先睹为快，看看Go 1.20给我们带来了哪些变化。(文末有彩蛋！)

安装方法：

```bash
$ go install golang.org/dl/go1.20rc1@latest
$ go1.20rc1 download
```

这是Go 1.20版本更新内容详解的第1篇，欢迎大家关注公众号，及时获取本系列最新更新。

## Go 1.20发布清单

和Go 1.19相比，改动内容适中，主要涉及语言(Language)、可移植性(Ports)、工具链(Go Tools)、运行时(Runtime)、编译器(Compiler)、汇编器(Assembler)、链接器(Linker)和核心库(Core library)等方面的优化。

我们逐个看看具体都有哪些变化。

### 语言变化

Go 1.20在语言层面带来了4个变化。

#### slice转数组

Go1.17在语言层面开始支持将slice转为指向数组的指针。

示例如下：

```go
s := make([]byte, 2, 4)
// 将s这个slice转为指向byte数组的指针s0
// 其中[0]byte里的0表示数组的长度，虽然长度为0，但值不等于nil
s0 := (*[0]byte)(s)      // s0 != nil
fmt.Printf("%T")
// 将s[1:]这个slice转为指向byte数组的指针s1
// s1指向的数组的长度为1
s1 := (*[1]byte)(s[1:])  // &s1[0] == &s[1]
// 将s这个slice转为指向byte数组的指针s2
// s2指向的数组的长度为2 
s2 := (*[2]byte)(s)      // &s2[0] == &s[0]
// 将s这个slice转为指向byte数组的指针s4
// s4指向的数组的长度为4 
s4 := (*[4]byte)(s)      // panics: len([4]byte) > len(s)
```

注意：slice转为指向数组的指针时，如果数组定义的长度超过了slice的长度，会抛panic。

所以上面`s4 := (*[4]byte)(s)`这行代码虽然可以编译通过，但是会出现runtime panic。

Go 1.20之前不支持将slice直接转为数组，如果要转，得先转为指向数组的指针，再转为数组，如下面代码所示：

```go
s := make([]byte, 2, 4)
s[0] = 100

s1 := (*[1]byte)(s[1:]) // &s1[0] == &s[1]
s2 := (*[2]byte)(s)     // &s2[0] == &s[0]
fmt.Printf("%T, %v, %p, %p\n", s1, s1[0], &s1[0], &s[1])
fmt.Printf("%T, %v, %v, %p\n", s2, s2[0], &s2[0], s)
// a1数组里元素的地址和s1指向的数组的元素地址不一样，a2同理
a1 := *s1
a2 := *s2
fmt.Printf("%T, %v, %p, %p\n", a1, a1[0], &a1[0], &s1[0])
fmt.Printf("%T, %v, %p, %p\n", a2, a2[0], &a2[1], &s2[1])
```

从Go 1.20开始，支持将slice直接转为数组，如下面代码所示：

```go
s := make([]byte, 2, 4)
s[0] = 100
s1 := [1]byte(s[1:])
s2 := [2]byte(s)
// s1数组里元素的地址和s指向的数组的元素地址不一样，s2同理
fmt.Printf("%T, %v, %p, %p\n", s1, s1[0], &s1[0], &s[1])
fmt.Printf("%T, %v, %v, %p\n", s2, s2[0], &s2[0], s)
```

总结：

* slice转为指向数组的指针后，这个指针会指向和slice相同的地址空间
* slice转为数组时，会把slice底层数组的值拷贝一份出来。转换后得到的数组的地址空间和slice底层数组空间不一样。

还有几个语法细节可以参考如下代码示例：

```go
var t []string
t0 := [0]string(t)       // ok for nil slice t
t1 := (*[0]string)(t)    // t1 == nil
t2 := (*[1]string)(t)    // panics: len([1]string) > len(t)

u := make([]byte, 0)
u0 := (*[0]byte)(u)      // u0 != nil
```

Go 1.20 extends this to allow conversions from a slice to an array: given a slice `x`, `[4]byte(x)` can now be written instead of `*(*[4]byte)(x)`.



#### unsafe

go标准库里的unsafe package定义了3个新的函数：

```go
func SliceData(slice []ArbitraryType) *ArbitraryType
func String(ptr *byte, len IntegerType) string
func StringData(str string) *byte
```

Go 1.17版本在unsafe package里引入过Slice函数，如下所示：

```go
func Slice(ptr *ArbitraryType, len IntegerType) []ArbitraryType
```

有了这4个函数，可以构造和解构slice和string，不需要依赖它们的真实表示。

具体细节可以参考：https://tip.golang.org/ref/spec#Package_unsafe



The [`unsafe` package](https://tip.golang.org/ref/spec/#Package_unsafe) defines three new functions `SliceData`, `String`, and `StringData`. Along with Go 1.17's `Slice`, these functions now provide the complete ability to construct and deconstruct slice and string values, without depending on their exact representation.



#### 值比较



The specification now defines that struct values are compared one field at a time, considering fields in the order they appear in the struct type definition, and stopping at the first mismatch. The specification could previously have been read as if all fields needed to be compared beyond the first mismatch. 

Similarly, the specification now defines that array values are compared one element at a time, in increasing index order. 



In both cases, the difference affects whether certain comparisons must panic. Existing programs are unchanged: the new spec wording describes what the implementations have always done.

#### Comparable类型

[Comparable types](https://tip.golang.org/ref/spec#Comparison_operators) (such as ordinary interfaces) may now satisfy `comparable` constraints, even if the type arguments are not strictly comparable (comparison may panic at runtime). This makes it possible to instantiate a type parameter constrained by `comparable` (e.g., a type parameter for a user-defined generic map key) with a non-strictly comparable type argument such as an interface type, or a composite type containing an interface type.



## 可移植性

### Darwin and iOS

Go 1.20将会成为支持macOS 10.13 High Sierra和10.14 Mojave的最后一个版本。

如果未来想在mac电脑上使用Go 1.21或者更新的Go版本，只能用macOS 10.15 Catalina和更新的macOS版本。

### FreeBSD/RISC-V

Go 1.20增加了对于RISC-V架构在FreeBSD操作系统的实验性支持。

```bash
GOOS=freebsd
GOARCH=riscv64
```

## 总结

下一篇会介绍Go 1.20在Go Tool工具链、运行时、编译器、汇编器、链接器和核心库的优化工作，有一些内容值得学习，欢迎大家保持关注。



## 开源地址

文章和示例代码开源在GitHub: [Go语言初级、中级和高级教程](https://github.com/jincheng9/go-tutorial)。

公众号：coding进阶。关注公众号可以获取最新Go面试题和技术栈。

个人网站：[Jincheng's Blog](https://jincheng9.github.io/)。

知乎：[无忌](https://www.zhihu.com/people/thucuhkwuji)。



## 福利

我为大家整理了一份后端开发学习资料礼包，包含编程语言入门到进阶知识(Go、C++、Python)、后端开发技术栈、面试题等。

关注公众号「coding进阶」，发送消息 **backend** 领取资料礼包，这份资料会不定期更新，加入我觉得有价值的资料。还可以发送消息「**进群**」，和同行一起交流学习，答疑解惑。



最后送上一个彩蛋，Go标准库的脑图，想学习Go标准库的可以参考这个来。

![](../../img/go-std-lib-mindmap.png)

## References

* https://tip.golang.org/doc/go1.20