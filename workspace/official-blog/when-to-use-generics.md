# 泛型最佳实践：Go泛型设计者教你如何用泛型

## 前言

Go泛型的设计者*Ian Lance Taylor*在官方博客网站上发表了一篇文章[when to use generics](https://go.dev/blog/when-generics)，详细说明了在什么场景下应该使用泛型，什么场景下不要使用泛型。这对于我们写出符合最佳实践的Go泛型代码非常有指导意义。

本人对原文在翻译的基础上做了一些表述上的优化，方便大家理解。



## 原文翻译

*Ian Lance Taylor*

*2022.04.14*

这篇博客汇总了我在2021年Google开源活动日和GopherCon会议上关于泛型的分享。

Go 1.18版本新增了一个重大功能：支持泛型编程。本文不会介绍什么是泛型以及如何使用泛型，而是把重点放在讲解Go编程实践中，什么时候应该使用泛型，什么时候不要使用泛型。

需要明确的是，我将会提供一些通用的指引，这并不是硬性规定，大家可以根据自己的判断来决定，但是如果你不确定如何使用泛型，那建议参考本文介绍的指引。

## 写代码

Go编程有一条通用准则：write Go programs by writing code, not by defining types. 

具体到泛型，如果你写代码的时候从定义类型参数约束(type parameter constraints)开始，那你可能搞错了方向。从编写函数开始，如果写的过程中发现使用类型参数更好，那再使用类型参数。

## 类型参数何时有用？

接下来我们看看在什么情况下，使用类型参数对我们写代码更有用。

### 使用Go内置的容器类型

如果函数使用了语言内置的容器类型(包括slice, map和channel)作为函数参数，并且函数代码对容器的处理逻辑并没有预设容器里的元素类型，那使用类型参数(type parameter)可能就会有用。

举个例子，我们要实现一个函数，该函数的入参是一个map，要返回该map的所有key组成的slice，key的类型可以是map支持的任意key类型。

```go
// MapKeys returns a slice of all the keys in m.
// The keys are not returned in any particular order.
func MapKeys[Key comparable, Val any](m map[Key]Val) []Key {
    s := make([]Key, 0, len(m))
    for k := range m {
        s = append(s, k)
    }
    return s
}
```

这段代码没有对map里key的类型做任何限定，并且没有用map里的value，因此这段代码适用于所有的map类型。这就是使用类型参数的一个很好的示例。

这种场景下，也可以使用反射(reflection)，但是反射是一种比较别扭的编程模型，在编译期没法做静态类型检查，并且会导致运行期的速度变慢。

### 实现通用的数据结构

对于通用的数据结构，类型参数也会有用。通用的数据结构类似于slice和map，但是并不是语言内置的数据结构，比如链表或者二叉树。

在没有泛型的时候，如果要实现通用的数据结构，有2种方案：

* 方案1：针对每个元素类型分别实现一个数据结构
* 方案2：使用interface类型

泛型相对方案1的优点是代码更精简，也更方便给其它模块调用。泛型相对方案2的优点是数据存储更高效，节约内存资源，并且可以在编译期做静态类型检查，避免代码里使用类型断言。

下面的例子就是使用类型参数实现的通用二叉树数据结构：

```go
// Tree is a binary tree.
type Tree[T any] struct {
    cmp  func(T, T) int
    root *node[T]
}

// A node in a Tree.
type node[T any] struct {
    left, right  *node[T]
    val          T
}

// find returns a pointer to the node containing val,
// or, if val is not present, a pointer to where it
// would be placed if added.
func (bt *Tree[T]) find(val T) **node[T] {
    pl := &bt.root
    for *pl != nil {
        switch cmp := bt.cmp(val, (*pl).val); {
        case cmp < 0:
            pl = &(*pl).left
        case cmp > 0:
            pl = &(*pl).right
        default:
            return pl
        }
    }
    return pl
}

// Insert inserts val into bt if not already there,
// and reports whether it was inserted.
func (bt *Tree[T]) Insert(val T) bool {
    pl := bt.find(val)
    if *pl != nil {
        return false
    }
    *pl = &node[T]{val: val}
    return true
}
```

二叉树的每个节点包含一个类型为`T`的变量`val`。当二叉树实例化的时候，需要传入类型实参，这个时候`val`的类型已经确定下来了，不会被存为interface类型。

这种场景使用类型参数是合理的，因为`Tree`是个通用的数据结构，包括方法里的代码实现都和`T`的类型无关。

`Tree`数据结构本身不需要知道如何比较二叉树节点上类型为`T`的变量`val`的大小，它有一个成员变量`cmp`来实现`val`大小的比较，`cmp`是一个函数类型变量，在二叉树初始化的时候被指定。因此二叉树上节点值的大小比较是`Tree`外部的一个函数来实现的，你可以在`find`方法的第4行看到对`cmp`的使用。

### 类型参数优先使用在函数而不是方法上

上面的 `Tree`数据结构示例阐述了另外一个通用准则：当你需要类似`cmp`的比较函数时，优先考虑使用函数而不是方法。

对于上面`Tree`类型，除了使用函数类型的成员变量`cmp`来比较`val`的大小之外，还有另外一种方案是要求类型`T`必须有一个`Compare`或者`Less`方法来做大小比较。要做到这一点，就需要定义一个类型约束(type constraint)用于限定类型`T`必须实现这个方法。

这造成的结果是即使`T`只是一个普通的int类型，那使用者也必须定义一个自己的int类型，实现类型约束里的方法(method)，然后把这个自定义的int类型作为类型实参传参给类型参数`T`。

但是如果我们参照上面`Tree`的代码实现，定义一个函数类型的成员变量`cmp`用来做`T`类型的大小比较，代码实现就比较简洁。

换句话说，把方法转为函数比给一个类型增加方法容易得多。因此对于通用的数据类型，优先考虑使用函数，而不是写一个必须有方法的类型限制。

### 不同类型需要实现公用方法

类型参数另一个有用的场景是不同的类型要实现一些公用方法，并且对于这些方法，不同类型的实现逻辑是一样的。

下面举个例子，Go标准库里有一个[sort](https://pkg.go.dev/sort)包，可以对存储不同数据类型的slice做排序，比如`Float64s(x)`可以对`[]float64`做排序，`Ints(x)`可以对`[]int`做排序。

同时sort包还可以对用户自定义的数据类型(比如结构体、自定义的int类型等)调用`sort.Sort()`做排序，只要该类型实现了`sort.Interface`这个接口类型里`Len()`、`Less()`和`Swap()`这3个方法即可。

下面我们对sort包可以使用泛型来做一些改造，就可以对存储不同数据类型的slice统一调用`sort.Sort()`来做排序，而不用专门为`[]int`调用`Ints(x)`，为`[]float64`调用`Float64s(x)`做差异化处理了，可以简化代码逻辑。

下面的代码实现了一个泛型的结构体类型`SliceFn`，这个结构体类型实现了`sort.Interface`。

```go
// SliceFn implements sort.Interface for a slice of T.
type SliceFn[T any] struct {
    s    []T
    less func(T, T) bool
}

func (s SliceFn[T]) Len() int {
    return len(s.s)
}
func (s SliceFn[T]) Swap(i, j int) {
    s.s[i], s.s[j] = s.s[j], s.s[i]
}
func (s SliceFn[T] Less(i, j int) bool {
    return s.less(s.s[i], s.s[j])
}
```

对于不同的slice类型， `Len` 和 `Swap` 方法的实现是一样的。`Less` 方法需要对slice里的2个元素做比较，比较逻辑实现在`SliceFn`里的成员变量`less`里头，`less`是一个函数类型的变量，在结构体初始化的时候进行传参赋值。这点和上面`Tree`这个二叉树通用数据结构的处理类似。

我们再将`sort.Sort`按照泛型风格封装为`SortFn`泛型函数，这样对于所有slice类型，我们都可以统一调用`SortFn`做排序。

```go
// SortFn sorts s in place using a comparison function.
func SortFn[T any](s []T, less func(T, T) bool) {
    sort.Sort(SliceFn[T]{s, cmp})
}
```

这和标准库里的[sort.Slice](https://pkg.go.dev/sort#Slice)很类似，只不过这里的`less`比较函数的参数是具体的值，而`sort.Slice`里比较函数`less`比较函数的参数是slice的下标索引。

这种场景使用类型参数比较合适，因为不同类型的`SliceFn`的方法实现逻辑都是一样的，只是`slice`里存储的元素的类型不一样而已。



## 类型参数何时不要用

现在我们谈谈类型参数不建议使用的场景。

### 不要把interface类型替换为类型参数

我们大家都知道Go语言有interface类型，interface支持某种意义上的泛型编程。

举个例子，被广泛使用的`io.Reader`接口提供了一种泛型机制用于读取数据，比如支持从文件和随机数生成器里读取数据。

如果你对某些类型的变量的操作只是调用该类型的方法，那就直接使用interface类型，不要使用类型参数。`io.Reader`从代码角度易于阅读且高效，没必要使用类型参数。

举个例子，有人可能会把下面第1个基于interface类型的`ReadSome`版本修改为第2个基于类型参数的版本。

```
func ReadSome(r io.Reader) ([]byte, error)

func ReadSome[T io.Reader](r T) ([]byte, error)
```

不要做这种修改，使用第1个基于interface的版本会让函数更容易编写和阅读，并且函数执行效率也几乎一样。

**注意**：尽管可以使用不同的方式来实现泛型，并且泛型的实现可能会随着时间的推移而发生变化，但是Go 1.18中泛型的实现在很多情况下对于类型为interface的变量和类型为类型参数的变量处理非常相似。这意味着使用类型参数通常并不会比使用interface快，所以不要单纯为了程序运行速度而把interface类型修改为类型参数，因为它可能并不会运行更快。

### 如果方法的实现不同，不要使用类型参数

当决定要用类型参数还是interface时，要考虑方法的逻辑实现。正如我们前面说的，如果方法的实现对于所有类型都一样，那就是用类型参数。相反，如果每个类型的方法实现是不同的，那就是用interface类型，不要用类型参数。

举个例子，从文件里`Read`的实现和从随机数生成器里`Read`的实现完全不一样，在这种场景下，可以定义一个`io.Reader`的interface类型，该类型包含有一个`Read`方法。文件和随机数生成器实现各自的`Read`方法。

### 在适当的时候可以使用反射(reflection)

Go有 [运行期反射](https://pkg.go.dev/reflect)。反射机制支持某种意义上的泛型编程，因为它允许你编写适用于任何类型的代码。如果某些操作需要支持以下场景，就可以考虑使用反射。

* 操作没有方法的类型，interface类型不适用。
* 每个类型的操作逻辑不一样，泛型不适用。

一个例子是[encoding/json](https://pkg.go.dev/encoding/json)包的实现。我们并不希望要求我们编码的每个类型都实现`MarshalJson`方法，因此我们不能使用interface类型。而且不同类型编码的逻辑不一样，因此我们不应该用泛型。

因此对于这种情况，[encoding/json](https://pkg.go.dev/encoding/json)使用了反射来实现。具体实现细节可以参考[源码](https://go.dev/src/encoding/json/encode.go)。



## 一个简单原则

总结一下，何时使用泛型可以简化为如下的一个简单原则。

如果你发现重复在写几乎完全一样的代码，唯一的区别是代码里使用的类型不一样，那就要考虑是否可以使用泛型来实现。



## 开源地址

文章和示例代码开源在GitHub: [Go语言初级、中级和高级教程](https://github.com/jincheng9/go-tutorial)。

公众号：coding进阶。关注公众号可以获取最新Go面试题和技术栈。

个人网站：[Jincheng's Blog](https://jincheng9.github.io/)。

知乎：[无忌](https://www.zhihu.com/people/thucuhkwuji)



## References

* Go Blog on When to Use Generics: https://go.dev/blog/when-generics
* Go Day 2021 on Google Open Source : https://www.youtube.com/watch?v=nr8EpUO9jhw
* GopherCon 2021: https://www.youtube.com/watch?v=Pa_e9EeCdy8