# Go Quiz: 从Go面试题看数值类型的自动推导

##  背景

Google工程师Valentin Deleplace出了几道关于Go数值类型的计算题，很有迷惑性，整体正确率不到50%，拿出来和大家分享下。

### 题目1

```go
var y = 5.2
const z = 2
fmt.Println(y / z)
```

* A: `2.6`

* B: `2.5 `

* C: `2`

* D: 编译错误

### 题目2

```go
const a = 7.0
var b = 2
fmt.Println(a / b)
```

* A: `3.5`
* B: `3 `
* C: 编译错误

### 题目3

```go
a := 40
f := float64(a/100.0)
fmt.Println(f)
```

* A: `0`

* B: `0.4`

* C: 编译错误

  


## 解析

这道题主要考察3个知识点：

* 对于变量而言，如果没有显式指定数据类型，编译器会根据赋值自动推导出确定的数据类型。

  整数的默认类型是`int`，浮点数的默认类型是`float64`，官方说明如下：

  > An untyped constant has a *default type* which is the type to which the constant is implicitly converted in contexts where a typed value is required, for instance, in a [short variable declaration](https://go.dev/ref/spec#Short_variable_declarations) such as `i := 0` where there is no explicit type. The default type of an untyped constant is `bool`, `rune`, `int`, `float64`, `complex128` or `string` respectively, depending on whether it is a boolean, rune, integer, floating-point, complex, or string constant.

* 对于常量而言，**如果没有显式指定数据类型**，编译器同样会推导出一个数据类型，**但是没有显式指定数据类型的常量在代码上下文里可以根据需要隐式转化为需要的数据类型进行计算**。

* Go不允许不同的数据类型做运算。当变量和**没有显式指定数据类型的常量**混合在一起运算时，如果常量转化成变量的类型不会损失精度，那常量会自动转化为变量的数据类型参与运算。如果常量转化成变量的类型会损失精度，那就会编译报错。

对于题目1：变量`y` 没有显式指定数据类型，但是根据后面的赋值`5.2`，编译器自动推导出变量`y`的数据类型是float64。常量`z`没有显式指定数据类型，编译器自动推导出的类型是`int`，但是在运算`y/z`时，因为`y`是float64类型，`z`转化为float64类型不会损失精度，所以`z`在运算时会自动转换为float64类型，所以本题的运算结果是`2.6`，答案是`A`。

对于题目2：变量`b`没有显式指定数据类型，根据后面的赋值`2`，编译器自动推导出变量`b`的数据类型是int。常量`a`没有显式指定数据类型，编译器自动推导出的类型是float64，但是在运算`a/b`时，因为`b`是int类型，`a`转换为int类型不会损失精度，所以`a`在运算时会自动转换为int类型参与计算，所以本题的结果是`7/2`，结果是`3`，答案是`B`。

对于题目3：变量`a`没有显式指定数据类型，根据后面的赋值`40`，编译器自动推导出变量`a`的数据类型是int。常量`100.0`没有显式指定数据类型，编译器自动推导出的类型是float64，但是在运算`a/100.0`时，因为`a`是int类型，`100.0`转换为int类型不会损失精度，所以`100.0`在运算时会自动转换为int类型参与计算，所以本题的结果是`40/100`，结果是`0`，答案是`A`。



##  思考题

题目1：

``` go
var (
    a = 1.0
    b = 2
)
fmt.Println(a / b)
```

题目2：

```go
const (
    x = 5.0
    y = 4
)
fmt.Println(x / y)
```

题目3：

```go
const t = 4.8
var u = 2
fmt.Println(t / u)
```

题目4：

```go
d := 9.0
const f int = 2
fmt.Println(d / f)
```

想知道答案的可以给公众号发送消息`data`获取答案。



## 总结

* 对于未指定数据类型的变量，编译器会自动推导出默认的数据类型，在参与运算时，变量始终用这个推导出来的数据类型参与运算，不会做任何隐式类型转换。

* 对于未指定数据类型的常量，编译器虽然也会自动推导出默认的数据类型，但是在参与运算时，常量可以根据代码的上下文，自动隐式转换为所需要的数据类型，只要不出现精度丢失即可。如果出现精度丢失，那就会编译报错。

  

## 开源地址

文章和示例代码开源在GitHub: [Go语言初级、中级和高级教程](https://github.com/jincheng9/go-tutorial)。

公众号：coding进阶。关注公众号可以获取最新Go面试题和技术栈。

个人网站：[Jincheng's Blog](https://jincheng9.github.io/)。

知乎：[无忌](https://www.zhihu.com/people/thucuhkwuji)。



## References

* https://twitter.com/val_deleplace/status/1530470260079480832
* https://twitter.com/val_deleplace/status/1529378134684057601
* https://twitter.com/val_deleplace/status/1529917910252146705
* https://twitter.com/val_deleplace/status/1530104838587092993
* https://twitter.com/val_deleplace/status/1529050473403142146
* https://twitter.com/val_deleplace/status/1531963904032747522
* https://go.dev/ref/spec#Constants
* https://stackoverflow.com/questions/61153803/how-does-the-constant-value-auto-type-work-in-golang