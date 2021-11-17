// example4.go
package main

import "fmt"

func initMap(data map[string]int) {
	data = make(map[string]int)
	fmt.Println("in function initMap, data == nil:", data == nil)
}

func main() {
	var data map[string]int
	fmt.Println("before init, data == nil:", data == nil)
	initMap(data)
	fmt.Println("after init, data == nil:", data == nil)
}
