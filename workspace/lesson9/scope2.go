package main

import "fmt"

func main() {
	a := 10
	{
		a := 5
		fmt.Println("a=", a)
	}
	fmt.Println("a=", a)
}