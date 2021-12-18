package main

import "fmt"

func testAddr() [2]int {
	return [2]int{1,2}
}

func main() {
	fmt.Println(testAddr()[:])
}
