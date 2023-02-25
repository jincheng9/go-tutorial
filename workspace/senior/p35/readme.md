# 一文读懂Go 1.21引入的PGO性能优化

## 背景

Go 1.21版本于2023年2月份正式发布，在这个版本里引入了PGO性能优化机制。

PGO的英文全称是Profile Guided Optimization，基本原理分为以下2个步骤：

* 先对程序做profile，得到一个.pgo文件
* 编译程序时启用pgo，编译器会根据.pgo文件里的内容对程序做性能优化

编译器会对程序做很多优化，包括大家熟知的内联优化(inline optimization)、逃逸分析(escape analysis)、常数传播(constant propagation)。

这些优化是编译器可以直接通过分析程序源代码来实现的。

但是有些优化是无法通过解析源代码来实现的。

比如一个函数里有很多if/else条件分支判断，我们可能希望编译器自动帮我们优化条件分支顺序，来加快条件分支的判断，提升程序性能。

但是，编译器可能是无法知道哪些条件分支进入的次数多，哪些条件分支进入的次数少，因为这个和程序的输入是有关系的。

这个时候，做编译器优化的人就想到了PGO: Profile Guided Optimization。

PGO的原理很简单，那就是先把程序跑起来，收集程序运行过程中的数据。然后编译器再根据收集到的程序运行的数据来分析程序的行为，进而做针对性的性能优化。

比如程序可以收集到哪些条件分支进入的次数更多，就把该条件分支的判断放在前面，这样可以减少条件判断的耗时，提升程序性能。

那Go语言如何使用PGO来优化程序的性能呢？我们接下来看看具体的例子。

## 实例

我们来实现一个web接口render，该接口以markdown文件作为输入，将markdown格式转换为html格式返回。

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

自己本地新建一个input.md文件，内容可以自定义，符合markdown语法即可。

我演示的例子里用到了https://raw.githubusercontent.com/golang/go/c16c2c49e2fa98ae551fc6335215fadd62d33542/README.md 这个markdown文件。

通过curl命令发送markdown文件的二进制内容给render接口。

```bash
$ curl --data-binary @input.md http://localhost:8080/render
<h1>The Go Programming Language</h1>
<p>Go is an open source programming language that makes it easy to build simple,
reliable, and efficient software.</p>
...
```

可以看到该接口返回了input.md文件内容对应的html格式。

### Profiling

Now that we have a working service, let’s collect a profile and rebuild with PGO to see if we get better performance.

In `main.go`, we imported [net/http/pprof](https://pkg.go.dev/net/http/pprof) which automatically adds a `/debug/pprof/profile` endpoint to the server for fetching a CPU profile.

Normally you want to collect a profile from your production environment so that the compiler gets a representative view of behavior in production. Since this example doesn’t have a “production” environment, we will create a simple program to generate load while we collect a profile. Copy the source of [this program](https://go.dev/play/p/yYH0kfsZcpL) to `load/main.go` and start the load generator (make sure the server is still running!).

```
$ go run example.com/markdown/load
```

While that is running, download a profile from the server:

```
$ curl -o cpu.pprof "http://localhost:8080/debug/pprof/profile?seconds=30"
```

Once this completes, kill the load generator and the server.

### Using the profile

We can ask the Go toolchain to build with PGO using the `-pgo` flag to `go build`. `-pgo` takes either the path to the profile to use, or `auto`, which will use the `default.pgo` file in the main package directory.

We recommending commiting `default.pgo` profiles to your repository. Storing profiles alongside your source code ensures that users automatically have access to the profile simply by fetching the repository (either via the version control system, or via `go get`) and that builds remain reproducible. In Go 1.20, `-pgo=off` is the default, so users still need to add `-pgo=auto`, but a future version of Go is expected to change the default to `-pgo=auto`, automatically giving anyone that builds the binary the benefit of PGO.

Let’s build:

```
$ mv cpu.pprof default.pgo
$ go build -pgo=auto -o markdown.withpgo.exe
```

### Evaluation

We will use a Go benchmark version of the load generator to evaluate the effect of PGO on performance. Copy [this benchmark](https://go.dev/play/p/6FnQmHfRjbh) to `load/bench_test.go`.

First, we will benchmark the server without PGO. Start that server:

```
$ ./markdown.nopgo.exe
```

While that is running, run several benchmark iterations:

```
$ go test example.com/markdown/load -bench=. -count=20 -source ../README.md > nopgo.txt
```

Once that completes, kill the original server and start the version with PGO:

```
$ ./markdown.withpgo.exe
```

While that is running, run several benchmark iterations:

```
$ go test example.com/markdown/load -bench=. -count=20 -source ../README.md > withpgo.txt
```

Once that completes, let’s compare the results:

```
$ go install golang.org/x/perf/cmd/benchstat@latest
$ benchstat nopgo.txt withpgo.txt
goos: linux
goarch: amd64
pkg: example.com/markdown/load
cpu: Intel(R) Xeon(R) W-2135 CPU @ 3.70GHz
        │  nopgo.txt  │            withpgo.txt             │
        │   sec/op    │   sec/op     vs base               │
Load-12   393.8µ ± 1%   383.6µ ± 1%  -2.59% (p=0.000 n=20)
```

The new version is around 2.6% faster! In Go 1.20, workloads typically get between 2% and 4% CPU usage improvements from enabling PGO. Profiles contain a wealth of information about application behavior and Go 1.20 just begins to crack the surface by using this information for inlining. Future releases will continue improving performance as more parts of the compiler take advantage of PGO.

## Next steps

In this example, after collecting a profile, we rebuilt our server using the exact same source code used in the original build. In a real-world scenario, there is always ongoing development. So we may collect a profile from production, which is running last week’s code, and use it to build with today’s source code. That is perfectly fine! PGO in Go can handle minor changes to source code without issue.

For much more information on using PGO, best practices and caveats to be aware of, please see the [profile-guided optimization user guide](https://go.dev/doc/pgo).

Please send us your feedback! PGO is still in preview and we’d love to hear about anything that is difficult to use, doesn’t work correctly, etc. Please file issues at https://go.dev/issue/new.

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