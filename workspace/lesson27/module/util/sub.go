package util2

import "fmt"

func init() {
    fmt.Println("sub init")
}

func Sub(a, b int) int {
    return a-b
}
