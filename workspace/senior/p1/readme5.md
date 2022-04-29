# go install安装的不同Go版本的可执行程序和源码存放在哪里

## 使用场景

有新的Go版本出来后，我们通常使用`go install`来进行安装，我们以安装go 1.18 beta 2版本为例进行说明。

使用`go install`安装`go1.18beta2`的最新版本

```bash
go install golang.org/dl/go1.18beta2@latest
```

执行这条命令的结果是会下载`go1.18beta2`这个可执行程序文件到本地，那会有几个疑问：

* `go1.18beta2` 可执行文件存放在哪里？
* `go1.18beta2`要怎么使用？
* `go1.18beta2`的源码存放在哪里？即`go1.18beta2`的`GOROOT`路径是什么？
* `go1.18beta2`的`GOPATH`路径是什么？



## `go1.18beta2` 可执行文件存放在哪里？

> ```
> go install [build flags] [packages]
> ```
>
> Install compiles and installs the packages named by the import paths.
>
> Executables are installed in the directory named by the GOBIN environment variable, which defaults to $GOPATH/bin or $HOME/go/bin if the GOPATH environment variable is not set. Executables in $GOROOT are installed in $GOROOT/bin or $GOTOOLDIR instead of $GOBIN.

* 如果设置了`GOPATH`环境变量，`go1.18beta2`会安装到`$GOPATH/bin`目录下

* 如果没有设置`GOPATH`环境变量，`go1.18beta2`会安装到`$HOME/go/bin`目录下

一般都会设置`GOPATH`环境变量，可以使用如下命令查看`GOPATH`的值

```bash
go env GOPATH
```

安装完之后可以去对应目录查看是否存在`go1.18beta2` 文件。



## `go1.18beta2`要怎么使用？

通过上面`go install`下载到的`go1.18beta2`可执行文件其实还不能直接使用。

比如执行`go1.18beta2 version`命令，会得到如下结果：

```bash
go1.18beta2: not downloaded. Run 'go1.18beta2 download' to install to /Users/xxx/sdk/go1.18beta2
```

因此还要使用`go1.18beta2 download`下载`go1.18beta2`版本的源码。

```bash
$ go1.18beta2 download
Downloaded   0.0% (    16384 / 143466224 bytes) ...
Downloaded   0.3% (   360448 / 143466224 bytes) ...
Downloaded   2.8% (  3964912 / 143466224 bytes) ...
Downloaded   5.4% (  7798736 / 143466224 bytes) ...
Downloaded   7.9% ( 11321264 / 143466224 bytes) ...
Downloaded  10.8% ( 15531952 / 143466224 bytes) ...
Downloaded  14.3% ( 20479856 / 143466224 bytes) ...
Downloaded  16.8% ( 24051536 / 143466224 bytes) ...
Downloaded  19.8% ( 28360496 / 143466224 bytes) ...
Downloaded  22.6% ( 32456496 / 143466224 bytes) ...
Downloaded  24.8% ( 35553088 / 143466224 bytes) ...
Downloaded  27.9% ( 40074960 / 143466224 bytes) ...
Downloaded  30.2% ( 43302592 / 143466224 bytes) ...
Downloaded  33.0% ( 47333024 / 143466224 bytes) ...
Downloaded  35.3% ( 50691712 / 143466224 bytes) ...
Downloaded  37.4% ( 53608048 / 143466224 bytes) ...
Downloaded  40.1% ( 57474640 / 143466224 bytes) ...
Downloaded  42.5% ( 60915248 / 143466224 bytes) ...
Downloaded  45.3% ( 65060368 / 143466224 bytes) ...
Downloaded  47.4% ( 68025856 / 143466224 bytes) ...
Downloaded  49.9% ( 71597536 / 143466224 bytes) ...
Downloaded  52.5% ( 75365840 / 143466224 bytes) ...
Downloaded  55.3% ( 79298000 / 143466224 bytes) ...
Downloaded  58.6% ( 84098432 / 143466224 bytes) ...
Downloaded  61.3% ( 87981424 / 143466224 bytes) ...
Downloaded  64.0% ( 91880784 / 143466224 bytes) ...
Downloaded  66.3% ( 95092080 / 143466224 bytes) ...
Downloaded  68.3% ( 97975584 / 143466224 bytes) ...
Downloaded  70.9% (101678336 / 143466224 bytes) ...
Downloaded  73.4% (105282784 / 143466224 bytes) ...
Downloaded  75.3% (108100816 / 143466224 bytes) ...
Downloaded  78.0% (111934640 / 143466224 bytes) ...
Downloaded  80.5% (115440784 / 143466224 bytes) ...
Downloaded  83.5% (119848048 / 143466224 bytes) ...
Downloaded  85.8% (123124864 / 143466224 bytes) ...
Downloaded  87.7% (125844544 / 143466224 bytes) ...
Downloaded  90.0% (129072176 / 143466224 bytes) ...
Downloaded  92.7% (133020688 / 143466224 bytes) ...
Downloaded  95.3% (136707056 / 143466224 bytes) ...
Downloaded  97.1% (139312096 / 143466224 bytes) ...
Downloaded  99.4% (142556112 / 143466224 bytes) ...
Downloaded 100.0% (143466224 / 143466224 bytes)
Unpacking /Users/xxx/sdk/go1.18beta2/go1.18beta2.darwin-amd64.tar.gz ...
Success. You may now run 'go1.18beta2'
```

下载完毕后，就可以使用`go1.18beta2`命令来测试和验证新版本的Go了。

```bash
$ go1.18beta2 version
go version go1.18beta2 darwin/amd64
```



## `go1.18beta2`的源码存放在哪里？

执行`go1.18beta2 download`之前，如果使用`go1.18beta2 version`命令会提示源码安装位置：

```bash
go1.18beta2: not downloaded. Run 'go1.18beta2 download' to install to /Users/xxx/sdk/go1.18beta2
```

执行`go1.18beta2 download`之后，也提示了源码安装位置：

```bash
Unpacking /Users/xxx/sdk/go1.18beta2/go1.18beta2.darwin-amd64.tar.gz ...
Success. You may now run 'go1.18beta2'
```

因此源码安装位置就在`$HOME/sdk`目录下。

执行`go1.18beta2 env GOROOT`也可以查看到源码地址：

```bash
$ go1.18beta2 env GOROOT
/Users/xxx/sdk/go1.18beta2
```



## `go1.18beta2`的`GOPATH`路径是什么？

默认情况下，使用`go install`装的不同版本Go的`GOPATH`默认路径都是同一个：`$HOME/go`，如下所示：

```bash
C02G21KCMD6M:go-tutorial$ go env GOPATH
/Users/xxx/go
C02G21KCMD6M:go-tutorial$ go1.18beta1 env GOPATH
/Users/xxx/go
C02G21KCMD6M:go-tutorial$ echo $HOME
/Users/xxx
```



## References

* https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies
* https://go.dev/dl/