# 一文读懂Go 1.20引入的PGO性能优化

## 背景

Go 1.20版本于2023年2月份正式发布，在这个版本里引入了PGO性能优化机制。

PGO的英文全称是Profile Guided Optimization，基本原理分为以下2个步骤：

* 先对程序做profiling，收集程序运行时的数据，生成profiling文件。
* 编译程序时启用PGO选项，编译器会根据.pgo文件里的内容对程序做性能优化。

我们都知道在编译程序的时候，编译器会对程序做很多优化，包括大家熟知的内联优化(inline optimization)、逃逸分析(escape analysis)、常数传播(constant propagation)。这些优化是编译器可以直接通过分析程序源代码来实现的。

但是有些优化是无法通过解析源代码来实现的。

比如一个函数里有很多if/else条件分支判断，我们可能希望编译器自动帮我们优化条件分支顺序，来加快条件分支的判断，提升程序性能。

但是，编译器可能是无法知道哪些条件分支进入的次数多，哪些条件分支进入的次数少，因为这个和程序的输入是有关系的。

这个时候，做编译器优化的人就想到了PGO: Profile Guided Optimization。

PGO的原理很简单，那就是先把程序跑起来，收集程序运行过程中的数据。然后编译器再根据收集到的程序运行时数据来分析程序的行为，进而做针对性的性能优化。

比如程序可以收集到哪些条件分支进入的次数更多，就把该条件分支的判断放在前面，这样可以减少条件判断的耗时，提升程序性能。

那Go语言如何使用PGO来优化程序的性能呢？我们接下来看看具体的例子。

## 示例

我们实现一个web接口`/render`，该接口以markdown文件的二进制格式作为输入，将markdown格式转换为html格式返回。

我们借助 [`gitlab.com/golang-commonmark/markdown`](https://pkg.go.dev/gitlab.com/golang-commonmark/markdown) 项目来实现该接口。

### 环境搭建

```bash
$ go mod init example.com/markdown
```

新建一个 `main.go`文件，代码如下：

```go
package main

import (
    "bytes"
    "io"
    "log"
    "net/http"
    _ "net/http/pprof"

    "gitlab.com/golang-commonmark/markdown"
)

func render(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
        return
    }

    src, err := io.ReadAll(r.Body)
    if err != nil {
        log.Printf("error reading body: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    md := markdown.New(
        markdown.XHTMLOutput(true),
        markdown.Typographer(true),
        markdown.Linkify(true),
        markdown.Tables(true),
    )

    var buf bytes.Buffer
    if err := md.Render(&buf, src); err != nil {
        log.Printf("error converting markdown: %v", err)
        http.Error(w, "Malformed markdown", http.StatusBadRequest)
        return
    }

    if _, err := io.Copy(w, &buf); err != nil {
        log.Printf("error writing response: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
}

func main() {
    http.HandleFunc("/render", render)
    log.Printf("Serving on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

编译和运行该程序：

```bash
$ go mod tidy
$ go build -o markdown.nopgo
$ ./markdown.nopgo
2023/02/25 22:30:51 Serving on port 8080...
```

程序主目录下新建`input.md`文件，内容可以自定义，符合markdown语法即可。

我演示的例子里用到了[input.md](https://raw.githubusercontent.com/golang/go/c16c2c49e2fa98ae551fc6335215fadd62d33542/README.md) 这个markdown文件。

通过curl命令发送markdown文件的二进制内容给`/render`接口。

```bash
$ curl --data-binary @input.md http://localhost:8080/render
<h1>The Go Programming Language</h1>
<p>Go is an open source programming language that makes it easy to build simple,
reliable, and efficient software.</p>
...
```

可以看到该接口返回了`input.md`文件内容对应的html格式。

### Profiling

那接下来我们给`main.go`程序做profiling，得到程序运行时的数据，然后通过PGO来做性能优化。

在`main.go`里，有import [net/http/pprof](https://pkg.go.dev/net/http/pprof) 这个库，它会在原来已有的web接口`/render`的基础上，新增一个新的web接口`/debug/pprof/profile`，我们可以通过请求这个profiling接口来获取程序运行时的数据。

* 在程序主目录下，新增load子目录，在load子目录下新增`main.go`的文件，`load/main.go`运行时会不断请求上面`./markdown.nogpo`启动的server的`/render`接口，来模拟程序实际运行时的情况。

  ```bash
  $ go run example.com/markdown/load
  ```

* 请求profiling接口来获取程序运行时数据。

  ```bash
  $ curl -o cpu.pprof "http://localhost:8080/debug/pprof/profile?seconds=30"
  ```

​		等待30秒，curl命令会结束，在程序主目录下会生成`cpu.pprof`文件。

**注意**：要使用Go 1.20版本去编译和运行程序。

### PGO优化程序

```bash
$ mv cpu.pprof default.pgo
$ go build -pgo=auto -o markdown.withpgo
```

`go build`编译程序的时候，启用`-pgo`选项。

`-pgo`既可以支持指定的profiling文件，也可以支持`auto`模式。

如果是`auto`模式，会自动寻找程序主目录下名为`default.pgo`的profiling文件。

Go官方推荐大家使用`auto`模式，而且把`default.pgo`文件也存放在程序主目录下维护，这样方便项目所有开发者使用`default.pgo`来对程序做性能优化。

Go 1.20版本里，`-pgo`选项的默认值是`off`，我们必须添加`-pgo=auto`来开启PGO优化。

未来的Go版本里，官方计划将`-pgo`选项的默认值设置为`auto`。

### 性能对比

在程序的子目录`load`下新增`bench_test.go`文件，`bench_test.go`里使用Go性能测试的Benchmark框架来给server做压力测试。

#### 未开启PGO优化的场景

启用未开启PGO优化的server程序：

```bash
$ ./markdown.nopgo
```

开启压力测试：

```bash
$ go test example.com/markdown/load -bench=. -count=20 -source ../input.md > nopgo.txt
```

#### 开启PGO优化的场景

启用开启了PGO优化的server程序：

```bash
$ ./markdown.withpgo
```

开启压力测试：

```bash
$ go test example.com/markdown/load -bench=. -count=20 -source ../input.md > withpgo.txt
```

#### 综合对比

通过上面压力测试得到的`nopgo.txt`和`withpgo.txt`来做性能比较。

```bash
$ go install golang.org/x/perf/cmd/benchstat@latest
$ benchstat nopgo.txt withpgo.txt
goos: darwin
goarch: amd64
pkg: example.com/markdown/load
cpu: Intel(R) Core(TM) i5-5250U CPU @ 1.60GHz
       │  nopgo.txt  │             withpgo.txt             │
       │   sec/op    │   sec/op     vs base                │
Load-4   447.3µ ± 7%   401.3µ ± 1%  -10.29% (p=0.000 n=20)
```

可以看到，使用PGO优化后，程序的性能提升了10.29%，这个提升效果非常可观。

在Go 1.20版本里，使用PGO之后，通常程序的性能可以提升2%-4%左右。

后续的版本里，编译器还会继续优化PGO机制，进一步提升程序的性能。

## 总结

Go 1.20版本引入了PGO来让编译器对程序做性能优化。PGO使用分2个步骤：

* 先得到一个profiling文件。
* 使用`go build`编译时开启PGO选项，通过profiling文件来指导编译器对程序做性能优化。

在生产环境里，我们可以收集近段时间的profiling数据，然后通过PGO去优化程序，以提升系统处理性能。

更多关于PGO的使用说明和最佳实践可以参考[profile-guided optimization user guide](https://go.dev/doc/pgo)。

源代码地址：[pgo optimization source code](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p35)。

## 推荐阅读

* [Go 1.20来了，看看都有哪些变化](https://mp.weixin.qq.com/s?__biz=Mzg2MTcwNjc1Mg==&mid=2247484693&idx=1&sn=9f84d42dfadb7319f8c4e4645893d218&chksm=ce124a7af965c36c63deafc09b9f2bfdae35bc8714aa2f76bca63e233f664b499bad742a8c3f&token=293290824&lang=zh_CN#rd)

* [Go面试题系列，看看你会几题](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=Mzg2MTcwNjc1Mg==&action=getalbum&album_id=2199553588283179010#wechat_redirect)

* [Go常见错误和最佳实践系列](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=Mzg2MTcwNjc1Mg==&action=getalbum&album_id=2549657749539028992#wechat_redirect)

  

## 开源地址

文章和示例代码开源在GitHub: [Go语言初级、中级和高级教程](https://github.com/jincheng9/go-tutorial)。

公众号：coding进阶。关注公众号可以获取最新Go面试题和技术栈。

个人网站：[Jincheng's Blog](https://jincheng9.github.io/)。

知乎：[无忌](https://www.zhihu.com/people/thucuhkwuji)。



## 福利

我为大家整理了一份后端开发学习资料礼包，包含编程语言入门到进阶知识(Go、C++、Python)、后端开发技术栈、面试题等。

关注公众号「coding进阶」，发送消息 **backend** 领取资料礼包，这份资料会不定期更新，加入我觉得有价值的资料。还可以发送消息「**进群**」，和同行一起交流学习，答疑解惑。



## References

* https://go.dev/blog/pgo-preview