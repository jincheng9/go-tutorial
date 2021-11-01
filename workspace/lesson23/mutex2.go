package main

import (
	"fmt"
	"sync"
)

var sum int = 0
var mutex sync.Mutex
/*多个goroutine同时访问add
sum是多个goroutine共享的
通过加互斥锁来保证并发安全
*/
func add(i int) {
	mutex.Lock()
	defer mutex.Unlock()
	sum += i
}

func main() {
	var wg sync.WaitGroup
	size := 100
	wg.Add(size)
	for i:=1; i<=size; i++ {
		i := i
		go func() {
			defer wg.Done()
			add(i)
		}()
	}
	wg.Wait()
	fmt.Printf("sum of 1 to %d is: %d\n", size, sum)
}