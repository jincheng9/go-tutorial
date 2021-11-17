// example5.go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	data := make(map[string]int)
	var p uintptr
	fmt.Println("data size:", unsafe.Sizeof(data))
	fmt.Println("pointer size:", unsafe.Sizeof(p))
}
