# Go十大常见错误第5篇：go语言Error管理

## 前言

这是Go十大常见错误系列的第5篇：go语言Error管理。素材来源于Go布道者，现Docker公司资深工程师[Teiva Harsanyi](https://teivah.medium.com/)。

本文涉及的源代码全部开源在：[Go十大常见错误源代码](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p28)，欢迎大家关注公众号，及时获取本系列最新更新。



## 场景

Go语言在错误处理(error handling)机制上经常被诟病。

在Go 1.13版本之前，Go标准库里只有一个用于构建error的`errors.New`函数，没有其它函数。

### pkg/errors包

由于Go标准库里errors包的功能比较少，所以很多人可能用过开源的[*pkg/errors*](https://github.com/pkg/errors)包来处理Go语言里的error。

比较早使用Go语言做开发，并且使用[*pkg/errors*](https://github.com/pkg/errors)包的开发者也会犯一些错误，下文会详细讲到。

`pkg/errors`包的代码风格很好，遵循了下面的error处理法则。

>  An error should be handled only **once**. Logging an error is handling an error. So an error should  either be logged or propagated.

翻译成中文就是：

> error只应该被处理一次，打印error也是对error的一种处理。所以对于error，要么打印出来，要么就把error返回传递给上一层。

很多开发者在日常开发中，如果某个函数里遇到了error，可能会先打印error，同时把error也返回给上层调用方，这就没有遵循上面的最佳实践。

我们接下来看一个具体的示例，代码逻辑是后台收到了一个RESTful的接口请求，触发了数据库报错。我们想打印如下的堆栈信息：

```
unable to serve HTTP POST request for customer 1234
 |_ unable to insert customer contract abcd
     |_ unable to commit transaction
```

假设我们使用`pkg/errors`包，我们可以使用如下代码来实现：

```go
func postHandler(customer Customer) Status {
	err := insert(customer.Contract)
	if err != nil {
		log.WithError(err).Errorf("unable to serve HTTP POST request for customer %s", customer.ID)
		return Status{ok: false}
	}
	return Status{ok: true}
}

func insert(contract Contract) error {
	err := dbQuery(contract)
	if err != nil {
		return errors.Wrapf(err, "unable to insert customer contract %s", contract.ID)
	}
	return nil
}

func dbQuery(contract Contract) error {
	// Do something then fail
	return errors.New("unable to commit transaction")
}
```

函数调用链是`postHandler` -> `insert` -> `dbQuery`。

* `dbQuery`使用`errors.New`函数创建error并返回给上层调用方。
* `insert`对`dbQuery`返回的error做了一层封装，添加了一些上下文信息，把error返回给上层调用方。
* `postHandler`打印`insert`返回的error。

函数调用链的每一层，要么返回error，要么打印error，遵循了上面提到的error处理法则。

### error判断

在业务逻辑里，我们经常会需要判断error类型，根据error的类型，决定下一步的操作：

* 比如可能做重试操作，直到成功。
* 比如可能直接打印错误日志，然后退出函数。

举个例子，假设我们使用了一个名为`db`的包，用来做数据库的读写操作。

在数据库负载比较高的情况下，调用`db`包里的方法可能会返回一个临时的`db.DBError`的错误，对于这种情况我们需要做重试。

那就可以使用如下的代码，先判断error的类型，然后根据具体的error类型做对应的处理。

```go
func postHandler(customer Customer) Status {
	err := insert(customer.Contract)
	if err != nil {
		switch errors.Cause(err).(type) {
		default:
			log.WithError(err).Errorf("unable to serve HTTP POST request for customer %s", customer.ID)
			return Status{ok: false}
		case *db.DBError:
			return retry(customer)
		}

	}
	return Status{ok: true}
}

func insert(contract Contract) error {
	err := db.dbQuery(contract)
	if err != nil {
		return errors.Wrapf(err, "unable to insert customer contract %s", contract.ID)
	}
	return nil
}
```

上面判断error的类型使用了`pkg/errors`包里的`errors.Cause`函数。

### 常见错误

对于上面的error判断，一个常见的错误是如下的代码：

```go
switch err.(type) {
default:
  log.WithError(err).Errorf("unable to serve HTTP POST request for customer %s", customer.ID)
  return Status{ok: false}
case *db.DBError:
  return retry(customer)
}
```

**可能的错误在哪里呢？**

上面代码示例里对error类型的判断使用了`err.(type)`，没有使用`errors.Cause(err).(type)`。

如果在业务函数调用链中有一个环节对`*db.DBError`做了封装，那`err.(type)`就无法匹配到`*db.DBError`，就永远不会触发重试。



## 推荐阅读

* [Go十大常见错误第1篇：未知枚举值](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484146&idx=1&sn=10fb12b643a2e37c090e5aa3bc583152&chksm=ce124d9df965c48bb954aeddabdff3db12738ded3875542250c5d0ef6cfd4417fc56580288b1&token=1912894792&lang=zh_CN#rd)

* [Go十大常见错误第2篇：benchmark性能测试的坑](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484163&idx=1&sn=b28d61c1f3ec9d914e698dce105ba5d1&chksm=ce124c6cf965c57a90bc85a5295ed9375103de20607b509f845583ff6686385df0ed96653d00&token=1912894792&lang=zh_CN#rd)

* [Go十大常见错误第3篇：go指针的性能问题和内存逃逸](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484247&idx=1&sn=faf716627afb00df646cecff023fb63c&chksm=ce124c38f965c52efd009a4c98691d56b5765dc7dce98aa49b226ad9274bd062d8d01e702e91&token=1899277735&lang=zh_CN#rd)

* [Go十大常见错误第4篇：break操作的注意事项](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484262&idx=1&sn=c1bea8af60444a4ef73c4d4d7a09d16d&chksm=ce124c09f965c51f3663ac9089a792d36c3685850e12695dd26d15a1a50f393b2d7c92b9983a&token=461369035&lang=zh_CN#rd)

* [Go面试题系列，看看你会几题？](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=Mzg2MTcwNjc1Mg==&action=getalbum&album_id=2199553588283179010#wechat_redirect)

  

## 开源地址

文章和示例代码开源在GitHub: [Go语言初级、中级和高级教程](https://github.com/jincheng9/go-tutorial)。

公众号：coding进阶。关注公众号可以获取最新Go面试题和技术栈。

个人网站：[Jincheng's Blog](https://jincheng9.github.io/)。

知乎：[无忌](https://www.zhihu.com/people/thucuhkwuji)。



## 福利

我为大家整理了一份后端开发学习资料礼包，包含编程语言入门到进阶知识(Go、C++、Python)、后端开发技术栈、面试题等。

关注公众号「coding进阶」，发送消息 **backend** 领取资料礼包，这份资料会不定期更新，加入我觉得有价值的资料。还可以发送消息「**进群**」，和同行一起交流学习，答疑解惑。



## References

* https://itnext.io/the-top-10-most-common-mistakes-ive-seen-in-go-projects-4b79d4f6cd65