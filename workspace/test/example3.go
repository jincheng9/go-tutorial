package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a = 1
	b := ^a
	fmt.Println(b)
	var c uint8 = 1
	d := ^c

	fmt.Println(reflect.TypeOf(d), d)
}
