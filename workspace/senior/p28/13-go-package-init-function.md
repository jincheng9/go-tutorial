# Go常见错误第13篇：init函数的常见错误和最佳实践

## 前言

这是Go常见错误系列的第13篇：init函数的常见错误和最佳实践。

素材来源于Go布道者，现Docker公司资深工程师[Teiva Harsanyi](https://teivah.github.io/)。

本文涉及的源代码全部开源在：[Go常见错误源代码](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p28)，欢迎大家关注公众号，及时获取本系列最新更新。

 

## 常见错误和最佳实践

很多Go语言开发者会错误地使用package里的init函数，导致代码难懂，维护困难。

我们先回顾下package里init函数的概念，然后讲解init函数的常见错误和最佳实践。

### init基本概念

Go语言里的init函数有如下特点：

* init函数没有参数，没有返回值。如果加了参数或返回值，会编译报错。
* 一个package下面的每个.go源文件都可以有自己的init函数。当这个package被import时，就会执行该package下的init函数。
* 一个.go源文件里可以有一个或者多个init函数，虽然函数签名完全一样，但是Go允许这么做。
* .go源文件里的全局常量和变量会先被编译器解析，然后再执行init函数。

### 示例1

我们来看如下的代码示例：

```go
package main

import "fmt"

func init() {
	fmt.Println("init")
}

func init() {
	fmt.Println(a)
}

func main() {
	fmt.Println("main")
}

var a = func() int {
	fmt.Println("var")
	return 0
}()
```

`go run main.go`执行这段程序的结果是：

```bash
var
init
0
main
```

全局变量`a`的定义虽然放在了最后面，但是先被编译器解析，然后执行init函数，最后执行main函数。

### 示例2

有2个package: `main`和`redis`，`main`这个package依赖了`redis`这个package。

```go
package main
 
import (
    "fmt"
 
    "redis"
)
 
func init() {
    // ...
}
 
func main() {
    err := redis.Store("foo", "bar")
    // ...
}
```



```go
package redis
 
// imports
 
func init() {
    // ...
}
 
func Store(key, value string) error {
    // ...
}
```

因为`main` import了`redis`，所以`redis`这个package里的init函数先执行，然后再执行`main`这个package里的init函数。

* 如果一个package下面有多个.go源文件，每个.go源文件里都有各自的init函数，那会按照.go源文件名的字典序执行init函数。比如有a.go和b.go这2个源文件，里面都有init函数，那a.go里的init函数比b.go里的init函数先执行。
* 如果一个.go源文件里有多个init函数，那按照代码里的先后顺序执行。

![img](https://drek4537l1klr.cloudfront.net/harsanyi/Figures/CH02_F02_Harsanyi.png)

* 我们在工程实践里，不要去依赖init函数的执行顺序。如果预设了init函数的执行顺序，通常是很危险的，也不是Go语言的最佳实践。因为源文件名是有可能被修改的。

* init函数不能被直接调用，否则会编译报错。

  ```go
  package main
   
  func init() {}
   
  func main() {
      init()
  }
  ```

  上面这段代码编译报错如下：

  ```go
  $ go build .
  ./main.go:6:2: undefined: init
  ```

到现在为止，大家对package里的init函数应该有了一个比较清晰的理解，接下来我们看看init函数的常见错误和最佳实践。

### init函数的错误用法

我们先看看init函数一种常见的不太好的用法。

```go
var db *sql.DB
 
func init() {
    dataSourceName :=
        os.Getenv("MYSQL_DATA_SOURCE_NAME")
    d, err := sql.Open("mysql", dataSourceName)
    if err != nil {
        log.Panic(err)
    }
    err = d.Ping()
    if err != nil {
        log.Panic(err)
    }
    db = d
}
```

上面的程序做了如下几个事情：

* 创建一个数据库连接实例。
* 对数据库做ping检查。
* 如果连接数据库和ping检查都通过的话，会把数据库连接实例赋值给全局变量`db`。

大家可以先思考下这段程序会有哪些问题。

* 第一，init函数里面做错误管理的方式是很有限的。比如，init函数没法返回error，因为init函数是不能有返回值的。那如果init函数出现了error要让外界感知的话，得主动触发panic，让程序停止。对于上面的示例程序，虽然init函数遇到错误时，表示数据库连接失败，去停止程序运行或许是可以的。但是在init函数里去创建数据库连接，如果失败的话，就不好做重试或者容错处理。试想，如果是在一个普通函数里去创建数据库连接，那这个普通函数可以在创建数据库连接失败的时候返回error信息，然后函数的调用者来决定做重试或者退出的操作。

* 第二，会影响代码的单元测试。因为init函数在测试代码执行之前就会运行，如果我们仅仅是想测试这个package里某个不需要做数据库连接的基础函数，那测试的时候还是会执行init函数，去创建数据库连接，这显然并不是我们想要的效果，增加了单元测试的复杂性。

* 第三，这段程序把数据库连接赋值给了全局变量。用全局变量会有一些潜在的风险，比如这个package里的其它函数可以修改这个全局变量的值，导致被误修改；一些和数据库连接无关的单元测试也得考虑这个全局变量。

那我们如何对上面的程序做修改来解决以上问题呢？参考如下代码：

```go
func createClient(dsn string) (*sql.DB, error) {
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }
    if err = db.Ping(); err != nil {
        return nil, err
    }
    return db, nil
}
```

通过这个函数来创建数据库连接就可以解决以上3个问题了。

- 错误处理可以交给createClient函数的调用者去管理，调用者可以选择退出程序或者重试。
- 单元测试既可以测试和数据库无关的基础函数，也可以测试createClient来检查数据库连接的代码实现。
- 没有暴露全局变量，数据库连接实例在createClient函数里面创建和返回。

### 何时使用init函数

init函数也并不是完全不建议用，在有些场景下是可以考虑使用的。比如Go的官方blog的[源码实现](https://cs.opensource.google/go/x/website/+/e0d934b4:blog/blog.go;l=32)就用到了init函数。

```go
func init() {
    redirect := func(w http.ResponseWriter, r *http.Request) {
        http.Redirect(w, r, "/", http.StatusFound)
    }
    http.HandleFunc("/blog", redirect)
    http.HandleFunc("/blog/", redirect)
 
    static := http.FileServer(http.Dir("static"))
    http.Handle("/favicon.ico", static)
    http.Handle("/fonts.css", static)
    http.Handle("/fonts/", static)
 
    http.Handle("/lib/godoc/", http.StripPrefix("/lib/godoc/",
        http.HandlerFunc(staticHandler)))
}
```

这段源码里，init函数不可能失败，因为http.HandleFunc只有在第2个handler参数为nil的时候才会panic，显然这段程序里http.HandleFunc的第2个handler参数都是合法值，所以init函数不会失败。

同时，这里也无需创建全局变量，而且这个函数也不会影响单元测试。

因此这是一个适合用init函数的场景示例。

## 总结

init函数要慎用，如果使用不当可能会带来问题，千万不要在代码里依赖同一package下不同.go文件init的执行顺序。

最后回顾下Go语言init函数的注意事项：

* init函数没有参数，没有返回值。如果加了参数或返回值，会编译报错。
* 一个package下面的每个.go源文件都可以有自己的init函数。当这个package被import时，就会执行该package下的init函数。
* 一个.go源文件里可以有一个或者多个init函数，虽然函数签名完全一样，但是Go允许这么做。
* .go源文件里的全局常量和变量会先被编译器解析，然后再执行init函数。



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
* https://github.com/jincheng9/go-tutorial/tree/main/workspace/lesson27