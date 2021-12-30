package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type SliceHeader struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}

func main() {

	var s1 []int
	s2 := make([]int, 0)
	// s4 := make([]int, 0)
	s5 := make([]int, 1, 10)
	headerS2 := (*reflect.SliceHeader)(unsafe.Pointer(&s2))
	// headerS2.Data = uintptr(unsafe.Pointer(&s2))

	headerS5 := (*reflect.SliceHeader)(unsafe.Pointer(&s5))
	// headerS5.Data = uintptr(unsafe.Pointer(&s5))

	fmt.Println(headerS2.Data, headerS5.Data)
	// fmt.Printf("s1 pointer:%+v, \ns2 pointer:%+v, \ns5 pointer:%+v, \n", (*reflect.SliceHeader)(unsafe.Pointer(&s1)), (*reflect.SliceHeader)(unsafe.Pointer(&s2)), (*reflect.SliceHeader)(unsafe.Pointer(&s5)))
	// fmt.Println(((*reflect.SliceHeader)(unsafe.Pointer(&s1))).Data)
	fmt.Printf("%v\n", ((*reflect.SliceHeader)(unsafe.Pointer(&s1))).Data == ((*reflect.SliceHeader)(unsafe.Pointer(&s2))).Data)
	// fmt.Printf("%v\n", (*(*SliceHeader)(unsafe.Pointer(&s2))).Data == (*(*SliceHeader)(unsafe.Pointer(&s4))).Data)
	fmt.Printf("%v\n", ((*reflect.SliceHeader)(unsafe.Pointer(&s2))).Data == ((*reflect.SliceHeader)(unsafe.Pointer(&s5))).Data)
	fmt.Printf("%v\n", ((*reflect.SliceHeader)(unsafe.Pointer(&s2))).Data)
	fmt.Printf("%v\n", ((*reflect.SliceHeader)(unsafe.Pointer(&s5))).Data)
	// fmt.Printf("%v\n", ((*SliceHeader)(unsafe.Pointer(&s5))).Data)
	// fmt.Printf("%p %p %p %p\n", s1, s2, s4, s5)

	// data := &s5[0]
	// fmt.Println(data)
}
