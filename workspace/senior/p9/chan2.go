package main

import "fmt"

func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	close(c)
	for i := 0; i < 10; i++ {
		x, ok := <-c
		fmt.Println(x, ok)
	}
}
