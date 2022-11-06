# 数组

## 一维数组

* 声明：数组的大小必须是常量，不能是变量，比如下面的语法里的size必须是常量

  * 语法

    ```go
    var variable_name [size] variable_type
    ```

  * 示例

    ```go
    var num_list [10] int
    ```

* 初始化

  * 指定数组大小

    ```go
    var float_num_list1 [5]float32 = [5]float32{1.0, 2.0, 3.0, 4.0, 5.0}
    var float_num_list2 = [5]float32{1.0, 2.0, 3.0, 4.0, 5.0}
    int_num_list := [3]int{1, 2, 3}
    for index, value := range float_num_list1 {
    	fmt.Println("[float_num_list1]index=", index, "value=", value)
    }
    
    for index, value := range float_num_list2 {
    	fmt.Println("[float_num_list2]index=", index, "value=", value)
    }
    
    for index, value := range int_num_list {
    	fmt.Println("[int_num_list]index=", index, "value=", value)
    }
    ```

  * 不显式指定数组大小，编译器根据赋的值自行推导

    ```go
    var balance1 []int = [...]int{1,2} // 等价于[2]int{1,2}
    var balance2 = [...]int{1,2,3}
    balance3 := [...]int{1, 2}
    fmt.Println("balance1=", balance1)
    fmt.Println("balance2=", balance2)
    fmt.Println("balance3=", balance3)
    ```

  * 指定数组大小情况下，特殊的初始化方式

    ```go
    balance := [5]int{1:10, 3:30} // 将数组下标为1和3的元素分别初始化为10和30
    fmt.Println(balance) // [0, 10, 0, 30, 0]
    ```

* 访问数组

  * 使用下标访问
  
    ```go
    balance := [5]int{1:10, 3:30} // 将数组下标为1和3的元素分别初始化为10和30
    fmt.Println(balance)
    
    num := balance[1]
    fmt.Println("num=", num)
    for i:=0; i<5; i++ {
    	fmt.Printf("balance[%d]=%d\n", i, balance[i])
    }
    ```
  
  * range遍历 
  
    ```go
    var float_num_list1 [5]float32 = [5]float32{1.0, 2.0, 3.0, 4.0, 5.0}
    for index := range float_num_list1 {
        // index是数组下标
        fmt.Println("[float_num_list1]index=", index) 
    }
    
    for index, value := range float_num_list1 {
        // index是数组下标，value是对应的数组元素
    	fmt.Println("[float_num_list1]index=", index, "value=", value)
    }
    ```
  
  * 获取数组长度len(array)
  
    ```go
    a := [...]int {1, 2, 3, 4, 5} 
    fmt.Println("array length=", len(a)) // array length=5
    ```

## 多维数组

* 声明：数组的大小必须是常量，不能是变量，比如下面语法里的size1，size2，...，sizeN必须是常量

  * 语法

    ```go
    var variable_name [size1][size2]...[sizeN] variable_type
    ```

  * 示例

    ```go
    var threeDimArray [2][3][4]int // 三维数组，大小是 2x3x4
    var twoDimArray [2][3] // 二维数组，大小是2x3
    ```

* 初始化

  * 初始化直接赋值

    ```go
    array1 := [2][3]int {
        {0, 1, 2},
        {3, 4, 5}, // 如果花括号}在下一行，这里必须有逗号。如果花括号在这一行可以不用逗号
    }
    ```

  * 初始化默认值，后续再赋值

    ```go
    array2 := [2][3]int{}
    array2[0][2] = 1
    array2[1][1] = 2
    fmt.Println("array2=", array2)
    ```

  * append赋值，只能对slice切片类型使用append，不能对数组使用append。参见后面lesson13里的[slice类型介绍](../lesson13)

    ```go
    twoDimArray := [][]int{}
    row1 := []int{1,2,3}
    row2 := []int{4,5}
    twoDimArray = append(twoDimArray, row1)
    fmt.Println("twoDimArray=", twoDimArray)
    
    twoDimArray = append(twoDimArray, row2)
    fmt.Println("twoDimArray=", twoDimArray)
    ```

* 访问二维数组

  * 数组下标遍历具体的元素
  
    ```go
    array1 := [2][3]int {
        {0, 1, 2},
        {3, 4, 5}}
    for i:=0; i<2; i++ {
        for j:=0; j<3; j++ {
            fmt.Printf("array1[%d][%d]=%d ", i, j, array1[i][j])
        }
        fmt.Println()
    }
    ```
  
  * 数组下标遍历某行元素
  
    ```go
    package main
    
    import "fmt"
    import "reflect"
    
    func main() {
        array := [2][3]int{{1, 2, 3}, {4, 5, 6}}
        for index := range array {
            // array[index]类型是一维数组
            fmt.Println(reflect.TypeOf(array[index])) 
            fmt.Printf("index=%d, value=%v\n", index, array[index])
        }
    }
    ```
    
    
    
  * range遍历
  
    ```go
    twoDimArray := [2][3]int {
        {0, 1, 2},
        {3, 4, 5}}
    for index := range twoDimArray {
        fmt.Printf("row %d is ", index) //index的值是0,1，表示二维数组的第1行和第2行
        fmt.Println(twoDimArray[index]) //twoDimArray[index]类型就是一维数组
    }
     for row_index, row_value := range twoDimArray {
        for col_index, col_value := range row_value {
            fmt.Printf("twoDimArray[%d][%d]=%d ", row_index, col_index, col_value)
        }
        fmt.Println()
    }
    ```
  
* 注意事项

  * slice类型的每一维度的大小可以不相同，比如下例里的第0行size是3，第1行size是2。如果直接访问twoDimArray\[2][2]会报错。slice类型的介绍参见[lesson13](../lesson13)

    ```go
    twoDimArray := [][]int{}
    row1 := []int{1,2,3}
    row2 := []int{4,5}
    twoDimArray = append(twoDimArray, row1)
    fmt.Println("twoDimArray=", twoDimArray)
    
    twoDimArray = append(twoDimArray, row2)
    fmt.Println("twoDimArray=", twoDimArray)
    ```
  
* 数组作为函数参数进行传递

  * 如果数组作为函数参数，实参和形参的定义必须相同，要么都是长度相同的数组，要么都是slice类型。如果实参和形参的类型一个是数组，一个是slice，或者实参和形参都是数组但是长度不一致都会编译报错
  
    ```go
    package main
    
    import "fmt"
    import "reflect"
    
    func sum(array [5]int, size int) int{
        sum := 0
        for i:=0; i<size; i++ {
            sum += array[i]
        }
        return sum
    }
    
    func sumSlice(array []int, size int) int{
        sum := 0
        for i:=0; i<size; i++ {
            sum += array[i]
        }
        return sum
    }
    
    func main() {
        a := [5]int {1, 2, 3, 4, 5} // a := [...]int{1, 2, 3, 4, 5}也可以去调用sum，编译器会自动推导出a的长度5
        fmt.Println("type of a:", reflect.TypeOf(a)) // type of a: [5]int
        ans := sum(a, 5)
        fmt.Println("ans=", ans)
        
        b := []int{1, 2, 3, 4, 5}
        ans2 := sumSlice(b, 5)
        fmt.Println("ans2=", ans2)
        
        array := [...]int {1}
        fmt.Println("type of array:", reflect.TypeOf(array)) // type of array: [1]int，是一个数组类型
        fmt.Println("array=", array)
    }
    ```
  
  * 值传递和引用传递
  
    * **Go语言里只有值传递，没有引用传递**。可以参考进阶篇文章[Go有引用传递么?](../senior/p3)。
  
    * 如果数组作为函数参数，在函数体内不能改变外部实参的值。如果使用数组作为形参，想修改实参的值，那就要传指向数组的指针
    
    * 如果slice作为函数参数，在函数体内可以改变外部实参的值，**但是这并不意味着slice是引用传递**，slice传参也是值传递。只不过slice这个结构里有一个指针指向底层的数组，实参把值拷贝给形参的时候，形参slice里的指针和外部实参slice的指针的值相同，也就指向了同一块数组内存空间，所以形参slice对数组元素做修改也会影响外部实参的值。
    
      ```go
      // changeArray无法改变实参数组的值
      func changeArray(array [3]int) {
          array[0] = 10
      }
      
      // changeArray2可以改变实参的值
      func changeArray2(array *[3]int) {
          array[0] = 10
      }
      
      // changeArray3可以改变实参的值
      func changeArray3(array []int) {
          array[0] = 10
      }
      
      param := [3]int{1,2,3}
      changeArray(param)
      fmt.Println("param=", param) // param= [1 2 3]
      changeArray2(&param)
      fmt.Println("param=", param) // param= [10 2 3]
      
      sliceArray := []int{1,2,3}
      changeArray3(sliceArray)
      fmt.Println("sliceArray=", sliceArray) // sliceArray= [10 2 3]
      ```
    
      
    
      
    
  
  

