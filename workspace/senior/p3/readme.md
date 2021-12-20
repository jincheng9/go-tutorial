# Go有引用变量和引用传递么？

# 前言

在Go中如果使用过map和channel，就会发现把map和channel作为函数参数传递，不需要在函数形参里对map和channel加指针标记*就可以在函数体内改变外部map和channel的值。

这会给人一种错觉：map和channel难道是类似C++的引用变量，函数传参的时候使用的是引用传递？

比如下面的例子：

```go
// example1.go
package main

import "fmt"

func changeMap(data map[string]interface{}) {
	data["c"] = 3
}

func main() {
	counter := map[string]interface{}{"a": 1, "b": 2}
	fmt.Println("begin:", counter)
	changeMap(counter)
	fmt.Println("after:", counter)
}
```

程序运行的结果是：

```go
begin: map[a:1 b:2]
after: map[a:1 b:2 c:3]
```

上面的例子里，函数changeMap改变了外部的map类型counter的值。

那map传参是使用的引用传递么？带着这个问题，我们先回顾下什么是引用变量和引用传递。



## 什么是引用变量(reference variable)和引用传递(pass-by-reference)

我们先回顾下C++里的引用变量和引用传递。看下面的例子：

```cpp
// example2.cpp
#include <iostream>

using namespace std;

/*函数changeValue使用引用传递*/
void changeValue(int &n) {
    n = 2;
}

int main() {
    int a = 1;
    /*
    b是引用变量，引用的是变量a
    */
    int &b = a;
    cout << "a=" << a << " address:" << &a << endl;
    cout << "b=" << b << " address:" << &b << endl;
    /*
    调用changeValue会改变外部实参a的值
    */
    changeValue(a);
    cout << "a=" << a << " address:" << &a << endl;
    cout << "b=" << b << " address:" << &b << endl;
}
```

程序的运行结果是：

```cpp
a=1 address:0x7ffee7aa776c
b=1 address:0x7ffee7aa776c
a=2 address:0x7ffee7aa776c
b=2 address:0x7ffee7aa776c
```

在这个例子里，变量b是引用变量，引用的是变量a。引用变量就好比是原变量的一个别名，引用变量和引用传递的特点如下：

* 引用变量和原变量的内存地址一样。就像上面的例子里引用变量b和原变量a的内存地址相同。
* 函数使用引用传递，可以改变外部实参的值。就像上面的例子里，changeValue函数使用了引用传递，改变了外部实参a的值。
* 对原变量的值的修改也会改变引用变量的值。就像上面的例子里，changeValue函数对a的修改，也改变了引用变量b的值。



## Go有引用变量(reference variable)和引用传递(pass-by-reference)么？

先给出结论：Go语言里没有引用变量和引用传递。

在Go语言里，不可能有2个变量有相同的内存地址，也就不存在引用变量了。

**注意**：这里说的是不可能2个变量有相同的内存地址，但是2个变量指向同一个内存地址是可以的，这2个是不一样的。参考下面的例子：

```go
// example3.go
package main

import "fmt"

func main() {
	a := 10
	var p1 *int = &a
	var p2 *int = &a
	fmt.Println("p1 value:", p1, " address:", &p1)
	fmt.Println("p2 value:", p2, " address:", &p2)
}
```

程序运行结果是：

```go
p1 value: 0xc0000ac008  address: 0xc0000ae018
p2 value: 0xc0000ac008  address: 0xc0000ae020
```

可以看出，变量p1和p2的值相同，都指向变量a的内存地址。但是变量p1和p2自己本身的内存地址是不一样的。而C++里的引用变量和原变量的内存地址是相同的。

因此，在Go语言里是不存在引用变量的，也就自然没有引用传递了。



## 有map不是使用引用传递的反例么

看下面的例子：

```go
// example4.go
package main

import "fmt"

func initMap(data map[string]int) {
	data = make(map[string]int)
	fmt.Println("in function initMap, data == nil:", data == nil)
}

func main() {
	var data map[string]int
	fmt.Println("before init, data == nil:", data == nil)
	initMap(data)
	fmt.Println("after init, data == nil:", data == nil)
}

```

大家可以先思考一会，想想程序运行结果是什么。







程序实际运行结果如下：

```go
before init, data == nil: true
in function initMap, data == nil: false
after init, data == nil: true
```

可以看出，函数initMap并没有改变外部实参data的值，因此也证明了map并不是引用变量。

那问题来了，为啥map作为函数参数不是使用的引用传递，但是在本文最开头举的例子里，却可以改变外部实参的值呢？



## map究竟是什么？

结论是：**map变量是指向runtime.hmap的指针**

当我们使用下面的代码初始化map的时候

```go
data := make(map[string]int)
```

Go编译器会把make调用转成对[runtime.makemap](https://golang.org/src/runtime/map.go#L298)的调用，我们来看看runtime.makemap的源代码实现。

```go
298  // makemap implements Go map creation for make(map[k]v, hint).
299  // If the compiler has determined that the map or the first bucket
300  // can be created on the stack, h and/or bucket may be non-nil.
301  // If h != nil, the map can be created directly in h.
302  // If h.buckets != nil, bucket pointed to can be used as the first bucket.
303  func makemap(t *maptype, hint int, h *hmap) *hmap {
304  	mem, overflow := math.MulUintptr(uintptr(hint), t.bucket.size)
305  	if overflow || mem > maxAlloc {
306  		hint = 0
307  	}
308  
309  	// initialize Hmap
310  	if h == nil {
311  		h = new(hmap)
312  	}
313  	h.hash0 = fastrand()
314  
315  	// Find the size parameter B which will hold the requested # of elements.
316  	// For hint < 0 overLoadFactor returns false since hint < bucketCnt.
317  	B := uint8(0)
318  	for overLoadFactor(hint, B) {
319  		B++
320  	}
321  	h.B = B
322  
323  	// allocate initial hash table
324  	// if B == 0, the buckets field is allocated lazily later (in mapassign)
325  	// If hint is large zeroing this memory could take a while.
326  	if h.B != 0 {
327  		var nextOverflow *bmap
328  		h.buckets, nextOverflow = makeBucketArray(t, h.B, nil)
329  		if nextOverflow != nil {
330  			h.extra = new(mapextra)
331  			h.extra.nextOverflow = nextOverflow
332  		}
333  	}
334  
335  	return h
336  }
```



从上面的源代码可以看出，runtime.makemap返回的是一个指向runtime.hmap结构的指针。

我们也可以通过下面的例子，来验证map变量到底是不是指针。

```go
// example5.go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	data := make(map[string]int)
	var p uintptr
	fmt.Println("data size:", unsafe.Sizeof(data))
	fmt.Println("pointer size:", unsafe.Sizeof(p))
}
```

程序运行结果是：

```go
data size: 8
pointer size: 8
```

map的size和指针的size一样，都是8个字节。



思考更为深入的读者，看到这里，可能还会有一个疑问：

既然map是指针，那为什么make()函数的说明里，有这么一句Unlike new, make's return type is the same as the type of its argument, not a pointer to it. 

> The make built-in function allocates and initializes an object of type slice, map, or chan (only). Like new, the first argument is a type, not a value. Unlike new, make's return type is the same as the type of its argument, not a pointer to it. The specification of the result depends on the type:

如果map是指针，那make返回的不应该是*map[string]int么，为啥官方文档里说的是not a pointer to it.

这里其实也有Go语言历史上的一个演变过程，看看Go作者之一Ian Taylor的说法：

> In the very early days what we call maps now
> were written as pointers, so you wrote *map[int]int. We moved away
> from that when we realized that no one ever wrote `map` without
> writing `*map`. That simplified many things but it left this issue
> behind as a complication.

所以，在Go语言早期，的确对于map是使用过指针形式的，但是最后Go设计者们发现，几乎没有人使用map不加指针，因此就直接去掉了形式上的指针符号*。



## 总结

map和channel，本质上都是指针，指向Go runtime结构。带着这个思路，我们再回顾下之前讲过的例子：

```go
// example4.go
package main

import "fmt"

func initMap(data map[string]int) {
	data = make(map[string]int)
	fmt.Println("in function initMap, data == nil:", data == nil)
}

func main() {
	var data map[string]int
	fmt.Println("before init, data == nil:", data == nil)
	initMap(data)
	fmt.Println("after init, data == nil:", data == nil)
}
```

既然map是一个指针，因此在函数initMap里，

```go
data = make(map[string]int)
```

这一句等于把data这个指针，进行了重新赋值，函数内部的data指针不再指向外部实参data对应的runtime.hmap结构体的内存地址。

因此在函数体内对data的修改，并没有影响外部实参data以及data对应的runtime.hmap结构体的值。

程序实际运行结果如下：

```go
before init, data == nil: true
in function initMap, data == nil: false
after init, data == nil: true
```



## 代码

相关代码和说明开源在GitHub：[Go有引用变量和引用传递么？](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p3)

也可以搜索公众号：coding进阶，查看更多Go知识。

![df](../../official-blog/qrcode_wechat.jpg) 

## References

* https://dave.cheney.net/2017/04/29/there-is-no-pass-by-reference-in-go
