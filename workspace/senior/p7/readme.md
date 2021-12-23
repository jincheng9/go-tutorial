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

   类型推导，可以帮助我们在写代码的时候不用传递类型实参，由编译器自行推导。

   **注意**：类型推导并不是永远都可行。

## Type parameters(类型参数)

```go
[P, Q constraint1, R constraint2]
```

这里

这里定义了一个类型参数列表，列表里可以包含一个或者多个类型参数。

`P，Q和R`都是类型参数，`contraint1`和`contraint2`都是类型限制(type constraint)。

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

这个例子，只能计算2个`float64`中的较小者。有泛型之前，如果我们要支持计算2个int或者其它数值类型的较小者，就需要实现新的函数、或者使用`interface{}`，或者使用`Refelect`。

对于这个场景，使用泛型代码更简洁，效率也更优。支持比较不同数值类型的泛型min函数实现如下：

```go
func min(T constraints.Ordered) (x, y T) T {
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
2. `min[int](2, 3)`是在对泛型函数`min`实例化(instantiation)，在编译期将泛型函数里的类型参数`T`替换为`int`。

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
func min(T constraints.Ordered) (x, y T) T {
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
func min(T constraints.Ordered) (x, y T) T {
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

**注意**：类型限制必须是`interface`类型。比如上例的`constraints.Ordered`就是一个`interface`类型。

### | 和 ~

`\|`: 表示取并集。比如下例的`Number`这个interface可以作为类型限制，用于限定类型参数必须是int，int32和int64这3种类型。

```go
type Number interface{
	int | int32 | int64
}
```

`~T`: `~` 是Go 1.18新增的符号，`~T`表示底层类型是T的所有类型。比如下例的`AnyString`这个interface可以作为类型限制，用于限定类型参数的底层类型必须是string。`string`本身以及下例的`MyString`都满足`AnyString`这个类型限制。

```go
type AnyString interface{
   ~string
}
type MyString string
```

类型限制有2个作用：

1. 用于约定有效的类型实参，不满足类型限制的类型实参会被编译器报错。
2. 如果某个类型限制里的所有类型都支持某个操作，那在代码里，类型参数就可以使用这个操作。

### constraint literals(类型限制字面值)

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

`constraints`包定义了一些常用的类型限制，整个包除了测试代码，就1个`constraints.go`文件，50行代码，源码地址：

https://github.com/golang/go/blob/master/src/constraints/constraints.go



## Type inference(类型推导)

我们看下面的代码示例：

```go
func min(T constraints.Ordered) (x, y T) T {
	if x < y {
		return x
	}
	return y
}

var a, b, m1, m2 float64
// 方式1：显示指定type argument
m1 = min[float64](a, b)
// 方式2：不指定type argument，让编译器自行推导
m2 = min(a, b)
```

方式2没有传递类型实参，编译器是根据函数实参`a`和`b`推导出`T`的类型。

并不是总是能推导出来。

> Constraint Type Inference: Deduce type arguments from type parameter constraints



## 使用场景

### 箴言

> Write code, don't design types.

在写Go代码的时候，对于泛型，Go泛型设计者`Ian Lance Taylor`建议不要一上来就定义type parameter和type constraint，如果你一上来就这么做，那大概率是搞错了泛型的最佳实践。

先写具体的代码逻辑，等意识到需要使用type parameter或者新的type constraint的时候，再加上type parameter和type constraint。

### 什么时候使用泛型？

* 函数使用到了slice, map, channel类型，但是slice, map, channel里的元素类型可能有多种。

* 通用的数据结构。比如支持下面的代码实现了一个支持任意数据类型的二叉树。

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

1. When just calling a method on the type argument

   ```go
   // good
   func ReadFour(r io.Reader) ([]byte, error)
   
   // bad
   func ReadFour[T io.Reader](r T) ([]byte, error)
   ```

   

2. 当函数或者方法对于不同类型的实现逻辑不一样时，不要用泛型。

3. 当不同类型的操作不一样时，不要用泛型。比如`encoding/json`就不适合用泛型实现，`encoding/json`使用了`Reflect`。



###  总结

> Avoid boilerplate.
>
> Corollary: Don't use type parameters prematurely; wait until you are about to write boilerplate code.

不要随便使用泛型，当你发现针对不同类型，会写出同样的代码逻辑时，才去使用泛型。



## 开源地址

GitHub: https://github.com/jincheng9/go-tutorial

公众号：coding进阶

## References

* [官方教程：Go泛型入门](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p6)
* [GopherCon 2021 Talk on Generics](https://www.youtube.com/watch?v=35eIxI_n5ZM&t=1755s)
* [Go Generics Proposal](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md)
* https://teivah.medium.com/when-to-use-generics-in-go-36d49c1aeda
* https://bitfieldconsulting.com/golang/generics