package main

import "fmt"
import "time"


func addData(ch chan int) {
	/*
	每3秒往通道ch里发送一次数据
	*/
	size := cap(ch)
	for i:=0; i<size; i++ {
		ch <- i
		time.Sleep(3*time.Second)
	}
	// 数据发送完毕，关闭通道
	close(ch)
}


func main() {
	ch := make(chan int, 10)
	// 开启一个goroutine，用于往通道ch里发送数据
	go addData(ch)

	/* 
	for循环取完channel里的值后，因为通道close了，再次获取会拿到对应数据类型的零值
	如果通道不close，for循环取完数据后就会阻塞报错
	*/
	for {
		value, ok := <-ch
		if ok {
			fmt.Println(value)
		} else {
			fmt.Println("finish")
			break
		}
	}
}