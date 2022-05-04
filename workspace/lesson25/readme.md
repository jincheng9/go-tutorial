# sync.Map

## 定义

Map是sync包里的一个结构体类型，定义如下

```go
type Map struct {
    // some fields
}
```

Go语言里普通map的读写不是并发安全的，sync.Map的读写是并发安全的。

sync.Map可以理解为类似一个map[interface{}]interface{}的结构，key可以类型不一样，value也可以类型不一样，多个goroutine对其进行读写不需要额外加锁。

**Go官方设计sync.Map主要满足以下2个场景的用途**

1. **每个key只写一次，其它对该key的操作都是读操作**

2. **多个goroutine同时读写map，但是每个goroutine只读写各自的keys**

以上2种场景，相对于对普通的map加Mutex或者RWMutex来实现并发安全，使用sync.Map不用在业务代码里加锁，会大幅减少锁竞争，提升性能。**其它更为常见的场景还是使用普通的Map，搭配Mutex或者RWMutex来使用**。

不能对sync.Map使用值传递方式进行函数调用。

sync.Map结构体类型有如下几个方法：

* Delete，删除map里的key，即使key不存在，执行Delete操作也没任何影响

  ```go
  func (m *Map) Delete(key interface{})
  ```

* Load，从map里取出key对应的value。如果key存在map里，返回值value就是对应的值，ok就是true。如果key不在map里，返回值value就是nil，ok就是false。

  ```go
  func (m* Map) Load(key interface{}) (value interface{}, ok bool)
  ```

* LoadAndDelete，删除map里的key。如果key存在map里，返回值value就是对应的值，loaded就是true。如果key不在map里，返回值value就是nil，loaded就是false。

  ```go
  func (m* Map) LoadAndDelete(key interface{}) (value interface{}, loaded bool)
  ```

* LoadOrStore，从map里取出key对应的value。如果key在map里不存在，就把LoadOrStrore函数调用传入的参数<key, value>存储到map里，并返回参数里的value。如果key在map里，那loaded是true，如果key不在map里，那loaded是false。

  ```go
  func (m* Map) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)
  ```

* Range，遍历map里的所有<key, value>对，把每个<key, value>对，都作为参数传递给**f**去调用，如果遍历执行过程中，**f**返回false，那range迭代就结束了。

  ```go
  func (m* Map) Range(f func(key, value interface{}) bool)
  ```

* Store，往map里插入<key, vaue>对，即使key已经存在于map里，也没有任何影响

  ```go
  func (m* Map) Store(key, value interface{})
  ```

Delete, Load, LoadAndDelete, LoadOrStore, Store的均摊时间复杂度是O(1)，Range的时间复杂度是O(N)

## 使用

* 初始化

  ```go
  var m1 sync.Map
  m2 := sync.Map{}
  ```

  

* 示例1：统计字符串里每个字符出现的次数

  ```go
  package main
  
  import (
      "fmt"
      "sync"
  )
  
  func main() {
      /*统计字符串里每个字符出现的次数*/
      m := sync.Map{}
      str := "abcabcd"
      for _, value := range str {
          temp, ok := m.Load(value)
          //fmt.Println(temp, ok)
          if !ok {
              m.Store(value, 1)
          } else {
              /*temp是个interface变量，要转int才能和1做加法*/
              m.Store(value, temp.(int)+1)
          }
      }
      
      /*使用sync.Map里的Range遍历map*/
      m.Range(func(key, value interface{}) bool{
          fmt.Println(key, value)
          return true
      })
  }
  ```

* 示例2：多个goroutine并发写sync.Map，不加锁。如果是普通的map，这么来写就会出现运行时错误“fatal error: concurrent map writes”

  ```go
  package main
  
  import (
      "fmt"
      "sync"
  )
  
  var m sync.Map
  
  /*
  sync.Map里每个key只写一次，属于场景1
  */
  func changeMap(key int) {
      m.Store(key, 1)
  }
  
  func main() {
      var wg sync.WaitGroup
      size := 2
      wg.Add(size)
      
      for i:=0; i<size; i++ {
          i := i
          go func() {
              defer wg.Done()
              changeMap(i)
          }()
      }
      wg.Wait()
      
      /*使用sync.Map里的Range遍历map*/
      m.Range(func(key, value interface{}) bool{
          fmt.Println(key, value)
          return true
      })
  }
  ```

## 注意事项

* sync.Map不支持len和cap函数

* 在评估要不要使用sync.Map的时候，先考察业务场景是否符合上面描述的场景1和2，符合再考虑用sync.Map，不符合就用普通map+Mutex或者RWMutex。

## References

https://pkg.go.dev/sync@go1.17.2#Map

