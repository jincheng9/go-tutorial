# Go标准库之cmd命令使用详解

## go命令

语法：go <command> [arguments]

go命令支持的命令有如下这些：

```go
bug         start a bug report
build       compile packages and dependencies
clean       remove object files and cached files
doc         show documentation for package or symbol
env         print Go environment information
fix         update packages to use new APIs
fmt         gofmt (reformat) package sources
generate    generate Go files by processing source
get         add dependencies to current module and install them
install     compile and install packages and dependencies
list        list packages or modules
mod         module maintenance
run         compile and run Go program
test        test packages
tool        run specified go tool
version     print Go version
vet         report likely mistakes in packages
```

如果不知道这些命令怎么用，可以使用go help <cmd>查看官方说明。

### 简单的

#### go bug(给go语言官方提bug)

命令语法：

```sh
go bug
```

在终端(terminal)执行这个命令，会自动跳转到GitHub上的go语言官方repo，自动开1个issue

#### go env(处理go环境变量)

```sh
go env [-json] [-u] [-w] [var ...]
```

-json：表示把结果以json格式展示，只在查看环境变量的时候使用，不能和-u, -w放在一起使用

-u: u表示unset，恢复环境变量的默认设置，后面必须有环境变量的名称，表示具体恢复哪个环境变量的值

-w: w表示write，设置环境变量的值，后面必须跟name=value的形式，表示把环境变量name的值设置为value

主要4个场景的用法

* 查看全部环境变量的值

  ```sh
  # 方法1
  go env
  # 方法2：以json形式展示结果
  go env -json
  ```

* 查看具体某个(1个或者多个)环境变量的值

  ```sh
  # 方法1
  go env GO111MODULE GOPATH
  # 方法2
  go env | grep GO111MODULE
  # 方法3
  go env | findstr GO111MODULE
  ```

* 设置某个(1个或者多个)环境变量的值

  ```sh
  go env -w GO111MODULE=on
  ```

* 恢复某个(1个或者多个)环境变量的值为默认值

  ```sh
  go env -u GO111MODULE
  ```

####  go version(查看可执行文件的go版本)

命令语法：

```go
go version [-m] [-v] [file ...]
```

go version既可以用来查看当前系统安装的go的版本号，也可以查看可执行文件是使用哪个版本的go编译出来的，命令里最后一个参数既可以是可执行文件，也可以是目录。

-m: m代指module，后面必须带上可执行文件或者目录作为参数，用于展示可执行文件依赖的模块的版本信息

-v: v代指verbose，表示打印更为详细的信息，最后一个参数是目录的时候才真正起作用

* 查看当前系统安装的go的版本号

  ```sh
  go version
  ```

* 查看可执行文件是使用哪个版本的go编译生成的

  ```sh
  go version binfile
  ```

  比如我们使用go build main.go 编译生成了一个可执行文件main，那就可以使用如下命令来查看main这个可执行文件是使用哪个版本的go编译出来的

  ```sh
  go version main
  ```

* 查看目录下的所有可执行文件是使用哪个版本的go编译生成的

  ```sh
  go version dir
  ```

* 查看可执行文件使用的go版本以及可执行文件依赖的模块信息

  ```sh
  go version -m binfile
  go version -m dir
  ```



### 代码检查

#### go fmt

#### go vet

命令语法：

```sh
go vet [-n] [-x] [-vettool prog] [build flags] [vet flags] [packages]
```

vet的中文含义是”审查“，因此这个命令就是对Go代码做检查，报告潜在可能的错误



#### go fix



### 编译运行

#### go build

#### go clean

#### go run



### 单元测试

#### go test



### 模块管理

#### go mod

#### go install

#### go get

更新go.sum?写个example

#### go list



## References

* https://pkg.go.dev/cmd/go
* https://github.com/hyper0x/go_command_tutorial
