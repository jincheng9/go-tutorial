package main

import (
	"fmt"
	"time"
)

func a() {
	var c1, c2, c3 chan int = make(chan int, 1), make(chan int, 1), make(chan int, 1)
	var i1, i2 int
	go func() {
		c1 <- 10
		i1 = <-c2
	}()
	time.Sleep(time.Second)
	select {
	case c2 <- i2:
		print("sent ", i2, " to c2\n")
	case i1 = <-c1:
		fmt.Print("received ", i1, " from c1\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			print("received ", i3, " from c3\n")
		} else {
			print("c3 is closed\n")
		}
	// case a[f()] = <-c4:
	// same as:
	// case t := <-c4
	//	a[f()] = t
	default:
		print("no communication\n")
	}
}

func b() {
	var ch chan int
	go func() {
		fmt.Println("b1")
		ch <- 10
		fmt.Println("b2")
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("func b")
}
func main() {
	b()
}
