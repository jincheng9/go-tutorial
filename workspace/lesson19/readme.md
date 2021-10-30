# 并发

## goroutine

* 定义：goroutine是轻量级的用户态线程，可以在代码里创建成千上万个goroutine来并发工作。如此多的goroutine是Go运行时来调度的。Go运行时会把goroutine的任务分配给CPU去执行。**注意**，goroutine不是我们通常理解的线程，线程是操作系统调度的。

* Go编程里不需要自己去写进程、线程和协程，想让某个任务并发执行，就把这个任务封装为一个函数，然后启动一个goroutine去执行这个函数就行了。

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

* 多个goroutine之间可以通过channel来通信

## channel

* 定义：channel是一种类型，默认值是nil。channel是引用类型，如果作为函数参数，是传引用。

  多个goroutine之间，可以通过channel来通信，一个goroutine可以发送数据到指定channel，其它goroutine可以从这个channel里接收数据。

  channel就像队列，满足FIFO原则，定义channel的时候必须指定channel要传递的元素类型。

* 语法：

  ```go
  /*channel_name是变量名，data_type是通道里的数据类型
  channel_size是channel通道缓冲区最多可以存放的元素个数，这个参数是可选的，不给就表示没有缓冲区
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

* 使用：channel有3种操作，发送数据，接收数据和关闭channel。发送和接收都是用**<-**符号

  * 发送值到通道：channel <- value

    ```go
    ch := make(chan int)
    ch <- 10 // 把10发送到ch里
    ```

  * 从通道接收值：value <- channel

    ```go
    ch := make(chan int)
    x := <-ch // 从通道ch里接收值，并复制给变量x
    <-ch // 从通道里接收值，不做其它处理
    ```

  * 关闭通道

    ```go
    ch := make(chan int)
    close(ch) // 关闭通道
    ```

* 缓冲区：channel默认没有缓冲区，可以在定义channel的时候指定缓冲区长度，也就是缓冲区可以存储的元素个数

* 遍历通道
* 单向通道：指定通道方向
* **注意**
  * 

## 并发同步和锁



## 原子操作



