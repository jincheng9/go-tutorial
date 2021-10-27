# 切片Slice

* 切片slice：切片是对数组的抽象。Go数组的长度在定义后是固定的，不可改变的。切片的长度和容量是不固定的，可以动态增加元素，切片的容量也会根据情况扩容

* 定义和初始化

  * 语法

    ```go
    var slice_var []data_type // 元素类型为data_type的切片
    var slice_var []data_type = make([]data_type, len, cap)// cap是切片容量，是make的可选参数
    var slice_var []data_type = make([]data_type, len)
    slice_var := []data_type{}
    slice_var := make([]data_type, len)
    ```

  * 示例

    ```go
    package main
    
    import "fmt"
    
    
    func printSlice(param []int) {
        fmt.Printf("slice len:%d, cap:%d, value:%v\n", len(param), cap(param), param)
    }
    
    func main() {
        slice1 := []int{1}
        slice2 := make([]int, 3, 100)
        printSlice(slice1)
        printSlice(slice2)
    }
    ```

    

* 默认值nil

  * 如果slice类型的变量定义后没有初始化赋值，那值就是默认值nil。

    **注意：**下例里的slice有赋值，所以slice!=nil。slice2没有赋值，slice2==nil

    ```go
    package main
    
    import "fmt"
    
    func printSlice(param []int) {
        fmt.Printf("param len:%d, cap:%d, value:%v\n", len(param), cap(param), param)
    }
    
    func main() {
        slice := []int{}
        var slice2 []int
        
        fmt.Println("slice==nil", slice==nil) // false
        printSlice(slice)
        
        fmt.Println("slice2==nil", slice2==nil) // true
        printSlice(slice2)
    }
    ```

* 切片截取：类似Python，使用冒号**:**来对数组或者切片做截取。切片后的变量是对原数组或者切片的**引用**，对新切片的修改也会影响原数组或者原切片。

  ```go
  package main
  
  import "fmt"
  import "reflect"
  
  
  func printSlice(param []int) {
      fmt.Printf("param len:%d, cap:%d, value:%v\n", len(param), cap(param), param)
  }
  
  func main() {
      slice := []int{}
      var slice2 []int
      
      fmt.Println("slice==nil", slice==nil) // false
      printSlice(slice)
      
      fmt.Println("slice2==nil", slice2==nil) // true
      printSlice(slice2)
      
      // 对数组做切片
      array := [3]int{1,2,3} // array是数组
      slice3 := array[1:3] // slice3是切片
      fmt.Println("slice3 type:", reflect.TypeOf(slice3))
      fmt.Println("slice3=", slice3) // slice3= [2 3]
      
      slice4 := slice3[1:2]
      fmt.Println("slice4=", slice4) // slice4= [3]
      
      /* slice5->slice4->slice3->array
      对slice5的修改，会影响到slice4, slice3和array
      */
      slice5 := slice4[:]
      fmt.Println("slice5=", slice5) // slice5= [3]
      
      slice5[0] = 10
      fmt.Println("array=", array) // array= [1 2 10]
      fmt.Println("slice3=", slice3) // slice3= [2 10]
      fmt.Println("slice4=", slice4) // slice4= [10]
      fmt.Println("slice5=", slice5) // slice5= [10]
  }
  ```

  

* len()和cap()函数：类似C++的vector里的size和capacity

  * len()：获取切片的长度，也就是实际存储了多少个元素
  * cap(): 获取切片的容量。如果切片的元素个数要超过当前容量，会自动扩容

* append()函数：通过append函数给切片加元素

  * append不改变原切片的值，比如下例里的append(slice, 4)并不会改变slice的值

  * 只能对切片使用append()函数，不能对数组使用append()

    ```go
    package main
    
    import "fmt"
    
    func main() {
        slice := []int{1, 2, 3}
        // 往原切片里加一个元素
        test := append(slice, 4)
        // append不会改变slice的值，除非把append的结果重新赋值给slice
        fmt.Println(slice) // [1 2 3]
        fmt.Println(test) // [1 2 3 4]
        
        // append添加切片
        temp := []int{1,2}
        test = append(test, temp...) // 注意，第2个参数有...结尾
        fmt.Println(test) // [1 2 3 4 1 2]
        
        /*下面对array数组做append就会报错:  first argument to append must be slice; have [3]int
        array := [3]int{1, 2, 3}
        array2 := append(array, 1)
        fmt.Println(array2)
        */
    }
    ```

    

* copy()函数：拷贝一个切片里的数据到另一个切片

* 指向切片的指针

  ```go
  ```

  

* 函数传参：slice如果是函数参数，则都是**引用**传递，函数体内对切片的修改会影响到实参
