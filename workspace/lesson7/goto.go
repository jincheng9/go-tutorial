package main

import "fmt"

func main() {
	LOOP: 
		println("Enter your age:")
		var age int
		_, err := fmt.Scan(&age) // 接受控制台输入
		if err != nil {
			println("error:", err)
			goto LOOP
		}
		if age < 18 {
			println("You are not eligible to vote!")
			goto LOOP
		} else {
			println("You are eligible to vote!")
		}
		println("all finish")
}