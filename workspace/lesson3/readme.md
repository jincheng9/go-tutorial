# 变量定义
* 全局变量：以下是全局变量的定义方法。全局变量允许声明后不使用
  * 方法1
  ```go 
  var name type = value
  ```
  * 方法2：注意，全局变量如果采用这个方式定义，那不能在全局范围内赋值，只能在函数体内给这个全局变量赋值
  ```go
  var name type // value will be defaulted to 0, false, "" based on the type
  /* 如果定义上面的全局变量，就不能紧接着对变量name做赋值，比如name = 10，会编译报错：
   syntax error: non-declaration statement outside function body
  *、
  ```
  * 方法3
  ```go
  var name = value 
  ```
  * 方法4
  ```go
  var (
  	v1 int = 10
  	v2 bool = true
  )
  var (
  	v5 int   // the value will be defaulted to 0
  	v6 bool  // the value will be defaulted to false
  )
  var (
  	v3 = 20
  	v4 = false
  )
  ```
  
* 局部变量：
  * 和全局变量的定义相比，多了以下定义方法
    * 方法5
    ```go
    name := value
    ```
    * 方法6
  	```go
  	var name type
  	name = value
  	```
  	
  * 局部变量定义后必须要被使用，否则编译报错
  
* 多变量定义：一次声明和定义多个变量

  * 全局变量

    * 方法1

      ```go
      var a, b, c int = 1, 2, 3
      ```

    * 方法2

      ```go
      var a, b, c bool
      ```

    * 方法3

      ```go
      var a, b, c = 1, 2, "str"
      ```

  * 局部变量：和全局变量相比，多了以下定义方法

    * 方法4

      ```go
      var a, b int
      a, b = 1, 2
      
      var c, d int
      c = 10
      d = 20
      ```

    * 方法5

      ```go
      a, b := 1, 2
      a1, b1 := 1, "str"
      ```

* 变量类型

  * 数值：整数，浮点数，复数

  * bool

  * 字符串

  * 指针：var a *int

    ```go
    num := 100
    var a * int = &num
    ```

  * 数组：var a []int

    ```go
    var a []int = []int{1,2}
    list := [6]int{1,2} //size为6的数组，前面2个元素是1和2，后面的是默认值0
    ```

  * map：var a map[string] int

    ```go
    dict := map[string] int{"a":1, "b":2}
    ```

  * 函数：var a func(string) int

    ```go
    function := func(str string) string {
      return str
    }
    result := function("hello fans")
    fmt.Println("result=", result)
    ```

  * 结构体:  var instance Struct

    ```go
    type Circle struct {
      redius float64
    }
    
    var c1 Circle
    c1.radius = 10.00
    ```

  * channel：var a chan int

  * 接口：var a error // error是接口

