# 变量
* 全局变量：以下是全局变量的定义方法，定义后全局变量在代码里可以不使用
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
	```go
	name := value
	```
	* 局部变量定义后必须要被使用，否则编译报错
