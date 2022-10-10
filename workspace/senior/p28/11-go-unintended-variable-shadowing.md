# Go常见错误第11篇：意外的变量遮蔽

## 前言

这是Go常见错误系列的第11篇：Go语言中意外的变量遮蔽。素材来源于Go布道者，现Docker公司资深工程师[Teiva Harsanyi](https://teivah.medium.com/)。

本文涉及的源代码全部开源在：[Go常见错误源代码](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p28)，欢迎大家关注公众号，及时获取本系列最新更新。



##  什么是变量遮蔽

变量遮蔽的英文原词是 variable shadowing，我们来看看维基百科上的定义：

> In [computer programming](https://en.wikipedia.org/wiki/Computer_programming), **variable shadowing** occurs when a variable declared within a certain [scope](https://en.wikipedia.org/wiki/Scope_(computer_science)) (decision block, method, or [inner class](https://en.wikipedia.org/wiki/Inner_class)) has the same name as a variable declared in an outer scope. At the level of [identifiers](https://en.wikipedia.org/wiki/Identifier_(computer_languages)) (names, rather than variables), this is known as [name masking](https://en.wikipedia.org/wiki/Name_masking). This outer variable is said to be shadowed by the inner variable, while the inner identifier is said to *mask* the outer identifier. This can lead to confusion, as it may be unclear which variable subsequent uses of the shadowed variable name refer to, which depends on the [name resolution](https://en.wikipedia.org/wiki/Name_resolution_(programming_languages)) rules of the language.

简单来说，如果某个作用域里声明了一个变量，同时在这个作用域的外层作用域又有一个相同名字的变量，就叫variable shadowing(变量遮蔽)。

```go
func test() {
	i := -100
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println(i)
}
```

比如上面这段代码，在for循环里面和外面都有一个变量`i`。

for循环里面`fmt.Println(i)`用到的变量`i`是for循环里面定义的变量`i`，for循环外面的`i`在for循环里面是不可见的，被遮蔽了。



## 常见错误

对于下面这段代码，大家思考下，看看会有什么问题：

```go
var client *http.Client
if tracing {
    client, err := createClientWithTracing()
    if err != nil {
        return err
    }
    log.Println(client)
} else {
    client, err := createDefaultClient()
    if err != nil {
        return err
    }
    log.Println(client)
}
// Use client
```

这段代码逻辑分3步：

* 首先定义了一个变量`client`
* 在后面的代码逻辑里，根据不同情况创建不同的client
* 最后使用赋值后的client做业务操作

但是，我们要注意到，在if/else里对`client`变量赋值时，使用了`:=`。

这个会直接创建一个新的局部变量`client`，而不是对我们最开始定义的`client`变量进行赋值，这就是variable shadowing现象。

这段代码带来的问题就是我们最开始定义的变量`client`的值会是`nil`，不符合我们的预期。

那我们应该怎么写代码，才能对我们最开始定义的`client`变量赋值呢？有以下2种解决方案。

### 解决方案1

```go
var client *http.Client
if tracing {
    c, err := createClientWithTracing()
    if err != nil {
        return err
    }
    client = c
} else {
    // Same logic
}
```

在if/else里定义一个临时变量`c`，然后把`c`赋值给变量`client`。

### 解决方案2

```
var client *http.Client
var err error
if tracing {
    client, err = createClientWithTracing()
    if err != nil {
        return err
    }
} else {
    // Same logic
}
```

直接先把`error`变量提前定义好，在if/else里直接用`=`做赋值，而不用`:=`。

### 方案区别

上面这2种方案其实都可以满足业务需求。我个人比较推荐方案2，主要理由如下：

* 代码会更精简，只需要直接对最终用到的变量做一次赋值即可。方案1里要做2次赋值，先赋值给临时变量`c`，再赋值给变量`client`。

* 可以对error统一处理。不需要在if/else里对返回的error做判断，方案2里我们可以直接在if/else外面对error做判断和处理，代码示例如下：

  ```go
  if tracing {
      client, err = createClientWithTracing()
  } else {
      client, err = createDefaultClient()
  }
  if err != nil {
      // Common error handling
  }
  ```



## 如何自动发现variable shadowing

靠人肉去排查还是容易遗漏的，Go工具链里有一个`shadow`命令可以帮助我们排查代码里潜在的variable shadowing问题。

* 第一步，安装`shadow`命令

  ```bash
  go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
  ```

* 第二步，使用`shadow`检查代码里是否有variable shadowing

  ```bash
  go vet -vettool=$(which shadow)
  ```

比如，我检查后的结果如下：

```bash
$ go vet -vettool=$(which shadow)
# example.com/shadow
./main.go:9:6: declaration of "i" shadows declaration at line 8
```

此外，shadow命令也可以单独使用，不需要结合`go vet`。shadow后面需要带上package名称或者.go源代码文件名。

```bash
$ shadow example.com/shadow
11-variable-shadowing/main.go:9:6: declaration of "i" shadows declaration at line 8
$ shadow main.go
11-variable-shadowing/main.go:9:6: declaration of "i" shadows declaration at line 8
```



## 总结

* 遇到variable shadowing的情况，我们需要小心，避免出现上述例子里的情况。
* 可以结合`shadow`工具做variable shadowing的自动检测。



## 推荐阅读

* [Go面试题系列，看看你会几题？](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=Mzg2MTcwNjc1Mg==&action=getalbum&album_id=2199553588283179010#wechat_redirect)

* [Go十大常见错误第1篇：未知枚举值](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484146&idx=1&sn=10fb12b643a2e37c090e5aa3bc583152&chksm=ce124d9df965c48bb954aeddabdff3db12738ded3875542250c5d0ef6cfd4417fc56580288b1&token=1912894792&lang=zh_CN#rd)

* [Go十大常见错误第2篇：benchmark性能测试的坑](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484163&idx=1&sn=b28d61c1f3ec9d914e698dce105ba5d1&chksm=ce124c6cf965c57a90bc85a5295ed9375103de20607b509f845583ff6686385df0ed96653d00&token=1912894792&lang=zh_CN#rd)

* [Go十大常见错误第3篇：go指针的性能问题和内存逃逸](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484247&idx=1&sn=faf716627afb00df646cecff023fb63c&chksm=ce124c38f965c52efd009a4c98691d56b5765dc7dce98aa49b226ad9274bd062d8d01e702e91&token=1899277735&lang=zh_CN#rd)

* [Go十大常见错误第4篇：break操作的注意事项](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484262&idx=1&sn=c1bea8af60444a4ef73c4d4d7a09d16d&chksm=ce124c09f965c51f3663ac9089a792d36c3685850e12695dd26d15a1a50f393b2d7c92b9983a&token=461369035&lang=zh_CN#rd)

* [Go十大常见错误第5篇：Go语言Error管理](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484274&idx=1&sn=711abea3c6fd5d15341ee1b34da8a160&chksm=ce124c1df965c50b3af84965f7ed30b574cd0b247ea6f77b944ec858bd43ee37f4c1554a5bce&token=1846351524&lang=zh_CN#rd)

* [Go十大常见错误第6篇：slice初始化常犯的错误](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484289&idx=1&sn=2b8171458cde4425b28fdf8f51df8d7c&chksm=ce124ceef965c5f8a14f5951457ce2ac0ecc4612cf2013957f1d818b6e74da7c803b9df1d394&token=1477304797&lang=zh_CN#rd)

* [Go十大常见错误第7篇：不使用-race选项做并发竞争检测](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484299&idx=1&sn=583c3470a76e93b0af0d5fc04fe29b55&chksm=ce124ce4f965c5f20de5887b113eab91f7c2654a941491a789e4ac53c298fbadb4367acee9bb&token=1918756920&lang=zh_CN#rd)

* [Go十大常见错误第8篇：并发编程中Context使用常见错误](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484317&idx=1&sn=474dad373684979fc96ba59182f08cf5&chksm=ce124cf2f965c5e4a29e313b4654faacef03e78da7aaf2ba6912d7b490a1df851a1bcbfec1c9&token=1918756920&lang=zh_CN#rd)

* [Go十大常见错误第9篇：使用文件名称作为函数输入](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484325&idx=1&sn=689c1b3823697cc583e1e818c4c76ee5&chksm=ce124ccaf965c5dce4e497f6251c5f0a8473b8e2ae3824bd72fe8c532d6dd84e6375c3990b3e&token=1266762504&lang=zh_CN#rd)

* [Go十大常见错误第10篇：Goroutine和循环变量一起使用的坑](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484335&idx=1&sn=cc8c6ceae72b30ec6f4d4e7b4367baca&chksm=ce124cc0f965c5d60410f977cdf31f127694fd0d49c35e2061ce8fb5fb9387bfa321196db438&token=1656737387&lang=zh_CN#rd)

  

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