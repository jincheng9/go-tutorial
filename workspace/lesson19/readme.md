# goroutine和channel

## goroutine

### 概念和语法

* 定义：goroutine是轻量级的用户态线程，可以在代码里创建成千上万个goroutine来并发工作。如此多的goroutine是Go运行时来调度的。Go运行时会把goroutine的任务分配给CPU去执行。**注意**，goroutine不是我们通常理解的线程，线程是操作系统调度的。

* Go编程里不需要自己在代码里写线程和协程，想让某个任务并发执行，就把这个任务封装为一个函数，然后启动一个goroutine去执行这个函数就行了。

* 语法：go 函数名([参数列表])，示例代码如下：

  ```go
  package main
  
  import "fmt"
  
  func hello() {
      fmt.Println("hello")
  }
  
  func main() {
    	/*开启一个goroutine去执行hello函数*/
  		go hello()
  		fmt.Println("main end")
  }
  ```

  

* Go会为main()函数创建一个默认的goroutine，如果main()函数结束了，那所有在main()中启动的goroutine都会立马结束。比如下面的例子：

  ```go
  package main
  
  import "fmt"
  
  func hello() {
      fmt.Println("hello")
  }
  
  func main() {
  	/*开启一个goroutine去执行hello函数*/
      go hello()
      fmt.Println("main end")
  }
  ```

  执行结果可能有以下3种：

  * main end // 只打印main end

  * main end // 先打印main end，再打印hello

    hello

  * hello // 先打印hello，再打印main end

    main end

  这是因为main函数的goroutine和hello这个goroutine是并发执行的，有可能main执行完了，hello还没执行，这个时候只打印main end。有可能hello先执行完，main后执行完，也可能反过来。所以共有3种情况。

### goroutine注意事项

**goroutine和闭包closure一起使用时要注意**，避免多个goroutine闭包使用同一个变量，否则goroutine执行的时候，这个变量的值可能已经被改了，和原来预期不符。比如下面例子：

```go
package main

import (
    "fmt"
    "sync"
    "time"
)


func worker(id int) {
    fmt.Printf("worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup
    /* wg跟踪10个goroutine */
    size := 10
    wg.Add(size)
    /* 开启10个goroutine并发执行 */
    for i:=0; i<size; i++ {
        go func() {
            defer wg.Done()
            worker(i)
        }()
    }
    /* Wait会一直阻塞，直到wg的计数器为0*/
    wg.Wait()
    fmt.Println("end")
}
```

在for循环里，用到了goroutine和闭包，每个闭包共享变量`i`，在闭包真正执行的时候，闭包里面用到的变量**i**的值可能已经被改了，所以闭包里调用worker的时候的传参i就不是想象中的从0到9。

有2种方法规避

* 方法1，把变量作为闭包的参数传给闭包

  ```go
  package main
  
  import (
      "fmt"
      "sync"
      "time"
  )
  
  
  func worker(id int) {
      fmt.Printf("worker %d starting\n", id)
      time.Sleep(time.Second)
      fmt.Printf("worker %d done\n", id)
  }
  
  func main() {
      var wg sync.WaitGroup
      /* wg跟踪10个goroutine */
      size := 10
      wg.Add(size)
      /* 开启10个goroutine并发执行 */
      for i:=0; i<size; i++ {
          go func(id int) {
              defer wg.Done()
              worker(id)
          }(i)
      }
      /* Wait会一直阻塞，直到wg的计数器为0*/
      wg.Wait()
      fmt.Println("end")
  }
  ```

  

* 方法2，在启动goroutine执行闭包前，定义一个新的变量**i**，这样每个闭包就可以用各自预期的变量值了

  ```go
  package main
  
  import (
      "fmt"
      "sync"
      "time"
  )
  
  
  func worker(id int) {
      fmt.Printf("worker %d starting\n", id)
      time.Sleep(time.Second)
      fmt.Printf("worker %d done\n", id)
  }
  
  func main() {
      var wg sync.WaitGroup
      /* wg跟踪10个goroutine */
      size := 10
      wg.Add(size)
      /* 开启10个goroutine并发执行 */
      for i:=0; i<size; i++ {
          /*定义一个新的变量*/
          i := i
          go func() {
              defer wg.Done()
              worker(i)
          }()
      }
      /* Wait会一直阻塞，直到wg的计数器为0*/
      wg.Wait()
      fmt.Println("end")
  }
  ```

* 多个goroutine之间可以通过channel来通信

## channel

### 概念和语法

* 定义：channel是一种类型，零值是nil。

  多个goroutine之间，可以通过channel来通信，一个goroutine可以发送数据到指定channel，其它goroutine可以从这个channel里接收数据。

  channel就像队列，满足FIFO原则，定义channel的时候必须指定channel要传递的元素类型。

* 语法：

  **未初始化的channel变量的值是nil，为nil的channel不能用于通信**。nil channel收发消息都会阻塞，可能引起死锁。

  ```go
  /*channel_name是变量名，data_type是通道里的数据类型
  channel_size是channel通道缓冲区的容量，表示最多可以存放的元素个数，这个参数是可选的，不给就表示没有缓冲区，通过cap()函数可以获取channel的容量
  */
  var channel_name chan data_type = make(chan data_type, [channel_size])
  ```

  ```go
  var ch1 chan int 
  var ch2 chan string
  var ch3 chan []int
  var ch4 chan struct_type // 可以往通道传递结构体变量
  
  ch5 := make(chan int)
  ch6 := make(chan string, 100)
  ch7 := make(chan []int)
  ch8 := make(chan struct_type)
  ```

### channel三种操作

channel有3种操作，发送数据，接收数据和关闭channel。发送和接收都是用`<-`符号

* 发送值到通道：channel <- value

  ```go
  ch := make(chan int)
  ch <- 10 // 把10发送到ch里
  ```

* 从通道接收值：value <- channel

  ```go
  ch := make(chan int)
  
  x := <-ch // 从通道ch里接收值，并赋值给变量x
  <-ch // 从通道里接收值，不做其它处理
  
  var y int
  y = <-ch // 从通道ch里接收值，并赋值给变量y
  ```

* 关闭通道: close(channel)，关闭nil channel会触发`panic: close of nil channel `

  ```go
  ch := make(chan int)
  close(ch) // 关闭通道
  ```

### channel缓冲区

channel默认没有缓冲区，可以在定义channel的时候指定缓冲区容量，也就是缓冲区最多可以存储的元素个数，通过内置函数**cap**可以获取到channel的容量。

#### 无缓冲区情况

channel无缓冲区的时候，往channel发送数据和从channel接收数据都会**阻塞**。

往channel发送数据的时候，必须有其它goroutine从channel里接收了数据，发送操作才可以成功，发送操作所在的goroutine才能继续往下执行。从channel里接收数据也是同理，必须有其它goroutine往channel里发送了数据，接收操作才可以成功，接收操作所在的goroutine才能继续往下执行。

```go
package main

import "fmt"
import "time"

type Cat struct {
	name string
	age int
}

func fetchChannel(ch chan Cat) {
	value := <- ch
	fmt.Printf("type: %T, value: %v\n", value, value)
}


func main() {
	ch := make(chan Cat)
	a := Cat{"yingduan", 1}
	// 启动一个goroutine，用于从ch这个通道里获取数据
	go fetchChannel(ch)
	// 往cha这个通道里发送数据
	ch <- a
	// main这个goroutine在这里等待2秒
	time.Sleep(2*time.Second)
	fmt.Println("end")
}
```

对于上面的例子，有2个点可以**思考**下

* 如果go fetchChannel(ch)和下面的 ch<-a这2行交换顺序会怎么样？

  Answer: 如果交换了顺序，main函数就会堵塞在ch<-a这一行，因为这个发送是阻塞的，不会往下执行，这个时候没有任何goroutine会从channel接收数据，错误信息如下：

  ```go
  fatal error: all goroutines are asleep - deadlock!
  ```

* 如果没有time.Sleep(2*time.Second)这一行，那程序运行结果会是怎么样？

  Answer: 可能main函数里的end和函数fetchChannel里的print内容都打印，**也可能只会打印main函数里的end**。因为fetchChannel里的value := <-ch执行之后，main里的ch<-a就不再阻塞，继续往下执行了，所以可能main里最后的fmt.Println比fetchChannel里的fmt.Printf先执行，main执行完之后程序就结束了，所有goroutine自动结束，就不再执行fetchChannel里的fmt.Printf了。main里加上time.Sleep就可以允许fetchChannel这个goroutine有足够的时间执行完成。

#### 有缓冲区情况

可以在初始化channel的时候通过make指定channel的缓冲区容量。

```go
ch := make(chan int, 100) // 定义了一个可以缓冲区容量为100的channel
```

对于有缓冲区的channel，对发送方而言：

* 如果缓冲区未满，那发送方发送数据到channel缓冲区后，就可以继续往下执行，不用阻塞等待接收方从channel里接收数据。
* 如果缓冲区已满，那发送方发送数据到channel会阻塞，直到接收方从channel里接收了数据，这样缓冲区才有空间存储发送方发送的数据，发送方所在goroutine才能继续往下执行。

对于接收方而言，在有值可以从channel接收之前，会一直阻塞。

```go
package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	// 下面2个发送操作不用阻塞等待接收方接收数据
	ch <- 10
	ch <- 20
	/*
	如果添加下面这行代码，就会一直阻塞，因为缓冲区已满，运行会报错
	fatal error: all goroutines are asleep - deadlock!
	
	ch <- 30
	*/
	
	fmt.Println(<-ch) // 10
	fmt.Println(<-ch) // 20
}
```

### 遍历通道channel

* range迭代从channel里不断取数据

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

  对于上面的例子，有个点可以思考下：

  * 如果删掉close(ch)这一行代码，结果会怎么样？

    Answer: 如果通道没有close，采用range从channel里循环取值，当channel里的值取完后，range会阻塞，如果没有继续往channel里发送值，go运行时会报错

    ```go
    fatal error: all goroutines are asleep - deadlock!
    ```

    

* for死循环不断获取channel里的数据，如果channel的值取完后，继续从channel里获取，会存在2种情况

  * 如果channel已经被close了，继续从channel里获取值会拿到对应channel里数据类型的零值
  * 如果channel没有被close，也不再继续往channel里发送数据，接收方会阻塞报错

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
  
  	/* 
  	for循环取完channel里的值后，因为通道close了，再次获取会拿到对应数据类型的零值
  	如果通道不close，for循环取完数据后就会阻塞报错
  	*/
  	for {
  		value, ok := <-ch
  		if ok {
  			fmt.Println(value)
  		} else {
  			fmt.Println("finish")
  			break
  		}
  	}
  }
  ```


### 单向通道

如果channel作为函数的形参，可以控制限制数据和channel之间的数据流向，控制只能往channel发送数据或者只能从channel接收数据。

不做限制的时候，channel是双向的，既可以往channel写数据，也可以从channel读数据。

* 语法

  ```go
  chan <- int // 只写，只能往channel写数据，不能从channel读数据
  <- chan int // 只读，只能从channel读数据，不能往channel写数据
  ```

* 实例

  ```go
  package main
  
  import "fmt"
  import "time"
  
  
  func write(ch chan<-int) {
  	/*
  	参数ch是只写channel，不能从channel读数据，否则编译报错
  	receive from send-only type chan<- int
  	*/
  	ch <- 10
  }
  
  
  func read(ch <-chan int) {
  	/*
  	参数ch是只读channel，不能往channel里写数据，否则编译报错
  	send to receive-only type <-chan int
  	*/
  	fmt.Println(<-ch)
  }
  
  func main() {
  	ch := make(chan int)
  	go write(ch)
  	go read(ch)
  
  	// 等待3秒，保证write和read这2个goroutine都可以执行完成
  	time.Sleep(3*time.Second)
  }
  ```


### channel注意事项

* channel被close后，如果再往channel里发送数据，会引发panic
* channel被close后，如果再次close，也会引发panic
* channel被close后，如果channel还有值，接收方可以一直从channel里获取值，直到channel里的值都已经取完。
* channel被close后，如果channel里没有值了，接收方继续从channel里取值，会得到channel里存的数据类型对应的默认零值，如果一直取值，就一直拿到零值。
* [从Go面试题看channel注意事项](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p9)


## References

* https://go.dev/doc/faq#closures_and_goroutines
* https://www.liwenzhou.com/posts/Go/14_concurrence/
* https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p9







