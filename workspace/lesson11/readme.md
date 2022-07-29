# 指针

## 指针的基础语法 

指针的值是指向的变量的内存地址。

* 语法

  ```go
  var var_name *var_type
  ```

* 示例

  ```go
  var intPtr *int
  ```

* 初始化

  ```go
  package main
  
  import "fmt"
  import "reflect"
  
  func main() {
      i := 10
      // 方式1
      var intPtr *int = &i
      fmt.Println("pointer value:", intPtr, " point to: ", *intPtr)
      fmt.Println("type of pointer:", reflect.TypeOf(intPtr))
      
      // 方式2
      intPtr2 := &i
      fmt.Println(*intPtr2)
      fmt.Println("type of pointer:", reflect.TypeOf(intPtr2))
      
      // 方式3
      var intPtr3 = &i;
      fmt.Println(*intPtr3)
      fmt.Println("type of pointer:", reflect.TypeOf(intPtr3))
      
      // 方式4
      var intPtr4 *int
      intPtr4 = &i
      fmt.Println(*intPtr4)
      fmt.Println("type of pointer:", reflect.TypeOf(intPtr4))
  }
  ```




## 默认值

* 不赋值的时候，默认值是nil

  ```go
  var intPtr5 *int    
  fmt.Println("intPtr5==nil:", intPtr5==nil) // intPtr5==nil: true
  ```



## 指向数组的指针

* 注意这里和C++不一样，C++的数组名就是指向数组首元素的地址，Go不是

  ```go
  array := [3]int{1,2,3}
  var arrayPtr *[3]int = &array // C++赋值就不用加&
  for i:=0; i<len(array); i++ {
    // arrayPtr[i]的值就是数组array里下标索引i对应的值
  	fmt.Printf("arrayPtr[%d]=%d\n", i, arrayPtr[i])
  }
  ```
  



## 指针数组

* 定义

  ```go
  var ptr [SIZE]*int // 指向int的指针数组，数组里有多个指针，每个都指向一个int
  ```

* 使用

  ```go
  package main
  
  import "fmt"
  
  const SIZE = 5
  
  func main() {
      var ptrArray [SIZE]*int
      a := [5]int{1,2,3,4,5}
      for i:=0; i<SIZE; i++ {
          ptrArray[i] = &a[i]
      }
      
      for i:=0; i<SIZE; i++ {
          fmt.Printf("%d ", *ptrArray[i])
      }
      fmt.Println()
  }
  ```

  

## 指向指针的指针

* 定义

  ```go
  var a int = 100
  var ptr1 *int = &a
  var ptr2 **int = &ptr1
  var ptr3 ***int = &ptr2
  ```

  

* 使用

  ```go
  package main
  
  import "fmt"
  
  func main() {
      var a int = 100
      var ptr1 *int = &a
      var ptr2 **int = &ptr1
      var ptr3 ***int = &ptr2
      
      fmt.Println("*ptr1=", *ptr1)
      fmt.Println("**ptr2=", **ptr2)
      fmt.Println("***ptr3=", ***ptr3)
  }
  ```

  

## 向函数传递指针参数

* 示例：通过指针参数修改实参的值

  ```go
  package main
  
  import "fmt"
  
  // 这个可以交换外部传入的2个实参的值
  func swap(a *int, b *int) {
      *a, *b = *b, *a
  }
  
  // 这个无法交换外部传入的2个实参的值
  func swap2(a *int, b *int) {
      a, b = b, a
  }
  
  
  func main() {
      a, b := 1, 2
      swap(&a, &b)
      fmt.Println("a=", a, " b=", b) // a= 2  b= 1
      
      swap2(&a, &b)
      fmt.Println("a=", a, " b=", b) // a= 2  b= 1
  }
  ```



## 指向结构体的指针

指向结构体的指针在访问结构体成员的时候使用点`.`，和C++里用箭头->不一样。具体参见[指向结构体的指针](./workspace/lesson12)

