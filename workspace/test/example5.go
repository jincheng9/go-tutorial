package main

import (
	"fmt"
	"time"
)

type ST struct {
	ST2
}

type ST2 struct{

}

func(st *ST2) close(a int) int {
	return a
}

func add(a, b  int) int{
	return a+b
}

func printDuration() {
	timeCost := time.Minute + time.Second + 200 * time.Microsecond
	fmt.Printf("%13v\n", timeCost)
	if timeCost > time.Minute {
		timeCost = timeCost.Truncate(time.Second)
	}
	fmt.Printf("%13v\n", timeCost)
}

func main() {
	s := ST{ST2{}}
	fmt.Println(s.close(10))

	sl := make([]int, 3, 10)
	fmt.Println(sl)

	printDuration()
}

