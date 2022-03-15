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

	a := string(97) // 根据ASCII码值转成对应的字符, "a"
	fmt.Println(a == "a")
}
