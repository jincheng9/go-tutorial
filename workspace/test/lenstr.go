package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "测试01"
	bs := []byte(str)
	rs := []rune(str)
	fmt.Println(len(str), len(bs), len(rs), utf8.RuneCountInString(str), utf8.RuneCount(bs))
}
