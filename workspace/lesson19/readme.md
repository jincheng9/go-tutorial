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

* 缓冲区：channel默认没有缓冲区，可以在定义channel的时候指定缓冲区长度，也就是缓冲区最多可以存储的元素个数

  * 无缓冲区： channel无缓冲区的时候，往channel发送数据和从channel接收数据都会**阻塞**。往channel发送数据的时候，必须有其它goroutine从channel里接收了数据，发送操作才可以成功，发送操作所在的goroutine才能继续往下执行。从channel里接收数据也是同理，必须有其它goroutine往channel里发送了数据，接收操作才可以成功，接收操作所在的goroutine才能继续往下执行。

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

      Answer: 可能end和函数fetchChannel里的print内容都打印，**也可能只会打印end**。因为fetchChannel里的value := <-ch执行之后，main里的ch<-a就不再阻塞，继续往下执行了，所以可能main里最后的fmt.Println比fetchChannel里的fmt.Printf先执行，main执行完之后程序就结束了，所有goroutine自动结束，就不再执行fetchChannel里的fmt.Printf了。main里加上time.Sleep就可以允许fetchChannel这个goroutine有足够的时间执行完成。

  * 有缓冲区

* 遍历通道
* 单向通道：指定通道方向
* **注意**
  * 

## 并发同步和锁



## 原子操作



