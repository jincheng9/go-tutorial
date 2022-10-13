# 一文读懂Go泛型设计和使用场景

## 前言

2021.12.14日，Go官方正式发布了支持泛型的Go 1.18beta1版本，这是Go语言自2007年诞生以来，最重大的功能变革。

泛型核心就3个概念：

1. Type parameters for functions and types

   类型参数，可以用于泛型函数以及泛型类型

2. Type sets defined by interfaces

   Go 1.18之前，interface用来定义方法集( a set of methods)。

   Go 1.18开始，还可以使用interface来定义类型集(a set of types)，作为类型参数的Type constraint(类型限制)

3. Type inference

   类型推导，可以帮助我们在写代码的时候不用传递类型实参(Type Arguments)，由编译器自行推导。

   **注意**：类型推导并不是永远都可行。

## Type parameters(类型参数)

```go
[P, Q constraint1, R constraint2]
```

这里定义了一个类型参数列表(type parameter list)，列表里可以包含一个或者多个类型参数。

`P，Q`和`R`都是类型参数，`contraint1`和`contraint2`都是类型限制(type constraint)。

* 类型参数列表使用方括号`[]`
* 类型参数建议首字母大写，用来表示它们是类型

先看一个简单示例：

```go
func min(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}
```

这个例子，只能计算2个`float64`中的较小者。有泛型之前，如果我们要支持计算2个int或者其它数值类型的较小者，就需要实现新的函数、或者使用`interface{}`，或者使用`Reflect`。

对于这个场景，使用泛型代码更简洁，效率也更优。支持比较不同数值类型的泛型min函数实现如下：

```go
func min[T constraints.Ordered] (x, y T) T {
	if x < y {
		return x
	}
	return y
}

// 调用泛型函数
m := min[int](2, 3)
```

注意：

1. 使用`constraints.Ordered`类型，需要`import constraints`。
2. `min[int](2, 3)`在编译期会对泛型函数`min`实例化(instantiation)，将泛型函数里的类型参数`T`替换为`int`，在运行期就是实例化之后的函数调用了，不是泛型函数调用了。

### instantiation(实例化)

泛型函数的实例化做2个事情

1. 把泛型函数的类型参数替换为类型实参(type argument)。

   比如上面的例子，min函数调用传递的类型实参是`int`，会把泛型函数的类型参数`T`替换为`int`

2. 检查类型实参是否满足泛型函数定义里的类型限制。

   对于上例，就是检查类型实参`int`是否满足类型限制`constraints.Ordered`。

任何一步失败了，那泛型函数的实例化就失败了，也就是泛型函数调用就失败了。

泛型函数实例化后就生成了一个非泛型函数，用于真正的函数执行。

上面的`min[int](2, 3)`调用还可以替换为如下代码：

```go
func min[T constraints.Ordered] (x, y T) T {
	if x < y {
		return x
	}
	return y
}

// 方式1
m := min[int](2, 3)
// 方式2
fmin := min[int]
m2 := fmin(2, 3)
```

`min[int](2, 3)`会被编译器解析成`(min[int])(2, 3)`，也就是

1. 先实例化得到一个非泛型函数
2. 然后再做真正的函数执行。

### generic types(泛型类型)

类型参数除了用于泛型函数之外，还可以用于Go的类型定义，来实现泛型类型(generic types)。

看如下代码示例，实现了一个泛型二叉树结构

```go
type Tree[T interface{}] struct {
	left, right *Tree[T]
	data T
}

func (t *Tree[T]) Lookup(x T) *Tree[T] 

var stringTree Tree[string]
```

二叉树节点存储的数据类型可能是多样的，有的二叉树存储`int`，有的存储`string`等等。

使用泛型，可以让`Tree`这个结构体类型支持二叉树节点存储不同的数据类型。

对于泛型类型的方法，需要在方法接收者声明对应的类型参数。比如上例里的Lookup方法，在指针接收者`*Tree[T]`里声明了类型参数`T`。

## type sets(类型集)

类型参数的类型限制约定了该类型参数允许的具体类型。

类型限制往往包含了多个具体类型，这些具体类型就构成了类型集。

```go
func min[T constraints.Ordered] (x, y T) T {
	if x < y {
		return x
	}
	return y
}
```

比如上面的例子，类型参数`T`的类型限制是`constraints.Ordered`，`contraints.Ordered`包含了非常多的具体类型，定义如下：

```go
// Ordered is a constraint that permits any ordered type: any type
// that supports the operators < <= >= >.
// If future releases of Go add new ordered types,
// this constraint will be modified to include them.
type Ordered interface {
  Integer | Float | ~string
}
```

`Integer`和`Float`也是定义在`constraints`这个包里的类型限制，

**类型参数列表不能用于方法，只能用于函数**。

```go
type Foo struct {}

func (Foo) bar[T any](t T) {}
```

上面的例子在结构体类型`Foo`的方法`bar`使用了类型参数列表，编译会报错：

```bash
./example1.go:30:15: methods cannot have type parameters
./example1.go:30:16: invalid AST: method must have no type parameters
```

个人认为Go的这个编译提示：`methods cannot have type paramters` 不是特别准确。

比如下面的例子，就是在方法`bar`里用到了类型参数`T`，改成`methods cannot have type paramter list`感觉会更好。

```go
type Foo[T any] struct {}

func (Foo[T]) bar(t T) {}
```

**注意**：类型限制必须是`interface`类型。比如上例的`constraints.Ordered`就是一个`interface`类型。

### | 和 ~

`|`: 表示取并集。比如下例的`Number`这个interface可以作为类型限制，用于限定类型参数必须是int，int32和int64这3种类型。

```go
type Number interface{
	int | int32 | int64
}
```

`~T`: `~` 是Go 1.18新增的符号，`~T`表示底层类型是T的所有类型。`~`的英文读作tilde。

* 例1：比如下例的`AnyString`这个interface可以作为类型限制，用于限定类型参数的底层类型必须是string。`string`本身以及下面的`MyString`都满足`AnyString`这个类型限制。

  ```go
  type AnyString interface{
     ~string
  }
  type MyString string
  ```

* 例2：再比如，我们定义一个新的类型限制叫`customConstraint`，用于限定底层类型为`int`并且实现了`String() string`方法的所有类型。下面的`customInt`就满足这个type constraint。

  ```go
  type customConstraint interface {
     ~int
     String() string
  }
  
  type customInt int
  
  func (i customInt) String() string {
     return strconv.Itoa(int(i))
  }
  ```

类型限制有2个作用：

1. 用于约定有效的类型实参，不满足类型限制的类型实参会被编译器报错。
2. 如果类型限制里的所有类型都支持某个操作，那在代码里，对应的类型参数就可以使用这个操作。

### constraint literals(类型限制字面值)

type constraint既可以提前定义好，也可以在type parameter list里直接定义，后者就叫constraint literals。

```go
[S interface{~[]E}, E interface{}]

[S ~[]E, E interface{}]

[S ~[]E, E any]
```

几个注意点：

* 可以直接在方括号[]里，直接定义类型限制，即使用类型限制字面值，比如上例。
* 在类型限制的位置，`interface{E}`也可以直接写为`E`，因此就可以理解`interface{~[]E}`可以写为`~[]E`。
* `any`是Go 1.18新增的预声明标识符，是`interface{}`的别名。

## constraints包

**update 2022.02.03**：Russ Cox在2022.02.03宣布Go 1.18正式版本会从标准库移除`constraints`包，所以这个章节大家可以快速跳过。

`constraints`包定义了一些常用的类型限制，整个包除了测试代码，就1个`constraints.go`文件，50行代码，源码地址：

https://github.com/golang/go/blob/master/src/constraints/constraints.go

包含的类型限制如下：

* `constraints.Signed`

  ```go
  type Signed interface {
  	~int | ~int8 | ~int16 | ~int32 | ~int64
  }
  ```

* `constraints.Unsigned`

  ```go
  type Unsigned interface {
  	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
  }
  ```

* `constraints.Integer`

  ```go
  type Integer interface {
  	Signed | Unsigned
  }
  ```

* `constraints.Float`

  ```go
  type Float interface {
  	~float32 | ~float64
  }
  ```

* `constraints.Complex`

  ```go
  type Complex interface {
  	~complex64 | ~complex128
  }
  ```

* `constraints.Ordered`

  ```go
  type Ordered interface {
  	Integer | Float | ~string
  }
  ```

## Type inference(类型推导)

我们看下面的代码示例：

```go
func min[T constraints.Ordered] (x, y T) T {
	if x < y {
		return x
	}
	return y
}

var a, b, m1, m2 float64
// 方式1：显式指定type argument
m1 = min[float64](a, b)
// 方式2：不指定type argument，让编译器自行推导
m2 = min(a, b)
```

方式2没有传递类型实参，编译器是根据函数实参`a`和`b`推导出类型实参。

类型推导可以让我们的代码更简洁，更具可读性。

Go泛型有2种类型推导：

1. function argument type inference: deduce type arguments from the types of the non-type arguments.

   通过函数的实参推导出来具体的类型。比如上面例子里的`m2 = min(a, b)`，就是根据`a`和`b`这2个函数实参

   推导出来`T`是`float64`。

2. constraint type inference: inferring a type argument from another type argument, based on type parameter constraints.

   通过已经确定的类型实参，推导出未知的类型实参。下面的代码示例里，根据函数实参2不能确定`E`是什么类型，但是可以确定`S`是`[]int32`，再结合类型限制里`S`的底层类型是`[]E`，可以推导出`E`是int32，int32满足`constraints.Integer`限制，因此推导成功。

   ```go
   type Point []int32
   
   func ScaleAndPrint(p Point) {
     r := Scale(p, 2)
     fmt.Println(r)
   }
   
   func Scale[S ~[]E, E constraints.Integer](s S, c E) S {
     r := make(S, len(s))
     for i, v := range s {
       r[i] = v * c
     }
     return r
   }
   ```

**类型推导并不是一定成功**，比如类型参数用在函数的返回值或者函数体内，这种情况就必须指定类型实参了。

```go
func test[T any] () T {
  var result T
  return result
}
func test[T any] () {
  var result T
  fmt.Println(result)
}
```

更深入了解type inference可以参考：https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#type-inference



## 使用场景

### 箴言

> Write code, don't design types.

在写Go代码的时候，对于泛型，Go泛型设计者`Ian Lance Taylor`建议不要一上来就定义type parameter和type constraint，如果你一上来就这么做，那就搞错了泛型的最佳实践。

先写具体的代码逻辑，等意识到需要使用type parameter或者定义新的type constraint的时候，再加上type parameter和type constraint。

### 什么时候使用泛型？

* 需要使用slice, map, channel类型，但是slice, map, channel里的元素类型可能有多种。

* 通用的数据结构，比如链表，二叉树等。下面的代码实现了一个支持任意数据类型的二叉树。

  ```go
  type Tree[T any] struct {
    cmp func(T, T) int
    root *node[T]
  }
  
  type node[T any] struct {
    left, right *node[T]
    data T
  }
  
  func (bt *Tree[T]) find(val T) **node[T] {
    pl := &bt.root
    for *pl != nil {
      switch cmp := bt.cmp(val, (*pl).data); {
        case cmp < 0 : pl = &(*pl).left
        case cmp > 0 : pl = &(*pl).right
      default: return pl
      }
    }
    return pl
  }
  ```

* 当一个方法的实现对所有类型都一样。

  ```go
  type SliceFn[T any] struct {
    s []T
    cmp func(T, T) bool
  }
  
  func (s SliceFn[T]) Len() int{return len(s.s)}
  func (s SliceFn[T]) Swap(i, j int) {
    s.s[i], s.s[j] = s.s[j], s.s[i]
  }
  func (s SliceFn[T]) Less(i, j int) bool {
    return s.cmp(s.s[i], s.s[j])
  }
  ```

### 什么时候不要使用泛型?

1. 只是单纯调用实参的方法时，不要用泛型。

   ```go
   // good
   func foo(w io.Writer) {
      b := getBytes()
      _, _ = w.Write(b)
   }
   
   // bad
   func foo[T io.Writer](w T) {
      b := getBytes()
      _, _ = w.Write(b)
   }
   ```

   比如上面的例子，单纯是调用`io.Writer`的`Write`方法，把内容写到指定地方。使用`interface`作为参数更合适，可读性更强。

2. 当函数或者方法或者具体的实现逻辑，对于不同类型不一样时，不要用泛型。比如`encoding/json`这个包使用了`reflect`，如果用泛型反而不合适。

   

###  总结

> Avoid boilerplate.
>
> Corollary: Don't use type parameters prematurely; wait until you are about to write boilerplate code.

不要随便使用泛型，Ian给的建议是：当你发现针对不同类型，会写出同样的代码逻辑时，才去使用泛型。也就是

`Avoid boilerplate code`。

Go语言里`interface`和`reflect`可以在某种程度上实现泛型，我们在处理多种类型的时候，要考虑具体的使用场景，切勿盲目用泛型。

想更加深入了解Go泛型设计原理的可以参考Go泛型设计作者Ian和Robert写的Go Proposal：

https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md



## 开源地址

文章和代码开源地址在GitHub: https://github.com/jincheng9/go-tutorial

公众号：coding进阶

个人网站：https://jincheng9.github.io/



## References

* [官方教程：Go泛型入门](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p6)
* [GopherCon 2021 Talk on Generics](https://www.youtube.com/watch?v=35eIxI_n5ZM&t=1755s)
* [Go Generics Proposal](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md)
* https://teivah.medium.com/when-to-use-generics-in-go-36d49c1aeda
* https://bitfieldconsulting.com/golang/generics