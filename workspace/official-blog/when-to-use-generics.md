# 官方指导：什么场景应该使用泛型

*Ian Lance Taylor*

*2022.04.14*

## 导言

这篇博客汇总了我在2021年Google开源活动日和GopherCon会议上关于泛型的分享。

Go 1.18版本新增了一个重大功能：支持泛型编程。本文不会介绍什么是泛型以及如何使用泛型，而是把重点放在讲解Go编程实践中，什么时候应该使用泛型，什么时候不要使用泛型。

需要明确的是，我将会提供一些通用的指引，这并不是硬性规定，大家可以根据自己的判断来决定，但是如果你不确定如何使用泛型，那建议参考本文介绍的指引。

## 写代码

Go编程有一条通用准则：write Go programs by writing code, not by defining types. 

具体到泛型，如果你写代码的时候从定义类型参数约束(type parameter constraints)开始，那你可能搞错了方向。从编写函数开始，如果写的过程中发现使用类型参数更好，那就再使用类型参数。

## 类型参数何时有用？

接下来我们看看在什么情况下，使用类型参数对我们写代码更有用。

### 使用内置的容器类型

如果函数使用了语言内置的容器类型(包括slice, map和channel)作为函数参数，并且函数代码对容器的处理逻辑并没有预设容器里的元素类型，那使用类型参数(type parameter)可能就会有用。

举个例子，我们要实现一个函数，该函数的入参是一个map，要返回该map的所有key组成的slice，key的类型可以是map支持的任意key类型。

```
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

### 通用的数据结构

对于通用的数据结构，类型参数也会有用。通用的数据结构类似于slice和map，但是并不是语言内置的数据结构，比如链表或者二叉树。

在没有泛型的时候，如果要实现通用的数据结构，有2种方案：

* 方案1：针对每个元素类型分别实现一个数据结构
* 方案2：使用interface类型

泛型相对方案1的优点是代码更精简，也更方便给其它模块调用。泛型相对方案2的优点是数据存储更高效，节约内存资源，并且可以在编译期做静态类型检查，避免代码里使用类型断言。

下面的例子就是使用类型参数实现的通用二叉树数据结构：

```
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



The `Tree` data structure does need to know how to compare values of the element type `T`; it uses a passed-in comparison function for that. You can see this on the fourth line of the `find` method, in the call to `bt.cmp`. Other than that, the type parameter doesn’t matter at all.

### For type parameters, prefer functions to methods

The `Tree` example illustrates another general guideline: when you need something like a comparison function, prefer a function to a method.

We could have defined the `Tree` type such that the element type is required to have a `Compare` or `Less` method. This would be done by writing a constraint that requires the method, meaning that any type argument used to instantiate the `Tree` type would need to have that method.

A consequence would be that anybody who wants to use `Tree` with a simple data type like `int` would have to define their own integer type and write their own comparison method. If we define `Tree` to take a comparison function, as in the code shown above, then it is easy to pass in the desired function. It’s just as easy to write that comparison function as it is to write a method.

If the `Tree` element type happens to already have a `Compare` method, then we can simply use a method expression like `ElementType.Compare` as the comparison function.

To put it another way, it is much simpler to turn a method into a function than it is to add a method to a type. So for general purpose data types, prefer a function rather than writing a constraint that requires a method.

### Implementing a common method

Another case where type parameters can be useful is when different types need to implement some common method, and the implementations for the different types all look the same.

For example, consider the standard library’s `sort.Interface`. It requires that a type implement three methods: `Len`, `Swap`, and `Less`.

Here is an example of a generic type `SliceFn` that implements `sort.Interface` for any slice type:

```
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

For any slice type, the `Len` and `Swap` methods are exactly the same. The `Less` method requires a comparison, which is the `Fn` part of the name `SliceFn`. As with the earlier `Tree` example, we will pass in a function when we create a `SliceFn`.

Here is how to use `SliceFn` to sort any slice using a comparison function:

```
// SortFn sorts s in place using a comparison function.
func SortFn[T any](s []T, less func(T, T) bool) {
    sort.Sort(SliceFn[T]{s, cmp})
}
```

This is similar to the standard library function `sort.Slice`, but the comparison function is written using values rather than slice indexes.

Using type parameters for this kind of code is appropriate because the methods look exactly the same for all slice types.

(I should mention that Go 1.19–not 1.18–will most likely include a generic function to sort a slice using a comparison function, and that generic function will most likely not use `sort.Interface`. See [proposal #47619](https://go.dev/issue/47619). But the general point is still true even if this specific example will most likely not be useful: it’s reasonable to use type parameters when you need to implement methods that look the same for all the relevant types.)

## 类型参数何时不要用

Now let’s talk about the other side of the question: when not to use type parameters.

### Don’t replace interface types with type parameters

As we all know, Go has interface types. Interface types permit a kind of generic programming.

For example, the widely used `io.Reader` interface provides a generic mechanism for reading data from any value that contains information (for example, a file) or that produces information (for example, a random number generator). If all you need to do with a value of some type is call a method on that value, use an interface type, not a type parameter. `io.Reader` is easy to read, efficient, and effective. There is no need to use a type parameter to read data from a value by calling the `Read` method.

For example, it might be tempting to change the first function signature here, which uses just an interface type, into the second version, which uses a type parameter.

```
func ReadSome(r io.Reader) ([]byte, error)

func ReadSome[T io.Reader](r T) ([]byte, error)
```

Don’t make that kind of change. Omitting the type parameter makes the function easier to write, easier to read, and the execution time will likely be the same.

It’s worth emphasizing the last point. While it’s possible to implement generics in several different ways, and implementations will change and improve over time, the implementation used in Go 1.18 will in many cases treat values whose type is a type parameter much like values whose type is an interface type. What this means is that using a type parameter will generally not be faster than using an interface type. So don’t change from interface types to type parameters just for speed, because it probably won’t run any faster.

### Don’t use type parameters if method implementations differ

When deciding whether to use a type parameter or an interface type, consider the implementation of the methods. Earlier we said that if the implementation of a method is the same for all types, use a type parameter. Inversely, if the implementation is different for each type, then use an interface type and write different method implementations, don’t use a type parameter.

For example, the implementation of `Read` from a file is nothing like the implementation of `Read` from a random number generator. That means that we should write two different `Read` methods, and use an interface type like `io.Reader`.

### Use reflection where appropriate

Go has [run time reflection](https://pkg.go.dev/reflect). Reflection permits a kind of generic programming, in that it permits you to write code that works with any type.

If some operation has to support even types that don’t have methods (so that interface types don’t help), and if the operation is different for each type (so that type parameters aren’t appropriate), use reflection.

An example of this is the [encoding/json](https://pkg.go.dev/encoding/json) package. We don’t want to require that every type that we encode have a `MarshalJSON` method, so we can’t use interface types. But encoding an interface type is nothing like encoding a struct type, so we shouldn’t use type parameters. Instead, the package uses reflection. The code is not simple, but it works. For details, see [the source code](https://go.dev/src/encoding/json/encode.go).

## 一个简单原则

总结一下，何时使用泛型可以简化为一个简单原则。

If you find yourself writing the exact same code multiple times, where the only difference between the copies is that the code uses different types, consider whether you can use a type parameter.

Another way to say this is that you should avoid type parameters until you notice that you are about the write the exact same code multiple times.



## 开源地址

文章和示例代码开源在GitHub: [Go语言初级、中级和高级教程](https://github.com/jincheng9/go-tutorial)。

公众号：coding进阶。关注公众号可以获取最新Go面试题和技术栈。

个人网站：[Jincheng's Blog](https://jincheng9.github.io/)。

知乎：[无忌](https://www.zhihu.com/people/thucuhkwuji)



## References

* Go Blog on When to Use Generics: https://go.dev/blog/when-generics
* Go Day 2021 on Google Open Source : https://www.youtube.com/watch?v=nr8EpUO9jhw
* GopherCon 2021: https://www.youtube.com/watch?v=Pa_e9EeCdy8