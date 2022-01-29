# range迭代

range可以用于for循环，对字符串，数组array，切片slice，集合map或通道channel进行迭代

## range对字符串string进行迭代

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

## range对数组array进行迭代

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

  

## range对切片slice进行迭代

* 一维切片：会根据切片的长度len()进行遍历

  ```go
  package main
  
  import "fmt"
  
  func main() {
      slice := []int{1,2,3}
      // 方式1
      for index := range slice {
          fmt.Printf("index=%d, value=%d\n", index, slice[index])
      }
      // 方式2
      for index, value := range slice {
          fmt.Printf("index=%d, value=%d\n", index, value)
      }
  }
  ```

  

* 二维切片：range遍历方式类似二维数组

  ```go
  package main
  
  import "fmt"
  import "reflect"
  
  func main() {
      slice := [][]int{{1,2}, {3, 4, 5}}
      fmt.Println(len(slice))
      // 方法1，拿到行索引
      for index := range slice{
          fmt.Printf("index=%d, type:%v, value=%v\n", index, reflect.TypeOf(slice[index]), slice[index])
      }
      
      // 方法2，拿到行索引和该行的值，每行都是一维切片
      for row_index, row_value := range slice{
          fmt.Printf("index=%d, type:%v, value=%v\n", row_index, reflect.TypeOf(row_value), row_value)
      }
      
      // 方法3，双重遍历，获取每个元素的值
      for row_index, row_value := range slice {
          for col_index, col_value := range row_value {
              fmt.Printf("slice[%d][%d]=%d ", row_index, col_index, col_value)
          }
          fmt.Println()
      }
  }
  ```

  

## range对集合map进行迭代

* 有如下2种方法可以遍历map，一种是拿到key，一种是拿到key,value

  ```go
  package main
  
  import "fmt"
  
  func main() {
      hash := map[string]int{"a":1}
      // 方法1，拿到key，再根据key获取value
      for key := range hash{
          fmt.Printf("key=%s, value=%d\n", key, hash[key])
      }
      
      // 方法2，同时拿到key和value
      for key, value := range hash{
          fmt.Printf("key=%s, value=%d\n", key, value)
      }
      
      /* nil map不能存放key-value键值对，比如下面的方式会报错：panic: assignment to entry in nil map
      var hash2 map[string]int 
      hash2["a"] = 1
      */
  }
  ```

  

## range对通道channel进行迭代

对channel进行range迭代，会循环从channel里取数据

```go
package main

import "fmt"
import "time"


func addData(ch chan int) {
	/*
	每3秒往通道ch里发送一次数据
	*/
	size := cap(ch)
	for i:=0; i<size; i++ {
		ch <- i
		time.Sleep(3*time.Second)
	}
	// 数据发送完毕，关闭通道
	close(ch)
}


func main() {
	ch := make(chan int, 10)
	// 开启一个goroutine，用于往通道ch里发送数据
	go addData(ch)

	/* range迭代从通道ch里获取数据
	通道close后，range迭代取完通道里的值后，循环会自动结束
	*/
	for i := range ch {
		fmt.Println(i)
	}
}
```

