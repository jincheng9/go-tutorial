## defer语义

## defer含义

* defer是延迟的意思，在Go里可以放在某个函数或者方法调用的前面，让该函数或方法延迟执行

* 语法：

  ```go
  defer function([parameter_list]) // 延迟执行函数
  defer method([parameter_list]) // 延迟执行方法
  ```

  defer本身是在某个函数体内执行，比如在函数A内调用了defer func_name()，只要defer func_name()这行代码被执行到了，那func_name这个函数就会**被延迟到函数A return或者panic之前执行**。

  **注意**：如果是函数是因为调用了`os.Exit()`而退出，那`defer`就不会被执行了。参见[Go语言里被defer的函数一定会执行么？](https://github.com/jincheng9/go-tutorial/tree/main/workspace/senior/p2) 

  ```go
  defer func_name([parameter_list])
  defer package_name.func_name([parameter_list]) // 例如defer fmt.Println("blabla")
  ```

* 如果在函数内调用了**多次defer**，那在函数return之前，defer的函数调用满足LIFO原则，先defer的函数后执行，后defer的函数先执行。比如在函数A内先后执行了defer f1(), defer f2(), defer f3()，那函数A return之前，会按照f3(), f2(), f1()的顺序执行，再return。

## defer的用途？

Answer：defer常用于成对的操作，比如文件打开后要关闭、锁的申请和释放、sync.WaitGroup跟踪的goroutine的计数器的释放等。为了确保资源被释放，可以结合defer一起使用，避免在代码的各种条件分支里去释放资源，容易遗漏和出错。

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



## 注意事项

1. defer后面跟的必须是函数或者方法调用，defer后面的表达式不能加括号。

   ```go
   defer (fmt.Println(1)) // 编译报错，因为defer后面跟的表达式不能加括号
   ```

2. 被defer的函数的参数在执行到defer语句的时候就被确定下来了。

   ```go
   func a() {
       i := 0
       defer fmt.Println(i) // 最终打印0
       i++
       return
   }
   ```

   上例中，被defer的函数fmt.Println的参数**i**在执行到defer这一行的时候，**i**的值是0，fmt.Println的参数就被确定下来是0了，因此最终打印的结果是0，而不是1。

3. 被defer的函数执行顺序满足LIFO原则，后defer的先执行。

   ```go
   func b() {
       for i := 0; i < 4; i++ {
           defer fmt.Print(i)
       }
   }
   ```

   上例中，输出的结果是3210，后defer的先执行。

4. 被defer的函数可以对defer语句所在的函数的命名返回值做读取和修改操作。

   ```go
   // f returns 42
   func f() (result int) {
   	defer func() {
   		// result is accessed after it was set to 6 by the return statement
   		result *= 7
   	}()
   	return 6
   }
   ```

   上例中，被defer的函数func对defer语句所在的函数**f**的命名返回值result做了修改操作。

   调用函数**f**，返回的结果是42。

   执行顺序是函数**f**先把要返回的值6赋值给result，然后执行被defer的函数func，result被修改为42，然后函数**f**返回result，也就是返回了42。

   官方说明如下：

   ```go
   Each time a "defer" statement executes, the function value and parameters to
   the call are evaluated as usual and saved anew but the actual function is not 
   invoked. Instead, deferred functions are invoked immediately before the 
   surrounding function returns, in the reverse order they were deferred. That
   is, if the surrounding function returns through an explicit return statement, 
   deferred functions are executed after any result parameters are set by that 
   return statement but before the function returns to its caller. If a deferred
   function value evaluates to nil, execution panics when the function is 
   invoked, not when the "defer" statement is executed.
   ```



## References

* https://go.dev/blog/defer-panic-and-recover
* https://golang.google.cn/ref/spec#Defer_statements

* https://stackoverflow.com/questions/52718143/is-golang-defer-statement-execute-before-or-after-return-statement
