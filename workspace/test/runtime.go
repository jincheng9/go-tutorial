package main

import (
	"fmt"
	"runtime"
)

func main() {
	num := runtime.GOMAXPROCS(0)
	nCPU := runtime.NumCPU()
	fmt.Println(num, nCPU)
}
