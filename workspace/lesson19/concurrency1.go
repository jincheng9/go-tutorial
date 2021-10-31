package main

import "fmt"
// import "time"
import "sync"

var sum int = 0
var wg sync.WaitGroup
var mutex sync.Mutex

func add() {
	defer wg.Done()
	mutex.Lock()
	sum++
	mutex.Unlock()
}

func main() {
	N := 100

	wg.Add(N)

	for i:=0; i<N; i++ {
		go add()
	}

	wg.Wait()
	
	fmt.Println("sum=", sum)	
}