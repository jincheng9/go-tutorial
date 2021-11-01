package main

import (
	"fmt"
	"sync"
)

var sum int = 0

/*多个goroutine同时访问add
sum是多个goroutine共享的
也就是多个goroutine同时对共享变量sum做写操作不是并发安全的
*/
func add(i int) {
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