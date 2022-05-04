// compare-and-swap.go
package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var dst int32 = 100
	oldValue := atomic.LoadInt32(&dst)
	var newValue int32 = 200
	// 先比较dst的值和oldValue的值，如果相等，就把dst的值替换为newValue
	swapped := atomic.CompareAndSwapInt32(&dst, oldValue, newValue)
	// 打印结果
	fmt.Printf("old value: %d, swapped value: %d, swapped success: %v\n", oldValue, dst, swapped)
}
