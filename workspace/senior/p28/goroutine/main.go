package main

import (
	"fmt"
	"sync"
)

func test() {
	var wg sync.WaitGroup
	wg.Add(3)
	ints := []int{1, 2, 3}
	for _, i := range ints {
		go func() {
			defer wg.Done()
			fmt.Printf("%v\n", i)
		}()
		// time.Sleep(time.Second * 2)
	}
	wg.Wait()
}

func main() {
	test()
}
