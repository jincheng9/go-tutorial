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

Go官方设计sync.Map主要满足以下2个场景的用途

1. 每个key只写一次，其它对该key的操作都是读操作

2. 多个goroutinet同时读写map，但是每个goroutine只读写各自的keys

以上2种场景，相对于对普通的map加Mutex或者RWMutex来实现并发安全，使用sync.Map会大幅减少锁竞争，提升性能。**笔者认为对于读多写少(包含场景1)以及场景2的情况可以使用sync.Map，其它更为常见的场景还是使用普通的Map，搭配Mutex或者RWMutex来使用**。

不能对sync.Map使用值传递方式进行函数调用。

sync.Map结构体类型有如下几个方法：

* Delete
* Load
* LoadAndDelete
* LoadOrStore
* Range
* Store

## 使用

* 初始化
* 示例

## References

https://pkg.go.dev/sync@go1.17.2#Map

