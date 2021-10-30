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