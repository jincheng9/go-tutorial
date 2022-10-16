# Go的程序结构
* 包声明

* 引入包

* 函数

* 变量

* 语句和表达式

* 注释

  ```go
  // hello.go
  // package declaration
  package main
  
  // import package
  import "fmt"
  
  // function
  func add(a, b int) int {
    return a+b
  }
  // global variable
  var g int = 100
  
  func main() {
    a, b := 1, 2
    res := add(a, b)
    fmt.Println("a=", a, "b=", b, "a+b=", res)
    fmt.Println("g=", g)
    fmt.Println("hello world!")
  }
  ```

  

# 注意事项
* func main()是程序开始执行的函数(但是如果有func init()函数，则会先执行init函数，再执行main函数)

* 源程序文件所在的目录名称与包名称没有直接关系，不需要一致。不过通常保持一致，这符合Go的编码规范。

* 源程序文件名与包名没有直接关系，不需要将源程序文件名与文件开头申明的包名保持一样，通常这2者是不一样的。

* 只有在源程序文件开头声明package main，并且有func main()定义，才能生成可执行程序，否则go run file.go会报错，报错内容:

  ```package command-line-arguments is not a main packagego
  package command-line-arguments is not a main package
  和
  runtime.main_main·f: function main is undeclared in the main package
  ```

# 编译和运行
Go是编译型语言
* 编译+运行分步执行 
    * go build hello.go
    * ./hello
* 编译+运行一步到位
    * go run hello.go 