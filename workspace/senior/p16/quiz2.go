// quiz2.go
package main

import "fmt"

func main() {
	c := make(chan int, 1)
	c <- 1
	close(c)
	close(c)
	fmt.Println("OK")
}
