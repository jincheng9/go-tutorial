package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	time1 := time.Now().UnixMilli()
	temp := strconv.FormatInt(time1, 10)
	time2 := time.Now().Format("20060102 15:04:05")
	fmt.Println(time1, temp, time2)
	var ch chan int
	i := <-ch
	fmt.Println(i)

	var j int
	j = <-ch
	fmt.Println(j)
}
