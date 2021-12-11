// example3.go
package main

import "fmt"

func main() {
	// a和b作为局部匿名结构体变量，只是临时一次性使用
	// 注意：a是把struct里的字段声明放在同一行，字段之间要用分号分割，否则编译报错
	a := struct{name string; age int}{"Alice", 16}
	fmt.Println(a)

	b := struct{
		school string
		city string
	}{"THU", "Beijing"}
	fmt.Println(b)
}
