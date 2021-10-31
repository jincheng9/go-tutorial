package main

import "fmt"



func buildMap(str string, m map[rune]int) {
	/*函数内对map变量m的修改会影响main里的实参mapping*/
	for _, value := range str {
		m[value]++
	}
}


func main() {
	mapping := map[rune]int{}
	str := "abc"
	buildMap(str, mapping)

	/*
	mapping的值被buildMap修改了
	*/
	for key, value := range mapping {
		fmt.Printf("key:%v, value:%d\n", key, value)
	}
}