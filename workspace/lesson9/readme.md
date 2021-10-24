# 变量作用域

* 局部变量

  * 函数内声明的变量，作用域只在函数体内。函数的参数和返回值也是局部变量

* 全局变量

  * 函数外声明的变量，全局变量作用域可以在当前的整个包甚至外部包(被导出后)使用

  * 全局变量和局部变量名称可以相同，但是函数内的局部变量会被优先考虑

    ```go
    package main
    
    import "fmt"
    
    
    var g int = 10
    
    func main() {
    	var g int = 20
    	fmt.Println("g=",g)	 // g=20
    }
    ```

    

* 函数形参

  * 函数定义中的参数，作为函数的局部变量来使用

* 花括号{}可以控制变量的作用域：和C++类似

  ```go
  package main
  
  import "fmt"
  
  func main() {
  	a := 10
  	{
  		a := 5
  		fmt.Println("a=", a) // a=5
  	}
  	fmt.Println("a=", a) // a=10
  }
  ```

  

