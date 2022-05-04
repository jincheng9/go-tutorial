// add.go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

// 多个goroutine并发读写sum，有并发冲突，最终计算得到的sum值是不准确的
func test1() {
	var sum int32 = 0
	N := 100
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func(i int32) {
			sum += i
			wg.Done()
		}(int32(i))
	}
	wg.Wait()
	fmt.Println("func test1, sum=", sum)
}

// 使用原子操作计算sum，没有并发冲突，最终计算得到sum的值是准确的
func test2() {
	var sum int32 = 0
	N := 100
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func(i int32) {
			atomic.AddInt32(&sum, i)
			wg.Done()
		}(int32(i))
	}
	wg.Wait()
	fmt.Println("func test2, sum=", sum)
}

func main() {
	test1()
	test2()
}
