// cpu.go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	// nCPU is the number of logical cpu
	nCPU := runtime.NumCPU()
	// num is the number of currrent GOMAXPROCS
	// default is the value of runtime.NumCPU()
	num := runtime.GOMAXPROCS(0)
	// 4 4
	fmt.Println(num, nCPU)
}
