package util2

import (
    "fmt"
    "project/util/strings"
)

func init() {
    fmt.Println("math init")
}

func Sum(a, b int) int {
    return a+b
}

func CallReverse(str string) string {
    result := strings.Reverse(str)
    return result
}
