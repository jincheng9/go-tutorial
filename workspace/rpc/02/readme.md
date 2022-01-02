# gRPC-Go入门教程

## protobuf简介

`Protocol Buffers(protobuf)`：与编程语言无关，与程序运行平台无关的**数据序列化协议**以及**接口定义语言**(IDL: interface definition language)。

要使用`protobuf`需要先理解几个概念：

* `protobuf`编译器`protoc`，用于编译`.proto`文件
  * 开源地址：https://github.com/protocolbuffers/protobuf

* 编程语言的`protobuf`插件，搭配`protoc`编译器，根据`.proto`文件生成对应编程语言的代码。

* `protobuf runtime library`：每个编程语言有各自的`protobuf runtime`，用于实现各自语言的`protobuf`协议。

* Go语言的`protobuf`插件和`runtime library`有过2个版本：

  * 第1个版本开源地址：https://github.com/golang/protobuf，包含有插件`proto-gen-go`，可以生成`xx.pb.go`和`xx_grpc.pb.go`。Go工程里导入该版本的`protobuf runtime`的方式如下：

    ```go
    import "github.com/golang/protobuf"
    ```

  * 第2个版本开源地址：https://github.com/protocolbuffers/protobuf-go，同样包含有插件`proto-gen-go`。不过该项目的`proto-gen-go`从`v1.20`版本开始，不再支持生成gRPC服务定义，也就是`xx_grpc.pb.go`文件。要生成gRPC服务定义需要使用`grpc-go`里的`progo-gen-go-grpc`插件。Go工程里导入该版本的`protobuf runtime`的方式如下：

    ```go
    import "google.golang.org/protobuf"
    ```

  推荐使用第2个版本，对protobuf的API做了优化和精简，并且把工程界限分清楚了：

  * 第一，把`protobuf`的Go实现都放在protobuf的项目里，而不是放在golang语言项目下面。
  * 第二，把`gRPC`的生成，放在`grpc-go`项目里，而不是和`protobuf runtime`混在一起。

  有的老项目可能使用了第1个版本的`protobuf runtime`，在老项目里开发新功能的时候也可以使用第2个版本`protobuf runtime`，支持2个版本在一个Go项目里共存。但是要**注意**：一个项目里同时使用2个版本必须保证第一个版本的版本号不低于`v1.4`。

  

## gRPC-Go简介

gRPC-Go: gRPC的Go语言实现，基于HTTP/2的RPC框架。

开源地址：https://github.com/grpc/grpc-go

Go项目里导入该模块的方式如下：

```go
import "google.golang.org/grpc"
```

`grpc-go`项目里还包含有`protoc-gen-go-grpc`插件，用于根据`.proto`文件生成`xx_grpc.pb.go`文件。



## 环境安装

分为3步：

* 安装Go

  * 步骤参考：https://go.dev/doc/install

* 安装Protobuf编译器`protoc`: 用于编译`.proto` 文件

  * 步骤参考：https://grpc.io/docs/protoc-installation/

  * 执行如下命令查看`protoc`的版本号，确认版本号是3+，用于支持protoc3

    ```bash
    protoc --version
    ```

* 安装`protoc`编译器的Go语言插件

  * `protoc-gen-go`插件：用于生成`xx.pb.go`文件

    ```bash
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    ```

  * `protoc-gen-go-grpc`插件：用于生成`xx_grpc.pb.go`文件

    ```bash
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    ```

**注意**：有的教程可能只让你安装`protoc-gen-go`，没有安装`protoc-gen-go-grpc`，那有2种情况：

* 使用的是第1个版本`github.com/golang/protobuf`的`protoc-gen-go`插件。
* 使用的是第2个版本`google.golang.org/protobuf`的`protoc-gen-go`插件并且`protoc-gen-go`版本号低于`v1.20`。从`v1.20`开始，第2个版本的`protoc-gen-go`插件不再支持生成gRPC服务定义。下面是官方说明：

> The v1.20 [`protoc-gen-go`](https://pkg.go.dev/google.golang.org/protobuf/cmd/protoc-gen-go) does not support generating gRPC service definitions. In the future, gRPC service generation will be supported by a new `protoc-gen-go-grpc` plugin provided by the Go gRPC project.
>
> The `github.com/golang/protobuf` version of `protoc-gen-go` continues to support gRPC and will continue to do so for the foreseeable future.



## 官方示例



## 工程开发



## proto3语法

https://developers.google.com/protocol-buffers/docs/proto3

https://developers.google.com/protocol-buffers/docs/reference/proto3-spec



## 开源地址

文章和示例代码开源地址在GitHub: [https://github.com/jincheng9/go-tutorial](https://github.com/jincheng9/go-tutorial)

公众号：coding进阶

个人网站：[https://jincheng9.github.io/](https://jincheng9.github.io/)

知乎：[https://www.zhihu.com/people/thucuhkwuji](https://www.zhihu.com/people/thucuhkwuji)



## References

* https://grpc.io/docs/languages/go/quickstart/

* https://github.com/protocolbuffers/protobuf-go/releases/tag/v1.20.0#v1.20-grpc-support
* https://stackoverflow.com/questions/64828054/differences-between-protoc-gen-go-and-protoc-gen-go-grpc
* https://github.com/golang/protobuf
* https://github.com/protocolbuffers/protobuf-go