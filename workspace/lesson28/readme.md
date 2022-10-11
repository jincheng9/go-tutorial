# panic和recover

panic和recover是Go的2个内置函数，用于程序运行期抛出异常(panic的功能)和处理异常(recover的功能)。

## panic

引发panic可以是程序显式调用panic函数，也可以是运行期错误，比如数组越界访问，除0等。

* 定义

  ```go
  func panic(interface{}) // 入参是一个空interface，无返回值
  ```

* 示例

  ```go
  panic(12)
  panic("invalid parameter")
  panic(Error("cannot parse"))
  ```

* 如果在函数F里，显式调用了panic或者函数F执行过程中出现运行期错误，那F的执行会终止，接下来会有以下行为依次产生：

  * F里被defer的函数会执行。
  * F的上一级函数，也就是调用F的函数，假设是函数E。对函数E而言，对F的调用就等同于调用了panic，函数E里被defer的函数被执行
  * 如果函数E还有上一级函数，就继续往上，每一级函数里被defer的函数都被执行，直到没有上一级函数。
  * 经过了以上步骤，panic的错误就会被抛出来，整个程序结束。

* 既然可以用panic来抛运行期异常，那就有相应办法可以捕获异常，让程序正常往下执行。Go通过结合内置函数recover和defer语义，来实现捕获运行期异常。

* **对于标准Go编译器，有些致命错误是无法被recover捕捉的，比如栈溢出(stack overflow)或者内存超限(out of memory)，遇到这种情况程序就会crash。**

## recover

recover是Go的内置函数，可以捕获panic异常。recover必须结合defer一起使用才能生效。

程序正常执行过程中，没有panic产生，这时调用recover函数会返回nil，除此之外，没有其它任何效果。

如果当前goroutine触发了panic，可以在代码的适当位置调用recover函数捕获异常，让程序继续正常执行，而不是异常终止。

* 定义

  ```go
  func recover() interface{} // 返回值是一个空interface，无入参
  ```

* 示例

  ```go
  package main
  
  import (
  	"fmt"
  )
  
  func a() {
  	defer func() {
  		/*捕获函数a内部的panic*/
  		r := recover()
  		fmt.Println("panic recover", r)
  	}()
  	panic(1)
  }
  
  func main() {
  	defer func() {
  		/*因为函数a的panic已经被函数a内部的recover捕获了
  		所以main里的recover捕获不到异常，r的值是nil*/
  		r := recover()
  		fmt.Println("main recover", r)
  	}()
  	a()
  	fmt.Println("main")
  }
  
  ```

  上例的执行结果是：

  ```go
  panic recover 1
  main
  main recover <nil>
  ```

  

* recover在以下几种情况返回nil

  * panic的参数是nil。这种情况recover捕获后，拿到的返回值也是nil。
  * goroutine没有panic产生。没有panic，那当然recover拿到的也就是nil了。
  * recover不是**在被defer的函数里面**被**直接调用**执行。

* 一个更复杂的示例

  ```go
  package main
  
  import "fmt"
  
  func main() {
      f()
      fmt.Println("Returned normally from f.")
  }
  
  func f() {
      defer func() {
          if r := recover(); r != nil {
              fmt.Println("Recovered in f", r)
          }
      }()
      fmt.Println("Calling g.")
      g(0)
      fmt.Println("Returned normally from g.")
  }
  
  func g(i int) {
      if i > 3 {
          fmt.Println("Panicking!")
          panic(fmt.Sprintf("%v", i))
      }
      defer fmt.Println("Defer in g", i)
      fmt.Println("Printing in g", i)
      g(i + 1)
  }
  ```

  大家可以下载[recover2.go](./recover2.go)代码，本地运行看看结果是否和预期相符。



## References

* https://go.dev/ref/spec#Run_time_panics
* https://go.dev/blog/defer-panic-and-recover
* https://go.dev/ref/spec#Handling_panics
* https://chai2010.gitbooks.io/advanced-go-programming-book/content/appendix/appendix-a-trap.html
* https://go101.org/article/control-flows-more.html
* https://go101.org/article/panic-and-recover-more.html
* https://draveness.me/golang/docs/part2-foundation/ch05-keyword/golang-panic-recover/
