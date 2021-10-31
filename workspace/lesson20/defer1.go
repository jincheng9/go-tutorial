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