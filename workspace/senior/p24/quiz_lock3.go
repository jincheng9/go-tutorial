// quiz_lock3.go

package main

import (
	"fmt"
	"sync"
)

var a sync.Mutex

func main() {
	a.Lock()
	fmt.Print("1, ")
	a.Unlock()
	fmt.Print("2, ")
	a.Unlock()
	fmt.Println("3")
}
