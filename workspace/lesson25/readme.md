# sync.Map

## 定义

Map是sync包里的一个结构体类型，定义如下

```go
type Map struct {
    // some fields
}
```

Go语言里普通map的读写不是并发安全的，sync.Map的读写是并发安全的。sync.Map可以理解为类似一个map[interface{}]interface{}的结构，key可以类型不一样，value也可以类型不一样，多个goroutine对其进行读写不需要额外加锁。

## 使用

## References

https://pkg.go.dev/sync@go1.17.2#Map