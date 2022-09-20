# 基础数据类型
## 数字

* 整数：int, uint8, uint16, uint32, uint64, int8, int16, int32, int64

* 浮点数：float32, float64

* 复数：
  * complex64：实部和虚部都是float32类型的值
  
    ```go
    var v complex64 = 1 + 0.5i
    ```
  
  * complex128：实部和虚部都是float64类型的值
  
    ```go
    var v complex128 = 1 + 0.5i
    ```
  
  * **注意**：虚部为1的情况，1不能省略，否则编译报错
  
    ```go
    var v complex64 = 1 + i // compile error: undefined i
    var v complex64 = 1 + 1i // correct
    ```
  
    

## 字符串：string

* len(str)函数可以获取字符串长度

    ```go
    package main
    
    import "fmt"
    
    func main() {
        str := "abcdgfg"
        fmt.Println(len(str)) // 7
    }
    ```

* **注意**：string是immutable的，不能在初始化string变量后，修改string里的值，除非对string变量重新赋值

    ```go
    package main
    
    import "fmt"
    
    func main() {
        str := "abc"
        str = "def" // ok
        /* 下面的就不行，编译报错：cannot assign to str[0] (strings are immutable)
        str[0] = "d"
        */
        fmt.Println(str)
    }
    ```

* 字符串里字符的访问可以通过str[index]下标索引或者range迭代的方式进行访问

    ```go
    package main
    
    import "fmt"
    
    func main() {
        str := "abc"
        /*下标访问*/
        size := len(str)
        for i:=0; i<size; i++ {
            fmt.Printf("%d ", str[i])
        }
        fmt.Println()
        
        /*range迭代访问*/
        for _, value := range str {
            fmt.Printf("%d ", value)
        }
        fmt.Println()
    }
    ```

* 不能对string里的某个字符取地址：如果s[i]是字符串s里的第i个字符，那&s[i]这种方式是非法的

    ```go
    // string3.go
    package main
    
    import "fmt"
    
    func main() {
    	str := "abc"
    	/*
    	the following code has compile error:
    	cannot take the address of str[0]
    	*/
    	fmt.Println(&str[0])
    }
    ```

* string可以使用 `:` 来做字符串截取

    **注意**：这里和[切片slice](../lesson13)的截取有区别
    
    * 字符串截取后赋值给新变量，对新变量的修改不会影响原字符串的值
    * 切片截取后复制给新变量，对新变量的修改会影响原切片的值
    
    ```go
    // string4.go
    package main
    
    import "fmt"
    
    func strTest() {
    	s := "abc"
    	fmt.Println(len(s)) // 3
    	s1 := s[:]
    	s2 := s[:1]
    	s3 := s[0:]
    	s4 := s[0:2]
    	fmt.Println(s, s1, s2, s3, s4) // abc abc a abc ab
    }
    
    func main() {
    	strTest()
    }
    ```
    

* string可以用`+`做字符串拼接

  ```go
  // string5.go
  package main
  
  import "fmt"
  
  func main() {
  	a := "ch"
  	b := "ina"
  	c := a + b
  	fmt.Println(c) // china
  }
  ```

* string的更多用法可以参考：https://yourbasic.org/golang/string-functions-reference-cheat-sheet/

## bool

值只能为`true`或`false`。



## 其它数字类型

* byte：等价于uint8，数据范围0-255，定义的时候超过这个范围会编译报错
* rune：等价于int32，数据范围-2147483648-2147483647
  * 字符串里的每一个字符的类型就是rune类型，或者说int32类型
* uint：在32位机器上等价于uint32，在64位机器上等价于uint64
* uintptr: 无符号整数，是内存地址的十进制整数表示形式，应用代码一般用不到（https://stackoverflow.com/questions/59042646/whats-the-difference-between-uint-and-uintptr-in-golang）

* reflect包的`TypeOf`函数或者`fmt.Printf`的`%T`可以用来获取变量的类型

    ```go
    var b byte = 10
    var c = 'a'
    fmt.Println(reflect.TypeOf(b)) // uint8
    fmt.Println(reflect.TypeOf(c)) // int32
    fmt.Printf("%T\n", c) // int32
    ```


## References

* https://gfw.go101.org/article/basic-types-and-value-literals.html
* https://www.callicoder.com/golang-basic-types-operators-type-conversion/
* https://yourbasic.org/golang/string-functions-reference-cheat-sheet/
