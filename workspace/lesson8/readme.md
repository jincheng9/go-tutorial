# 函数，闭包和方法

## 函数定义

```go
func name([parameter list]) [return_types] {
  do sth
}
```



### 无参数

```go
func name() int {
  do sth
}
```



### 无返回值

```go
func name(a int) {
  do sth
}
```



### 返回1个值

```go
func name(a int) int {
  do sth
}
```



### 返回多个值

```go
func name(a int) (int, string) {
  do sth
}
func name(a, b int) (int, string) {
  do sth
}
func name(a int, b string)(int, string) {
  do sth
}
func name(a, b int, c, d string) (int, string) {
  do sth
}
```

### 给返回值命名

```go
// func2.go
package main

import "fmt"

/*
函数add的返回值有2个，类型是int，标识符分别是c和d
可以在函数体内直接给c和d赋值，return后面可以带，也可以不带返回值
*/
func addAndSub(a int, b int) (c int, d int) {
	c = a + b
	d = a - b
	return // 这一行写为 return c, d 也可以
}

func main() {
	a1, b1 := 1, 2
	c1, d1 := addAndSub(a1, b1)
	/*输出结果是：3 -1*/
	fmt.Println(c1, d1)
}
```

**注意**：

* 函数的参数列表不允许部分形参有命名，部分形参没命名，如果违背这个原则，就会报如下的编译错误。
* 函数的返回值列表不允许部分返回值变量有命名，部分返回值变量没命名，如果违背这个原则，就会报如下的编译错误。

```bash
syntax error: mixed named and unnamed function parameters
```

一句话总结：要么都不命名，要么都命名(都命名的情况下，允许形参或者返回值变量使用`_`作为命名)。



## nil函数

函数也是一种类型，函数变量的默认值是`nil`，执行`nil`函数会引发panic

```go
var f func()
// f是一个函数类型，值是nil
// 编译正常，运行报错panic: runtime error: invalid memory address or nil pointer dereference
f() 
```





## 函数参数传递

**Go里的函数传参只有值传递这一种方式**：和C++里的传值一样，参加下例里的swap

```go
package main


func add(a, b int, c, d string) (int, string) {
	return a+b, c+d
}

func swap(a int, b int) {
	println("[func|swap]a=", a, "b=", b)
	a, b = b, a
	println("[func|swap]a=", a, "b=", b)
}

func swapRef(pa *int, pb *int) {
	println("[func|swapRef]a=", *pa, "b=", *pb)
	var temp = *pa
	*pa = *pb
	*pb = temp
	println("[func|swapRef]a=", *pa, "b=", *pb)
}

func main() {
	a, b := 1, 2
	c, d := "c", "d"
	res1, res2 := add(a, b, c, d)
	println("res1=", res1, "res2=", res2)

	println("[func|main]a=", a, "b=", b)
	swap(a, b)
	println("[func|main]a=", a, "b=", b)

	println("[func|main]a=", a, "b=", b)
	swapRef(&a, &b)
	println("[func|main]a=", a, "b=", b)	
}
```

虽然swap函数无法改变外部实参的值，swapRef函数可以改变外部实参的值，但是swap和swapRef函数其实都是值传递，细节区别是：

* swap是直接把变量a和b的值拷贝一份给形参
* swapRef是把变量a和b的地址拷贝一份给形参

所以，要清楚这2个其实都是值传递，Go里的函数传参也只有值传递这一种方式，并没有像C++那样的引用变量和引用传递。

后续学习了Go里的map等变量类型，可以参考这篇文章[Go有引用变量和引用传递么？](../senior/p3)

## 函数高级用法

函数作为其它函数的实参：函数定义后可以作为另一个函数的实参，比如下例的函数realFunc作为函数calValue的实参

```go
package main

import "fmt"
import "math"

// define function getSquareRoot1
func getSquareRoot1(x float64) float64 {
	return math.Sqrt(x)
}

// deffine a function variable
var getSquareRoot2 = func(x float64) float64 {
	return math.Sqrt(x)
}

// define a function type
type callback_func func(int) int


// function calValue accepts a function variable cb as its second argument
func calValue(x int, cb callback_func) int{
	fmt.Println("[func|calValue]")
	return cb(x)
}

func realFunc(x int) int {
	fmt.Println("[func|realFunc]callback function")
	return x*x
}

func main() {
	num := 100.00
	result1 := getSquareRoot1(num)
	result2 := getSquareRoot2(num)
	fmt.Println("result1=", result1)
	fmt.Println("result2=", result2)

	value := 81
	result3 := calValue(value, realFunc) // use function realFunc as argument of calValue
	fmt.Println("result3=", result3)
}
```

### 闭包：匿名函数。顾名思义就是没有函数名。

参考下面的代码示例：

```go
// func3.go
package main

import "fmt"

func main() {
	/*
		定义2个匿名函数，也就是闭包。
		闭包可以直接调用，也可以赋值给一个变量，后续调用
	*/
	result1 := func(a int, b int) int {
		return a + b
	}(1, 2)

	var sub = func(a int, b int) int {
		return a - b
	}
	result2 := sub(1, 2)
	/*输出结果：3 -1*/
	fmt.Println(result1, result2)
}
```

## 方法：类似C++ class里的方法，只是go没有class的概念。

* 定义：function_name是类型var_data_type的实例的方法

  ```go
  func (var_name var_data_type) function_name([parameter_list])[return type] {
    do sth
  }
  ```

  

* 示例：getArea是Circle的方法，Circle的实例可以调用该方法

```go
package main

import "fmt"

type Circle struct {
	radius float64
}

func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}

/*
changeRadius和changeRadius2的区别是后者可以改变变量c的成员radius的值，前者不能改变
*/
func (c Circle) changeRadius(radius float64) {
	c.radius = radius
}

func (c *Circle) changeRadius2(radius float64) {
	c.radius = radius
}

func (c Circle) addRadius(x float64) float64{
	return c.radius + x
}

func main() {
	var c Circle
	c.radius = 10
	fmt.Println("radius=", c.radius, "area=", c.getArea())	//10, 314

	c.changeRadius(20)
	fmt.Println("radius=", c.radius, "area=", c.getArea())	//10, 314	

	c.changeRadius2(20)
	fmt.Println("radius=", c.radius, "area=", c.getArea())	//20, 1256

	result := c.addRadius(3.6)
	fmt.Println("radius=", c.radius, "result=", result) // 20, 23.6
}
```



## References

* https://yourbasic.org/golang/named-return-values-parameters/
* https://golangshowcase.com/question/mixed-named-and-unnamed-parameters-in-golang
* https://www.geeksforgeeks.org/named-return-parameters-in-golang/