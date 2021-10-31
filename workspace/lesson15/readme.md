# map集合

## 定义

* 语法

  ```go
  var map_var map[key_data_type]value_data_type = map[key_data_type]value_data_type{}
  
  var map_var = map[key_data_type]value_data_type{}
  
  map_var := map[key_data_type]value_data_type{}
  
  map_var := make(map[key_data_type]value_data_type)
  ```

* 示例

  ```go
  package main
  
  import "fmt"
  
  func main() {
      var dict map[string]int = map[string]int{}
      dict["a"] = 1
      fmt.Println(dict)
      
      var dict2 = map[string]int{}
      dict2["b"] = 2
      fmt.Println(dict2)
      
      dict3 := map[string]int{"test":0}
      dict3["c"] = 3
      fmt.Println(dict2)
      
      dict4 := make(map[string]int)
      dict4["d"] = 4
      fmt.Println(dict4)
  }
  ```

  

## 使用

* 判断key在map里是否存在

  * 语法。

    ```go
    value, is_exist := map[key]
    ```

    如果key存在，那is_exist就是true, value是对应的值。否则is_exist就是false, value是map的value数据类型的默认值。

    **注意**: 如果key不存在，通过map[key]访问不会给map自动插入这个新key。C++是会自动插入新key的，两个语言不一样。如果确定key存在，可以直接使用map[key]拿到value。

  * 示例

    ```go
    package main
    
    import "fmt"
    
    func main() {
        // 构造一个map
        str := "aba"
        dict := map[rune]int{}
        for _, value := range str{
            dict[value]++
        }
        fmt.Println(dict) // map[97:2 98:1]
        
        // 访问map里不存在的key，并不会像C++一样自动往map里插入这个新key
        value, ok := dict['z']
        fmt.Println(value, ok) // 0 false
        fmt.Println(dict) // map[97:2 98:1]
        
        // 访问map里已有的key
        value2 := dict['a']
        fmt.Println(value2) // 2
    }
    ```

    

* 遍历map：使用range迭代，参见[lesson14](../lesson14)

* map作为函数参数，是传引用。示例如下：

  ```go
  package main
  
  import "fmt"
  
  
  
  func buildMap(str string, m map[rune]int) {
  	/*函数内对map变量m的修改会影响main里的实参mapping*/
  	for _, value := range str {
  		m[value]++
  	}
  }
  
  
  func main() {
  	mapping := map[rune]int{}
  	str := "abc"
  	buildMap(str, mapping)
  
  	/*
  	mapping的值被buildMap修改了
  	*/
  	for key, value := range mapping {
  		fmt.Printf("key:%v, value:%d\n", key, value)
  	}
  }
  ```

  

## delete函数

* 删除key，参数为map和对应的key。允许删除一个不存在的key，对map无任何影响。

  ```go
  package main
  
  import "fmt"
  
  func main() {
      dict :=  map[string]int{"a":1, "b":2}
      fmt.Println(dict) // map[a:1 b:2]
      
      // 删除"a"这个key
      delete(dict, "a")
      fmt.Println(dict) // map[b:2]
      
      // 删除"c"这个不在的key，对map结果无影响
      delete(dict, "c")
      fmt.Println(dict) // map[b:2]
  }
  ```

  



