package main

import (
	"fmt"
	"time"
)

func main() {
	// print current time
	fmt.Println(time.Now())

	// NewTimer creates a new Timer that will send
	// the current time on its channel after at least duration d.
	timer := time.NewTimer(5 * time.Second)

	// print current time
	fmt.Println(<-timer.C)
}
