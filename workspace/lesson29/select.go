package main

import (
	"fmt"
	"time"
)

func a() {
	var c1, c2, c3 chan int = make(chan int), make(chan int), make(chan int)
	var i1, i2 int
	go func() {
		c1 <- 10
		i1 = <-c2
	}()
	for {
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
		default:
			print("no communication\n")
		}
		time.Sleep(2*time.Second)
	}
}


func b() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	go func() {
		for i:=0; i<10; i++ {
			ch1 <- i
			ch2 <- i
		}
	}()
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch1:
			fmt.Printf("receive %d from channel 1\n", x)
		case y := <-ch2:
			fmt.Printf("receive %d from channel 2\n", y)
		}
	}
}

func main() {
	//a()
	b()
}
