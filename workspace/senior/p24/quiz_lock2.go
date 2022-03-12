// quiz_lock2.go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m sync.Mutex
	fmt.Print("1, ")
	m.Lock()

	go func() {
		time.Sleep(200 * time.Millisecond)
		m.Unlock()
	}()

	m.Lock()
	fmt.Println("2")
}
