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