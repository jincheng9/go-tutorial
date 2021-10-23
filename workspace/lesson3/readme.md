# 变量定义
* 全局变量：以下是全局变量的定义方法。全局变量允许声明后不使用
  * 方法1
  ```go 
  var name type = value
  ```
  * 方法2 
  ```go
  var name type // the value will be defaulted to 0, false, "" based on the type
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

      

