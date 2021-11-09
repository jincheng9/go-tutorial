## defer关键字

* defer是延迟的意思，在Go里可以放在某个函数或者方法调用的前面，让该函数或方法延迟执行

* 语法：defer本身是在某个函数体内执行，比如在函数A内调用了defer func_name()，只要defer func_name()这行代码被执行到了，那func_name这个函数就会**被延迟到函数A return之前执行，并且一定会执行**。 

  ```go
  defer func_name([parameter_list])
  defer package_name.func_name([parameter_list]) // 例如defer fmt.Println("blabla")
  ```

* 如果在函数内调用了**多次defer**，那在函数return之前，defer的函数调用满足LIFO原则，先defer的函数后执行，后defer的函数先执行。比如在函数A内先后执行了defer f1(), defer f2(), defer f3()，那函数A return之前，会按照f3(), f2(), f1()的顺序执行，再return。

* defer的用途？

  Answer：defer常用于成对的操作，比如文件打开后要关闭，sync.WaitGroup跟踪的goroutine的申请和释放等。为了确保资源被释放，可以结合defer一起使用，避免在代码的各种条件分支里去释放资源，容易遗漏和出错。

* 示例1

  ```go
  package main
  
  import (
  	"fmt"
  	"sync"
  )
  
  
  var wg sync.WaitGroup
  
  
  func sumN(N int) {
  	// 调用defer wg.Done()确保sumN执行完之后，可以对wg的计数器减1
  	defer wg.Done()
  	sum := 0
  	for i:=1; i<=N; i++ {
  		sum += i
  	}
  	fmt.Printf("sum from 1 to %d is %d\n", N, sum)
  }
  
  func main() {
  	// 设置wg跟踪的计数器数量为1
  	wg.Add(1)
  	// 开启sumN这个goroutine去计算1到100的和
  	go sumN(100)
  	// Wait会一直等待，直到wg的计数器为0
  	wg.Wait()
  	
  	fmt.Println("finish")		
  }
  ```

* defer结合goroutine和闭包一起使用，可以让任务函数内部不用关心Go并发里的同步原语，更多内容可以参考[goroutine](./workspace/lesson19)和[sync.WaitGroup](./workspace/lesson21)

  ```go
  package main
  
  import (
      "fmt"
      "sync"
  )
  
  func worker(id int) {
      fmt.Println(id)
  }
  
  func main() {
      var wg sync.WaitGroup
      size := 10
      wg.Add(size)
      
      for i:=0; i<size; i++ {
          i := i 
          /*把worker的调用和defer放在一个闭包里
          这样worker函数内部就不用使用WaitGroup了
          */
          go func() {
              defer wg.Done()
              worker(i)
          }()
      }
      
      wg.Wait()
  }
  ```




## References

* https://go.dev/blog/defer-panic-and-recover
