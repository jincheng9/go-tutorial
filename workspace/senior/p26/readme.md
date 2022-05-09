# rGo Quiz: 从Go面试题看变量的零值和初始化赋值的注意事项

##  背景

Google工程师Valentin Deleplace出了1道关于变量初始化的题目，本来以为很简单，没想到回答正确率不到30%，拿出来和大家分享下。

### 题目

```go
// quiz.go
package main

import "fmt"

func main() {
	var a *int
	*a = 5.0
	fmt.Println(*a)
}
```

* A: `5`

* B: `5.0 `

* C: `panic`

* D: 编译错误




## 解析

这道题主要考察2个知识点：

* 变量零值。题目中`a`是一个指针类型的变量，`var a *int`这行代码没有对变量`a`初始化赋值，所以变量`a`的值是零值，指针的零值是`nil`，所以`a`的值是`nil`。

* 变量初始化赋值。如果对`int`类型的变量赋值为浮点数`5.0`是合法的，因为`5.0`是untyped float constant，是可以在不损失精度的情况下转换为`5`的。如果是赋值为`5.1`那就非法了，因为要损失小数点后面的精度，编译报错如下：

  ```bash
  ./quiz1.go:8:7: cannot use 5.1 (untyped float constant) as int value in assignment (truncated)
  ```

所以本题答案是`C`，编译的时候不会报错，但是运行的时候因为`a` 的值是`nil`，对`nil`做`*`操作就会引发panic，具体panic内容为：`panic: runtime error: invalid memory address or nil pointer dereference`。



##  思考题

题目1：

``` go
// quiz1.go
package main

import "fmt"

func main() {
	var a *int = new(int)
	*a = 5.0
	fmt.Println(*a)
}
```

题目2：

```go
// quiz2.go
package main

import "fmt"

func main() {
	var a *int = new(int)
	var b float32 = 5.0
	*a = b
	fmt.Println(*a)
}
```

题目3：

```go
// quiz3.go
package main

import "fmt"

func main() {
	var a *int = new(int)
	*a = 5.1
	fmt.Println(*a)
}
```

想知道答案的可以给公众号发送消息`init`获取答案。



## 总结

Go语言里不同类型的变量的零值不一样，给大家总结了各个类型的变量的零值：

* 数值：所有数值类型的零值都是0

  * 整数，零值是0。byte, rune, uintptr也是整数类型，所以零值也是0。
  * 浮点数，零值是0
  * 复数，零值是0+0i
  * 整数类型的变量是可以用`untyped float constant`进行赋值的，只要不损失精度即可。

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



## 开源地址

文章和示例代码开源在GitHub: [Go语言初级、中级和高级教程](https://github.com/jincheng9/go-tutorial)。

公众号：coding进阶。关注公众号可以获取最新Go面试题和技术栈。

个人网站：[Jincheng's Blog](https://jincheng9.github.io/)。

知乎：[无忌](https://www.zhihu.com/people/thucuhkwuji)。



## References

* https://twitter.com/val_deleplace/status/1522193301704781827
* https://github.com/jincheng9/go-tutorial/tree/main/workspace/lesson3