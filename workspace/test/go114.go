package main

import "fmt"

func main() {
	defer func() {
		defer func() {
			if r := recover(); r != nil {
				result, ok := r.(int)
				fmt.Println(result, ok)
			}
		}()

		if r := recover(); r != nil {
			result, ok := r.(int)
			fmt.Println(result, ok)
			panic(r)
		}
	}()
	panic(1)
}
