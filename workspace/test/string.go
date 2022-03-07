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
}
