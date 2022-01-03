package main

import (
	"fmt"
	"reflect"
)

func main() {
	for i := 0; i < 2; i++ {
		fmt.Println(010 + 10i)
	}
	var a = 6_9
	fmt.Println(a, reflect.TypeOf(a))
}
