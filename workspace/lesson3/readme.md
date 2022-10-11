# 变量定义
## 全局变量

函数外定义的变量叫全局变量，以下是全局变量的定义方法。

* 方法1
```go 
var name type = value
```
* 方法2：注意，全局变量如果采用这个方式定义，那不能在全局范围内赋值，只能在函数体内给这个全局变量赋值
```go
var name type // value will be defaulted to 0, false, "" based on the type

/* 如果定义上面的全局变量，就不能紧接着在下一行通过name=value的方式对变量name做赋值，
比如name = 10，会编译报错：
 syntax error: non-declaration statement outside function body
*/
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

* **全局变量允许声明后不使用**，编译不会报错。

  

## 局部变量

函数内定义的变量叫局部变量。

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
	
* **局部变量定义后必须要被使用，否则编译报错**，报错内容为`declared but not used`。

  

## 多变量定义：

一次声明和定义多个变量

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



## 变量类型及其零值

* 零值：英文叫[zero vaue](https://go.dev/ref/spec#The_zero_value)，没有显式初始化的变量，Go编译器会给一个默认值，也叫零值。

* 数值：所有数值类型的零值都是0

  * 整数，零值是0。byte, rune, uintptr也是整数类型，所以零值也是0。
  * 浮点数，零值是0
  * 复数，零值是0+0i

* bool，零值是false

* 字符串，零值是空串""

* 指针：var a *int，零值是nil

  ```go
  num := 100
  var a * int = &num
  ```

* 切片：var a []int，零值是nil

  ```go
  var a []int = []int{1,2}
  list := [6]int{1,2} //size为6的数组，前面2个元素是1和2，后面的是默认值0
  ```

* map：var a map[string] int，零值是nil

  ```go
  dict := map[string] int{"a":1, "b":2}
  ```

* 函数：var a func(string) int，零值是nil

  ```go
  function := func(str string) string {
    return str
  }
  result := function("hello fans")
  fmt.Println("result=", result)
  ```

* channel：var a chan int，通道channel，零值是nil

  ```go
  var a chan int = make(chan int)
  var b = make(chan string)
  c := make(chan bool)
  ```

* 接口：var a interface_type，接口interface，零值是nil

  ```go
  type Animal interface {
    speak()
  }
  
  type Cat struct {
    name string
    age int
  }
  
  func(cat Cat) speak() {
    fmt.Println("miao...")
  }
  
  // 定义一个接口变量a
  var a Animal = Cat{"gaffe", 1}
  a.speak() // miao...
  ```

* 结构体:  var instance StructName，结构体里每个field的零值是对应field的类型的零值

  ```go
  type Circle struct {
    radius float64
  }
  
  var c1 Circle
  c1.radius = 10.00
  ```


## References

* https://go.dev/ref/spec#The_zero_value

