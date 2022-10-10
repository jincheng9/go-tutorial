# 一文读懂Go匿名结构体使用场景

## 前言

匿名行为在go语言里非常常见，比如：

* 匿名函数：也就是我们熟知的闭包(Closure)
* 结构体里的匿名字段(Anonymous Fields)
* 匿名结构体(Anonymous Structs)

匿名行为的设计带来了一些理解上的困难，但是熟悉了匿名设计的使用后，你会发现匿名设计在某些特定场景可以帮助大家写出更简洁、更优雅、更高效和更安全的代码。

## 什么是匿名结构体

匿名结构体：顾名思义，就是结构体没有命名。比如下面的代码示例：

```go
// example1.go
package main

import (
	"fmt"
)

func main() {
	a := struct{name string; age int}{"bob", 10}
	b := struct{
		school string
		city string
	}{"THU", "Beijing"}
	fmt.Println(a, b)
}
```

在这个例子里，我们定义了2个变量a和b，它们都是匿名结构体变量。

## 常见的使用场景

### 全局变量组合

有时候我们会在程序里定义若干全局变量，有些全局变量的含义是互相关联的，这个时候我们可以使用匿名结构体把这些关联的全局变量组合在一起。

```go
// example2.go
package main

import "fmt"

// DBConfig 声明全局匿名结构体变量
var DBConfig struct {
	user string
	pwd string
	host string
	port int
	db string
}

// SysConfig 全局匿名结构体变量也可以在声明的时候直接初始化赋值
var SysConfig = struct{
	sysName string
	mode string
}{"tutorial", "debug"}

func main() {
	// 给匿名结构体变量DBConfig赋值
	DBConfig.user = "root"
	DBConfig.pwd = "root"
	DBConfig.host = "127.0.0.1"
	DBConfig.port = 3306
	DBConfig.db = "test_db"
	fmt.Println(DBConfig)
}
```

对全局匿名结构体变量完成赋值后，后续代码都可以使用这个匿名结构体变量。

注意：如果你的程序对于某个全局的结构体要创建多个变量，就不能用匿名结构体了。



### 局部变量组合

全局变量可以组合，局部变量当然也可以组合了。

如果在局部作用域(比如函数或者方法体内)里，某些变量的含义互相关联，就可以组合到一个结构体里。

同时这个结构体只是临时一次性使用，不需要创建这个结构体的多个变量，那就可以使用匿名结构体。

```go
// example3.go
package main

import "fmt"

func main() {
	// a和b作为局部匿名结构体变量，只是临时一次性使用
	// 注意：a是把struct里的字段声明放在同一行，字段之间要用分号分割，否则编译报错
	a := struct{name string; age int}{"Alice", 16}
	fmt.Println(a)

	b := struct{
		school string
		city string
	}{"THU", "Beijing"}
	fmt.Println(b)
}
```



### 构建测试数据

匿名结构体可以和切片结合起来使用，通常用于创建一组测试数据。

```go
// example4.go
package main

import "fmt"

// 测试数据
var indexRuneTests = []struct {
	s    string
	rune rune
	out  int
}{
	{"a A x", 'A', 2},
	{"some_text=some_value", '=', 9},
	{"☺a", 'a', 3},
	{"a☻☺b", '☺', 4},
}

func main() {
	fmt.Println(indexRuneTests)
}
```



### 嵌套锁(embed lock)

我们经常遇到多个goroutine要操作共享变量，为了并发安全，需要对共享变量的读写加锁。

这个时候通常需要定义一个和共享变量配套的锁来保护共享变量。

匿名结构体和匿名字段相结合，可以写出更优雅的代码来保护匿名结构体里的共享变量，实现并发安全。

```go
// example5.go
package main

import (
	"fmt"
	"sync"
)

// hits 匿名结构体变量
// 这里同时用到了匿名结构体和匿名字段, sync.Mutex是匿名字段
// 因为匿名结构体嵌套了sync.Mutex，所以就有了sync.Mutex的Lock和Unlock方法
var hits struct {
	sync.Mutex
	n int
}

func main() {
	var wg sync.WaitGroup
	N := 100
	// 启动100个goroutine对匿名结构体的成员n同时做读写操作
	wg.Add(N)
	for i:=0; i<100; i++ {
		go func() {
			defer wg.Done()
			hits.Lock()
			defer hits.Unlock()
			hits.n++
		}()
	}
	wg.Wait()
	fmt.Println(hits.n) // 100
}
```



### HTTP处理函数中JSON序列化和反序列化

我们在处理http请求时，通常会和JSON数据打交道。

比如post请求的content-type使用application/json时，服务器接收过来的json数据是key:value格式，不同key的value的类型可以不一样，可能是数字、字符串、数组等，因此会遇到使用`json.Unmarshal`和`map[string]interface{}`来接收JSON反序列化后的数据。

但是使用map[string]interface{}有几个问题：

* 没有类型检查：比如json的某个value本来预期是string类型，但是请求传过来的是bool类型，使用json.Unmarshal解析到map[string]interface{}是不会报错的，因为空interface可以接受任何类型数据。
* map是模糊的：Unmarshal后得到了map，我们还得判断这个key在map里是否存在。否则拿不存在的key的value，得到的可能是给nil值，如果不做检查，直接对nil指针做*操作，会引发panic。
* 代码比较冗长：得先判断key是否存在，如果存在，要显式转换成对应的数据类型，并且还得判断转换是否成功。代码会比较冗长。

这个时候我们就可以使用匿名结构体来接收反序列化后的数据，代码会更简洁。参见如下代码示例：

```go
// example6.go
// 请求命令：curl -X POST -H "content-type: application/json" http://localhost:4000/user -d '{"name":"John", "age":111}'
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	// data 匿名结构体变量，用来接收http请求发送过来的json数据
	data := struct{
		Name string `json:"name"`
		Age int	`json:"age"`
	}{}
	// 把json数据反序列化到data变量里
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(data)
	fmt.Fprint(w, "Hello!")
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "index")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/user", AddUser)
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
```



## 总结

匿名结构体可以让我们不用先定义结构体类型，再定义结构体变量。让结构体的定义和变量的定义可以结合在一起，一次性完成。

匿名结构体有以下应用场景：

* 组合变量：

  * 全局匿名结构体：把有关联的全局变量组合在一起
  * 局部匿名结构体：临时一次性使用

* 构建测试数据：

  * 匿名结构体+切片，构造一组测试数据

* 嵌套锁：把多个goroutine的共享访问数据和保护共享数据的锁组合在一个匿名结构体内，代码更优雅。

* HTTP处理函数中Json序列化和反序列化：

  * 匿名结构体变量用来接收http请求数据
  * 和map[string]interface{}相比，代码更简洁，更安全

  

## 开源地址

文档和代码开源地址：https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p5

也欢迎大家关注公众号：coding进阶，学习更多Go、微服务和云原生架构相关知识。



## References

* https://go.dev/talks/2012/10things.slide#1
* https://qvault.io/golang/anonymous-structs-golang/
* https://gist.github.com/aodin/9493190