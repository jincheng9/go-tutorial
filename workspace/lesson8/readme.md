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

  * 函数作为另一个函数的实参：函数定义后可以作为另一个函数的实参
  * 闭包
  * 方法