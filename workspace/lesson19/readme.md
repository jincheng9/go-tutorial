# 并发

## goroutine

* 定义：goroutine是轻量级的用户态线程，可以在代码里创建成千上万个goroutine来并发工作。如此多的goroutine是Go运行时来调度的。Go运行时会把goroutine的任务分配给CPU去执行。**注意**，goroutine不是我们通常理解的线程，线程是操作系统调度的。

* Go编程里不需要自己去写进程、线程和协程，想让某个任务并发执行，就把这个任务封装为一个函数，然后启动一个goroutine去执行这个函数就行了。

* 语法：go 函数名([参数列表])，示例代码如下：

  ```
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

* 定义

* 使用

  

## 并发同步和锁



## 原子操作



