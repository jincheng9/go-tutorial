package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	s := []int{3, 1, 5, 7}
	size := len(s)
	var wg sync.WaitGroup
	wg.Add(size)
	for _, value := range s {
		go func(n int) {
			defer wg.Done()
			time.Sleep(time.Duration(n) * time.Second)
			fmt.Println(n)
		}(value)
	}
	wg.Wait()
}
