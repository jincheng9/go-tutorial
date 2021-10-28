# range迭代

* range可以用于for循环，对字符串，数组array，切片slice，集合map或通道channel进行迭代

* range对字符串string进行迭代

  * 有2种方法可以对string进行range遍历，一种是只拿到字符串的下标索引，一种是同时拿到下标索引和对应的值

    ```go
    package main
    
    import "fmt"
    
    func main() {
        str := "abcdgfg"
        // 方法1：可以通过range拿到字符串的下标索引
        for index := range(str) {
            fmt.Printf("index:%d, value:%d\n", index, str[index])
        }
        fmt.Println()
        
        // 方法2：可以通过range拿到字符串的下标索引和对应的值
        for index, value := range(str) {
            fmt.Println("index=", index, ", value=", value)
        }
        fmt.Println()
        
        // 也可以直接通过len获取字符串长度进行遍历
        for index:=0; index<len(str); index++ {
            fmt.Printf("index:%d, value:%d\n", index, str[index])
        }
    }
    ```

* range对数组array进行迭代

  * 一维数组

    ```go
    package main
    
    import "fmt"
    
    const SIZE = 4
    
    func main() {
        /*
        注意：数组的大小不能用变量，比如下面的SIZE必须是常量，如果是变量就会编译报错
        non-constant array bound size
        */
        array := [SIZE]int{1, 2, 3} 
        
        // 方法1：只拿到数组的下标索引
        for index := range array {
            fmt.Printf("index=%d value=%d ", index, array[index])
        }
        fmt.Println()
        
        // 方法2：同时拿到数组的下标索引和对应的值
        for index, value:= range array {
            fmt.Printf("index=%d value=%d ", index, value)
        }
        fmt.Println()
    }
    ```

    

  * 二维数组

    ```go
    package main
    
    import "fmt"
    import "reflect"
    
    func main() {
        array := [2][3]int{{1, 2, 3}, {4, 5, 6}}
        // 只拿到行的索引
        for index := range array {
            // array[index]类型是一维数组
            fmt.Println(reflect.TypeOf(array[index]))
            fmt.Printf("index=%d, value=%v\n", index, array[index])
        }
        
        // 拿到行索引和该行的数据
        for row_index, row_value := range array {
            fmt.Println(row_index, reflect.TypeOf(row_value), row_value)
        }
        
        // 双重遍历，拿到每个元素的值
        for row_index, row_value := range array {
            for col_index, col_value := range row_value {
                fmt.Printf("array[%d][%d]=%d ", row_index, col_index, col_value)
            }
            fmt.Println()
        }
    }
    ```

    

* range对切片slice进行迭代

  * 一维切片
  * 二维切片

* range集合map

* range通道channel

