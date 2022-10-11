# 官方教程：Go泛型入门

## 前言

本周Go官方重磅发布了Go 1.18 beta 1版本，正式支持泛型。作为Go语言诞生12年以来最大的功能变革，官方配套推出了一个非常细致的Go泛型入门基础教程，通俗易懂。

本人对Go官方教程在翻译的基础上做了一些表述上的优化，以飨读者。



## 教程内容

这个教程主要介绍Go泛型的基础知识。通过泛型，你可以声明和使用泛型函数，在调用函数的时候，允许使用不同类型的参数作为函数实参。

在这个教程里，我们先声明2个简单的非泛型函数，然后在一个泛型函数里实现这2个函数的逻辑。

接下来通过以下几个部分来进行讲解：

1. 为你的代码创建一个目录
2. 实现非泛型函数
3. 实现一个泛型函数来处理不同类型
4. 调用泛型函数的时候移除类型实参
5. 声明类型限制(type constraint)

**注意**：关于Go的其它教程，大家可以参考[https://go.dev/doc/tutorial/](https://go.dev/doc/tutorial/)。

**注意**：大家可以使用Go playground的Go dev branch模式来编写和运行你的泛型代码，地址[https://go.dev/play/?v=gotip](https://go.dev/play/?v=gotip)。



## 准备工作

* 安装Go 1.18 Beta 1或者更新的版本。安装指引可以参考[下面的介绍](#安装和使用beta版本)。
* 有一个代码编辑工具。任何文本编辑器都可以。
* 有一个命令行终端。Go可以运行在Linux，Mac上的任何命令行终端，也可以运行在Windows的PowerShell或者cmd之上。

### 安装和使用beta版本

这个教程需要使用Go 1.18 Beta 1版本了的泛型功能。使用如下步骤，安装beta版本 

1. 使用下面的命令安装beta版本

   ```sh
   $ go install golang.org/dl/go1.18beta1@latest
   ```

2. 运行如下命令来下载更新

   ```sh
   $ go1.18beta1 download
   ```

   **注意**：如果在MAC或者Linux上执行`go1.18beta1`提示`command not found`，需要设置`bash`或者`zsh`对应的profile环境变量文件。`bash`设置在`~/.bash_profile`文件里，内容为：

   ```bash
   export GOROOT=/usr/local/opt/go/libexec
   export GOPATH=$HOME/go
   export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
   ```

   `GOROOT`和`GOPATH`的值可以通过`go env`命令查看，设置完后执行`source ~/.bash_profile`让设置生效，再执行`go1.18beta1`就不报错了。

3. 使用beta版本的go命令，不要去使用release版本的go命令

   你可以通过直接使用`go1.18beta1`命令或者给`go1.18beta1`起一个简单的别名

   * 直接使用`go1.18beta1`命令

     ```sh
     $ go1.18beta1 version
     ```

   * 给`go1.18beta1`命令起一个别名

     ```sh
     $ alias go=go1.18beta1
     $ go version
     ```

   下面的教程都假设你已经把`go1.18beta1`命令设置了别名`go`。



##   为你的代码创建一个目录

首先创建一个目录用于存放你写的代码。

1. 打开一个命令行终端，切换到你的`home`目录

   * 在Linux或者Mac上执行如下命令(Linux或者Mac上只需要执行`cd`就可以进入到`home`目录)

     ```bash
     cd
     ```

   * 在Windows上执行如下命令

     ```po
     C:\> cd %HOMEPATH%
     ```

2. 在命令行终端，创建一个名为`generics`的目录

   ```bas
   $ mkdir generics
   $ cd generics
   ```

3. 创建一个go module

   运行`go mod init`命令，来给你的项目设置module路径

   ```bash
   $ go mod init example/generics
   ```

   **注意**：对于生产代码，你可以根据项目实际情况来指定module路径，如果想了解更多，可以参考[Go Module依赖管理](https://go.dev/doc/modules/managing-dependencies)。

接下来，我们来使用map写一些简单的代码。



## 实现非泛型函数

在这个步骤，你要实现2个函数，每个函数都是把map里`<key, value>`对应的所有value相加，返回总和。

你需要声明2个函数，因为你要处理2种不同类型的map，一个map存储的value是int64类型，一个map存储的value是float64类型。

### 代码实现

1. 打开你的代码编辑器，在`generics`目录创建文件`main.go`，你的代码将实现在这个文件里。

2. 进入`main.go`，在文件最开头，写包声明

   ```go
   package main
   ```

   一个独立的可执行程序总是声明在`package main`里，这点和库不一样。

3. 在包声明的下面，写如下代码

   ```go
   // SumInts adds together the values of m.
   func SumInts(m map[string]int64) int64 {
       var s int64
       for _, v := range m {
           s += v
       }
       return s
   }
   
   // SumFloats adds together the values of m.
   func SumFloats(m map[string]float64) float64 {
       var s float64
       for _, v := range m {
           s += v
       }
       return s
   }
   ```

   在上面的代码里，我们定义了2个函数，用于计算map里value的总和

   * SumInts计算value为int64类型的总和
   * SumFloats计算value为float64类型的总和

4. 在`main.go`的`package main`声明下面，实现main函数，用于初始化2个map，并把它们作为参数传递给我们实现的2个函数。

   ```go
   func main() {
   // Initialize a map for the integer values
   ints := map[string]int64{
       "first": 34,
       "second": 12,
   }
   
   // Initialize a map for the float values
   floats := map[string]float64{
       "first": 35.98,
       "second": 26.99,
   }
   
   fmt.Printf("Non-Generic Sums: %v and %v\n",
       SumInts(ints),
       SumFloats(floats))
   }
   ```

   在这段代码里，我们做了如下几个事情

   * 初始化2个map，每个map都有2个记录
   * 调用`SumInts`和`SumFloats`来分别计算2个map的value的总和
   * 打印结果

5. 在`main.go`的`package main`下面，添加`import fmt`，上面代码里调用的打印函数需要`fmt`这个package。

6. 保存`main.go`。

### 代码运行

在`main.go`所在目录下，运行如下命令

```bash
$ go run .
Non-Generic Sums: 46 and 62.97
```

使用泛型，我们只需要实现1个函数就可以计算2个不同类型map的value总和。接下来，我们会展示如何实现这个泛型函数。



## 实现一个泛型函数来处理不同类型

这个章节，我们会实现一个泛型函数，该泛型函数既可以接收value为int类型的map作为参数，也可以接收value为float类型的map作为参数，这样我们就不用为不同类型的map分别实现各自的函数了。

函数要支持这种泛型行为，需要有2个**前提条件**：

1. 对于函数而言，需要一种方式来声明这个函数到底支持哪些类型的参数
2. 对于函数调用方而言，需要一种方式来指定传给函数的到底是int类型的map还是float类型的map

为了满足以上前提条件：

1. 在声明函数的时候，除了需要像普通函数一样添加函数参数之外，还要声明**类型参数**(type parameters)。这些类型参数让函数能够实现泛型行为，让函数可以处理不同类型的参数。
2. 在函数调用的时候，除了需要像普通函数调用一样传递实参之外，还需要指定泛型函数的类型参数对应的**类型实参**(type arguments)。

每个类型参数都有一个**类型限制**(type constraint)，类型限制就好比类型参数的meta类型，每个类型限制会指明函数调用时该类型参数允许的类型实参。

尽管一个类型参数的类型限制是一系列类型的集合，但是在编译期，类型参数只会表示一种具体的类型，也就是函数调用方实际使用的类型实参。如果类型实参的类型不满足类型参数的类型限制，编译就会失败。

记住：一个类型参数一定要支持代码里对该类型所做的所有操作。例如你的函数代码试图对某个类型参数执行`string`操作，比如按照下标索引取值，但是这个类型参数的类型限制包括了数字类型，那代码就会编译失败。

在接下来的代码里，我们会使用类型限制来允许value为int类型和float类型的map作为函数的入参。

### 代码实现

1. 在上面实现的`SumInts`和`SumFloats`后面，添加如下函数

   ```go
   // SumIntsOrFloats sums the values of map m. It supports both int64 and float64
   // as types for map values.
   func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
       var s V
       for _, v := range m {
           s += v
       }
       return s
   }
   ```

   在这段代码里，我们做了如下事情：

   * 声明函数`SumIntsOrFloats`，它有2个类型参数`K`和`V`(在[ ]里面)，一个函数参数`m`，类型是`map[K]V`，返回返回类型是`V`。
   * 类型参数K的类型限制是`comparable`。`comparable`限制是Go里预声明的。它可以接受任何能做`==`和`!=`操作的类型。Go语言里map的key必须是`comparable`的，因此类型参数K的类型限制使用`comparable`是很有必要的，这也可以确保调用方使用了合法的类型作为map的key。
   * 类型参数V的类型限制是`int64`和`float64`的并集，|表示取并集，也就是`int64`和`float64`的任一个都可以满足该类型限制，可以作为函数调用方使用的类型实参。
   * 函数参数m的类型是`map[K]V`。我们知道`map[K]V`是一个合法的map类型，因为K是一个comparable的类型。如果我们不声明K为comparable，那编译器会拒绝对`map[K]V`的引用。

2. 在`main.go`已有代码后面，添加如下代码

   ```go
   fmt.Printf("Generic Sums: %v and %v\n",
       SumIntsOrFloats[string, int64](ints),
       SumIntsOrFloats[string, float64](floats))
   ```

   在这段代码里：

   * 调用了上面定义的泛型函数，传递了2种类型的map作为函数的实参。

   * 函数调用时指明了类型实参(方括号[ ]里面的类型名称)，用于替换调用的函数的类型实参。

     在接下来的内容里，你会经常看到调用函数时，会省略掉类型实参，因为Go通常(不是一定)可以根据你的代码推断出类型实参。

   * 打印函数的返回值。

### 代码运行

在`main.go`所在目录下，运行如下命令：

```bash
$ go run .
Non-Generic Sums: 46 and 62.97
Generic Sums: 46 and 62.97
```

编译器会自动把函数里的类型参数替换函数调用里指定的类型实参，在很多场景里，我们可以忽略掉这些类型实参，因为编译器可以进行自动推导。



## 调用泛型函数的时候移除类型实参

在这个章节，我们会添加一个修改版本的泛型函数调用，通过移除函数调用时的类型实参，让函数调用更为简洁。

我们在函数调用时可以移除类型实参是因为编译器可以自动推导出来，编译器是根据函数调用时传的函数实参类型做的推导判断。

**注意**：**类型实参的自动推导并不是永远可行的**。比如，你调用的泛型函数没有形参，不需要传递实参，那编译器就不能根据实参自动推导，需要在函数调用时在方括号[]里显式指定类型实参。

### 代码实现

* 在`main.go`已有代码后面，添加如下代码

  ```go
  fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
      SumIntsOrFloats(ints),
      SumIntsOrFloats(floats))
  ```

  在这段代码里，我们调用了泛型函数，忽略了类型实参，交给编译器进行自动类型推导。

### 代码运行

在`main.go`所在目录下，运行如下命令：

```go
$ go run .
Non-Generic Sums: 46 and 62.97
Generic Sums: 46 and 62.97
Generic Sums, type parameters inferred: 46 and 62.97
```

接下来，我们会进一步简化泛型函数。我们可以把int和float类型的并集做成一个可以复用的类型限制。



## 声明类型限制(type constraint)

在最后这个章节，我们会把泛型函数里的类型限制以接口(interface)的形式做定义，这样类型限制就可以在很多地方被复用。声明类型限制可以帮助精简代码，特别是在类型限制很复杂的场景下。

我们可以声明一个类型限制(type constraint)为接口(interface)类型。这样的类型限制可以允许任何实现了该接口的类型作为泛型函数的类型实参。例如，你声明了一个有3个方法的类型限制接口，然后把这个类型限制接口用于泛型函数的某个类型参数的类型限制，那函数调用时的类型实参必须要实现了接口里的所有方法。

类型限制接口也可以指代特定类型，在下面大家可以看到具体使用。

### 代码实现

1. 在`main`函数上面，`import`语句下面，添加如下代码用于声明一个类型限制

   ```go
   type Number interface {
       int64 | float64
   }
   ```

   在这段代码里，我们

   * 声明了一个名为Number的接口类型用于类型限制
   * 在接口定义里，声明了int64和float64的并集

   我们把原本来函数声明里的 int64和float64的并集改造成了一个新的类型限制接口Number，当我们需要限制类型参数为int64或float64时，就可以使用Number这个类型限制来代替`int64 | float64`的写法。

2. 在已有的函数下面，添加一个新的SumNumbers泛型函数

   ```go
   // SumNumbers sums the values of map m. Its supports both integers
   // and floats as map values.
   func SumNumbers[K comparable, V Number](m map[K]V) V {
       var s V
       for _, v := range m {
           s += v
       }
       return s
   }
   ```

   在这段代码里

   * 我们定义了一个新的泛型函数，函数逻辑和之前定义过的泛型函数`SumIntsOrFloats`完全一样，只不过对于类型参数V，我们使用了Number来作为类型限制。和之前一样，我们把类型参数用于函数形参和函数返回类型。

3. 在`main.go`已有代码后面，添加如下代码

   ```go
   fmt.Printf("Generic Sums with Constraint: %v and %v\n",
       SumNumbers(ints),
       SumNumbers(floats))
   ```

   在这段代码里

   * 我们对2个map都调用`SumNumbers`，打印每次函数调用的返回值。

     和上面一样，在这个泛型函数调用里，我们忽略了类型实参(方括号[]里面的类型名称)，Go编译器根据函数实参进行自动类型推导。

###  代码运行

在`main.go`所在目录下，运行如下命令：

```go
$ go run .
Non-Generic Sums: 46 and 62.97
Generic Sums: 46 and 62.97
Generic Sums, type parameters inferred: 46 and 62.97
Generic Sums with Constraint: 46 and 62.97
```



## 结论

目前为止，我们已经学完了Go泛型的入门知识。

如果你想继续实验，可以扩展Number接口，来支持更多的数字类型。

建议接下来了解的主题：

* Go Tour：https://go.dev/tour/welcome/1，非常棒的Go基础入门指引，一步一步教会你入门Go。
* 可以在Effective Go：https://go.dev/doc/effective_go 和 How to Write Go code：https://go.dev/doc/code 找到写Go代码的最佳实践。



## 完整代码

```go
// example6.go
package main

import "fmt"

type Number interface {
    int64 | float64
}

func main() {
    // Initialize a map for the integer values
    ints := map[string]int64{
        "first": 34,
        "second": 12,
    }

    // Initialize a map for the float values
    floats := map[string]float64{
        "first": 35.98,
        "second": 26.99,
    }

    fmt.Printf("Non-Generic Sums: %v and %v\n",
        SumInts(ints),
        SumFloats(floats))

    fmt.Printf("Generic Sums: %v and %v\n",
        SumIntsOrFloats[string, int64](ints),
        SumIntsOrFloats[string, float64](floats))

    fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
        SumIntsOrFloats(ints),
        SumIntsOrFloats(floats))

    fmt.Printf("Generic Sums with Constraint: %v and %v\n",
        SumNumbers(ints),
        SumNumbers(floats))
}

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
    var s int64
    for _, v := range m {
        s += v
    }
    return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
    var s float64
    for _, v := range m {
        s += v
    }
    return s
}

// SumIntsOrFloats sums the values of map m. It supports both floats and integers
// as map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}

// SumNumbers sums the values of map m. Its supports both integers
// and floats as map values.
func SumNumbers[K comparable, V Number](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}
```



## 开源地址

文档和代码开源地址：https://github.com/jincheng9/go-tutorial

也欢迎大家关注公众号：coding进阶，学习更多Go、微服务和云原生架构相关知识。



## References

* https://go.dev/doc/tutorial/generics







