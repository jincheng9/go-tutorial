package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	for range nums {
		fmt.Println(nums)
	}
}
