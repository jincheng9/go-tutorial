// store.go
package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var sum int32 = 100
	var newValue int32 = 200
	// 将sum的值修改为newValue
	atomic.StoreInt32(&sum, newValue)
	// 读取修改后的sum值
	result := atomic.LoadInt32(&sum)
	// 打印结果
	fmt.Println("result=", result)
}
