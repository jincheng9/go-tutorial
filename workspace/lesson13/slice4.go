package main

import "fmt"


func main() {
	a := []int{1, 2}
	b := make([]int, 1, 3) // 切片b的长度是1

	copy(b, a) // 只拷贝1个元素到b里
	fmt.Println("a=", a) // a= [1 2]
	fmt.Println("b=", b) // b= [1]
}