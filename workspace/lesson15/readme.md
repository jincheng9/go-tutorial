# map集合

## 定义

Go语言里的map底层是通过**hash**实现的，是一种**无序**的基于<key, value>对组成的数据结构，key是唯一的，类似python的dict。

map必须初始化后才能写map。

如果只是声明map，但没有初始化，只能读，不能写。参考下面的例子的说明：

```go
package main

import "fmt"

func main() {
	var counter map[string]int
	/*
	map没有初始化，读map相当于读了一个空map
	下例中：value是int的零值0，ok是false
	*/
	value, ok := counter["a"]
	fmt.Println(value, ok)

	/*counter没有初始化，给counter赋值会在运行时报错
	  panic: assignment to entry in nil map
	*/
	counter["a"] = 1
	fmt.Println(counter)
}
```



* 语法

  ```go
  var map_var map[key_data_type]value_data_type = map[key_data_type]value_data_type{}
  
  var map_var = map[key_data_type]value_data_type{}
  
  map_var := map[key_data_type]value_data_type{}
  
  /*cap是map容量，超过后会自动扩容*/
  map_var := make(map[key_data_type]value_data_type, [cap]) 
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

    如果key存在，那is_exist就是true, value是对应的值。否则is_exist就是false, value是map的value数据类型的零值。

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

* len(map)：通过内置的len()函数可以获取map里<key, value>对的数量

  ```go
  counter := make(map[string]int)
  fmt.Println(len(counter))
  counter["a"] = 1
  fmt.Println(len(counter))
  ```

  

* map作为函数形参，可以在函数体内部改变外部实参的值，原理参见[Go有引用传递么？](../senior/p3)。示例如下：

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

  

## 注意事项

* key必须支持==和!=比较，才能用作map的key。

  因此切片slice，函数类型function，集合map，不能用作map的key

* map不是并发安全的，并发读写要加锁

  

