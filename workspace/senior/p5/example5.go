// example5.go
package main

import (
	"fmt"
	"sync"
)

// hits 匿名结构体变量
// 这里同时用到了匿名结构体和匿名字段, sync.Mutex是匿名字段
// 因为匿名结构体嵌套了sync.Mutex，所以就有了sync.Mutex的Lock和Unlock方法
var hits struct {
	sync.Mutex
	n int
}

func main() {
	var wg sync.WaitGroup
	N := 100
	// 启动100个goroutine对匿名结构体的成员n同时做读写操作
	wg.Add(N)
	for i:=0; i<100; i++ {
		go func() {
			defer wg.Done()
			hits.Lock()
			defer hits.Unlock()
			hits.n++
		}()
	}
	wg.Wait()
	fmt.Println(hits.n) // 100
}