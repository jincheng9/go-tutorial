// quiz_lock1.go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Mutex
	fmt.Print("1, ")
	m.Lock()
	m.Lock()
	m.Unlock()
	fmt.Println("2")
}
