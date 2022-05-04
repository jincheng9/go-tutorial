// load.go
package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var sum int32 = 100
	result := atomic.LoadInt32(&sum)
	fmt.Println("result=", result)
}
