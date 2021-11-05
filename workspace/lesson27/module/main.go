package main

import (
    "fmt"
    "project/util"
    "project/util/strings"
    "sync"
)
func init() {
    fmt.Println("main init")
}
func main() {
    var result int
    var wg sync.WaitGroup
    wg.Add(1)
    go func() {
        defer wg.Done()
        result = util2.Sub(1, 2)
    }()
    wg.Wait()
    fmt.Println(result)
    fmt.Println(util2.Sub(1,2))

    str := "abc"
    fmt.Println(strings.Reverse(str))

    str2 := "cba"
    fmt.Println(util2.CallReverse(str2))

    str2 = "test"
    fmt.Println(str2)
}