# select语义

## 含义

select语义是和channel绑定在一起使用的，select可以实现从多个channel收发数据。

语法上和switch类似，有case分支和default分支，只不过select的每个case后面跟的是channel的收发操作。

在执行select语句的时候，如果当下那个时间点没有一个case满足条件，就会走default分支。

至多只能有一个 default分支。

如果没有default分支，select语句就会阻塞，直到某一个case满足条件。

如果select里任何case和default分支都没有，就会一直阻塞。

如果多个case同时满足，select会随机选一个case执行。

## 语法

```go
func d() {
	select {
	case ch1 <- data1: // send data to channel
		do sth
	case var_name = <-ch2 : // receive data from channel
		do sth
	case data, ok := <-ch3:
		do sth
	default:
		do sth
	}
}
```

语法上和[switch](../lesson6)的一些区别：

* select关键字和后面的{ 之间，不能有表达式或者语句。
* 没有fallthrough语句
* 每个case关键字后面跟的必须是channel的发送或者接收操作
* 允许多个case分支使用相同的channel，case分支后的语句甚至可以重复

## 示例

```go
func b() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	go func() {
		for i:=0; i<10; i++ {
			ch1 <- i
			ch2 <- i
		}
	}()
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch1:
			fmt.Printf("receive %d from channel 1\n", x)
		case y := <-ch2:
			fmt.Printf("receive %d from channel 2\n", y)
		}
	}
}
```

上面的示例里，select语句会for循环执行10次，每次select语句都会随机从channel1或者channel2里接收一个值并打印。可以把[示例代码](./select.go)下载到本地运行看结果。

## References

* https://golang.google.cn/ref/spec#Select_statements
* https://gobyexample.com/select
* https://go101.org/article/channel.html
* https://medium.com/a-journey-with-go/go-ordering-in-select-statements-fd0ff80fd8d6