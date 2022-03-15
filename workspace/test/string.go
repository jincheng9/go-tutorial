package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "live-id"
	upperStr := strings.ToUpper(str)
	result := strings.HasPrefix(upperStr, "LIVE")
	fmt.Println(upperStr, result)

	a := string(97) // 把Unicode值转成字符串, "a"
	fmt.Println(a, a == "a")

	b := string(0x5403) // 0x5403对应的字符是中文"吃"
	fmt.Println(b)
}
