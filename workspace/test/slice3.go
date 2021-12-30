package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s2 := make([]int, 0)
	s4 := make([]int, 0)
	s5 := make([]int, 1, 10)
	s6 := make([]int, 1, 10)
	headerS2 := (*reflect.SliceHeader)(unsafe.Pointer(&s2))
	headerS4 := (*reflect.SliceHeader)(unsafe.Pointer(&s4))
	headerS5 := (*reflect.SliceHeader)(unsafe.Pointer(&s5))
	headerS6 := (*reflect.SliceHeader)(unsafe.Pointer(&s6))
	fmt.Println(headerS2.Data, headerS4.Data, headerS5.Data, headerS6.Data)

	// headerS2.Data = uintptr(unsafe.Pointer(&s2))

	// headerS5 := (*reflect.SliceHeader)(unsafe.Pointer(&s5))
	// headerS5.Data = uintptr(unsafe.Pointer(&s5))
	// fmt.Printf("%p %p %p %p\n", s2, s4, s5, s6)
	// fmt.Println(headerS2.Data, headerS5.Data)
	// fmt.Printf("%p %p %p\n", s2, s4, s5)
	// var temp = reflect.SliceHeader{Data: uintptr(unsafe.Pointer(&s5)), Len: len(s5), Cap: cap(s5)}
	// fmt.Printf("%v, %p, %p\n", temp, &s5[0], &s5)
}
