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