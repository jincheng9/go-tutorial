# 类型转换

* 语法

  ```go
  type_name(expression)
  ```

* 示例

  ```go
  package main
  
  import "fmt"
  
  func main() {
      total_weight := 100
      num := 12
      // total_weight和num都是整数，相除结果还是整数
      fmt.Println("average=", total_weight/num) //  average= 8
      
      // 转成float32再相除，结果就是准确值了
      fmt.Println("average=", float32(total_weight)/float32(num)) // average= 8.333333
      
      /* 注意，float32只能和float32做运算，否则会报错，比如下例里float32和int相加，编译报错:
      invalid operation: float32(total_weight) + num (mismatched types float32 and int)
     
      res := float32(total_weight) + num
      fmt.Println(res)
      */
  }
  ```

  

* **注意**：Go不支持隐式类型转换，要做数据类型转换必须按照type_name(expression)方式做显式的类型转换

  ```go
  package main
  
  import "fmt"
  
  
  func main() {
      num := 10
      var f float32 = float32(num)
      fmt.Println(f) // 10
      
      /*
      不支持隐式类型转换，比如下例想隐式讲num这个int类型转换为float32就会编译报错:
       cannot use num (type int) as type float32 in assignment
       
      var f float32 = num
      */
  }
  ```

  