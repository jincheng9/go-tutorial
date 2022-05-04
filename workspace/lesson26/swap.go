// swap.go
package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var newValue int32 = 200
	var dst int32 = 100
	// 把dst的值替换为newValue
	old := atomic.SwapInt32(&dst, newValue)
	// 打印结果
	fmt.Println("old value: ", old, " new value:", dst)
}
