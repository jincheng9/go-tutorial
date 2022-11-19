# Go常见错误第15篇：interface使用的常见错误和最佳实践

## 前言

这是Go常见错误系列的第15篇：interface使用的常见错误和最佳实践。

素材来源于Go布道者，现Docker公司资深工程师[Teiva Harsanyi](https://teivah.github.io/)。

本文涉及的源代码全部开源在：[Go常见错误源代码](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p28)，欢迎大家关注公众号，及时获取本系列最新更新。

 

## 常见错误和最佳实践

Interface是Go语言里的核心功能，但是在日常开发中，经常会出现interface被乱用的情况，代码过度抽象，或者抽象不合理，导致代码晦涩难懂。

本文先带大家回顾下interface的重要概念，然后讲解使用interface的常见错误和最佳实践。

### interface重要概念回顾

interface里面包含了若干个方法，大家可以理解为一个interface代表了一类群体的共同行为。

结构体要实现interface不需要类似implement的关键字，只要该结构体实现了interface里的所有方法即可。

我们拿Go语言里的io标准库来说明interface的强大之处。io标准库包含了2个interface：

* io.Reader：表示从某个数据源读数据
* io.Writer：表示写数据到目标位置，比如写到指定文件或者数据库

##### Figure 2.3 io.Reader reads from a data source and fills a byte slice, whereas io.Writer writes to a target from a byte slice.

![img](https://drek4537l1klr.cloudfront.net/harsanyi/Figures/CH02_F03_Harsanyi.png)

io.Reader这个interface里只有一个Read方法：

```
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

> Read reads up to len(p) bytes into p. It returns the number of bytes read (0 <= n <= len(p)) and any error encountered. Even if Read returns n < len(p), it may use all of p as scratch space during the call. If some data is available but not len(p) bytes, Read conventionally returns what is available instead of waiting for more.

如果某个结构体要实现io.Reader，需要实现Read方法。这个方法要包含以下逻辑：

* 入参：接受元素类型为byte的slice作为方法的入参。
* 方法逻辑：把Reader对象里的数据读出来赋值给p。比如Reader对象可能是一个strings.Reader，那调用Read方法就是把string的值赋值给p。
* 返回值：要么返回读到的字节数，要么返回error。

io.Writer这个interface里只有一个Write方法：

```
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

> Write writes len(p) bytes from p to the underlying data stream. It returns the number of bytes written from p (0 <= n <= len(p)) and any error encountered that caused the write to stop early. Write must return a non-nil error if it returns n < len(p). Write must not modify the slice data, even temporarily.

如果某个结构体要实现io.Writer，需要实现Write方法。这个方法要包含以下逻辑：

* 入参：接受元素类型为byte的slice作为方法的入参。
* 方法逻辑：把p的值写入到Writer对象。比如Writer对象可能是一个os.File类型，那调用Write方法就是把p的值写入到文件里。
* 返回值：要么返回写入的字节数，要么返回error。

这2个函数看起来非常抽象，很多Go初级开发者都不太理解，为啥要设计这样2个interface？

假设我们要实现一个函数，功能是拷贝一个文件的内容到另一个文件。

* 方式1：这个函数用2个*os.Files作为参数，来从一个文件读内容，写入到另一个文件

  ```go
  func copySourceToDest(source *io.File, dest *io.File) error {
      // ...
  }
  ```

* 方式2：使用io.Reader和io.Writer作为参数。由于os.Fil实现了io.Reader和io.Writer，所以os.File也可以作为下面函数的参数，传参给source和dest。

  ```go
  func copySourceToDest(source io.Reader, dest io.Writer) error {
      // ...
  }
  ```

  方法2的实现会更通用一些，source既可以是文件，也可以是字符串对象(strings.Reader)，dest既可以是文件，也可以是其它数据库对象(比如我们自己实现一个io.Writer，Write方法是把数据写入到数据库)。

在设计interface的时候要考虑到简洁性，如果interface里定义的方法很多，那这个interface的抽象就会不太好。

引用Go语言设计者Rob Pike在Gopherfest 2015上的技术分享[Go Proverbs with Rob Pike](https://www.youtube.com/watch?v=PAAkCSZUG1c&t=318s)中关于interface的说明：

> The bigger the interface, the weaker the abstraction.

当然，我们也可以把多个interface结合为一个interface，在有些场景下是可以方便代码编写的。

比如io.ReaderWriter就结合了io.Reader和io.Writer的方法。

```
type ReadWriter interface {
    Reader
    Writer
}
```

### 何时使用interface

When should we create interfaces in Go? Let’s look at three concrete use cases where interfaces are usually considered to bring value. Note that the goal isn’t to be exhaustive because the more cases we add, the more they would depend on the context. However, these three cases should give us a general idea:

- Common behavior
- Decoupling
- Restricting behavior

#### Common behavior

The first option we will discuss is to use interfaces when multiple types implement a common behavior. In such a case, we can factor out the behavior inside an interface. If we look at the standard library, we can find many examples of such a use case. For example, sorting a collection can be factored out via three methods:

- Retrieving the number of elements in the collection
- Reporting whether one element must be sorted before another
- Swapping two elements

Hence, the following interface was added to the sort package:

```
type Interface interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}
```

This interface has a strong potential for reusability because it encompasses the common behavior to sort any collection that is index-based.

Throughout the sort package, we can find dozens of implementations. If at some point we compute a collection of integers, for example, and we want to sort it, are we necessarily interested in the implementation type? Is it important whether the sorting algorithm is a merge sort or a quicksort? In many cases, we don’t care. Hence, the sorting behavior can be abstracted, and we can depend on the sort.Interface.

Finding the right abstraction to factor out a behavior can also bring many benefits. For example, the sort package provides utility functions that also rely on sort.Interface, such as checking whether a collection is already sorted. For instance,

```
func IsSorted(data Interface) bool {
    n := data.Len()
    for i := n - 1; i > 0; i-- {
        if data.Less(i, i-1) {
            return false
        }
    }
    return true
}
```

Because sort.Interface is the right level of abstraction, it makes it highly valuable.

Let’s now see another main use case when using interfaces.

#### Decoupling

Another important use case is about decoupling our code from an implementation. If we rely on an abstraction instead of a concrete implementation, the implementation itself can be replaced with another without even having to change our code. This is the Liskov Substitution Principle (the *L* in Robert C. Martin’s SOLID design principles).

One benefit of decoupling can be related to unit testing. Let’s assume we want to implement a CreateNewCustomer method that creates a new customer and stores it. We decide to rely on the concrete implementation directly (let’s say a mysql.Store struct):

```
type CustomerService struct {
    store mysql.Store
}
 
func (cs CustomerService) CreateNewCustomer(id string) error {
    customer := Customer{id: id}
    return cs.store.StoreCustomer(customer)
}
```



Now, what if we want to test this method? Because customerService relies on the actual implementation to store a Customer, we are obliged to test it through integration tests, which requires spinning up a MySQL instance (unless we use an alternative technique such as go-sqlmock, but this isn’t the scope of this section). Although integration tests are helpful, that’s not always what we want to do. To give us more flexibility, we should decouple CustomerService from the actual implementation, which can be done via an interface like so:

```
type customerStorer interface {
    StoreCustomer(Customer) error
}
 
type CustomerService struct {
    storer customerStorer
}
 
func (cs CustomerService) CreateNewCustomer(id string) error {
    customer := Customer{id: id}
    return cs.storer.StoreCustomer(customer)
}
```

Because storing a customer is now done via an interface, this gives us more flexibility in how we want to test the method. For instance, we can

- Use the concrete implementation via integration tests
- Use a mock (or any kind of test double) via unit tests
- Or both

Let’s now discuss another use case: to restrict a behavior.

#### Restricting behavior

The last use case we will discuss can be pretty counterintuitive at first sight. It’s about restricting a type to a specific behavior. Let’s imagine we implement a custom configuration package to deal with dynamic configuration. We create a specific container for int configurations via an IntConfig struct that also exposes two methods: Get and Set. Here’s how that code would look:

```
type IntConfig struct {
    // ...
}
 
func (c *IntConfig) Get() int {
    // Retrieve configuration
}
 
func (c *IntConfig) Set(value int) {
    // Update configuration
}
```

Now, suppose we receive an IntConfig that holds some specific configuration, such as a threshold. Yet, in our code, we are only interested in retrieving the configuration value, and we want to prevent updating it. How can we enforce that, semantically, this configuration is read-only, if we don’t want to change our configuration package? By creating an abstraction that restricts the behavior to retrieving only a config value:

```
type intConfigGetter interface {
    Get() int
}
```

Then, in our code, we can rely on intConfigGetter instead of the concrete implementation:

```
type Foo struct {
    threshold intConfigGetter
}
 
func NewFoo(threshold intConfigGetter) Foo {
    return Foo{threshold: threshold}
}
 
func (f Foo) Bar()  {
    threshold := f.threshold.Get()
    // ...
}
```

In this example, the configuration getter is injected into the NewFoo factory method. It doesn’t impact a client of this function because it can still pass an IntConfig struct as it implements intConfigGetter. Then, we can only read the configuration in the Bar method, not modify it. Therefore, we can also use interfaces to restrict a type to a specific behavior for various reasons, such as semantics enforcement.

In this section, we saw three potential use cases where interfaces are generally considered as bringing value: factoring out a common behavior, creating some decoupling, and restricting a type to a certain behavior. Again, this list isn’t exhaustive, but it should give us a general understanding of when interfaces are helpful in Go.

Now, let’s finish this section and discuss the problems with interface pollution.

### 乱用Interface的场景

It’s fairly common to see interfaces being overused in Go projects. Perhaps the developer’s background was C# or Java, and they found it natural to create interfaces before concrete types. However, this isn’t how things should work in Go.

As we discussed, interfaces are made to create abstractions. And the main caveat when programming meets abstractions is remembering that abstractions *should be discovered, not created*. What does this mean? It means we shouldn’t start creating abstractions in our code if there is no immediate reason to do so. We shouldn’t design with interfaces but wait for a concrete need. Said differently, we should create an interface when we need it, not when we foresee that we could need it.

What’s the main problem if we overuse interfaces? The answer is that they make the code flow more complex. Adding a useless level of indirection doesn’t bring any value; it creates a worthless abstraction making the code more difficult to read, understand, and reason about. If we don’t have a strong reason for adding an interface and it’s unclear how an interface makes a code better, we should challenge this interface’s purpose. Why not call the implementation directly?

##### NOTE

We may also experience performance overhead when calling a method through an interface. It requires a lookup in a hash table’s data structure to find the concrete type an interface points to. But this isn’t an issue in many contexts as the overhead is minimal.

In summary, we should be cautious when creating abstractions in our code—abstractions should be discovered, not created. It’s common for us, software developers, to overengineer our code by trying to guess what the perfect level of abstraction is, based on what we think we might need later. This process should be avoided because, in most cases, it pollutes our code with unnecessary abstractions, making it more complex to read.

Don’t design with interfaces, discover them.

—Rob Pike

Let’s not try to solve a problem abstractly but solve what has to be solved now. Last, but not least, if it’s unclear how an interface makes the code better, we should probably consider removing it to make our code simpler.

The following section continues with this thread and discusses a common interface mistake: creating interfaces on the producer side.



## 推荐阅读

* [Go面试题系列，看看你会几题？](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=Mzg2MTcwNjc1Mg==&action=getalbum&album_id=2199553588283179010#wechat_redirect)

* [Go常见错误第1篇：未知枚举值](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484146&idx=1&sn=10fb12b643a2e37c090e5aa3bc583152&chksm=ce124d9df965c48bb954aeddabdff3db12738ded3875542250c5d0ef6cfd4417fc56580288b1&token=1912894792&lang=zh_CN#rd)

* [Go常见错误第2篇：benchmark性能测试的坑](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484163&idx=1&sn=b28d61c1f3ec9d914e698dce105ba5d1&chksm=ce124c6cf965c57a90bc85a5295ed9375103de20607b509f845583ff6686385df0ed96653d00&token=1912894792&lang=zh_CN#rd)

* [Go常见错误第3篇：go指针的性能问题和内存逃逸](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484247&idx=1&sn=faf716627afb00df646cecff023fb63c&chksm=ce124c38f965c52efd009a4c98691d56b5765dc7dce98aa49b226ad9274bd062d8d01e702e91&token=1899277735&lang=zh_CN#rd)

* [Go常见错误第4篇：break操作的注意事项](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484262&idx=1&sn=c1bea8af60444a4ef73c4d4d7a09d16d&chksm=ce124c09f965c51f3663ac9089a792d36c3685850e12695dd26d15a1a50f393b2d7c92b9983a&token=461369035&lang=zh_CN#rd)

* [Go常见错误第5篇：Go语言Error管理](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484274&idx=1&sn=711abea3c6fd5d15341ee1b34da8a160&chksm=ce124c1df965c50b3af84965f7ed30b574cd0b247ea6f77b944ec858bd43ee37f4c1554a5bce&token=1846351524&lang=zh_CN#rd)

* [Go常见错误第6篇：slice初始化常犯的错误](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484289&idx=1&sn=2b8171458cde4425b28fdf8f51df8d7c&chksm=ce124ceef965c5f8a14f5951457ce2ac0ecc4612cf2013957f1d818b6e74da7c803b9df1d394&token=1477304797&lang=zh_CN#rd)

* [Go常见错误第7篇：不使用-race选项做并发竞争检测](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484299&idx=1&sn=583c3470a76e93b0af0d5fc04fe29b55&chksm=ce124ce4f965c5f20de5887b113eab91f7c2654a941491a789e4ac53c298fbadb4367acee9bb&token=1918756920&lang=zh_CN#rd)

* [Go常见错误第8篇：并发编程中Context使用常见错误](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484317&idx=1&sn=474dad373684979fc96ba59182f08cf5&chksm=ce124cf2f965c5e4a29e313b4654faacef03e78da7aaf2ba6912d7b490a1df851a1bcbfec1c9&token=1918756920&lang=zh_CN#rd)

* [Go常见错误第9篇：使用文件名称作为函数输入](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484325&idx=1&sn=689c1b3823697cc583e1e818c4c76ee5&chksm=ce124ccaf965c5dce4e497f6251c5f0a8473b8e2ae3824bd72fe8c532d6dd84e6375c3990b3e&token=1266762504&lang=zh_CN#rd)

* [Go常见错误第10篇：Goroutine和循环变量一起使用的坑](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484335&idx=1&sn=cc8c6ceae72b30ec6f4d4e7b4367baca&chksm=ce124cc0f965c5d60410f977cdf31f127694fd0d49c35e2061ce8fb5fb9387bfa321196db438&token=1656737387&lang=zh_CN#rd)

* [Go常见错误第11篇：意外的变量遮蔽(variable shadowing)](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484519&idx=1&sn=00f13bdb95bd8f7c4eb1582a4c981991&chksm=ce124b08f965c21e28ea3a1c67b3b501fe45ba1a2c4af5048cad3a2d5d8303b6fc0192235b6d&token=1762934632&lang=zh_CN#rd)

* [Go常见错误第12篇：如何破解箭头型代码](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484539&idx=1&sn=189ff2e8fdb4a7d18620f9367128d6c4&chksm=ce124b14f965c20269c766a0f18f98cc8b9e340317669034a074d39843f2f9a94f3c9caf18da&token=329552886&lang=zh_CN#rd)

* [Go常见错误第13篇：init函数的常见错误和最佳实践](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484553&idx=1&sn=a4de11c452157193ae4381ab3555c42c&chksm=ce124be6f965c2f0885b8acf21c867e82d09807eba7be7890c54c01acf2b6fc413267e90d371&token=2029492652&lang=zh_CN#rd)

* [Go常见错误第14篇：过度使用getter和setter方法 ](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484568&idx=1&sn=2f078aa6561093691b4aeae58de44830&chksm=ce124bf7f965c2e17b99896393dfb684d5868156ea4e0b324bc86e2dcc596ea55d5e6c562c29&token=1629431500&lang=zh_CN#rd)

  

## 开源地址

文章和示例代码开源在GitHub: [Go语言初级、中级和高级教程](https://github.com/jincheng9/go-tutorial)。

公众号：coding进阶。关注公众号可以获取最新Go面试题和技术栈。

个人网站：[Jincheng's Blog](https://jincheng9.github.io/)。

知乎：[无忌](https://www.zhihu.com/people/thucuhkwuji)。



## 福利

我为大家整理了一份后端开发学习资料礼包，包含编程语言入门到进阶知识(Go、C++、Python)、后端开发技术栈、面试题等。

关注公众号「coding进阶」，发送消息 **backend** 领取资料礼包，这份资料会不定期更新，加入我觉得有价值的资料。

发送消息「**进群**」，和同行一起交流学习，答疑解惑。



## References

* https://livebook.manning.com/book/100-go-mistakes-how-to-avoid-them/chapter-2/
* https://github.com/jincheng9/go-tutorial/tree/main/workspace/lesson18
* https://bbs.huaweicloud.com/blogs/348512