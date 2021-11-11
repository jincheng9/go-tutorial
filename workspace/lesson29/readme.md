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



## 示例



## References

* https://golang.google.cn/ref/spec#Select_statements
* https://gobyexample.com/select