package main

import "fmt"

func change1(param []int) {
	param[0] = 100             // 这个会改变外部切片的值
	param = append(param, 200) // append不会改变外部切片的值
}

func change2(param *[]int) {
	*param = append(*param, 300) // 传切片指针，通过这种方式append可以改变外部切片的值
}

func main() {
	slice := make([]int, 2, 100)
	fmt.Println(slice) // [0, 0]

	change1(slice)
	fmt.Println(slice) // [100, 0]

	change2(&slice)
	fmt.Println(slice) // [100, 0, 300]
}
