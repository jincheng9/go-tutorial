# 切片Slice

## 概念

切片slice：切片是对数组的抽象。Go数组的长度在定义后是固定的，不可改变的。

切片的长度和容量是不固定的，可以动态增加元素，切片的容量也会根据情况自动扩容

* 切片的底层数据结构

  ```go
  type slice struct {
  	array unsafe.Pointer
  	len   int
  	cap   int
  }
  ```

  切片slice是个struct结构体，里面实际有个指针array，类型是unsafe.Pointer，也就是个指针，指向存放数据的数组。

  len是切片的长度，cap是切片的容量。

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

    

* 零值nil

  * 如果slice类型的变量定义后没有初始化赋值，那值就是默认值nil。对于nil切片，len和cap函数执行结果都是0。

    **注意**：下例里的slice有赋值，所以slice!=nil。slice2没有赋值，slice2==nil

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

## 切片的使用

切片访问：对切片的访问，类似数组一样，可以用下标索引或者range迭代的方式进行。可以参考[lesson10](./workspace/lesson10)和[lesson14](./workspace/lesson14)

```go
package main

import "fmt"

func main() {
	slice := make([]int, 3, 10)
	/*下标访问切片*/
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	for i:=0; i<len(slice); i++ {
		fmt.Printf("slice[%d]=%d\n", i, slice[i])		
	}

	/*range迭代访问切片*/
	for index, value := range slice {
		fmt.Printf("slice[%d]=%d\n", index, value)
	}
}
```

### 切片截取

切片截取`:`类似Python，使用冒号`:`来对数组或者切片做截取。

冒号`:`截取后的新slice变量底层有个指针，会指向原数组或者原切片的数组空间，对新切片的修改也会影响原数组或者原切片。

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

### 切片常用的几个函数

len()和cap()函数：类似C++的vector里的size和capacity

* len()：获取切片的长度，也就是实际存储了多少个元素
* cap(): 获取切片的容量。如果切片的元素个数要超过当前容量，会自动扩容

append()：通过append函数给切片加元素

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
      
      // 通过append给切片添加切片
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


copy()：拷贝一个切片里的数据到另一个切片

* 语法

  ```go
  copy(dstSlice, srcSlice) // 把srcSlice切片里的元素拷贝到dstSlice切片里
  ```


​	**注意事项**：只从源切片srcSlice拷贝min(len(srcSlice), len(dstSlice))个元素到目标切片dstSlice里。如果dstSlice的长度是0，那一个都不会从srcSlice拷贝到dstSlice里。如果dstSlice的长度M小于srcSlice的长度N，则只会拷贝srcSlice里的前M个元素到目标切片dstSlice里。

```go
package main

import "fmt"


func main() {
	a := []int{1, 2}
	b := make([]int, 1, 3) // 切片b的长度是1

	copy(b, a) // 只拷贝1个元素到b里
	fmt.Println("a=", a) // a= [1 2]
	fmt.Println("b=", b) // b= [1]
}
```

### 函数传参

* slice切片如果是函数参数，函数体内对切片底层数组的修改会影响到实参。比如下例里的change1函数第一行

* **如果在函数体内通过append直接对切片添加新元素，不会改变外部切片的值**，比如下例里的change1函数第2行。但是如果函数使用切片指针作为参数，在函数体内可以通过切片指针修改外部切片的值，比如下例里的change2函数

  ```go
  package main
  
  import "fmt"
  
  
  func change1(param []int) {
  	param[0] = 100 // 这个会改变外部切片的值
  	param = append(param, 200) // append不会改变外部切片的值
  }
  
  func change2(param *[]int) {
  	*param = append(*param, 300) // 传切片指针，通过这种方式append可以改变外部切片的值
  }
  
  func main() {
  	slice := make([]int, 2, 100)
  	fmt.Println(slice) // [0, 0]
  
  	change1(slice)
  	fmt.Println(slice) // [100, 0]
  
  	change2(&slice)
  	fmt.Println(slice) // [100, 0, 300]
  }
  ```


## 切片的底层原理

1. [Go Quiz: 从Go面试题看slice的底层原理和注意事项](https://github.com/jincheng9/go-tutorial/blob/main/workspace/senior/p8)
2. [Go Quiz: 从Go面试题搞懂slice range遍历的坑](https://github.com/jincheng9/go-tutorial/blob/main/workspace/senior/p13)
