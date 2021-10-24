# 函数

* 函数定义

  ```go
  func name([parameter list]) [return_types] {
    do sth
  }
  ```

  

  * 无参数

    ```go
    func name() int {
      do sth
    }
    ```

    

  * 无返回值

    ```go
    func name(a int) {
      do sth
    }
    ```

    

  * 返回1个值

    ```go
    func name(a int) int {
      do sth
    }
    ```

    

  * 返回多个值

    ```go
    func name(a int) (int, string) {
      do sth
    }
    func name(a b int) (int, string) {
      do sth
    }
    func name(a int, b string)(int, string) {
      do sth
    }
    func name(a, b int, c, d string) (int, string) {
      do sth
    }
    ```

    

* 函数参数

  * 值传递：和C++里的传值一样，参加下例里的swap

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

  * 引用传递：和 C++里的传指针一样，参见下例里的swapRef

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

    

* 函数高级用法

  * 函数作为其它函数的实参：函数定义后可以作为另一个函数的实参，比如下例的函数realFunc作为函数calValue的实参

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

  * 闭包

  * 方法：类似C++ class里的方法，只是go没有class的概念。

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

    