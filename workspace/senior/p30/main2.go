package main

import (
	"fmt"
)

// send-only channel
func testSendChan(c chan<- int) {
	c <- 20
}

// receive-only channel
func testRecvChan(c <-chan int) {
	result := <-c
	fmt.Println("result:", result)
}

func main() {
	ch := make(chan int, 3)
	testSendChan(ch)
	testRecvChan(ch)
}
