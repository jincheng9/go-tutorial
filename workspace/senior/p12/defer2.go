package main

import "fmt"

func main() {
	defer func() {
		r := recover()
		fmt.Println(r)
	}()
	var f func(int)
	defer f(1)
	f = func(a int) {}
}
