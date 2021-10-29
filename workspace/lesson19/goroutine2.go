package main

import "fmt"

type Cat struct {
	name string
	age int
}

func fetchChannel(ch chan Cat) {
	fmt.Println(<-ch)
}

func main() {
	ch := make(chan Cat)
	a := Cat{"yingduan", 1}
	go fetchChannel(ch)
	ch <- a
	fmt.Println("end")
}